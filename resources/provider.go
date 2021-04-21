package resources

import (
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Provider() *provider.Provider {
	return &provider.Provider{
		Name:      "aws",
		Configure: client.Configure,
		ResourceMap: map[string]*schema.Table{
			"compute.disks":            ComputeDisks(),
			"resources.groups":         ResourceGroups(),
			"keyvault.vaults":          KeyVaultVaults(),
			"storage.accounts":         StorageAccounts(),
			"mysql.servers":            MySQLServers(),
			"postgresql.servers":       PostgresqlServers(),
			"sql.servers":              SQLServers(),
			"network.virtual_networks": NetworkVirtualNetworks(),
		},
		Config: func() interface{} {
			return &client.Config{}
		},
		DefaultConfigGenerator: func() (string, error) {
			return client.DefaultConfig, nil
		},
	}

}
