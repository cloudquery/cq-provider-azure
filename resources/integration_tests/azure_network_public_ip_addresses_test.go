package integration_tests

import (
	"fmt"
	"github.com/cloudquery/cq-provider-azure/resources/services/network"
	"testing"

	"github.com/Masterminds/squirrel"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationNetworkPublicIPAddresses(t *testing.T) {
	awsTestIntegrationHelper(t, network.NetworkPublicIPAddresses(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: network.NetworkPublicIPAddresses().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s-%s-ip", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"provisioning_state":          "Succeeded",
					"idle_timeout_in_minutes":     float64(4),
					"tags":                        map[string]interface{}{"environment": "Production"},
					"type":                        "Microsoft.Network/publicIPAddresses",
					"sku_name":                    "Basic",
					"sku_tier":                    "Regional",
					"public_ip_allocation_method": "Static",
					"public_ip_address_version":   "IPv4",
				},
			}},
		}
	})
}
