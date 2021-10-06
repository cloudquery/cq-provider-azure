package resources_test

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/cloudquery/faker/v3"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"
)

func fakeSkipFields(tst *testing.T, data interface{}, skipFields []string) {
	skipMap := make(map[string]struct{})
	for _, s := range skipFields {
		skipMap[s] = struct{}{}
	}
	v := reflect.ValueOf(data)
	ind := reflect.Indirect(v)
	s := ind.Type()

	for i := 0; i < s.NumField(); i++ {
		if _, ok := skipMap[s.Field(i).Name]; !ok {
			ifc := ind.Field(i).Interface()
			if err := faker.FakeData(&ifc); err != nil {
				tst.Fatal(err)
			}
			ind.Field(i).Set(reflect.ValueOf(ifc))
		}
	}
}

func createADGroupsTestServer(t *testing.T) (*msgraph.GraphServiceRequestBuilder, error) {
	var group msgraph.Group
	fakeSkipFields(t, &group, []string{
		"Conversations",
		"Threads",
		"CalendarView",
		"Events",
		"Calendar",
		"Drive",
		"Drives",
		"Sites",
		"Extensions",
		"Planner",
		"Onenote",
	})
	mux := httprouter.New()
	mux.GET("/v1.0/groups", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		groups := []msgraph.Group{
			group,
		}

		value, err := json.Marshal(groups)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}

		resp := msgraph.Paging{
			NextLink: "",
			Value:    value,
		}

		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}

		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})

	ts := httptest.NewTLSServer(mux)
	u, _ := url.Parse(ts.URL)

	transport := http.Transport{
		DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, network, u.Host)
		},
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Transport: &transport,
	}

	svc := msgraph.NewClient(&client)

	return svc, nil
}

func TestADGroups(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: resources.AdGroups(),
		Config: client.Config{
			Subscriptions: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			graph, err := createADGroupsTestServer(t)
			if err != nil {
				return nil, err
			}
			c := client.NewAzureClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{"testProject"})
			c.SetSubscriptionServices("testProject", services.Services{
				Graph: graph,
			})
			return c, nil
		},
	}
	providertest.TestResource(t, resources.Provider, resource)
}
