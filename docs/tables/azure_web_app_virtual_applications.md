
# Table: azure_web_app_virtual_applications
VirtualApplication virtual application in an app
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|app_cq_id|uuid|Unique ID of azure_web_apps table (FK)|
|app_id|text|ID of web app|
|virtual_path|text|Virtual path|
|physical_path|text|Physical path|
|preload_enabled|boolean|otherwise, <code>false</code>|
