//go:generate mockgen -destination=./mocks/tenants.go -package=mocks . TenantsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/Azure/go-autorest/autorest"
)

type Tenants struct {
	SubscriptionID string
	Tenants        TenantsClient
}

type TenantsClient interface {
	ListComplete(ctx context.Context) (result subscription.TenantListResultIterator, err error)
}

func NewTenantsClient(subscriptionId string, auth autorest.Authorizer) Tenants {
	c := subscription.NewTenantsClient()
	c.Authorizer = auth
	return Tenants{
		SubscriptionID: subscriptionId,
		Tenants:        c,
	}
}
