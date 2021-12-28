package integration_tests

import (
	"github.com/cloudquery/cq-provider-azure/resources/services/security"
	"testing"

	"github.com/Masterminds/squirrel"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSecurityPricings(t *testing.T) {
	awsTestIntegrationHelper(t, security.SecurityPricings(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: security.SecurityPricings().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{
					squirrel.Eq{"name": "VirtualMachines"},
					squirrel.Eq{"pricing_properties_tier": "Free"},
				})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"resource_type": "Microsoft.Security/pricings",
				},
			}},
		}
	})
}
