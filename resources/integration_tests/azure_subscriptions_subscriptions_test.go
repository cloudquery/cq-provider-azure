package integration_tests

import (
	"github.com/cloudquery/cq-provider-azure/resources/services/subscription"
	"testing"

	"github.com/Masterminds/squirrel"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSubscriptionsSubscriptions(t *testing.T) {
	table := subscription.SubscriptionSubscriptions()

	awsTestIntegrationHelper(t, table, []string{}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"state": "Enabled",
				},
			}},
		}
	})
}
