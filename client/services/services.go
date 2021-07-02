//go:generate mockgen -destination=./mocks/services.go -package=mocks  . DisksClient,GroupsClient,KeyClient,VaultClient,StorageAccountClient,StorageContainerClient,SqlServerClient,SqlDatabaseClient,MySQLServerClient,MySQLConfigurationClient,VirtualNetworksClient
//go:generate mockgen -destination=./mocks/ad_groups.go -package=mocks . ADGroupsClient
//go:generate mockgen -destination=./mocks/ad_service_principals.go -package=mocks . ADServicePrinicpals
//go:generate mockgen -destination=./mocks/ad_users.go -package=mocks . ADUsersClient
//go:generate mockgen -destination=./mocks/security_auto_provisioning_settings.go -package=mocks . SecurityAutoProvisioningSettingsClient
//go:generate mockgen -destination=./mocks/security.go -package=mocks . SecurityContactsClient,SecurityPricingsClient,SecuritySettingsClient
//go:generate mockgen -destination=./mocks/postgresql.go -package=mocks . PostgresqlConfigurationClient,PostgresqlServerClient,PostgresqlFirewallRuleClient
package services

import "github.com/Azure/go-autorest/autorest"

type Services struct {
	AD         AD
	Compute    ComputeClient
	KeyVault   KeyVaultClient
	MySQL      MySQL
	Network    NetworksClient
	PostgreSQL PostgreSQL
	Resources  ResourcesClient
	Security   SecurityClient
	SQL        SQLClient
	Storage    StorageClient
}

func InitServices(subscriptionId string, auth autorest.Authorizer) Services {
	return Services{
		AD:         NewADClient(subscriptionId, auth),
		Compute:    NewComputeClient(subscriptionId, auth),
		KeyVault:   NewKeyVaultClient(subscriptionId, auth),
		MySQL:      NewMySQLClient(subscriptionId, auth),
		Network:    NewNetworksClient(subscriptionId, auth),
		PostgreSQL: NewPostgresClient(subscriptionId, auth),
		Resources:  NewResourcesClient(subscriptionId, auth),
		Security:   NewSecurityClient(subscriptionId, auth),
		SQL:        NewSQLClient(subscriptionId, auth),
		Storage:    NewStorageClient(subscriptionId, auth),
	}
}
