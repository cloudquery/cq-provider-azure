
# Table: azure_ad_group_team_operations

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|group_cq_id|uuid|Unique CloudQuery ID of azure_ad_groups table (FK)|
|id|text||
|operation_type|text||
|created_date_time|timestamp without time zone||
|status|text||
|last_action_date_time|timestamp without time zone||
|attempts_count|bigint||
|target_resource_id|text||
|target_resource_location|text||
|error_code|text||
|error_message|text||
