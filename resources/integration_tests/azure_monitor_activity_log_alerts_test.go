package integration_tests

import (
	"fmt"
	"github.com/cloudquery/cq-provider-azure/resources/services/monitor"
	"testing"

	"github.com/Masterminds/squirrel"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationMonitorActivityLogAlerts(t *testing.T) {
	awsTestIntegrationHelper(t, monitor.MonitorActivityLogAlerts(), []string{
		"azure_monitor_activity_log_alerts.tf",
		"azure_storage_accounts.tf",
		"networks.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: monitor.MonitorActivityLogAlerts().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s-%s-activitylogalert", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"enabled":     true,
					"description": "This alert will monitor a specific storage account updates.",
					"type":        "Microsoft.Insights/ActivityLogAlerts",
					"location":    "Global",
					"tags": map[string]interface{}{
						"test": "test",
					},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_monitor_activity_log_alert_conditions",
					ForeignKeyName: "activity_log_alert_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"field":  "category",
							"equals": "Recommendation",
						},
					}},
				},
				{
					Name:           "azure_monitor_activity_log_alert_conditions",
					ForeignKeyName: "activity_log_alert_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"field":  "operationName",
							"equals": "Microsoft.Storage/storageAccounts/write",
						},
					}},
				},
				{
					Name:           "azure_monitor_activity_log_alert_conditions",
					ForeignKeyName: "activity_log_alert_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"field": "resourceId",
						},
					}},
				},
				{
					Name:           "azure_monitor_activity_log_alert_action_groups",
					ForeignKeyName: "activity_log_alert_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"webhook_properties": map[string]interface{}{"from": "terraform"},
						},
					},
					},
				},
			},
		}
	})
}
