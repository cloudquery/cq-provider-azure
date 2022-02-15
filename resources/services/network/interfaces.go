// package network

// import (
// 	"context"
// 	"fmt"

// 	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
// 	"github.com/cloudquery/cq-provider-azure/client"
// 	"github.com/cloudquery/cq-provider-sdk/provider/schema"
// )

// func NetworkInterfaces() *schema.Table {
// 	return &schema.Table{
// 		Name:          "azure_network_interfaces",
// 		Description:   "Interfaces network interfaces.",
// 		Resolver:      fetchNetworkInterfaces,
// 		Multiplex:     client.SubscriptionMultiplex,
// 		DeleteFilter:  client.DeleteSubscriptionFilter,
// 		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
// 		IgnoreInTests: true,
// 		Columns: []schema.Column{
// 			{
// 				Name:        "subscription_id",
// 				Description: "Azure subscription id",
// 				Type:        schema.TypeString,
// 				Resolver:    client.ResolveAzureSubscription,
// 			},
// 			{
// 				Name:        "id",
// 				Description: "Resource Id",
// 				Type:        schema.TypeString,
// 				Resolver:    schema.PathResolver("ID"),
// 			},
// 			{
// 				Name:        "name",
// 				Description: "Resource name",
// 				Type:        schema.TypeString,
// 			},
// 			{
// 				Name:        "location",
// 				Description: "Location where the resource is stored",
// 				Type:        schema.TypeString,
// 			},
// 			{
// 				Name:        "type",
// 				Description: "Resource type",
// 				Type:        schema.TypeString,
// 			},
// 			{
// 				Name:        "tags",
// 				Description: "Resource tags.",
// 				Type:        schema.TypeJSON,
// 			},
// 			{
// 				Name:          "extended_location_name",
// 				Description:   "The name of the extended location.",
// 				Type:          schema.TypeString,
// 				Resolver:      schema.PathResolver("ExtendedLocation.Name"),
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:          "extended_location_type",
// 				Description:   "The type of the extended location.",
// 				Type:          schema.TypeString,
// 				Resolver:      schema.PathResolver("ExtendedLocation.Type"),
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:          "dns_settings_applied_dns_servers",
// 				Description:   "list of applied dns servers",
// 				Type:          schema.TypeStringArray,
// 				Resolver:      resolveNetworkInterfaceDnsSettingsAppliedDnsServers,
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:          "dns_settings_dns_servers",
// 				Description:   "list of dns servers",
// 				Type:          schema.TypeStringArray,
// 				Resolver:      resolveNetworkInterfaceDnsSettingsDnsServers,
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:          "dns_settings_internal_dns_name_label",
// 				Description:   "Relative DNS name for this NIC used for internal communications between VMs in the same virtual network.",
// 				Type:          schema.TypeString,
// 				Resolver:      schema.PathResolver("InterfacePropertiesFormat.DNSSettings.InternalDNSNameLabel"),
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:          "dns_settings_internal_fqdn",
// 				Description:   "Fully qualified DNS name supporting internal communications between VMs in the same virtual network.",
// 				Type:          schema.TypeString,
// 				Resolver:      schema.PathResolver("InterfacePropertiesFormat.DNSSettings.InternalFqdn"),
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:        "dns_settings_internal_domain_name_suffix",
// 				Description: "Suffix for internal domain name",
// 				Type:        schema.TypeString,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.DNSSettings.InternalDomainNameSuffix"),
// 			},
// 			{
// 				Name:        "dscp_resource_id",
// 				Description: "resource id for dns settings dscp",
// 				Type:        schema.TypeString,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.DscpConfiguration.ID"),
// 			},
// 			{
// 				Name:        "enable_accelerated_networking",
// 				Description: "If the network interface is accelerated networking enabled.",
// 				Type:        schema.TypeBool,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.EnableAcceleratedNetworking"),
// 			},
// 			{
// 				Name:        "enable_ip_forwarding",
// 				Description: "Indicates whether IP forwarding is enabled on this network interface.",
// 				Type:        schema.TypeBool,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.EnableIPForwarding"),
// 			},
// 			{
// 				Name:        "hosted_workloads",
// 				Description: "A list of references to linked BareMetal resources.",
// 				Type:        schema.TypeStringArray,
// 				Resolver:    resolveNetworkInterfaceHostedWorkloads,
// 			},
// 			{
// 				Name:        "mac_address",
// 				Description: "The MAC address of the network interface.",
// 				Type:        schema.TypeString,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.MacAddress"),
// 			},
// 			{
// 				Name:        "migration_phase",
// 				Description: "Migration phase of Network Interface resource.",
// 				Type:        schema.TypeString,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.MigrationPhase"),
// 			},
// 			{
// 				Name:          "security_group_resource_guid",
// 				Description:   "The resource GUID property of the network interface security group resource.",
// 				Type:          schema.TypeString,
// 				Resolver:      schema.PathResolver("InterfacePropertiesFormat.NetworkSecurityGroup.SecurityGroupPropertiesFormat.ResourceGUID"),
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:        "security_group_provisioning_state",
// 				Description: "The provisioning state of the network interface security group resource",
// 				Type:        schema.TypeString,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.NetworkSecurityGroup.SecurityGroupPropertiesFormat.ProvisioningState"),
// 			},
// 			{
// 				Name:          "security_group_etag",
// 				Description:   "A unique read-only string that changes whenever the resource is updated.",
// 				Type:          schema.TypeString,
// 				Resolver:      schema.PathResolver("InterfacePropertiesFormat.NetworkSecurityGroup.Etag"),
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:          "security_group_id",
// 				Description:   "Resource ID.",
// 				Type:          schema.TypeString,
// 				Resolver:      schema.PathResolver("InterfacePropertiesFormat.NetworkSecurityGroup.ID"),
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:          "security_group_name",
// 				Description:   "Resource name.",
// 				Type:          schema.TypeString,
// 				Resolver:      schema.PathResolver("InterfacePropertiesFormat.NetworkSecurityGroup.Name"),
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:          "security_group_type",
// 				Description:   "Resource type.",
// 				Type:          schema.TypeString,
// 				Resolver:      schema.PathResolver("InterfacePropertiesFormat.NetworkSecurityGroup.Type"),
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:          "security_group_location",
// 				Description:   "Resource location.",
// 				Type:          schema.TypeString,
// 				Resolver:      schema.PathResolver("InterfacePropertiesFormat.NetworkSecurityGroup.Location"),
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:          "security_group_tags",
// 				Description:   "Resource tags.",
// 				Type:          schema.TypeJSON,
// 				Resolver:      schema.PathResolver("InterfacePropertiesFormat.NetworkSecurityGroup.Tags"),
// 				IgnoreInTests: true,
// 			},
// 			{
// 				Name:        "primary",
// 				Description: "Whether this is a primary network interface on a virtual machine.",
// 				Type:        schema.TypeBool,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.Primary"),
// 			},
// 			{
// 				Name:        "private_endpoint_id",
// 				Description: "Resource id of the private endpoint",
// 				Type:        schema.TypeString,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.PrivateEndpoint.ID"),
// 			},
// 			{
// 				Name:        "private_link_service_id",
// 				Description: "Resource id of the private link service",
// 				Type:        schema.TypeString,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.PrivateLinkService.ID"),
// 			},
// 			{
// 				Name:        "provisioning_state",
// 				Description: "The provisioning state of the network interface resource Possible values include: 'PrivateEndpointConnectionProvisioningStateSucceeded', 'PrivateEndpointConnectionProvisioningStateCreating', 'PrivateEndpointConnectionProvisioningStateDeleting', 'PrivateEndpointConnectionProvisioningStateFailed'",
// 				Type:        schema.TypeString,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.ProvisioningState"),
// 			},
// 			{
// 				Name:        "resource_guild",
// 				Description: "The resource GUID property of the network interface resource.",
// 				Type:        schema.TypeString,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.ResourceGUID"),
// 			},
// 			{
// 				Name:        "virtual_machine",
// 				Description: "resource id for virtual machine",
// 				Type:        schema.TypeString,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.VirtualMachine.ID"),
// 			},
// 			{
// 				Name:        "tap_configurations",
// 				Description: "A list of TapConfigurations of the network interface.",
// 				Type:        schema.TypeJSON,
// 				Resolver:    schema.PathResolver("InterfacePropertiesFormat.TapConfigurations"),
// 			},
// 		},
// 		Relations: []*schema.Table{
// 			{
// 				Name:          "azure_network_interface_ip_configurations",
// 				Description:   "NetworkInterface IP Configurations. ",
// 				Resolver:      fetchNetworkInterfaceIPConfigurations,
// 				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"interface_cq_id", "id"}},
// 				IgnoreInTests: true,
// 				Columns: []schema.Column{
// 					{
// 						Name:        "interface_cq_id",
// 						Description: "Unique CloudQuery ID of azure_network_interface table (FK)",
// 						Type:        schema.TypeUUID,
// 						Resolver:    schema.ParentIdResolver,
// 					},
// 					{
// 						Name:        "id",
// 						Description: "Resource Id",
// 						Type:        schema.TypeString,
// 						Resolver:    schema.PathResolver("ID"),
// 					},
// 					{
// 						Name:        "name",
// 						Description: "Resource name",
// 						Type:        schema.TypeString,
// 					},
// 					{
// 						Name:        "etag",
// 						Description: "A unique read-only string that changes whenever the resource is updated.",
// 						Type:        schema.TypeString,
// 					},
// 					{
// 						Name:        "type",
// 						Description: "Resource type",
// 						Type:        schema.TypeString,
// 					},
// 					{
// 						Name:        "primary",
// 						Description: "Whether this is a primary network interface on a virtual machine.",
// 						Type:        schema.TypeBool,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.Primary"),
// 					},
// 					{
// 						Name:        "application_gateway_backend_address_pools",
// 						Description: "The reference to ApplicationGatewayBackendAddressPool resource.",
// 						Type:        schema.TypeJSON,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.ApplicationGatewayBackendAddressPools"),
// 					},
// 					{
// 						Name:        "application_security_groups",
// 						Description: "Application security groups in which the IP configuration is included.",
// 						Type:        schema.TypeJSON,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.ApplicationSecurityGroups"),
// 					},
// 					{
// 						Name:        "load_balancer_backend_address_pools",
// 						Description: "The reference to LoadBalancerBackendAddressPool resource.",
// 						Type:        schema.TypeJSON,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.LoadBalancerBackendAddressPools"),
// 					},
// 					{
// 						Name:        "load_balancer_inbound_nat_rules",
// 						Description: "A list of references of LoadBalancerInboundNatRules.",
// 						Type:        schema.TypeJSON,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.LoadBalancerInboundNatRules"),
// 					},
// 					{
// 						Name:        "private_ip_address",
// 						Description: "Private IP address of the IP configuration.",
// 						Type:        schema.TypeString,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PrivateIPAddress"),
// 					},
// 					{
// 						Name:        "private_ip_address_version",
// 						Description: "Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4. Possible values include: 'IPVersionIPv4', 'IPVersionIPv6",
// 						Type:        schema.TypeString,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PrivateIPAddressVersion"),
// 					},
// 					{
// 						Name:        "private_ip_allocation_method",
// 						Description: "Private IP address allocation method.",
// 						Type:        schema.TypeString,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PrivateIPAllocationMethod"),
// 					},
// 					{
// 						Name:        "private_link_connection_properties",
// 						Description: "PrivateLinkConnection properties for the network interface.",
// 						Type:        schema.TypeString,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PrivateLinkConnectionProperties"),
// 					},
// 					{
// 						Name:        "provisioning_state",
// 						Description: "READ-ONLY; The provisioning state of the network interface IP configuration.",
// 						Type:        schema.TypeString,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.ProvisioningState"),
// 					},
// 					{
// 						Name:        "public_ip_address",
// 						Description: "Public IP address bound to the IP configuration.",
// 						Type:        schema.TypeString,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PublicIPAddress"),
// 					},
// 					{
// 						Name:        "subnet_id",
// 						Description: "subnet ID of network interface ip configuration",
// 						Type:        schema.TypeString,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.Subnet.ID"),
// 					},
// 					{
// 						Name:        "virtual_network_taps",
// 						Description: "The reference to Virtual Network Taps.",
// 						Type:        schema.TypeJSON,
// 						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.VirtualNetworkTaps"),
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// // ====================================================================================================================
// //                                               Table Resolver Functions
// // ====================================================================================================================

