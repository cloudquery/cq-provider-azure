// //go:generate mockgen -destination=./mocks/network.go -package=mocks . VirtualNetworksClient,SecurityGroupsClient,WatchersClient,PublicIPAddressesClient,InterfacesClient
// package network

// import (
// 	"testing"

// 	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
// 	"github.com/cloudquery/cq-provider-azure/client"
// 	"github.com/cloudquery/cq-provider-azure/client/services"
// 	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
// 	"github.com/cloudquery/faker/v3"
// 	"github.com/golang/mock/gomock"
// )

// func buildNetworkInterfacesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
// 	interfaceSvc := mocks.NewMockInterfacesClient(ctrl)
// 	s := services.Services{
// 		Network: services.NetworksClient{
// 			Interfaces: interfaceSvc,
// 		},
// 	}
// 	interface := network.Interface{}
// 	err := faker.FakeData(&interface)
// 	if err != nil {
// 		t.Errorf("failed building mock %s", err)
// 	}
// 	id := client.FakeResourceGroup + "/" + *interface.ID
// 	interface.ID = &id
// 	page := network.InterfaceListResult{Value: &[]network.Interface{interface}}
// 	interfaceSvc.EXPECT().ListAll(gomock.Any()).Return(page, nil)

// 	return s
// }

// func TestNetworkInterfaces(t *testing.T) {
// 	client.AzureMockTestHelper(t, NetworkInterfaces(), buildNetworkInterfacesMock, client.TestOptions{})
// }

package network

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2020-11-01/network"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func buildNetworkInterfacesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	ic := mocks.NewMockInterfacesClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			Interfaces: ic,
		},
	}

	i := network.Interface{
		InterfacePropertiesFormat: &network.InterfacePropertiesFormat{
			NetworkSecurityGroup: &network.SecurityGroup{},
			PrivateEndpoint:      &network.PrivateEndpoint{},
			IPConfigurations:     &[]network.InterfaceIPConfiguration{},
			TapConfigurations:    &[]network.InterfaceTapConfiguration{},
			PrivateLinkService:   &network.PrivateLinkService{},
		},
	}

	require.Nil(t, faker.FakeData(&i.ID))
	require.Nil(t, faker.FakeData(&i.Etag))
	require.Nil(t, faker.FakeData(&i.ExtendedLocation))
	require.Nil(t, faker.FakeData(&i.Location))
	require.Nil(t, faker.FakeData(&i.Name))
	require.Nil(t, faker.FakeData(&i.Tags))
	require.Nil(t, faker.FakeData(&i.Type))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.DNSSettings))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.DscpConfiguration))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.HostedWorkloads))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.ResourceGUID))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.MacAddress))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.VirtualMachine))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.Primary))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.EnableAcceleratedNetworking))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.EnableIPForwarding))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.ProvisioningState))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.MigrationPhase))
	require.Nil(t, faker.FakeData(&i.InterfacePropertiesFormat.NicType))
	fakeId := client.FakeResourceGroup + "/" + *i.ID
	i.ID = &fakeId

	page := network.NewInterfaceListResultPage(network.InterfaceListResult{Value: &[]network.Interface{i}}, func(ctx context.Context, result network.InterfaceListResult) (network.InterfaceListResult, error) {
		return network.InterfaceListResult{}, nil
	})
	ic.EXPECT().ListAll(gomock.Any()).Return(page, nil)
	return s
}

func TestNetworkInterfaces(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkInterfaces(), buildNetworkInterfacesMock, client.TestOptions{})
}
