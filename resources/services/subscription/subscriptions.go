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
				Name:        "id",
				Description: "The fully qualified ID for the subscription",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "display_name",
				Description: "The subscription display name",
				Type:        schema.TypeString,
			},
			{
				Name:        "state",
				Description: "The subscription state",
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
				Name:        "authorization_source",
				Description: "The authorization source of the request",
				Type:        schema.TypeString,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSubscriptionSubscriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Subscriptions
	m, err := svc.Subscriptions.Get(ctx, svc.SubscriptionID)
	if err != nil {
		return diag.WrapError(err)
	}
	res <- m
	return nil
}