// func fetchNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
// 	svc := meta.(*client.Client).Services().Network.Interfaces
// 	response, err := svc.ListAll(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	for response.NotDone() {
// 		res <- response.Values()
// 		if err := response.NextWithContext(ctx); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
// func fetchNetworkInterfaceIPConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
// 	ni, ok := parent.Item.(network.Interface)
// 	if !ok {
// 		return fmt.Errorf("expected to have network.Interface but got %T", parent.Item)
// 	}
// 	res <- *ni.IPConfigurations
// 	return nil
// }
// func resolveNetworkInterfaceDnsSettingsAppliedDnsServers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
// 	ni, ok := resource.Item.(network.Interface)
// 	if !ok {
// 		return fmt.Errorf("expected to have network.Interface but got %T", resource.Item)
// 	}
// 	if ni.DNSSettings == nil || ni.DNSSettings.AppliedDNSServers == nil {
// 		return nil
// 	}
// 	appliedDnsServers := make([]string, 0, len(*ni.DNSSettings.AppliedDNSServers))
// 	for _, a := range *ni.DNSSettings.AppliedDNSServers {
// 		appliedDnsServers = append(appliedDnsServers, a)
// 		fmt.Println("hereeeee", a)
// 		fmt.Println("hereeeeeeeee,", &a)
// 	}
// 	return resource.Set(c.Name, appliedDnsServers)
// }
// func resolveNetworkInterfaceDnsSettingsDnsServers(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
// 	ni, ok := resource.Item.(network.Interface)
// 	if !ok {
// 		return fmt.Errorf("expected to have network.Interface but got %T", resource.Item)
// 	}
// 	if ni.DNSSettings == nil || ni.DNSSettings.DNSServers == nil {
// 		return nil
// 	}
// 	return resource.Set(c.Name, ni.DNSSettings.DNSServers)
// }
// func resolveNetworkInterfaceHostedWorkloads(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
// 	ni, ok := resource.Item.(network.Interface)
// 	if !ok {
// 		return fmt.Errorf("expected to have network.Interface but got %T", resource.Item)
// 	}
// 	if ni.HostedWorkloads == nil {
// 		return nil
// 	}
// 	return resource.Set(c.Name, ni.HostedWorkloads)
// }

