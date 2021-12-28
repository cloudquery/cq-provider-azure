package integration_tests

import (
	"fmt"
	"github.com/cloudquery/cq-provider-azure/resources/services/web"
	"testing"

	"github.com/Masterminds/squirrel"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationWebApps(t *testing.T) {
	awsTestIntegrationHelper(t, web.WebApps(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: web.WebApps().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("as-%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"reserved":                     false,
					"is_xenon":                     false,
					"hyper_v":                      false,
					"key_vault_reference_identity": "SystemAssigned",
					"https_only":                   false,
					"redundancy_mode":              "None",
					"storage_account_required":     false,
					"identity_type":                "SystemAssigned",
					"kind":                         "app",
					"type":                         "Microsoft.Web/sites",
					"container_size":               float64(0),
					"tags": map[string]interface{}{
						"test": "test",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_web_app_host_name_ssl_states",
					ForeignKeyName: "app_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"host_type": "Standard",
							"name":      fmt.Sprintf("as-%s%s.azurewebsites.net", res.Prefix, res.Suffix),
							"ssl_state": "Disabled",
						},
					}},
				},
				{
					Name:           "azure_web_app_host_name_ssl_states",
					ForeignKeyName: "app_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"host_type": "Repository",
							"name":      fmt.Sprintf("as-%s%s.scm.azurewebsites.net", res.Prefix, res.Suffix),
							"ssl_state": "Disabled",
						},
					}},
				},
				{
					Name:           "azure_web_app_publishing_profiles",
					ForeignKeyName: "app_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 2,
						Data: map[string]interface{}{
							"publish_url": fmt.Sprintf("as-%s%s.scm.azurewebsites.net:443", res.Prefix, res.Suffix),
							"user_name":   fmt.Sprintf("$as-%s%s", res.Prefix, res.Suffix),
						},
					}},
				},
				{
					Name:           "azure_web_app_auth_settings",
					ForeignKeyName: "app_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name":                            "authsettings",
							"default_provider":                "MicrosoftAccount",
							"token_store_enabled":             false,
							"token_refresh_extension_hours":   float64(72),
							"microsoft_account_client_id":     "x",
							"microsoft_account_client_secret": "x",
							"type":                            "Microsoft.Web/sites/config",
							"enabled":                         true,
						},
					}},
				},
			},
		}
	})
}
