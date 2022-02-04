package sql

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SqlManagedInstances() *schema.Table {
	return &schema.Table{
		Name:         "azure_sql_managed_instances",
		Description:  "ManagedInstance an Azure SQL managed instance.",
		Resolver:     fetchSqlManagedInstances,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "identity_principal_id",
				Description: "PrincipalID - READ-ONLY; The Azure Active Directory principal id.",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_type",
				Description: "Type - The identity type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "TenantID - READ-ONLY; The Azure Active Directory tenant id.",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "sku_name",
				Description: "Name - The name of the SKU, typically, a letter + Number code, e.g",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "Tier - The tier or edition of the particular SKU, e.g",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "sku_size",
				Description: "Size - Size of the particular SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Size"),
			},
			{
				Name:        "sku_family",
				Description: "Family - If the service has different generations of hardware, for the same SKU, then that can be captured here.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Family"),
			},
			{
				Name:        "sku_capacity",
				Description: "Capacity - Capacity of the particular SKU.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Sku.Capacity"),
			},
			{
				Name:        "provisioning_state",
				Description: "ProvisioningState - READ-ONLY; Possible values include: 'ProvisioningState1Creating', 'ProvisioningState1Deleting', 'ProvisioningState1Updating', 'ProvisioningState1Unknown', 'ProvisioningState1Succeeded', 'ProvisioningState1Failed'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.ProvisioningState"),
			},
			{
				Name:        "managed_instance_create_mode",
				Description: "ManagedInstanceCreateMode - Specifies the mode of database creation.  Default: Regular instance creation.  Restore: Creates an instance by restoring a set of backups to specific point in time",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.ManagedInstanceCreateMode"),
			},
			{
				Name:        "fully_qualified_domain_name",
				Description: "FullyQualifiedDomainName - READ-ONLY; The fully qualified domain name of the managed instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.FullyQualifiedDomainName"),
			},
			{
				Name:        "administrator_login",
				Description: "AdministratorLogin - Administrator username for the managed instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.AdministratorLogin"),
			},
			{
				Name:        "administrator_login_password",
				Description: "AdministratorLoginPassword - The administrator login password (required for managed instance creation).",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.AdministratorLoginPassword"),
			},
			{
				Name:        "subnet_id",
				Description: "SubnetID - Subnet resource ID for the managed instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.SubnetID"),
			},
			{
				Name:        "state",
				Description: "State - READ-ONLY; The state of the managed instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.State"),
			},
			{
				Name:        "license_type",
				Description: "LicenseType - The license type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.LicenseType"),
			},
			{
				Name:        "v_cores",
				Description: "VCores - The number of vCores",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.VCores"),
			},
			{
				Name:        "storage_size_in_gb",
				Description: "StorageSizeInGB - Storage size in GB",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.StorageSizeInGB"),
			},
			{
				Name:        "collation",
				Description: "Collation - Collation of the managed instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.Collation"),
			},
			{
				Name:        "dns_zone",
				Description: "DNSZone - READ-ONLY; The Dns Zone that the managed instance is in.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.DNSZone"),
			},
			{
				Name:        "dns_zone_partner",
				Description: "DNSZonePartner - The resource id of another managed instance whose DNS zone this managed instance will share after creation.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.DNSZonePartner"),
			},
			{
				Name:        "public_data_endpoint_enabled",
				Description: "PublicDataEndpointEnabled - Whether or not the public data endpoint is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.PublicDataEndpointEnabled"),
			},
			{
				Name:        "source_managed_instance_id",
				Description: "SourceManagedInstanceID - The resource identifier of the source managed instance associated with create operation of this instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.SourceManagedInstanceID"),
			},
			{
				Name:     "restore_point_in_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ManagedInstanceProperties.RestorePointInTime.Time"),
			},
			{
				Name:        "proxy_override",
				Description: "ProxyOverride - Connection type used for connecting to the instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.ProxyOverride"),
			},
			{
				Name:        "timezone_id",
				Description: "TimezoneID - Id of the timezone",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.TimezoneID"),
			},
			{
				Name:        "instance_pool_id",
				Description: "InstancePoolID - The Id of the instance pool this managed server belongs to.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.InstancePoolID"),
			},
			{
				Name:        "maintenance_configuration_id",
				Description: "MaintenanceConfigurationID - Specifies maintenance configuration id to apply to this managed instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.MaintenanceConfigurationID"),
			},
			{
				Name:        "minimal_tls_version",
				Description: "MinimalTLSVersion - Minimal TLS version",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.MinimalTLSVersion"),
			},
			{
				Name:        "storage_account_type",
				Description: "StorageAccountType - The storage account type used to store backups for this instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.StorageAccountType"),
			},
			{
				Name:        "zone_redundant",
				Description: "ZoneRedundant - Whether or not the multi-az is enabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ManagedInstanceProperties.ZoneRedundant"),
			},
			{
				Name:        "location",
				Description: "Location - Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Tags - Resource tags.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "id",
				Description: "ID - READ-ONLY; Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Name - READ-ONLY; Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Type - READ-ONLY; Resource type.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_sql_managed_instance_private_endpoint_connections",
				Description: "ManagedInstancePecProperty a private endpoint connection under a managed instance",
				Resolver:    fetchSqlManagedInstancePrivateEndpointConnections,
				Columns: []schema.Column{
					{
						Name:        "managed_instance_cq_id",
						Description: "Unique CloudQuery ID of azure_sql_managed_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "id",
						Description: "ID - READ-ONLY; Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "private_endpoint_id",
						Description: "ID - Resource id of the private endpoint.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateEndpoint.ID"),
					},
					{
						Name:        "private_link_service_connection_state_status",
						Description: "Status - The private link service connection status.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Status"),
					},
					{
						Name:        "private_link_service_connection_state_description",
						Description: "Description - The private link service connection description.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.Description"),
					},
					{
						Name:        "private_link_service_connection_state_actions_required",
						Description: "ActionsRequired - READ-ONLY; The private link service connection description.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.PrivateLinkServiceConnectionState.ActionsRequired"),
					},
					{
						Name:        "provisioning_state",
						Description: "ProvisioningState - READ-ONLY; State of the Private Endpoint Connection.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Properties.ProvisioningState"),
					},
				},
			},
			{
				Name:        "azure_sql_managed_instance_vulnerability_assessments",
				Description: "ManagedInstanceVulnerabilityAssessment a managed instance vulnerability assessment.",
				Resolver:    fetchSqlManagedInstanceVulnerabilityAssessments,
				Columns: []schema.Column{
					{
						Name:        "managed_instance_cq_id",
						Description: "Unique CloudQuery ID of azure_sql_managed_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "storage_container_path",
						Description: "StorageContainerPath - A blob storage container path to hold the scan results (e.g",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ManagedInstanceVulnerabilityAssessmentProperties.StorageContainerPath"),
					},
					{
						Name:        "storage_container_sas_key",
						Description: "StorageContainerSasKey - A shared access signature (SAS Key) that has read and write access to the blob container specified in 'storageContainerPath' parameter",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ManagedInstanceVulnerabilityAssessmentProperties.StorageContainerSasKey"),
					},
					{
						Name:        "storage_account_access_key",
						Description: "StorageAccountAccessKey - Specifies the identifier key of the storage account for vulnerability assessment scan results",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ManagedInstanceVulnerabilityAssessmentProperties.StorageAccountAccessKey"),
					},
					{
						Name:        "recurring_scans_is_enabled",
						Description: "IsEnabled - Recurring scans state.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ManagedInstanceVulnerabilityAssessmentProperties.RecurringScans.IsEnabled"),
					},
					{
						Name:        "recurring_scans_email_subscription_admins",
						Description: "EmailSubscriptionAdmins - Specifies that the schedule scan notification will be is sent to the subscription administrators.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ManagedInstanceVulnerabilityAssessmentProperties.RecurringScans.EmailSubscriptionAdmins"),
					},
					{
						Name:        "recurring_scans_emails",
						Description: "Emails - Specifies an array of e-mail addresses to which the scan notification is sent.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ManagedInstanceVulnerabilityAssessmentProperties.RecurringScans.Emails"),
					},
					{
						Name:        "id",
						Description: "ID - READ-ONLY; Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Name - READ-ONLY; Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Type - READ-ONLY; Resource type.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_sql_managed_instance_encryption_protectors",
				Description: "ManagedInstanceEncryptionProtector the managed instance encryption protector.",
				Resolver:    fetchSqlManagedInstanceEncryptionProtectors,
				Columns: []schema.Column{
					{
						Name:        "managed_instance_cq_id",
						Description: "Unique CloudQuery ID of azure_sql_managed_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "kind",
						Description: "Kind - READ-ONLY; Kind of encryption protector",
						Type:        schema.TypeString,
					},
					{
						Name:        "server_key_name",
						Description: "ServerKeyName - The name of the managed instance key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ManagedInstanceEncryptionProtectorProperties.ServerKeyName"),
					},
					{
						Name:        "server_key_type",
						Description: "ServerKeyType - The encryption protector type like 'ServiceManaged', 'AzureKeyVault'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ManagedInstanceEncryptionProtectorProperties.ServerKeyType"),
					},
					{
						Name:        "uri",
						Description: "URI - READ-ONLY; The URI of the server key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ManagedInstanceEncryptionProtectorProperties.URI"),
					},
					{
						Name:        "thumbprint",
						Description: "Thumbprint - READ-ONLY; Thumbprint of the server key.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ManagedInstanceEncryptionProtectorProperties.Thumbprint"),
					},
					{
						Name:        "id",
						Description: "ID - READ-ONLY; Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
					{
						Name:        "name",
						Description: "Name - READ-ONLY; Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Type - READ-ONLY; Resource type.",
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

func fetchSqlManagedInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ManagedInstances
	servers, err := svc.List(ctx)
	if err != nil {
		return err
	}
	for servers.NotDone() {
		res <- servers.Values()
		if err := servers.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
func fetchSqlManagedInstancePrivateEndpointConnections(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	server, ok := parent.Item.(sql.ManagedInstance)
	if !ok {
		return fmt.Errorf("not an sql.ManagedInstance instance: %T", parent.Item)
	}
	if server.PrivateEndpointConnections != nil {
		res <- *server.PrivateEndpointConnections
	}
	return nil
}
func fetchSqlManagedInstanceVulnerabilityAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ManagedInstanceVulnerabilityAssessments
	s, ok := parent.Item.(sql.ManagedInstance)
	if !ok {
		return fmt.Errorf("not an sql.ManagedInstance instance: %T", parent.Item)
	}
	details, err := client.ParseResourceID(*s.ID)
	if err != nil {
		return err
	}
	result, err := svc.ListByInstance(ctx, details.ResourceGroup, *s.Name)
	if err != nil {
		return err
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
func fetchSqlManagedInstanceEncryptionProtectors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	svc := meta.(*client.Client).Services().SQL.ManagedInstanceEncryptionProtectors
	s, ok := parent.Item.(sql.ManagedInstance)
	if !ok {
		return fmt.Errorf("not an sql.ManagedInstance instance: %T", parent.Item)
	}
	details, err := client.ParseResourceID(*s.ID)
	if err != nil {
		return err
	}
	result, err := svc.ListByInstance(ctx, details.ResourceGroup, *s.Name)
	if err != nil {
		return err
	}
	for result.NotDone() {
		res <- result.Values()
		if err := result.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
