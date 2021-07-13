
# Table: azure_web_app_auto_heal_rules_triggers_status_codes_ranges
StatusCodesRangeBasedTrigger trigger based on range of status codes
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|app_cq_id|uuid|Unique ID of azure_web_apps table (FK)|
|app_id|text|ID of web app|
|status_codes|text|HTTP status code|
|path|text||
|count|integer|Request Count|
|time_interval|text|Time interval|
