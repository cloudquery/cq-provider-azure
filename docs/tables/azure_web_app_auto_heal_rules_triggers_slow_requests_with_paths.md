
# Table: azure_web_app_auto_heal_rules_triggers_slow_requests_with_paths
SlowRequestsBasedTrigger trigger based on request execution time
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|app_cq_id|uuid|Unique ID of azure_web_apps table (FK)|
|app_id|text|ID of web app|
|time_taken|text|Time taken|
|path|text|Request Path|
|count|integer|Request Count|
|time_interval|text|Time interval|
