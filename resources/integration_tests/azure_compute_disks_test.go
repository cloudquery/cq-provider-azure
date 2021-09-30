package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeDisks(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ComputeDisks(), []string{
		"azure_compute_disks.tf",
		//"azure_keyvault_vaults.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.ComputeDisks().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("disks-%s-%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"type": "Microsoft.Compute/disks",
				},
			}},
			//Relations: []*providertest.ResourceIntegrationVerification{
			//	{
			//		Name:           "azure_compute_disk_encryption_settings",
			//		ForeignKeyName: "disk_cq_id",
			//		ExpectedValues: []providertest.ExpectedValue{{
			//			Count: 1,
			//			Data: map[string]interface{}{
			//				"protocol": "Http",
			//			},
			//		}},
			//	},
			//},
		}
	})
}
