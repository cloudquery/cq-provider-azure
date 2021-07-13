
# Table: azure_web_app_connection_strings
ConnStringInfo database connection string information
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|app_cq_id|uuid|Unique ID of azure_web_apps table (FK)|
|app_id|text|ID of web app|
|name|text|Name of connection string|
|connection_string|text|Connection string value|
|type|text|Type of database Possible values include: 'MySQL', 'SQLServer', 'SQLAzure', 'Custom', 'NotificationHub', 'ServiceBus', 'EventHub', 'APIHub', 'DocDb', 'RedisCache', 'PostgreSQL'|
