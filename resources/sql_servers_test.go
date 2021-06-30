package resources_test

import (
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"
	"github.com/cloudquery/cq-provider-azure/client/services"
	"github.com/cloudquery/cq-provider-azure/client/services/mocks"
	"github.com/cloudquery/cq-provider-azure/resources"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"
)

func buildSQLServerMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	serverSvc := mocks.NewMockSqlServerClient(ctrl)
	databaseSvc := mocks.NewMockSqlDatabaseClient(ctrl)
	firewallSvc := mocks.NewMockSQLFirewallClient(ctrl)
	s := services.Services{
		SQL: services.SQLClient{
			Database: databaseSvc,
			Firewall: firewallSvc,
			Servers:  serverSvc,
		},
	}
	server := sql.Server{}
	if err := faker.FakeData(&server); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	name := "testServer"
	server.Name = &name
	server.ID = &fakeResourceGroup
	page := sql.ServerListResult{Value: &[]sql.Server{server}}
	serverSvc.EXPECT().List(gomock.Any()).Return(page, nil)
	database := sql.Database{}
	if err := faker.FakeData(&database); err != nil {
		t.Errorf("failed building mock %s", err)
	}
	databaseSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name, "true", "").Return(sql.DatabaseListResult{Value: &[]sql.Database{database}}, nil)

	var rule sql.FirewallRule
	if err := faker.FakeData(&rule); err != nil {
		t.Fatal(err)
	}
	firewallSvc.EXPECT().ListByServer(gomock.Any(), "test", *server.Name).Return(
		sql.FirewallRuleListResult{Value: &[]sql.FirewallRule{rule}}, nil,
	)
	return s
}

func TestSQLServers(t *testing.T) {
	azureTestHelper(t, resources.SQLServers(), buildSQLServerMock)
}
