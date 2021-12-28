package integration_tests

import (
	"fmt"
	"github.com/cloudquery/cq-provider-azure/resources/services/resources"
	"testing"

	"github.com/Masterminds/squirrel"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationResourcesGroups(t *testing.T) {
	table := resources.ResourcesGroups()

	awsTestIntegrationHelper(t, table, []string{}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"name": fmt.Sprintf("resource-group-%s%s", res.Prefix, res.Suffix),
					"tags": map[string]interface{}{
						"Type":   "integration_test",
						"TestId": res.Suffix,
					},
				},
			}},
		}
	})
}
