package integration_tests

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-azure/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"testing"
)

func TestIntegrationKeyvaultVaults(t *testing.T) {
	awsTestIntegrationHelper(t, resources.KeyvaultVaults(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.KeyvaultVaults().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("%s-%s-vm", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"hardware_profile_vm_size": "Standard_B1ls",
					"computer_name":            "hostname",
					"admin_username":           "testadmin",
					"type":                     "Microsoft.Compute/virtualMachines",
					"windows_configuration_provision_vm_agent": true,
					"allow_extension_operations":               true,
					"require_guest_provision_signal":           true,
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "azure_compute_virtual_machine_network_interfaces",
					ForeignKeyName: "virtual_machine_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"network_interface_reference_properties_primary": false,
						},
					}},
				},
				{
					Name:           "azure_compute_virtual_machine_resources",
					ForeignKeyName: "virtual_machine_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"name": fmt.Sprintf("vm-extension-%s-%s", res.Prefix, res.Suffix),
							"type": "Microsoft.Compute/virtualMachines/extensions",
							"virtual_machine_extension_properties": map[string]interface{}{
								"type": "CustomScript",
								"settings": map[string]interface{}{
									"commandToExecute": "hostname && uptime",
								},
								"publisher":               "Microsoft.Azure.Extensions",
								"typeHandlerVersion":      "2.0",
								"autoUpgradeMinorVersion": false,
							},
							"tags": map[string]interface{}{
								"test": "test",
							},
						},
					}},
				},
			},
		}
	})
}
