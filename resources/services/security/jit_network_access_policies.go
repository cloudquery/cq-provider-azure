package security

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/security/mgmt/v3.0/security"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SecurityJitNetworkAccessPolicies() *schema.Table {
	return &schema.Table{
		Name:         "azure_security_jit_network_access_policies",
		Description:  "Just in Time network access policy",
		Resolver:     fetchSecurityJitNetworkAccessPolicies,
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
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "Kind of the resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Location where the resource is stored",
				Type:        schema.TypeString,
			},
			{
				Name:        "provisioning_state",
				Description: "Gets the provisioning state of the Just-in-Time policy.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("JitNetworkAccessPolicyProperties.ProvisioningState"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_security_jit_network_access_policy_virtual_machines",
				Description: "Describes virtual machine attached to current jit policy",
				Resolver:    fetchSecurityJitNetworkAccessPolicyVirtualMachines,
				Columns: []schema.Column{
					{
						Name:        "jit_network_access_policy_cq_id",
						Description: "Unique CloudQuery ID of azure_security_jit_network_access_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "Resource ID of the virtual machine that is linked to this policy",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "public_ip_address",
						Description: "Public IP address of the Azure Firewall that is linked to this policy, if applicable",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("PublicIPAddress"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_security_jit_network_access_policy_virtual_machine_ports",
						Description: "Describes port rule of vm",
						Resolver:    fetchSecurityJitNetworkAccessPolicyVirtualMachinePorts,
						Columns: []schema.Column{
							{
								Name:        "jit_network_access_policy_virtual_machine_cq_id",
								Description: "Unique CloudQuery ID of azure_security_jit_network_access_policy_virtual_machines table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "number",
								Description: "Port number",
								Type:        schema.TypeInt,
							},
							{
								Name:        "protocol",
								Description: "Possible values include: 'TCP', 'UDP', 'All'",
								Type:        schema.TypeString,
							},
							{
								Name:        "allowed_source_address_prefix",
								Description: "Mutually exclusive with the \"allowedSourceAddressPrefixes\" parameter",
								Type:        schema.TypeString,
							},
							{
								Name:        "allowed_source_address_prefixes",
								Description: "Mutually exclusive with the \"allowedSourceAddressPrefix\" parameter.",
								Type:        schema.TypeStringArray,
							},
							{
								Name:        "max_request_access_duration",
								Description: "Maximum duration requests can be made for",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "azure_security_jit_network_access_policy_requests",
				Description: "Describes jit network access policy access request",
				Resolver:    fetchSecurityJitNetworkAccessPolicyRequests,
				Columns: []schema.Column{
					{
						Name:        "jit_network_access_policy_cq_id",
						Description: "Unique CloudQuery ID of azure_security_jit_network_access_policies table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:     "virtual_machines",
						Type:     schema.TypeStringArray,
						Resolver: resolveSecurityJitNetworkAccessPolicyRequestsVirtualMachines,
					},
					{
						Name:     "start_time_utc",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("StartTimeUtc.Time"),
					},
					{
						Name:        "requestor",
						Description: "The identity of the person who made the request",
						Type:        schema.TypeString,
					},
					{
						Name:        "justification",
						Description: "The justification for making the initiate request",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSecurityJitNetworkAccessPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Security.JitNetworkAccessPolicies
	response, err := svc.List(ctx)
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
func fetchSecurityJitNetworkAccessPolicyVirtualMachines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	policy, ok := parent.Item.(security.JitNetworkAccessPolicy)
	if !ok {
		return fmt.Errorf("expected to have security.JitNetworkAccessPolicy but got %T", parent.Item)
	}
	if policy.VirtualMachines == nil {
		return nil
	}

	res <- *policy.VirtualMachines
	return nil
}
func fetchSecurityJitNetworkAccessPolicyVirtualMachinePorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	vm, ok := parent.Item.(security.JitNetworkAccessPolicyVirtualMachine)
	if !ok {
		return fmt.Errorf("expected to have security.JitNetworkAccessPolicyVirtualMachine but got %T", parent.Item)
	}
	if vm.Ports == nil {
		return nil
	}

	res <- *vm.Ports

	return nil
}
func fetchSecurityJitNetworkAccessPolicyRequests(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	policy, ok := parent.Item.(security.JitNetworkAccessPolicy)
	if !ok {
		return fmt.Errorf("expected to have security.JitNetworkAccessPolicy but got %T", parent.Item)
	}
	if policy.Requests == nil {
		return nil
	}
	res <- *policy.Requests
	return nil
}
func resolveSecurityJitNetworkAccessPolicyRequestsVirtualMachines(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	vm, ok := resource.Item.(security.JitNetworkAccessRequest)
	if !ok {
		return fmt.Errorf("expected to have security.JitNetworkAccessPolicyVirtualMachine but got %T", resource.Item)
	}
	if vm.VirtualMachines == nil {
		return nil
	}

	result := make([]string, 0, len(*vm.VirtualMachines))
	for _, v := range *vm.VirtualMachines {
		result = append(result, *v.ID)
	}

	return resource.Set(c.Name, result)
}
