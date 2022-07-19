
# Table: azure_servicebus_namespace_private_endpoint_connections
PrivateEndpointConnection properties of the PrivateEndpointConnection
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|namespace_cq_id|uuid|Unique CloudQuery ID of azure_servicebus_namespaces table (FK)|
|private_endpoint_id|text|The ARM identifier for Private Endpoint|
|status|text|Status of the connection|
|status_description|text|Description of the connection state|
|provisioning_state|text|Provisioning state of the Private Endpoint Connection|
|system_data_created_by|text|The identity that created the resource|
|system_data_created_by_type|text|The type of identity that created the resource|
|system_data_created_at_time|timestamp without time zone||
|system_data_last_modified_by|text|The identity that last modified the resource|
|system_data_last_modified_by_type|text|The type of identity that last modified the resource|
|system_data_last_modified_at_time|timestamp without time zone||
|id|text|Resource Id|
|name|text|Resource name|
|type|text|Resource type|
