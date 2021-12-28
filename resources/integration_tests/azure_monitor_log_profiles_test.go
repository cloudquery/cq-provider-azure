package integration_tests

import (
	"fmt"
	"github.com/cloudquery/cq-provider-azure/resources/services/monitor"
	"testing"

	"github.com/Masterminds/squirrel"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationMonitorLogProfiles(t *testing.T) {
	awsTestIntegrationHelper(t, monitor.MonitorLogProfiles(), []string{
		"azure_monitor_log_profiles.tf",
		"azure_storage_accounts.tf",
		"networks.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: monitor.MonitorLogProfiles().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s-%s-log-profile", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"retention_policy_days": float64(7),
					"categories": []interface{}{
						"Delete",
						"Action",
						"Write",
					},
					"retention_policy_enabled": true,
				},
			}},
		}
	})
}
