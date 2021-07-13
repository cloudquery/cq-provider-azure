package resources

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v4.0/sql"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func SQLDatabases() *schema.Table {
	return &schema.Table{
		Name:         "azure_sql_databases",
		Description:  "Azure sql database",
		Resolver:     fetchSqlDatabases,
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
				Name:        "sku_name",
				Description: "The name of the SKU, typically, a letter + Number code, eg P3",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Name"),
			},
			{
				Name:        "sku_tier",
				Description: "The tier or edition of the particular SKU, eg Basic, Premium",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Tier"),
			},
			{
				Name:        "sku_size",
				Description: "Size of the particular SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Size"),
			},
			{
				Name:        "sku_family",
				Description: "If the service has different generations of hardware, for the same SKU, then that can be captured here",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Sku.Family"),
			},
			{
				Name:        "sku_capacity",
				Description: "Capacity of the particular SKU",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Sku.Capacity"),
			},
			{
				Name:        "kind",
				Description: "Kind of database This is metadata used for the Azure portal experience",
				Type:        schema.TypeString,
			},
			{
				Name:        "managed_by",
				Description: "Resource that manages the database",
				Type:        schema.TypeString,
			},
			{
				Name:        "create_mode",
				Description: "Specifies the mode of database creation  Default: regular database creation  Copy: creates a database as a copy of an existing database sourceDatabaseId must be specified as the resource ID of the source database  Secondary: creates a database as a secondary replica of an existing database sourceDatabaseId must be specified as the resource ID of the existing primary database  PointInTimeRestore: Creates a database by restoring a point in time backup of an existing database sourceDatabaseId must be specified as the resource ID of the existing database, and restorePointInTime must be specified  Recovery: Creates a database by restoring a geo-replicated backup sourceDatabaseId must be specified as the recoverable database resource ID to restore  Restore: Creates a database by restoring a backup of a deleted database sourceDatabaseId must be specified If sourceDatabaseId is the database's original resource ID, then sourceDatabaseDeletionDate must be specified Otherwise sourceDatabaseId must be the restorable dropped database resource ID and sourceDatabaseDeletionDate is ignored restorePointInTime may also be specified to restore from an earlier point in time  RestoreLongTermRetentionBackup: Creates a database by restoring from a long term retention vault recoveryServicesRecoveryPointResourceId must be specified as the recovery point resource ID  Copy, Secondary, and RestoreLongTermRetentionBackup are not supported for DataWarehouse edition Possible values include: 'CreateModeDefault', 'CreateModeCopy', 'CreateModeSecondary', 'CreateModePointInTimeRestore', 'CreateModeRestore', 'CreateModeRecovery', 'CreateModeRestoreExternalBackup', 'CreateModeRestoreExternalBackupSecondary', 'CreateModeRestoreLongTermRetentionBackup', 'CreateModeOnlineSecondary'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CreateMode"),
			},
			{
				Name:        "collation",
				Description: "The collation of the database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.Collation"),
			},
			{
				Name:        "max_size_bytes",
				Description: "The max size of the database expressed in bytes",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DatabaseProperties.MaxSizeBytes"),
			},
			{
				Name:        "sample_name",
				Description: "The name of the sample schema to apply when creating this database Possible values include: 'AdventureWorksLT', 'WideWorldImportersStd', 'WideWorldImportersFull'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.SampleName"),
			},
			{
				Name:        "elastic_pool_id",
				Description: "The resource identifier of the elastic pool containing this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.ElasticPoolID"),
			},
			{
				Name:        "source_database_id",
				Description: "The resource identifier of the source database associated with create operation of this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.SourceDatabaseID"),
			},
			{
				Name:        "status",
				Description: "The status of the database Possible values include: 'DatabaseStatusOnline', 'DatabaseStatusRestoring', 'DatabaseStatusRecoveryPending', 'DatabaseStatusRecovering', 'DatabaseStatusSuspect', 'DatabaseStatusOffline', 'DatabaseStatusStandby', 'DatabaseStatusShutdown', 'DatabaseStatusEmergencyMode', 'DatabaseStatusAutoClosed', 'DatabaseStatusCopying', 'DatabaseStatusCreating', 'DatabaseStatusInaccessible', 'DatabaseStatusOfflineSecondary', 'DatabaseStatusPausing', 'DatabaseStatusPaused', 'DatabaseStatusResuming', 'DatabaseStatusScaling', 'DatabaseStatusOfflineChangingDwPerformanceTiers', 'DatabaseStatusOnlineChangingDwPerformanceTiers', 'DatabaseStatusDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.Status"),
			},
			{
				Name:        "database_id",
				Description: "The ID of the database",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("DatabaseProperties.DatabaseID"),
			},
			{
				Name:        "creation_date_time",
				Description: "The creation date of the database (ISO8601 format)",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DatabaseProperties.CreationDate.Time"),
			},
			{
				Name:        "current_service_objective_name",
				Description: "The current service level objective name of the database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentServiceObjectiveName"),
			},
			{
				Name:        "requested_service_objective_name",
				Description: "The requested service level objective name of the database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.RequestedServiceObjectiveName"),
			},
			{
				Name:        "default_secondary_location",
				Description: "The default secondary region for this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.DefaultSecondaryLocation"),
			},
			{
				Name:        "failover_group_id",
				Description: "Failover Group resource identifier that this database belongs to",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.FailoverGroupID"),
			},
			{
				Name:        "restore_point_in_time",
				Description: "Specifies the point in time (ISO8601 format) of the source database that will be restored to create the new database",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DatabaseProperties.RestorePointInTime.Time"),
			},
			{
				Name:        "source_database_deletion_date_time",
				Description: "Specifies the time that the database was deleted",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DatabaseProperties.SourceDatabaseDeletionDate.Time"),
			},
			{
				Name:        "recovery_services_recovery_point_id",
				Description: "The resource identifier of the recovery point associated with create operation of this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.RecoveryServicesRecoveryPointID"),
			},
			{
				Name:        "long_term_retention_backup_resource_id",
				Description: "The resource identifier of the long term retention backup associated with create operation of this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.LongTermRetentionBackupResourceID"),
			},
			{
				Name:        "recoverable_database_id",
				Description: "The resource identifier of the recoverable database associated with create operation of this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.RecoverableDatabaseID"),
			},
			{
				Name:        "restorable_dropped_database_id",
				Description: "The resource identifier of the restorable dropped database associated with create operation of this database",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.RestorableDroppedDatabaseID"),
			},
			{
				Name:        "catalog_collation",
				Description: "Collation of the metadata catalog Possible values include: 'DATABASEDEFAULT', 'SQLLatin1GeneralCP1CIAS'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CatalogCollation"),
			},
			{
				Name:        "zone_redundant",
				Description: "Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DatabaseProperties.ZoneRedundant"),
			},
			{
				Name:        "license_type",
				Description: "The license type to apply for this database `LicenseIncluded` if you need a license, or `BasePrice` if you have a license and are eligible for the Azure Hybrid Benefit Possible values include: 'LicenseIncluded', 'BasePrice'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.LicenseType"),
			},
			{
				Name:        "max_log_size_bytes",
				Description: "The max log size for this database",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DatabaseProperties.MaxLogSizeBytes"),
			},
			{
				Name:        "earliest_restore_date_time",
				Description: "This records the earliest start date and time that restore is available for this database (ISO8601 format)",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("DatabaseProperties.EarliestRestoreDate.Time"),
			},
			{
				Name:        "read_scale",
				Description: "The state of read-only routing If enabled, connections that have application intent set to readonly in their connection string may be routed to a readonly secondary replica in the same region Possible values include: 'DatabaseReadScaleEnabled', 'DatabaseReadScaleDisabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.ReadScale"),
			},
			{
				Name:        "high_availability_replica_count",
				Description: "The number of secondary replicas associated with the database that are used to provide high availability",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DatabaseProperties.HighAvailabilityReplicaCount"),
			},
			{
				Name:        "secondary_type",
				Description: "The secondary type of the database if it is a secondary  Valid values are Geo and Named Possible values include: 'Geo', 'Named'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.SecondaryType"),
			},
			{
				Name:        "current_sku_name",
				Description: "The name of the SKU, typically, a letter + Number code, eg P3",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentSku.Name"),
			},
			{
				Name:        "current_sku_tier",
				Description: "The tier or edition of the particular SKU, eg Basic, Premium",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentSku.Tier"),
			},
			{
				Name:        "current_sku_size",
				Description: "Size of the particular SKU",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentSku.Size"),
			},
			{
				Name:        "current_sku_family",
				Description: "If the service has different generations of hardware, for the same SKU, then that can be captured here",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentSku.Family"),
			},
			{
				Name:        "current_sku_capacity",
				Description: "Capacity of the particular SKU",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DatabaseProperties.CurrentSku.Capacity"),
			},
			{
				Name:        "auto_pause_delay",
				Description: "Time in minutes after which database is automatically paused A value of -1 means that automatic pause is disabled",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("DatabaseProperties.AutoPauseDelay"),
			},
			{
				Name:        "storage_account_type",
				Description: "The storage account type used to store backups for this database Possible values include: 'GRS', 'LRS', 'ZRS'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.StorageAccountType"),
			},
			{
				Name:        "min_capacity",
				Description: "Minimal capacity that database will always have allocated, if not paused",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("DatabaseProperties.MinCapacity"),
			},
			{
				Name:     "paused_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.PausedDate.Time"),
			},
			{
				Name:     "resumed_date_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DatabaseProperties.ResumedDate.Time"),
			},
			{
				Name:        "maintenance_configuration_id",
				Description: "Maintenance configuration id assigned to the database This configuration defines the period when the maintenance updates will occur",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DatabaseProperties.MaintenanceConfigurationID"),
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
				Name:        "id",
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
			{
				Name:        "azure_sql_database_db_blob_auditing_policies",
				Description: "Database blob auditing policy",
				Resolver:    fetchSqlDatabaseDbBlobAuditingPolicies,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"database_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique ID of azure_sql_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "kind",
						Description: "Resource kind",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "Specifies the state of the policy If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required Possible values include: 'BlobAuditingPolicyStateEnabled', 'BlobAuditingPolicyStateDisabled'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.State"),
					},
					{
						Name:        "storage_endpoint",
						Description: "Specifies the blob storage endpoint (eg https://MyAccountblobcorewindowsnet) If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled is required",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.StorageEndpoint"),
					},
					{
						Name:        "storage_account_access_key",
						Description: "Specifies the identifier key of the auditing storage account If state is Enabled and storageEndpoint is specified, not specifying the storageAccountAccessKey will use SQL server system-assigned managed identity to access the storage Prerequisites for using managed identity authentication: 1 Assign SQL Server a system-assigned managed identity in Azure Active Directory (AAD) 2 Grant SQL Server identity access to the storage account by adding 'Storage Blob Data Contributor' RBAC role to the server identity For more information, see [Auditing to storage using Managed Identity authentication](https://gomicrosoftcom/fwlink/?linkid=2114355)",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.StorageAccountAccessKey"),
					},
					{
						Name:        "retention_days",
						Description: "Specifies the number of days to keep in the audit logs in the storage account",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.RetentionDays"),
					},
					{
						Name:        "audit_actions_and_groups",
						Description: "Specifies the Actions-Groups and Actions to audit  The recommended set of action groups to use is the following combination - this will audit all the queries and stored procedures executed against the database, as well as successful and failed logins:  BATCH_COMPLETED_GROUP, SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP, FAILED_DATABASE_AUTHENTICATION_GROUP  This above combination is also the set that is configured by default when enabling auditing from the Azure portal  The supported action groups to audit are (note: choose only specific groups that cover your auditing needs Using unnecessary groups could lead to very large quantities of audit records):  APPLICATION_ROLE_CHANGE_PASSWORD_GROUP BACKUP_RESTORE_GROUP DATABASE_LOGOUT_GROUP DATABASE_OBJECT_CHANGE_GROUP DATABASE_OBJECT_OWNERSHIP_CHANGE_GROUP DATABASE_OBJECT_PERMISSION_CHANGE_GROUP DATABASE_OPERATION_GROUP DATABASE_PERMISSION_CHANGE_GROUP DATABASE_PRINCIPAL_CHANGE_GROUP DATABASE_PRINCIPAL_IMPERSONATION_GROUP DATABASE_ROLE_MEMBER_CHANGE_GROUP FAILED_DATABASE_AUTHENTICATION_GROUP SCHEMA_OBJECT_ACCESS_GROUP SCHEMA_OBJECT_CHANGE_GROUP SCHEMA_OBJECT_OWNERSHIP_CHANGE_GROUP SCHEMA_OBJECT_PERMISSION_CHANGE_GROUP SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP USER_CHANGE_PASSWORD_GROUP BATCH_STARTED_GROUP BATCH_COMPLETED_GROUP  These are groups that cover all sql statements and stored procedures executed against the database, and should not be used in combination with other groups as this will result in duplicate audit logs  For more information, see [Database-Level Audit Action Groups](https://docsmicrosoftcom/en-us/sql/relational-databases/security/auditing/sql-server-audit-action-groups-and-actions#database-level-audit-action-groups)  For Database auditing policy, specific Actions can also be specified (note that Actions cannot be specified for Server auditing policy) The supported actions to audit are: SELECT UPDATE INSERT DELETE EXECUTE RECEIVE REFERENCES  The general form for defining an action to be audited is: {action} ON {object} BY {principal}  Note that <object> in the above format can refer to an object like a table, view, or stored procedure, or an entire database or schema For the latter cases, the forms DATABASE::{db_name} and SCHEMA::{schema_name} are used, respectively  For example: SELECT on dbomyTable by public SELECT on DATABASE::myDatabase by public SELECT on SCHEMA::mySchema by public  For more information, see [Database-Level Audit Actions](https://docsmicrosoftcom/en-us/sql/relational-databases/security/auditing/sql-server-audit-action-groups-and-actions#database-level-audit-actions)",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.AuditActionsAndGroups"),
					},
					{
						Name:        "storage_account_subscription_id",
						Description: "Specifies the blob storage subscription Id",
						Type:        schema.TypeUUID,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.StorageAccountSubscriptionID"),
					},
					{
						Name:        "is_storage_secondary_key_in_use",
						Description: "Specifies whether storageAccountAccessKey value is the storage's secondary key",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.IsStorageSecondaryKeyInUse"),
					},
					{
						Name:        "is_azure_monitor_target_enabled",
						Description: "Specifies whether audit events are sent to Azure Monitor In order to send the events to Azure Monitor, specify 'state' as 'Enabled' and 'isAzureMonitorTargetEnabled' as true  When using REST API to configure auditing, Diagnostic Settings with 'SQLSecurityAuditEvents' diagnostic logs category on the database should be also created Note that for server level audit you should use the 'master' database as {databaseName}  Diagnostic Settings URI format: PUT https://managementazurecom/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/MicrosoftSql/servers/{serverName}/databases/{databaseName}/providers/microsoftinsights/diagnosticSettings/{settingsName}?api-version=2017-05-01-preview  For more information, see [Diagnostic Settings REST API](https://gomicrosoftcom/fwlink/?linkid=2033207) or [Diagnostic Settings PowerShell](https://gomicrosoftcom/fwlink/?linkid=2033043)",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.IsAzureMonitorTargetEnabled"),
					},
					{
						Name:        "queue_delay_ms",
						Description: "Specifies the amount of time in milliseconds that can elapse before audit actions are forced to be processed The default minimum value is 1000 (1 second) The maximum is 2,147,483,647",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("DatabaseBlobAuditingPolicyProperties.QueueDelayMs"),
					},
					{
						Name:        "id",
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
				Name:        "azure_sql_database_db_threat_detection_policies",
				Description: "Contains information about a database Threat Detection policy.",
				Resolver:    fetchSqlDatabaseDbThreatDetectionPolicies,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"database_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique ID of azure_sql_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "location",
						Description: "The geo-location where the resource lives",
						Type:        schema.TypeString,
					},
					{
						Name:        "kind",
						Description: "Resource kind",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "Specifies the state of the policy If state is Enabled, storageEndpoint and storageAccountAccessKey are required Possible values include: 'SecurityAlertPolicyStateNew', 'SecurityAlertPolicyStateEnabled', 'SecurityAlertPolicyStateDisabled'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.State"),
					},
					{
						Name:        "disabled_alerts",
						Description: "Sql_Injection_Vulnerability; Access_Anomaly; Data_Exfiltration; Unsafe_Action",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.DisabledAlerts"),
					},
					{
						Name:        "email_addresses",
						Description: "Specifies the semicolon-separated list of e-mail addresses to which the alert is sent",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.EmailAddresses"),
					},
					{
						Name:        "email_account_admins",
						Description: "Specifies that the alert is sent to the account administrators Possible values include: 'SecurityAlertPolicyEmailAccountAdminsEnabled', 'SecurityAlertPolicyEmailAccountAdminsDisabled'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.EmailAccountAdmins"),
					},
					{
						Name:        "storage_endpoint",
						Description: "Specifies the blob storage endpoint (eg https://MyAccountblobcorewindowsnet) This blob storage will hold all Threat Detection audit logs If state is Enabled, storageEndpoint is required",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.StorageEndpoint"),
					},
					{
						Name:        "storage_account_access_key",
						Description: "Specifies the identifier key of the Threat Detection audit storage account If state is Enabled, storageAccountAccessKey is required",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.StorageAccountAccessKey"),
					},
					{
						Name:        "retention_days",
						Description: "Specifies the number of days to keep in the Threat Detection audit logs",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.RetentionDays"),
					},
					{
						Name:        "use_server_default",
						Description: "Specifies whether to use the default server policy Possible values include: 'SecurityAlertPolicyUseServerDefaultEnabled', 'SecurityAlertPolicyUseServerDefaultDisabled'",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseSecurityAlertPolicyProperties.UseServerDefault"),
					},
					{
						Name:        "id",
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
				Name:        "azure_sql_database_db_vulnerability_assessments",
				Description: "DatabaseVulnerabilityAssessment a database vulnerability assessment",
				Resolver:    fetchSqlDatabaseDbVulnerabilityAssessments,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"database_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:        "database_cq_id",
						Description: "Unique ID of azure_sql_databases table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "storage_container_path",
						Description: "A blob storage container path to hold the scan results (eg https://myStorageblobcorewindowsnet/VaScans/)  It is required if server level vulnerability assessment policy doesn't set",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.StorageContainerPath"),
					},
					{
						Name:        "storage_container_sas_key",
						Description: "A shared access signature (SAS Key) that has read and write access to the blob container specified in 'storageContainerPath' parameter If 'storageAccountAccessKey' isn't specified, StorageContainerSasKey is required",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.StorageContainerSasKey"),
					},
					{
						Name:        "storage_account_access_key",
						Description: "Specifies the identifier key of the storage account for vulnerability assessment scan results If 'StorageContainerSasKey' isn't specified, storageAccountAccessKey is required",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.StorageAccountAccessKey"),
					},
					{
						Name:        "recurring_scans_is_enabled",
						Description: "Recurring scans state",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.RecurringScans.IsEnabled"),
					},
					{
						Name:        "recurring_scans_email_subscription_admins",
						Description: "Specifies that the schedule scan notification will be is sent to the subscription administrators",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.RecurringScans.EmailSubscriptionAdmins"),
					},
					{
						Name:        "recurring_scans_emails",
						Description: "Specifies an array of e-mail addresses to which the scan notification is sent",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("DatabaseVulnerabilityAssessmentProperties.RecurringScans.Emails"),
					},
					{
						Name:        "id",
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
func fetchSqlDatabases(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.Databases
	server := parent.Item.(sql.Server)
	resourceDetails, err := client.ParseResourceID(*server.ID)
	if err != nil {
		return err
	}
	databases, err := svc.ListByServer(ctx, resourceDetails.ResourceGroup, *server.Name)
	if err != nil {
		return err
	}
	for databases.NotDone() {
		res <- databases.Values()
		if err := databases.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}

func fetchSqlDatabaseDbBlobAuditingPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.DatabaseBlobAuditingPolicies
	database := parent.Item.(sql.Database)
	details, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	serverName := parent.Parent.Get("name").(*string)
	result, err := svc.ListByDatabase(ctx, details.ResourceGroup, *serverName, *database.Name)
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

func fetchSqlDatabaseDbThreatDetectionPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.DatabaseThreatDetectionPolicies
	database := parent.Item.(sql.Database)
	details, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	serverName := parent.Parent.Get("name").(*string)
	result, err := svc.Get(ctx, details.ResourceGroup, *serverName, *database.Name)
	if err != nil {
		return err
	}
	res <- result
	return nil
}

func fetchSqlDatabaseDbVulnerabilityAssessments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().SQL.DatabaseVulnerabilityAssessments
	database := parent.Item.(sql.Database)
	details, err := client.ParseResourceID(*database.ID)
	if err != nil {
		return err
	}
	serverName := parent.Parent.Get("name").(*string)
	result, err := svc.ListByDatabase(ctx, details.ResourceGroup, *serverName, *database.Name)
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
