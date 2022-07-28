package subscription

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource subscriptions --config gen.hcl --output .
func Subscriptions() *schema.Table {
	return &schema.Table{
		Name:         "azure_subscription_subscriptions",
		Description:  "Azure subscription information",
		Resolver:     fetchSubscriptionSubscriptions,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "authorization_source",
				Description: "The authorization source of the request",
				Type:        schema.TypeString,
			},
			{
				Name:        "location_placement_id",
				Description: "The subscription location placement ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubscriptionPolicies.LocationPlacementID"),
			},
			{
				Name:        "quota_id",
				Description: "The subscription quota ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubscriptionPolicies.QuotaID"),
			},
			{
				Name:        "spending_limit",
				Description: "The subscription spending limit",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SubscriptionPolicies.SpendingLimit"),
			},
			{
				Name:        "tags",
				Description: "The tags attached to the subscription",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "display_name",
				Description: "The subscription display name",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The fully qualified ID for the subscription",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "state",
				Description: "The subscription state",
				Type:        schema.TypeString,
			},
			{
				Name:        "tenant_id",
				Description: "The subscription tenant ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TenantID"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_subscription_subscription_managed_by_tenants",
				Description: "Information about a tenant managing the subscription",
				Resolver:    fetchSubscriptionSubscriptionManagedByTenants,
				Columns: []schema.Column{
					{
						Name:        "subscription_cq_id",
						Description: "Unique CloudQuery ID of azure_subscription_subscriptions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "tenant_id",
						Description: "The tenant ID of the managing tenant",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("TenantID"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSubscriptionSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Subscriptions
	pager := svc.Subscriptions.NewListPager(nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
		for _, v := range nextResult.Value {
			res <- v
		}
	}
	return nil
}
func fetchSubscriptionSubscriptionManagedByTenants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
