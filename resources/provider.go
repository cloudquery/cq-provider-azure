package resources

import (
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "azure",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"ad.users":                 AdUsers(),
			"ad.groups":                AdGroups(),
			"compute.disks":            ComputeDisks(),
			"keyvault.vaults":          KeyVaultVaults(),
			"mysql.servers":            MySQLServers(),
			"network.virtual_networks": NetworkVirtualNetworks(),
			"postgresql.servers":       PostgresqlServers(),
			"resources.groups":         ResourcesGroups(),
			"sql.servers":              SQLServers(),
			"storage.accounts":         StorageAccounts(),
		},
		Config: func() provider.Config {
			return &client.Config{}
		},
	}

}
