package resources

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-08-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func NetworkWatchers() *schema.Table {
	return &schema.Table{
		Name:        "azure_network_watchers",
		Description: "Azure mysql server",
		Resolver:    fetchNetworkWatchers,
		Multiplex:   client.SubscriptionMultiplex,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:     "flow_log_status",
				Type:     schema.TypeBool,
				Resolver: resolveNetworkWatcherFlowLogStatus,
			},
			{
				Name:        "etag",
				Description: "A unique read-only string that changes whenever the resource is updated",
				Type:        schema.TypeString,
			},
			{
				Name:        "watcher_properties_format_provisioning_state",
				Description: "The provisioning state of the network watcher resource Possible values include: 'Succeeded', 'Updating', 'Deleting', 'Failed'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("WatcherPropertiesFormat.ProvisioningState"),
			},
			{
				Name:        "resource_id",
				Description: "Resource ID",
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
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchNetworkWatchers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Network.Watchers
	result, err := svc.ListAll(ctx)
	if err != nil {
		return err
	}
	res <- *result.Value
	return nil
}
func resolveNetworkWatcherFlowLogStatus(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(network.Watcher)
	if !ok {
		return fmt.Errorf("expected to have network.Watcher but got %T", p)
	}
	resourceDetails, err := client.ParseResourceID(*p.ID)
	svc := meta.(*client.Client).Services().Network.Watchers
	result, err := svc.GetFlowLogStatus(ctx, resourceDetails.ResourceGroup, *p.Name, network.FlowLogStatusParameters{})
	if err != nil {
		return err
	}

	var properties network.FlowLogInformation
	switch svc.(type) {
	case network.WatchersClient:
		if properties, err = result.Result(svc.(network.WatchersClient)); err != nil {
			return err
		}
	default:
		if properties, err = result.Result(network.WatchersClient{}); err != nil {
			return err
		}
	}
	if properties.FlowLogProperties == nil {
		return nil
	}
	return resource.Set(c.Name, properties.FlowLogProperties.Enabled)
}
