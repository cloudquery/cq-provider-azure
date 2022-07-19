package frontdoor

import (
	"context"
	"reflect"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func fakeBasicRouteConfiguration(v reflect.Value) (interface{}, error) {
	var config frontdoor.RouteConfiguration
	if err := faker.FakeData(&config); err != nil {
		return nil, err
	}
	return frontdoor.BasicRouteConfiguration(config), nil
}

func buildFrontDoorsServices(t *testing.T, ctrl *gomock.Controller) services.Services {
	if err := faker.AddProvider("routeConfigurationOverride", fakeBasicRouteConfiguration); err != nil {
		t.Fatal(err)
	}
	if err := faker.AddProvider("routeConfiguration", fakeBasicRouteConfiguration); err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := faker.RemoveProvider("routeConfigurationOverride"); err != nil {
			t.Fatal(err)
		}
		if err := faker.RemoveProvider("routeConfiguration"); err != nil {
			t.Fatal(err)
		}
	}()

	m := mocks.NewMockFrontDoorClient(ctrl)
	var frontDoor frontdoor.FrontDoor

	faker.SetIgnoreInterface(true)
	defer faker.SetIgnoreInterface(false)

	if err := faker.FakeData(&frontDoor); err != nil {
		t.Fatal(err)
	}
	id := client.FakeResourceGroup + "/" + *frontDoor.ID
	frontDoor.ID = &id
	m.EXPECT().List(gomock.Any()).Return(
		frontdoor.NewListResultPage(
			frontdoor.ListResult{Value: &[]frontdoor.FrontDoor{frontDoor}},
			func(c context.Context, lr frontdoor.ListResult) (frontdoor.ListResult, error) {
				return frontdoor.ListResult{}, nil
			},
		),
		nil,
	)

	return services.Services{FrontDoor: m}
}

func TestFrontDoors(t *testing.T) {
	table := FrontDoors()
	client.AzureMockTestHelper(t, table, buildFrontDoorsServices, client.TestOptions{})
}
