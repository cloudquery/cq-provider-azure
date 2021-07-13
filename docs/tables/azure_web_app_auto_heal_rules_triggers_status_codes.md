
# Table: azure_web_app_auto_heal_rules_triggers_status_codes
StatusCodesBasedTrigger trigger based on status code
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|app_cq_id|uuid|Unique ID of azure_web_apps table (FK)|
|app_id|text|ID of web app|
|status|integer|HTTP status code|
|sub_status|integer|Request Sub Status|
|win32_status|integer|Win32 error code|
|path|text|Request Path|
|count|integer|Request Count|
|time_interval|text|Time interval|
