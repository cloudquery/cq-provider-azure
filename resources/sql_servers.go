package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/sql/mgmt/2014-04-01/sql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SQLServers() *schema.Table {
	return &schema.Table{
		Name:         "azure_sql_servers",
		Description:  "Azure sql server",
		Resolver:     fetchSqlServers,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "kind",
				Description: "Kind of sql server  This is metadata used for the Azure portal experience",
				Type:        schema.TypeString,
			},
			{
				Name:        "fully_qualified_domain_name",
				Description: "The fully qualified domain name of the server",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.FullyQualifiedDomainName"),
			},
			{
				Name:        "version",
				Description: "The version of the server Possible values include: 'TwoFullStopZero', 'OneTwoFullStopZero'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.Version"),
			},
			{
				Name:        "administrator_login",
				Description: "Administrator username for the server Can only be specified when the server is being created (and is required for creation)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.AdministratorLogin"),
			},
			{
				Name:        "administrator_login_password",
				Description: "The administrator login password (required for server creation)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.AdministratorLoginPassword"),
			},
			{
				Name:        "external_administrator_sid",
				Description: "The ID of the Active Azure Directory object with admin permissions on this server Legacy parameter, always null To check for Active Directory admin, query /servers/{serverName}/administrators",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("ServerProperties.ExternalAdministratorSid"),
			},
			{
				Name:        "external_administrator_login",
				Description: "The display name of the Azure Active Directory object with admin permissions on this server Legacy parameter, always null To check for Active Directory admin, query /servers/{serverName}/administrators",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.ExternalAdministratorLogin"),
			},
			{
				Name:        "state",
				Description: "The state of the server Possible values include: 'ServerStateReady', 'ServerStateDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ServerProperties.State"),
			},
			{
				Name:        "location",
				Description: "Resource location",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "resource_id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			SQLDatabases(),
			{
				Name:        "azure_sql_server_firewall_rules",
				Description: "The list of server firewall rules.",
				Resolver:    fetchSqlServerFirewallRules,
				Columns: []schema.Column{
					{
						Name:        "server_id",
						Description: "Unique ID of azure_sql_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "kind",
						Description: "Kind of server that contains this firewall rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "location",
						Description: "Location of the server that contains this firewall rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "start_ip_address",
						Description: "The start IP address of the firewall rule Must be IPv4 format Use value '0000' to represent all Azure-internal IP addresses",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FirewallRuleProperties.StartIPAddress"),
					},
					{
						Name:        "end_ip_address",
						Description: "The end IP address of the firewall rule Must be IPv4 format Must be greater than or equal to startIpAddress Use value '0000' to represent all Azure-internal IP addresses",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FirewallRuleProperties.EndIPAddress"),
					},
					{
						Name:        "resource_id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_sql_server_admins",
				Description: "ServerAzureADAdministrator an server Active Directory Administrator",
				Resolver:    fetchSqlServerAdmins,
				Columns: []schema.Column{
					{
						Name:        "server_id",
						Description: "Unique ID of azure_sql_servers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "administrator_type",
						Description: "The type of administrator",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ServerAdministratorProperties.AdministratorType"),
					},
					{
						Name:        "login",
						Description: "The server administrator login value",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ServerAdministratorProperties.Login"),
					},
					{
						Name:        "sid",
						Description: "The server administrator Sid (Secure ID)",
						Type:        schema.TypeUUID,
						Resolver:    schema.PathResolver("ServerAdministratorProperties.Sid"),
					},
					{
						Name:        "tenant_id",
						Description: "The server Active Directory Administrator tenant id",
						Type:        schema.TypeUUID,
						Resolver:    schema.PathResolver("ServerAdministratorProperties.TenantID"),
					},
					{
						Name:        "resource_id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchSqlServers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.Servers
	servers, err := svc.List(ctx)
	if err != nil {
		return err
	}
	res <- *servers.Value
	return nil
}

func fetchSqlServerFirewallRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.Firewall
	s := parent.Item.(sql.Server)
	details, err := client.ParseResourceID(*s.ID)
	if err != nil {
		return err
	}
	result, err := svc.ListByServer(ctx, details.ResourceGroup, *s.Name)
	if err != nil {
		return err
	}
	if result.Value != nil {
		res <- *result.Value
	}
	return nil
}

func fetchSqlServerAdmins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ServerAdmins
	s := parent.Item.(sql.Server)
	details, err := client.ParseResourceID(*s.ID)
	if err != nil {
		return err
	}
	result, err := svc.ListByServer(ctx, details.ResourceGroup, *s.Name)
	if err != nil {
		return err
	}
	if result.Value != nil {
		res <- *result.Value
	}
	return nil
}
