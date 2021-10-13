package resources_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/hashicorp/go-hclog"
	"github.com/julienschmidt/httprouter"
	msgraph "github.com/yaegashi/msgraph.go/v1.0"

	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/resources"
)

func createADApplicationsTestServer(t *testing.T) (*msgraph.GraphServiceRequestBuilder, error) {
	var application msgraph.Application
	fakeSkipFields(t, &application, []string{})
	mux := httprouter.New()
	mux.GET("/v1.0/applications", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		groups := []msgraph.Application{
			application,
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
	client := createTestClient(u.Host)
	svc := msgraph.NewClient(&client)
	return svc, nil
}

func TestADApplications(t *testing.T) {
	resource := providertest.ResourceTestData{
		Table: resources.AdApplications(),
		Config: client.Config{
			Subscriptions: []string{"testProject"},
		},
		Configure: func(logger hclog.Logger, _ interface{}) (schema.ClientMeta, error) {
			graph, err := createADApplicationsTestServer(t)
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
