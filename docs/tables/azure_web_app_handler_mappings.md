
# Table: azure_web_app_handler_mappings
HandlerMapping the IIS handler mappings used to define which handler processes HTTP requests with certain extension For example, it is used to configure php-cgiexe process to handle all HTTP requests with *php extension
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|app_cq_id|uuid|Unique ID of azure_web_apps table (FK)|
|app_id|text|ID of web app|
|extension|text|Requests with this extension will be handled using the specified FastCGI application|
|script_processor|text|The absolute path to the FastCGI application|
|arguments|text|Command-line arguments to be passed to the script processor|
