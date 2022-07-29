package monitor

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/2020-09-01/monitor/mgmt/insights"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildActivityLogAlertsMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	alertsSvc := mocks.NewMockActivityLogAlertsClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			ActivityLogAlerts: alertsSvc,
		},
	}
	alert := insights.ActivityLogAlertResource{}
	err := faker.FakeData(&alert)
	if err != nil {
		t.Errorf("failed building mock %s", err)
	}

	page := insights.ActivityLogAlertList{Value: &[]insights.ActivityLogAlertResource{alert}}
	alertsSvc.EXPECT().ListBySubscriptionID(gomock.Any()).Return(page, nil)

	return s
}

func TestActivityLogAlerts(t *testing.T) {
	client.AzureMockTestHelper(t, MonitorActivityLogAlerts(), buildActivityLogAlertsMock, client.TestOptions{})
}
