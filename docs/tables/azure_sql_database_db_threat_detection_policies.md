
# Table: azure_sql_database_db_threat_detection_policies
Contains information about a database Threat Detection policy.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique ID of azure_sql_databases table (FK)|
|location|text|The geo-location where the resource lives|
|kind|text|Resource kind|
|state|text|Specifies the state of the policy If state is Enabled, storageEndpoint and storageAccountAccessKey are required Possible values include: 'SecurityAlertPolicyStateNew', 'SecurityAlertPolicyStateEnabled', 'SecurityAlertPolicyStateDisabled'|
|disabled_alerts|text|Sql_Injection_Vulnerability; Access_Anomaly; Data_Exfiltration; Unsafe_Action|
|email_addresses|text|Specifies the semicolon-separated list of e-mail addresses to which the alert is sent|
|email_account_admins|text|Specifies that the alert is sent to the account administrators Possible values include: 'SecurityAlertPolicyEmailAccountAdminsEnabled', 'SecurityAlertPolicyEmailAccountAdminsDisabled'|
|storage_endpoint|text|Specifies the blob storage endpoint (eg https://MyAccountblobcorewindowsnet) This blob storage will hold all Threat Detection audit logs If state is Enabled, storageEndpoint is required|
|storage_account_access_key|text|Specifies the identifier key of the Threat Detection audit storage account If state is Enabled, storageAccountAccessKey is required|
|retention_days|integer|Specifies the number of days to keep in the Threat Detection audit logs|
|use_server_default|text|Specifies whether to use the default server policy Possible values include: 'SecurityAlertPolicyUseServerDefaultEnabled', 'SecurityAlertPolicyUseServerDefaultDisabled'|
|id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
