package subscription

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildTenantsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockTenantsClient(ctrl)

	var subscriptionID string
	if err := faker.FakeData(&subscriptionID); err != nil {
		t.Fatal(err)
	}

	var model subscription.TenantIDDescription
	if err := faker.FakeData(&model); err != nil {
		t.Fatal(err)
	}
	m.EXPECT().ListComplete(gomock.Any()).Return(
		subscription.NewTenantListResultIterator(
			subscription.NewTenantListResultPage(subscription.TenantListResult{Value: &[]subscription.TenantIDDescription{
				model,
			}},
				func(c context.Context, lr subscription.TenantListResult) (subscription.TenantListResult, error) {
					return subscription.TenantListResult{}, nil
				},
			),
		),
		nil,
	)

	return services.Services{
		Subscriptions: services.Subscriptions{
			SubscriptionID: subscriptionID,
			Tenants:        m,
		},
	}
}

func TestTenants(t *testing.T) {
	client.AzureMockTestHelper(t, Tenants(), buildTenantsMock, client.TestOptions{})
}
