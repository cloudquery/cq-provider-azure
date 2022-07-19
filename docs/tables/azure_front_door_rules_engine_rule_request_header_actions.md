
# Table: azure_front_door_rules_engine_rule_request_header_actions
A list of header actions to apply from the request from AFD to the origin.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|front_door_rules_engine_rule_cq_id|uuid|Unique CloudQuery ID of azure_front_door_rules_engine_rules table (FK)|
|action_type|text|Which type of manipulation to apply to the header|
|name|text|The name of the header the action will apply to|
|value|text|The value to update the given header name with|
