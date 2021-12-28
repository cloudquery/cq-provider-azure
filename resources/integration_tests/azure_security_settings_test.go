package integration_tests

import (
	"github.com/cloudquery/cq-provider-azure/resources/services/security"
	"testing"

	"github.com/Masterminds/squirrel"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSecuritySettings(t *testing.T) {
	awsTestIntegrationHelper(t, security.SecuritySettings(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: security.SecuritySettings().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{
					squirrel.Eq{"name": "MCAS"},
					squirrel.Eq{"enabled": true},
				})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"resource_type": "Microsoft.Security/settings",
				},
			}},
		}
	})
}
