//go:generate mockgen -destination=./mocks/subscriptions.go -package=mocks . SubscriptionsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/Azure/go-autorest/autorest"
)

type Subscriptions struct {
	SubscriptionID string
	Subscriptions  SubscriptionsClient
	Tenants        TenantsClient
}

type SubscriptionsClient interface {
	Get(ctx context.Context, subscriptionID string) (result subscription.Model, err error)
	ListLocations(ctx context.Context, subscriptionID string) (result subscription.LocationListResult, err error)
}

type TenantsClient interface {
	ListComplete(ctx context.Context) (result subscription.TenantListResultIterator, err error)
}

func NewSubscriptionsClient(subscriptionId string, auth autorest.Authorizer) Subscriptions {
	s := subscription.NewSubscriptionsClient()
	s.Authorizer = auth

	t := subscription.NewTenantsClient()
	t.Authorizer = auth

	return Subscriptions{
		SubscriptionID: subscriptionId,
		Subscriptions:  s,
		Tenants:        t,
	}
}
