package subscription

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSubscriptionsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockSubscriptionsClient(ctrl)

	var subscriptionID string
	if err := faker.FakeData(&subscriptionID); err != nil {
		t.Fatal(err)
	}

	var model subscription.Model
	if err := faker.FakeData(&model); err != nil {
		t.Fatal(err)
	}
	rg := client.FakeResourceGroup
	model.SubscriptionID = &rg
	m.EXPECT().Get(gomock.Any(), subscriptionID).Return(model, nil)

	return services.Services{
		Subscriptions: services.Subscriptions{
			SubscriptionID: subscriptionID,
			Subscriptions:  m,
		},
	}
}

func TestSubscriptions(t *testing.T) {
	client.AzureMockTestHelper(t, Subscriptions(), buildSubscriptionsMock, client.TestOptions{})
}
