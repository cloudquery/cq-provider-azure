package integration_tests

import (
	"github.com/cloudquery/cq-provider-azure/resources/services/security"
	"testing"

	"github.com/Masterminds/squirrel"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationSecurityAutoProvisioningSettings(t *testing.T) {
	awsTestIntegrationHelper(t, security.SecurityAutoProvisioningSettings(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: security.SecurityAutoProvisioningSettings().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"auto_provision": "On"})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"auto_provision": "On",
					"resource_type":  "Microsoft.Security/autoProvisioningSettings",
				},
			}},
		}
	})
}
