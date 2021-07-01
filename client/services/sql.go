package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/Azure/go-autorest/autorest"
)

type SQLClient struct {
	Database                     SqlDatabaseClient
	DatabaseBlobAuditingPolicies SQLDatabaseBlobAuditingPoliciesClient
	Firewall                     SQLFirewallClient
	ServerAdmins                 SQLServerAdminClient
	Servers                      SqlServerClient
}

func NewSQLClient(subscriptionId string, auth autorest.Authorizer) SQLClient {
	servers := sql.NewServersClient(subscriptionId)
	servers.Authorizer = auth
	database := sql.NewDatabasesClient(subscriptionId)
	database.Authorizer = auth
	dbap := sql.NewDatabaseBlobAuditingPoliciesClient(subscriptionId)
	dbap.Authorizer = auth
	firewall := sql.NewFirewallRulesClient(subscriptionId)
	firewall.Authorizer = auth
	return SQLClient{
		Database:                     database,
		DatabaseBlobAuditingPolicies: dbap,
		Firewall:                     firewall,
		Servers:                      servers,
	}
}

type SqlServerClient interface {
	List(ctx context.Context) (result sql.ServerListResultPage, err error)
}

type SQLFirewallClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.FirewallRuleListResult, err error)
}

type SQLServerAdminClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.AdministratorListResultPage, err error)
}

type SqlDatabaseClient interface {
	ListByServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.DatabaseListResultPage, err error)
}

type SQLDatabaseBlobAuditingPoliciesClient interface {
	ListByDatabase(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result sql.DatabaseBlobAuditingPolicyListResultPage, err error)
}
