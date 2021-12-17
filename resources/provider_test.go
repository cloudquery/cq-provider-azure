package resources_test

import (
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/cq-provider-sdk/logging"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/golang/mock/gomock"
	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/assert"
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

func TestTableIdentifiers(t *testing.T) {
	t.Parallel()
	for _, res := range resources.Provider().ResourceMap {
		res := res
		t.Run(res.Name, func(t *testing.T) {
			testTableIdentifiers(t, res)
		})
	}
}

func testTableIdentifiers(t *testing.T, table *schema.Table) {
	const maxIdentifierLength = 63 // maximum allowed identifier length is 63 bytes https://www.postgresql.org/docs/13/limits.html

	assert.NotEmpty(t, table.Name)
	assert.LessOrEqual(t, len(table.Name), maxIdentifierLength, "Table name too long")

	for _, c := range table.Columns {
		assert.NotEmpty(t, c.Name)
		assert.LessOrEqual(t, len(c.Name), maxIdentifierLength, "Column name too long:", c.Name)
	}

	for _, res := range table.Relations {
		res := res
		t.Run(res.Name, func(t *testing.T) {
			testTableIdentifiers(t, res)
		})
	}
}
