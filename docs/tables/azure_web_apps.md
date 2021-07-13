
# Table: azure_web_apps
Site a web app, a mobile app backend, or an API app
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|subscription_id|text|Azure subscription id|
|state|text|Current state of the app|
|host_names|text[]|Hostnames associated with the app|
|repository_site_name|text|Name of the repository site|
|usage_state|text|State indicating whether the app has exceeded its quota usage Read-only Possible values include: 'UsageStateNormal', 'UsageStateExceeded'|
|enabled|boolean|otherwise, <code>false</code> Setting this value to false disables the app (takes the app offline)|
|enabled_host_names|text[]|Enabled hostnames for the appHostnames need to be assigned (see HostNames) AND enabled Otherwise, the app is not served on those hostnames|
|availability_state|text|Management information availability state for the app Possible values include: 'Normal', 'Limited', 'DisasterRecoveryMode'|
|server_farm_id|text|Resource ID of the associated App Service plan, formatted as: "/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/MicrosoftWeb/serverfarms/{appServicePlanName}"|
|reserved|boolean|otherwise, <code>false</code>|
|is_xenon|boolean|Obsolete: Hyper-V sandbox|
|hyper_v|boolean|Hyper-V sandbox|
|last_modified_time_utc_time|timestamp without time zone||
|number_of_workers|integer|Number of workers|
|default_documents|text[]|Default documents|
|net_framework_version|text|NET Framework version|
|php_version|text|Version of PHP|
|python_version|text|Version of Python|
|node_version|text|Version of Nodejs|
|power_shell_version|text|Version of PowerShell|
|linux_fx_version|text|Linux App Framework and version|
|windows_fx_version|text|Xenon App Framework and version|
|request_tracing_enabled|boolean|otherwise, <code>false</code>|
|request_tracing_expiration_time|timestamp without time zone||
|remote_debugging_enabled|boolean|otherwise, <code>false</code>|
|remote_debugging_version|text|Remote debugging version|
|http_logging_enabled|boolean|otherwise, <code>false</code>|
|logs_directory_size_limit|integer|HTTP logs directory size limit|
|detailed_error_logging_enabled|boolean|otherwise, <code>false</code>|
|publishing_username|text|Publishing user name|
|app_settings|jsonb|Application settings|
|azure_storage_accounts|jsonb|List of Azure Storage Accounts|
|machine_key_validation|text|MachineKey validation|
|machine_key_validation_key|text|Validation key|
|machine_key_decryption|text|Algorithm used for decryption|
|machine_key_decryption_key|text|Decryption key|
|document_root|text|Document root|
|scm_type|text|SCM type Possible values include: 'ScmTypeNone', 'ScmTypeDropbox', 'ScmTypeTfs', 'ScmTypeLocalGit', 'ScmTypeGitHub', 'ScmTypeCodePlexGit', 'ScmTypeCodePlexHg', 'ScmTypeBitbucketGit', 'ScmTypeBitbucketHg', 'ScmTypeExternalGit', 'ScmTypeExternalHg', 'ScmTypeOneDrive', 'ScmTypeVSO', 'ScmTypeVSTSRM'|
|use32_bit_worker_process|boolean|otherwise, <code>false</code>|
|web_sockets_enabled|boolean|otherwise, <code>false</code>|
|always_on|boolean|otherwise, <code>false</code>|
|java_version|text|Java version|
|java_container|text|Java container|
|java_container_version|text|Java container version|
|app_command_line|text|App command line to launch|
|managed_pipeline_mode|text|Managed pipeline mode Possible values include: 'Integrated', 'Classic'|
|load_balancing|text|Site load balancing Possible values include: 'WeightedRoundRobin', 'LeastRequests', 'LeastResponseTime', 'WeightedTotalTraffic', 'RequestHash', 'PerSiteRoundRobin'|
|limits_max_percentage_cpu|float|Maximum allowed CPU usage percentage|
|limits_max_memory_in_mb|bigint|Maximum allowed memory usage in MB|
|limits_max_disk_size_in_mb|bigint|Maximum allowed disk size usage in MB|
|auto_heal_enabled|boolean|otherwise, <code>false</code>|
|auto_heal_rules_triggers_requests_count|integer|Request Count|
|auto_heal_rules_triggers_requests_time_interval|text|Time interval|
|auto_heal_rules_triggers_private_bytes_in_kb|integer|A rule based on private bytes|
|auto_heal_rules_triggers_slow_requests_time_taken|text|Time taken|
|auto_heal_rules_triggers_slow_requests_path|text|Request Path|
|auto_heal_rules_triggers_slow_requests_count|integer|Request Count|
|auto_heal_rules_triggers_slow_requests_time_interval|text|Time interval|
|auto_heal_rules_actions_action_type|text|Predefined action to be taken Possible values include: 'Recycle', 'LogEvent', 'CustomAction'|
|auto_heal_rules_actions_custom_action_exe|text|Executable to be run|
|auto_heal_rules_actions_custom_action_parameters|text|Parameters for the executable|
|auto_heal_rules_actions_min_process_execution_time|text|Minimum time the process must execute before taking the action|
|tracing_options|text|Tracing options|
|vnet_name|text|Virtual Network name|
|vnet_route_all_enabled|boolean|Virtual Network Route All enabled This causes all outbound traffic to have Virtual Network Security Groups and User Defined Routes applied|
|vnet_private_ports_count|integer|The number of private ports assigned to this app These will be assigned dynamically on runtime|
|cors_allowed_origins|text[]|Gets or sets the list of origins that should be allowed to make cross-origin calls (for example: http://examplecom:12345) Use "*" to allow all|
|cors_support_credentials|boolean|Gets or sets whether CORS requests with credentials are allowed See https://developermozillaorg/en-US/docs/Web/HTTP/CORS#Requests_with_credentials for more details|
|push_push_settings_properties_is_push_enabled|boolean|Gets or sets a flag indicating whether the Push endpoint is enabled|
|push_push_settings_properties_tag_whitelist_json|text|Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint|
|push_push_settings_properties_tags_requiring_auth|text|Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint Tags can consist of alphanumeric characters and the following: '_', '@', '#', '', ':', '-' Validation should be performed at the PushRequestHandler|
|push_push_settings_properties_dynamic_tags_json|text|Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint|
|push_id|text|Resource Id|
|push_name|text|Resource Name|
|push_kind|text|Kind of resource|
|push_type|text|Resource type|
|api_definition_url|text|The URL of the API definition|
|api_management_config_id|text|APIM-Api Identifier|
|auto_swap_slot_name|text|Auto-swap slot name|
|local_my_sql_enabled|boolean|otherwise, <code>false</code>|
|managed_service_identity_id|integer|Managed Service Identity Id|
|x_managed_service_identity_id|integer|Explicit Managed Service Identity Id|
|scm_ip_security_restrictions_use_main|boolean|IP security restrictions for scm to use main|
|http_20_enabled|boolean|Http20Enabled: configures a web site to allow clients to connect over http20|
|min_tls_version|text|MinTlsVersion: configures the minimum version of TLS required for SSL requests Possible values include: 'OneFullStopZero', 'OneFullStopOne', 'OneFullStopTwo'|
|scm_min_tls_version|text|ScmMinTlsVersion: configures the minimum version of TLS required for SSL requests for SCM site Possible values include: 'OneFullStopZero', 'OneFullStopOne', 'OneFullStopTwo'|
|ftps_state|text|State of FTP / FTPS service Possible values include: 'AllAllowed', 'FtpsOnly', 'Disabled'|
|pre_warmed_instance_count|integer|Number of preWarmed instances This setting only applies to the Consumption and Elastic Plans|
|function_app_scale_limit|integer|Maximum number of workers that a site can scale out to This setting only applies to the Consumption and Elastic Premium Plans|
|health_check_path|text|Health check path|
|functions_runtime_scale_monitoring_enabled|boolean|Gets or sets a value indicating whether functions runtime scale monitoring is enabled When enabled, the ScaleController will not monitor event sources directly, but will instead call to the runtime to get scale status|
|website_time_zone|text|Sets the time zone a site uses for generating timestamps Compatible with Linux and Windows App Service Setting the WEBSITE_TIME_ZONE app setting takes precedence over this config For Linux, expects tz database values https://wwwianaorg/time-zones (for a quick reference see https://enwikipediaorg/wiki/List_of_tz_database_time_zones) For Windows, expects one of the time zones listed under HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows NT\CurrentVersion\Time Zones|
|minimum_elastic_instance_count|integer|Number of minimum instance count for a site This setting only applies to the Elastic Plans|
|traffic_manager_host_names|text[]|Azure Traffic Manager hostnames associated with the app Read-only|
|scm_site_also_stopped|boolean|otherwise, <code>false</code> The default is <code>false</code>|
|target_swap_slot|text|Specifies which deployment slot this app will swap into Read-only|
|hosting_environment_profile_id|text|Resource ID of the App Service Environment|
|hosting_environment_profile_name|text|Name of the App Service Environment|
|hosting_environment_profile_type|text|Resource type of the App Service Environment|
|client_affinity_enabled|boolean|<code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance Default is <code>true</code>|
|client_cert_enabled|boolean|otherwise, <code>false</code> Default is <code>false</code>|
|client_cert_mode|text|This composes with ClientCertEnabled setting - ClientCertEnabled: false means ClientCert is ignored - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted Possible values include: 'Required', 'Optional', 'OptionalInteractiveUser'|
|client_cert_exclusion_paths|text|client certificate authentication comma-separated exclusion paths|
|host_names_disabled|boolean|otherwise, <code>false</code>  If <code>true</code>, the app is only accessible via API management process|
|custom_domain_verification_id|text|Unique identifier that verifies the custom domains assigned to the app Customer will add this id to a txt record for verification|
|outbound_ip_addresses|text|List of IP addresses that the app uses for outbound connections (eg database access) Includes VIPs from tenants that site can be hosted with current settings Read-only|
|possible_outbound_ip_addresses|text|List of IP addresses that the app uses for outbound connections (eg database access) Includes VIPs from all tenants except dataComponent Read-only|
|container_size|integer|Size of the function container|
|daily_memory_time_quota|integer|Maximum allowed daily memory-time quota (applicable on dynamic apps only)|
|suspended_till_time|timestamp without time zone||
|max_number_of_workers|integer|Maximum number of workers This only applies to Functions container|
|cloning_info_correlation_id|uuid|Correlation ID of cloning operation This ID ties multiple cloning operations together to use the same snapshot|
|cloning_info_overwrite|boolean|otherwise, <code>false</code>|
|cloning_info_clone_custom_host_names|boolean|otherwise, <code>false</code>|
|cloning_info_clone_source_control|boolean|otherwise, <code>false</code>|
|cloning_info_source_web_app_id|text|ARM resource ID of the source app App resource ID is of the form /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftWeb/sites/{siteName} for production slots and /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftWeb/sites/{siteName}/slots/{slotName} for other slots|
|cloning_info_source_web_app_location|text|Location of source app ex: West US or North Europe|
|cloning_info_hosting_environment|text|App Service Environment|
|cloning_info_app_settings_overrides|jsonb|Application setting overrides for cloned app If specified, these settings override the settings cloned from source app Otherwise, application settings from source app are retained|
|cloning_info_configure_load_balancing|boolean|<code>true</code> to configure load balancing for source and destination app|
|cloning_info_traffic_manager_profile_id|text|ARM resource ID of the Traffic Manager profile to use, if it exists Traffic Manager resource ID is of the form /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftNetwork/trafficManagerProfiles/{profileName}|
|cloning_info_traffic_manager_profile_name|text|Name of Traffic Manager profile to create This is only needed if Traffic Manager profile does not already exist|
|resource_group|text|Name of the resource group the app belongs to Read-only|
|is_default_container|boolean|<code>true</code> if the app is a default container; otherwise, <code>false</code>|
|default_host_name|text|Default hostname of the app Read-only|
|slot_swap_status_timestamp_utc_time|timestamp without time zone||
|slot_swap_status_source_slot_name|text|The source slot of the last swap operation|
|slot_swap_status_destination_slot_name|text|The destination slot of the last swap operation|
|key_vault_reference_identity|text|Identity to use for Key Vault Reference authentication|
|https_only|boolean|HttpsOnly: configures a web site to accept only https requests Issues redirect for http requests|
|redundancy_mode|text|Site redundancy mode Possible values include: 'RedundancyModeNone', 'RedundancyModeManual', 'RedundancyModeFailover', 'RedundancyModeActiveActive', 'RedundancyModeGeoRedundant'|
|in_progress_operation_id|uuid|Specifies an operation id if this site has a pending operation|
|storage_account_required|boolean|Checks if Customer provided storage account is required|
|identity_type|text|Type of managed service identity Possible values include: 'ManagedServiceIdentityTypeSystemAssigned', 'ManagedServiceIdentityTypeUserAssigned', 'ManagedServiceIdentityTypeSystemAssignedUserAssigned', 'ManagedServiceIdentityTypeNone'|
|identity_tenant_id|text|Tenant of managed service identity|
|identity_principal_id|text|Principal Id of managed service identity|
|identity_user_assigned_identities|jsonb|The list of user assigned identities associated with the resource The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/MicrosoftManagedIdentity/userAssignedIdentities/{identityName}|
|id|text|Resource Id|
|name|text|Resource Name|
|kind|text|Kind of resource|
|location|text|Resource Location|
|type|text|Resource type|
|tags|jsonb|Resource tags|
