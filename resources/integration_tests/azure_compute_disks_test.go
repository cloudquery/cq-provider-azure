package integration_tests

import (
	"fmt"
	"github.com/cloudquery/cq-provider-azure/resources/services/compute"
	"testing"

	"github.com/Masterminds/squirrel"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeDisks(t *testing.T) {
	awsTestIntegrationHelper(t, compute.ComputeDisks(), []string{
		"azure_compute_disks.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: compute.ComputeDisks().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("disks-%s-%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"type": "Microsoft.Compute/disks",
				},
			}},
		}
	})
}
