package resources_test

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"reflect"
	"testing"

	"github.com/cloudquery/faker/v3"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
)

const testSubscriptionID = "test_sub"

func azureTestHelper(t *testing.T, table *schema.Table, builder func(*testing.T, *gomock.Controller) services.Services) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	providertest.TestResource(t, resources.Provider, providertest.ResourceTestData{
		Table:  table,
		Config: client.Config{Subscriptions: []string{testSubscriptionID}},
		Configure: func(logger hclog.Logger, i interface{}) (schema.ClientMeta, error) {
			c := client.NewAzureClient(logging.New(&hclog.LoggerOptions{
				Level: hclog.Warn,
			}), []string{testSubscriptionID})
			c.SetSubscriptionServices(testSubscriptionID, builder(t, ctrl))
			return c, nil
		},
	})
}

func createTestClient(adress string) http.Client {
	transport := http.Transport{
		DialContext: func(ctx context.Context, network string, addr string) (net.Conn, error) {
			var d net.Dialer
			return d.DialContext(ctx, network, adress)
		},
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	return http.Client{
		Transport: &transport,
	}
}

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
