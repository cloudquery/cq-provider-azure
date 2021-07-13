
# Table: azure_sql_server_devops_audit_settings
ServerDevOpsAuditingSettings a server DevOps auditing settings
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_cq_id|uuid|Unique ID of azure_sql_servers table (FK)|
|created_by|text|A string identifier for the identity that created the resource|
|created_by_type|text|The type of identity that created the resource: <User|Application|ManagedIdentity|Key> Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'|
|created_at_time|timestamp without time zone|The timestamp of resource creation (UTC).|
|last_modified_by|text|A string identifier for the identity that last modified the resource|
|last_modified_by_type|text|The type of identity that last modified the resource: <User|Application|ManagedIdentity|Key> Possible values include: 'User', 'Application', 'ManagedIdentity', 'Key'|
|last_modified_at_time|timestamp without time zone|The timestamp of last modification (UTC).|
|is_azure_monitor_target_enabled|boolean|Specifies whether DevOps audit events are sent to Azure Monitor In order to send the events to Azure Monitor, specify 'State' as 'Enabled' and 'IsAzureMonitorTargetEnabled' as true  When using REST API to configure DevOps audit, Diagnostic Settings with 'DevOpsOperationsAudit' diagnostic logs category on the master database should be also created  Diagnostic Settings URI format: PUT https://managementazurecom/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/MicrosoftSql/servers/{serverName}/databases/master/providers/microsoftinsights/diagnosticSettings/{settingsName}?api-version=2017-05-01-preview  For more information, see [Diagnostic Settings REST API](https://gomicrosoftcom/fwlink/?linkid=2033207) or [Diagnostic Settings PowerShell](https://gomicrosoftcom/fwlink/?linkid=2033043)|
|state|text|Specifies the state of the audit If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled are required Possible values include: 'BlobAuditingPolicyStateEnabled', 'BlobAuditingPolicyStateDisabled'|
|storage_endpoint|text|Specifies the blob storage endpoint (eg https://MyAccountblobcorewindowsnet) If state is Enabled, storageEndpoint or isAzureMonitorTargetEnabled is required|
|storage_account_access_key|text|Specifies the identifier key of the auditing storage account If state is Enabled and storageEndpoint is specified, not specifying the storageAccountAccessKey will use SQL server system-assigned managed identity to access the storage Prerequisites for using managed identity authentication: 1 Assign SQL Server a system-assigned managed identity in Azure Active Directory (AAD) 2 Grant SQL Server identity access to the storage account by adding 'Storage Blob Data Contributor' RBAC role to the server identity For more information, see [Auditing to storage using Managed Identity authentication](https://gomicrosoftcom/fwlink/?linkid=2114355)|
|storage_account_subscription_id|uuid|Specifies the blob storage subscription Id|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