package network

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkInterfaces() *schema.Table {
	return &schema.Table{
		Name:         "azure_network_interfaces",
		Description:  "Azure Network Interface",
		Resolver:     fetchNetworkInterfaces,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "extended_location_name",
				Description: "The name of the extended location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Name"),
			},
			{
				Name:        "extended_location_type",
				Description: "The type of the extended location",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ExtendedLocation.Type"),
			},
			{
				Name:        "id",
				Description: "Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "location",
				Description: "Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "name",
				Description: "Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "dns_settings_applied_dns_servers",
				Description: "The servers that are part of the same availability set.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.DNSSettings.AppliedDNSServers"),
			},
			{
				Name:        "dns_settings_dns_servers",
				Description: "List of DNS servers IP addresses.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.DNSSettings.DNSServers"),
			},
			{
				Name:        "dns_settings_internal_dns_name_label",
				Description: "The internal dns name label.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.DNSSettings.InternalDNSNameLabel"),
			},
			{
				Name:        "dns_settings_internal_domain_name_suffix",
				Description: "The internal domain name suffix.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.DNSSettings.InternalDomainNameSuffix"),
			},
			{
				Name:        "dns_settings_internal_fqdn",
				Description: "Fully qualified DNS name supporting internal communications between VMs in the same virtual network.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.DNSSettings.InternalFqdn"),
			},

			{
				Name:        "dscp_configuration_id",
				Description: "A reference to the dscp configuration to which the network interface is linked.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.DscpConfiguration.ID"),
			},
			{
				Name:        "enable_accelerated_networking",
				Description: "If the network interface is accelerated networking enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.EnableAcceleratedNetworking"),
			},
			{
				Name:        "enable_ip_forwarding",
				Description: "Indicates whether IP forwarding is enabled on this network interface.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.EnableIPForwarding"),
			},
			{
				Name:        "hosted_workloads",
				Description: "List of references to linked BareMetal resources.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.HostedWorkloads"),
			},
			// {
			// 	Name:        "ip_configurations",
			// 	Description: "A list of IPConfigurations of the network interface.",
			// 	Type:        schema.TypeJSON,
			// 	Resolver:    resolveNetworkInterfaceIPConfigurations,
			// },
			{
				Name:        "mac_address",
				Description: "The MAC address of the network interface.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.MacAddress"),
			},
			{
				Name:        "migration_phase",
				Description: "Migration phase of Network Interface resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.MigrationPhase"),
			},
			{
				Name:        "network_security_group",
				Description: "The reference to the NetworkSecurityGroup resource.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkInterfaceNetworkSecurityGroup,
			},
			{
				Name:        "nic_type",
				Description: "Type of Network Interface resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.NicType"),
			},
			{
				Name:        "primary",
				Description: "Whether this is a primary network interface on a virtual machine.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.Primary"),
			},
			{
				Name:        "private_endpoint",
				Description: "A reference to the private endpoint to which the network interface is linked.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkInterfacePrivateEndpoint,
			},
			{
				Name:        "private_link_service",
				Description: "Privatelinkservice of the network interface resource.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkInterfacePrivateLinkService,
			},
			{
				Name:        "provisioning_state",
				Description: "The provisioning state of the network interface resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "resource_guid",
				Description: "The provisioning state of the network interface resource.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.ResourceGUID"),
			},
			{
				Name:        "tap_configurations",
				Description: "A list of TapConfigurations of the network interface.",
				Type:        schema.TypeJSON,
				Resolver:    resolveNetworkInterfaceTapConfigurations,
			},
			{
				Name:        "virtual_machine_id",
				Description: "The reference to a virtual machine.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InterfacePropertiesFormat.VirtualMachine.ID"),
			},
			{
				Name:        "tags",
				Description: "Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "type",
				Description: "Resource type.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "azure_network_interface_ip_configurations",
				Description:   "NetworkInterface IP Configurations. ",
				Resolver:      fetchNetworkInterfaceIPConfigurations,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"interface_cq_id", "id"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "interface_cq_id",
						Description: "Unique CloudQuery ID of azure_network_interface table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Resource Id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "etag",
						Description: "A unique read-only string that changes whenever the resource is updated.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
					{
						Name:        "primary",
						Description: "Whether this is a primary network interface on a virtual machine.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.Primary"),
					},
					{
						Name:        "application_gateway_backend_address_pools",
						Description: "The reference to ApplicationGatewayBackendAddressPool resource.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.ApplicationGatewayBackendAddressPools"),
					},
					{
						Name:        "application_security_groups",
						Description: "Application security groups in which the IP configuration is included.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.ApplicationSecurityGroups"),
					},
					{
						Name:        "load_balancer_backend_address_pools",
						Description: "The reference to LoadBalancerBackendAddressPool resource.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.LoadBalancerBackendAddressPools"),
					},
					{
						Name:        "load_balancer_inbound_nat_rules",
						Description: "A list of references of LoadBalancerInboundNatRules.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.LoadBalancerInboundNatRules"),
					},
					{
						Name:        "private_ip_address",
						Description: "Private IP address of the IP configuration.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PrivateIPAddress"),
					},
					{
						Name:        "private_ip_address_version",
						Description: "Whether the specific IP configuration is IPv4 or IPv6. Default is IPv4. Possible values include: 'IPVersionIPv4', 'IPVersionIPv6",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PrivateIPAddressVersion"),
					},
					{
						Name:        "private_ip_allocation_method",
						Description: "Private IP address allocation method.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PrivateIPAllocationMethod"),
					},
					{
						Name:        "private_link_connection_properties",
						Description: "PrivateLinkConnection properties for the network interface.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PrivateLinkConnectionProperties"),
					},
					{
						Name:        "provisioning_state",
						Description: "READ-ONLY; The provisioning state of the network interface IP configuration.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.ProvisioningState"),
					},
					{
						Name:        "public_ip_address",
						Description: "Public IP address bound to the IP configuration.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.PublicIPAddress"),
					},
					{
						Name:        "subnet_id",
						Description: "subnet ID of network interface ip configuration",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.Subnet.ID"),
					},
					{
						Name:        "virtual_network_taps",
						Description: "The reference to Virtual Network Taps.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("InterfaceIPConfigurationPropertiesFormat.VirtualNetworkTaps"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Network.Interfaces
	response, err := svc.ListAll(ctx)
	if err != nil {
		return err
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}

func fetchNetworkInterfaceIPConfigurations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	ni, ok := parent.Item.(network.Interface)
	if !ok {
		return fmt.Errorf("expected to have network.Interface but got %T", parent.Item)
	}
	res <- *ni.IPConfigurations
	return nil
}

