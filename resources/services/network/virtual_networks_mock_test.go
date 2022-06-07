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
)

func fakeSubnet(t *testing.T) network.Subnet {
	sb := network.Subnet{
		SubnetPropertiesFormat: &network.SubnetPropertiesFormat{},
	}
	if err := faker.FakeDataSkipFields(&sb, []string{"ProvisioningState", "SubnetPropertiesFormat"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(sb.SubnetPropertiesFormat, []string{"ApplicationGatewayIPConfigurations",
		"RouteTable",
		"NetworkSecurityGroup",
		"ServiceEndpointPolicies",
		"PrivateEndpoints",
		"IPConfigurations",
		"IPConfigurationProfiles",
		"ProvisioningState",
		"PrivateEndpointNetworkPolicies",
		"PrivateLinkServiceNetworkPolicies"}); err != nil {
		t.Fatal(err)
	}
	return sb
}

func buildNetworkVirtualNetworksMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	n := mocks.NewMockVirtualNetworksClient(ctrl)
	ngc := mocks.NewMockVirtualNetworkGatewaysClient(ctrl)
	s := services.Services{
		Network: services.NetworksClient{
			VirtualNetworks:        n,
			VirtualNetworkGateways: ngc,
		},
	}

	vn := network.VirtualNetwork{
		VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
			Subnets: &[]network.Subnet{
				fakeSubnet(t),
			},
		},
	}
	if err := faker.FakeDataSkipFields(&vn, []string{"VirtualNetworkPropertiesFormat"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeDataSkipFields(vn.VirtualNetworkPropertiesFormat, []string{"Subnets", "ProvisioningState"}); err != nil {
		t.Fatal(err)
	}
	fakeId := client.FakeResourceGroup + "/" + *vn.ID
	vn.ID = &fakeId
	vn.DhcpOptions.DNSServers = &[]string{faker.IPv4()}

	vng := network.VirtualNetworkGateway{}
	if err := faker.FakeData(&vng); err != nil {
		t.Fatal(err)
	}
	fakeVngID := client.FakeResourceGroup + "/" + *vng.ID
	vng.ID = &fakeVngID

	vngc := network.VirtualNetworkGatewayConnectionListEntity{}
	if err := faker.FakeDataSkipFields(&vngc, []string{"VirtualNetworkGatewayConnectionListEntityPropertiesFormat"}); err != nil {
		t.Fatal(err)
	}
	if err := faker.FakeData(&vngc.VirtualNetworkGatewayConnectionListEntityPropertiesFormat); err != nil {
		t.Fatal(err)
	}
	fakeVngcID := client.FakeResourceGroup + "/" + *vngc.ID
	vngc.ID = &fakeVngcID

	vnp := network.NewVirtualNetworkListResultPage(network.VirtualNetworkListResult{Value: &[]network.VirtualNetwork{vn}}, func(ctx context.Context, result network.VirtualNetworkListResult) (network.VirtualNetworkListResult, error) {
		return network.VirtualNetworkListResult{}, nil
	})
	vngp := network.NewVirtualNetworkGatewayListResultPage(network.VirtualNetworkGatewayListResult{Value: &[]network.VirtualNetworkGateway{vng}}, func(ctx context.Context, result network.VirtualNetworkGatewayListResult) (network.VirtualNetworkGatewayListResult, error) {
		return network.VirtualNetworkGatewayListResult{}, nil
	})
	vngcp := network.NewVirtualNetworkGatewayListConnectionsResultPage(network.VirtualNetworkGatewayListConnectionsResult{Value: &[]network.VirtualNetworkGatewayConnectionListEntity{vngc}}, func(ctx context.Context, result network.VirtualNetworkGatewayListConnectionsResult) (network.VirtualNetworkGatewayListConnectionsResult, error) {
		return network.VirtualNetworkGatewayListConnectionsResult{}, nil
	})
	n.EXPECT().ListAll(gomock.Any()).Return(vnp, nil)
	ngc.EXPECT().List(gomock.Any(), gomock.Any()).Return(vngp, nil)
	ngc.EXPECT().ListConnections(gomock.Any(), gomock.Any(), gomock.Any()).Return(vngcp, nil)
	return s
}

func TestNetworkVirtualNetworks(t *testing.T) {
	client.AzureMockTestHelper(t, NetworkVirtualNetworks(), buildNetworkVirtualNetworksMock, client.TestOptions{})
}
