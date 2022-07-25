//go:generate mockgen -destination=./mocks/subscriptions.go -package=mocks . SubscriptionsClient,TenantsClient
package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
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
	NewListPager(options *armsubscriptions.TenantsClientListOptions) *runtime.Pager[armsubscriptions.TenantsClientListResponse]
}

func NewSubscriptionsClient(subscriptionId string, auth autorest.Authorizer, azCred azcore.TokenCredential) (Subscriptions, error) {
	s := subscription.NewSubscriptionsClient()
	s.Authorizer = auth

	t, err := armsubscriptions.NewTenantsClient(azCred, nil)
	if err != nil {
		return Subscriptions{}, err
	}

	return Subscriptions{
		SubscriptionID: subscriptionId,
		Subscriptions:  s,
		Tenants:        t,
	}, nil
}
