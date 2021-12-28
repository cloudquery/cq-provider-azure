package integration_tests

import (
	"fmt"
	"github.com/cloudquery/cq-provider-azure/resources/services/network"
	"testing"

	"github.com/Masterminds/squirrel"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationNetworkWatchers(t *testing.T) {
	awsTestIntegrationHelper(t, network.NetworkWatchers(), []string{
		"networks.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: network.NetworkWatchers().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s-%s-nw", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"provisioning_state": "Succeeded",
					"type":               "Microsoft.Network/networkWatchers",
					"tags": map[string]interface{}{
						"test": "test",
					},
				},
			}},
		}
	})
}
