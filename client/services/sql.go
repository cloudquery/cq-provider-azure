package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"
	"github.com/Azure/go-autorest/autorest"
)

type SQLClient struct {
	Database SqlDatabaseClient
	Firewall SQLFirewallClient
	Servers  SqlServerClient
}

func NewSQLClient(subscriptionId string, auth autorest.Authorizer) SQLClient {
	servers := sql.NewServersClient(subscriptionId)
	servers.Authorizer = auth
	database := sql.NewDatabasesClient(subscriptionId)
	database.Authorizer = auth
	firewall := sql.NewFirewallRulesClient(subscriptionId)
	firewall.Authorizer = auth
	return SQLClient{
		Database: database,
		Firewall: firewall,
		Servers:  servers,
	}
}

type SqlServerClient interface {
	List(ctx context.Context) (result sql.ServerListResult, err error)
}

type SQLFirewallClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.FirewallRuleListResult, err error)
}

type SqlDatabaseClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string, expand string, filter string) (result sql.DatabaseListResult, err error)
}
