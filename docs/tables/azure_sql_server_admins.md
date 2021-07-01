
# Table: azure_sql_server_admins
ServerAzureADAdministrator an server Active Directory Administrator
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|server_id|uuid|Unique ID of azure_sql_servers table (FK)|
|administrator_type|text|The type of administrator|
|login|text|The server administrator login value|
|sid|uuid|The server administrator Sid (Secure ID)|
|tenant_id|uuid|The server Active Directory Administrator tenant id|
|resource_id|text|Resource ID|
|name|text|Resource name|
|type|text|Resource type|