// func resolveNetworkInterfaceIPConfigurations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
// 	p, ok := resource.Item.(network.Interface)
// 	if !ok {
// 		return fmt.Errorf("expected to have network.Interface but got %T", resource.Item)
// 	}

// 	if p.InterfacePropertiesFormat == nil ||
// 		p.InterfacePropertiesFormat.IPConfigurations == nil {
// 		return nil
// 	}

// 	out, err := json.Marshal(p.InterfacePropertiesFormat.IPConfigurations)
// 	if err != nil {
// 		return err
// 	}
// 	return resource.Set(c.Name, out)
// }
func resolveNetworkInterfaceNetworkSecurityGroup(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.Interface)
	if !ok {
		return fmt.Errorf("expected to have network.Interface but got %T", resource.Item)
	}

	if p.InterfacePropertiesFormat == nil ||
		p.InterfacePropertiesFormat.NetworkSecurityGroup == nil {
		return nil
	}

	out, err := json.Marshal(p.InterfacePropertiesFormat.NetworkSecurityGroup)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
func resolveNetworkInterfacePrivateEndpoint(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.Interface)
	if !ok {
		return fmt.Errorf("expected to have network.Interface but got %T", resource.Item)
	}

	if p.InterfacePropertiesFormat == nil ||
		p.InterfacePropertiesFormat.PrivateEndpoint == nil {
		return nil
	}

	out, err := json.Marshal(p.InterfacePropertiesFormat.PrivateEndpoint)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
func resolveNetworkInterfacePrivateLinkService(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.Interface)
	if !ok {
		return fmt.Errorf("expected to have network.Interface but got %T", resource.Item)
	}

	if p.InterfacePropertiesFormat == nil ||
		p.InterfacePropertiesFormat.PrivateLinkService == nil {
		return nil
	}

	out, err := json.Marshal(p.InterfacePropertiesFormat.PrivateLinkService)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
func resolveNetworkInterfaceTapConfigurations(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.Interface)
	if !ok {
		return fmt.Errorf("expected to have network.Interface but got %T", resource.Item)
	}

	if p.InterfacePropertiesFormat == nil ||
		p.InterfacePropertiesFormat.TapConfigurations == nil {
		return nil
	}

	out, err := json.Marshal(p.InterfacePropertiesFormat.TapConfigurations)
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out)
}
