package subscription

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource tenants --config resources/services/subscription/gen.hcl --output .
func Tenants() *schema.Table {
	return &schema.Table{
		Name:         "azure_subscription_tenants",
		Description:  "Azure tenant information",
		Resolver:     fetchSubscriptionTenants,
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
				Description: "The fully qualified ID of the tenant",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "tenant_id",
				Description: "The tenant ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("TenantID"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchSubscriptionTenants(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().Subscriptions
	m, err := svc.Tenants.ListComplete(ctx)
	if err != nil {
		return diag.WrapError(err)
	}
	for m.NotDone() {
		res <- m.Value()
		err = m.NextWithContext(ctx)
		if err != nil {
			return diag.WrapError(err)
		}
	}
	return nil
}
