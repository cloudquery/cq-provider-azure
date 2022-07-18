package servicebus

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/servicebus/mgmt/2021-06-01-preview/servicebus"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildServicebusNamespaces(t *testing.T, ctrl *gomock.Controller) services.Services {
	m := mocks.NewMockNamespacesClient(ctrl)
	tp := mocks.NewMockTopicsClient(ctrl)
	var n servicebus.SBNamespace
	if err := faker.FakeData(&n); err != nil {
		t.Fatal(err)
	}
	id := client.FakeResourceGroup
	n.ID = &id
	m.EXPECT().List(gomock.Any()).Return(
		servicebus.NewSBNamespaceListResultPage(
			servicebus.SBNamespaceListResult{Value: &[]servicebus.SBNamespace{n}},
			func(c context.Context, slr servicebus.SBNamespaceListResult) (servicebus.SBNamespaceListResult, error) {
				return servicebus.SBNamespaceListResult{}, nil
			},
		),
		nil,
	)

	var topic servicebus.SBTopic
	if err := faker.FakeData(&topic); err != nil {
		t.Fatal(err)
	}
	tp.EXPECT().ListByNamespace(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(
		servicebus.NewSBTopicListResultPage(
			servicebus.SBTopicListResult{Value: &[]servicebus.SBTopic{topic}},
			func(c context.Context, slr servicebus.SBTopicListResult) (servicebus.SBTopicListResult, error) {
				return servicebus.SBTopicListResult{}, nil
			},
		),
		nil,
	)

	return services.Services{
		Servicebus: services.ServicebusClient{
			Namespaces: m,
			Topics:     tp,
		},
	}
}

func TestServicebusNamespaces(t *testing.T) {
	client.AzureMockTestHelper(t, Namespaces(), buildServicebusNamespaces, client.TestOptions{})
}
