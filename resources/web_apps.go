package resources

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/web/mgmt/2020-12-01/web"
	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func WebApps() *schema.Table {
	return &schema.Table{
		Name:        "azure_web_apps",
		Description: "Site a web app, a mobile app backend, or an API app",
		Resolver:    fetchWebApps,
		Multiplex:   client.SubscriptionMultiplex,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"subscription_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "state",
				Description: "Current state of the app",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.State"),
			},
			{
				Name:        "host_names",
				Description: "Hostnames associated with the app",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteProperties.HostNames"),
			},
			{
				Name:        "repository_site_name",
				Description: "Name of the repository site",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.RepositorySiteName"),
			},
			{
				Name:        "usage_state",
				Description: "State indicating whether the app has exceeded its quota usage Read-only Possible values include: 'UsageStateNormal', 'UsageStateExceeded'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.UsageState"),
			},
			{
				Name:        "enabled",
				Description: "otherwise, <code>false</code> Setting this value to false disables the app (takes the app offline)",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.Enabled"),
			},
			{
				Name:        "enabled_host_names",
				Description: "Enabled hostnames for the appHostnames need to be assigned (see HostNames) AND enabled Otherwise, the app is not served on those hostnames",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteProperties.EnabledHostNames"),
			},
			{
				Name:        "availability_state",
				Description: "Management information availability state for the app Possible values include: 'Normal', 'Limited', 'DisasterRecoveryMode'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.AvailabilityState"),
			},
			{
				Name:        "server_farm_id",
				Description: "Resource ID of the associated App Service plan, formatted as: \"/subscriptions/{subscriptionID}/resourceGroups/{groupName}/providers/MicrosoftWeb/serverfarms/{appServicePlanName}\"",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.ServerFarmID"),
			},
			{
				Name:        "reserved",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.Reserved"),
			},
			{
				Name:        "is_xenon",
				Description: "Obsolete: Hyper-V sandbox",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.IsXenon"),
			},
			{
				Name:        "hyper_v",
				Description: "Hyper-V sandbox",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.HyperV"),
			},
			{
				Name:     "last_modified_time_utc_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SiteProperties.LastModifiedTimeUtc.Time"),
			},
			{
				Name:        "number_of_workers",
				Description: "Number of workers",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.NumberOfWorkers"),
			},
			{
				Name:        "default_documents",
				Description: "Default documents",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.DefaultDocuments"),
			},
			{
				Name:        "net_framework_version",
				Description: "NET Framework version",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.NetFrameworkVersion"),
			},
			{
				Name:        "php_version",
				Description: "Version of PHP",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.PhpVersion"),
			},
			{
				Name:        "python_version",
				Description: "Version of Python",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.PythonVersion"),
			},
			{
				Name:        "node_version",
				Description: "Version of Nodejs",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.NodeVersion"),
			},
			{
				Name:        "power_shell_version",
				Description: "Version of PowerShell",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.PowerShellVersion"),
			},
			{
				Name:        "linux_fx_version",
				Description: "Linux App Framework and version",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.LinuxFxVersion"),
			},
			{
				Name:        "windows_fx_version",
				Description: "Xenon App Framework and version",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.WindowsFxVersion"),
			},
			{
				Name:        "request_tracing_enabled",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.RequestTracingEnabled"),
			},
			{
				Name:     "request_tracing_expiration_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SiteProperties.SiteConfig.RequestTracingExpirationTime.Time"),
			},
			{
				Name:        "remote_debugging_enabled",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.RemoteDebuggingEnabled"),
			},
			{
				Name:        "remote_debugging_version",
				Description: "Remote debugging version",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.RemoteDebuggingVersion"),
			},
			{
				Name:        "http_logging_enabled",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.HTTPLoggingEnabled"),
			},
			{
				Name:        "logs_directory_size_limit",
				Description: "HTTP logs directory size limit",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.LogsDirectorySizeLimit"),
			},
			{
				Name:        "detailed_error_logging_enabled",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.DetailedErrorLoggingEnabled"),
			},
			{
				Name:        "publishing_username",
				Description: "Publishing user name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.PublishingUsername"),
			},
			{
				Name:        "app_settings",
				Description: "Application settings",
				Type:        schema.TypeJSON,
				Resolver:    resolveWebAppAppSettings,
			},
			{
				Name:        "azure_storage_accounts",
				Description: "List of Azure Storage Accounts",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AzureStorageAccounts"),
			},
			{
				Name:        "machine_key_validation",
				Description: "MachineKey validation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.MachineKey.Validation"),
			},
			{
				Name:        "machine_key_validation_key",
				Description: "Validation key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.MachineKey.ValidationKey"),
			},
			{
				Name:        "machine_key_decryption",
				Description: "Algorithm used for decryption",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.MachineKey.Decryption"),
			},
			{
				Name:        "machine_key_decryption_key",
				Description: "Decryption key",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.MachineKey.DecryptionKey"),
			},
			{
				Name:        "document_root",
				Description: "Document root",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.DocumentRoot"),
			},
			{
				Name:        "scm_type",
				Description: "SCM type Possible values include: 'ScmTypeNone', 'ScmTypeDropbox', 'ScmTypeTfs', 'ScmTypeLocalGit', 'ScmTypeGitHub', 'ScmTypeCodePlexGit', 'ScmTypeCodePlexHg', 'ScmTypeBitbucketGit', 'ScmTypeBitbucketHg', 'ScmTypeExternalGit', 'ScmTypeExternalHg', 'ScmTypeOneDrive', 'ScmTypeVSO', 'ScmTypeVSTSRM'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.ScmType"),
			},
			{
				Name:        "use32_bit_worker_process",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Use32BitWorkerProcess"),
			},
			{
				Name:        "web_sockets_enabled",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.WebSocketsEnabled"),
			},
			{
				Name:        "always_on",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AlwaysOn"),
			},
			{
				Name:        "java_version",
				Description: "Java version",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.JavaVersion"),
			},
			{
				Name:        "java_container",
				Description: "Java container",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.JavaContainer"),
			},
			{
				Name:        "java_container_version",
				Description: "Java container version",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.JavaContainerVersion"),
			},
			{
				Name:        "app_command_line",
				Description: "App command line to launch",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AppCommandLine"),
			},
			{
				Name:        "managed_pipeline_mode",
				Description: "Managed pipeline mode Possible values include: 'Integrated', 'Classic'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.ManagedPipelineMode"),
			},
			{
				Name:        "load_balancing",
				Description: "Site load balancing Possible values include: 'WeightedRoundRobin', 'LeastRequests', 'LeastResponseTime', 'WeightedTotalTraffic', 'RequestHash', 'PerSiteRoundRobin'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.LoadBalancing"),
			},
			{
				Name:        "limits_max_percentage_cpu",
				Description: "Maximum allowed CPU usage percentage",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Limits.MaxPercentageCPU"),
			},
			{
				Name:        "limits_max_memory_in_mb",
				Description: "Maximum allowed memory usage in MB",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Limits.MaxMemoryInMb"),
			},
			{
				Name:        "limits_max_disk_size_in_mb",
				Description: "Maximum allowed disk size usage in MB",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Limits.MaxDiskSizeInMb"),
			},
			{
				Name:        "auto_heal_enabled",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealEnabled"),
			},
			{
				Name:        "auto_heal_rules_triggers_requests_count",
				Description: "Request Count",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealRules.Triggers.Requests.Count"),
			},
			{
				Name:        "auto_heal_rules_triggers_requests_time_interval",
				Description: "Time interval",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealRules.Triggers.Requests.TimeInterval"),
			},
			{
				Name:        "auto_heal_rules_triggers_private_bytes_in_kb",
				Description: "A rule based on private bytes",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealRules.Triggers.PrivateBytesInKB"),
			},
			{
				Name:        "auto_heal_rules_triggers_slow_requests_time_taken",
				Description: "Time taken",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealRules.Triggers.SlowRequests.TimeTaken"),
			},
			{
				Name:        "auto_heal_rules_triggers_slow_requests_path",
				Description: "Request Path",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealRules.Triggers.SlowRequests.Path"),
			},
			{
				Name:        "auto_heal_rules_triggers_slow_requests_count",
				Description: "Request Count",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealRules.Triggers.SlowRequests.Count"),
			},
			{
				Name:        "auto_heal_rules_triggers_slow_requests_time_interval",
				Description: "Time interval",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealRules.Triggers.SlowRequests.TimeInterval"),
			},
			{
				Name:        "auto_heal_rules_actions_action_type",
				Description: "Predefined action to be taken Possible values include: 'Recycle', 'LogEvent', 'CustomAction'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealRules.Actions.ActionType"),
			},
			{
				Name:        "auto_heal_rules_actions_custom_action_exe",
				Description: "Executable to be run",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealRules.Actions.CustomAction.Exe"),
			},
			{
				Name:        "auto_heal_rules_actions_custom_action_parameters",
				Description: "Parameters for the executable",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealRules.Actions.CustomAction.Parameters"),
			},
			{
				Name:        "auto_heal_rules_actions_min_process_execution_time",
				Description: "Minimum time the process must execute before taking the action",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoHealRules.Actions.MinProcessExecutionTime"),
			},
			{
				Name:        "tracing_options",
				Description: "Tracing options",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.TracingOptions"),
			},
			{
				Name:        "vnet_name",
				Description: "Virtual Network name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.VnetName"),
			},
			{
				Name:        "vnet_route_all_enabled",
				Description: "Virtual Network Route All enabled This causes all outbound traffic to have Virtual Network Security Groups and User Defined Routes applied",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.VnetRouteAllEnabled"),
			},
			{
				Name:        "vnet_private_ports_count",
				Description: "The number of private ports assigned to this app These will be assigned dynamically on runtime",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.VnetPrivatePortsCount"),
			},
			{
				Name:        "cors_allowed_origins",
				Description: "Gets or sets the list of origins that should be allowed to make cross-origin calls (for example: http://examplecom:12345) Use \"*\" to allow all",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Cors.AllowedOrigins"),
			},
			{
				Name:        "cors_support_credentials",
				Description: "Gets or sets whether CORS requests with credentials are allowed See https://developermozillaorg/en-US/docs/Web/HTTP/CORS#Requests_with_credentials for more details",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Cors.SupportCredentials"),
			},
			{
				Name:        "push_push_settings_properties_is_push_enabled",
				Description: "Gets or sets a flag indicating whether the Push endpoint is enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Push.PushSettingsProperties.IsPushEnabled"),
			},
			{
				Name:        "push_push_settings_properties_tag_whitelist_json",
				Description: "Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Push.PushSettingsProperties.TagWhitelistJSON"),
			},
			{
				Name:        "push_push_settings_properties_tags_requiring_auth",
				Description: "Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint Tags can consist of alphanumeric characters and the following: '_', '@', '#', '', ':', '-' Validation should be performed at the PushRequestHandler",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Push.PushSettingsProperties.TagsRequiringAuth"),
			},
			{
				Name:        "push_push_settings_properties_dynamic_tags_json",
				Description: "Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Push.PushSettingsProperties.DynamicTagsJSON"),
			},
			{
				Name:        "push_id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Push.ID"),
			},
			{
				Name:        "push_name",
				Description: "Resource Name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Push.Name"),
			},
			{
				Name:        "push_kind",
				Description: "Kind of resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Push.Kind"),
			},
			{
				Name:        "push_type",
				Description: "Resource type",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.Push.Type"),
			},
			{
				Name:        "api_definition_url",
				Description: "The URL of the API definition",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.APIDefinition.URL"),
			},
			{
				Name:        "api_management_config_id",
				Description: "APIM-Api Identifier",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.APIManagementConfig.ID"),
			},
			{
				Name:        "auto_swap_slot_name",
				Description: "Auto-swap slot name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.AutoSwapSlotName"),
			},
			{
				Name:        "local_my_sql_enabled",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.LocalMySQLEnabled"),
			},
			{
				Name:        "managed_service_identity_id",
				Description: "Managed Service Identity Id",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.ManagedServiceIdentityID"),
			},
			{
				Name:        "x_managed_service_identity_id",
				Description: "Explicit Managed Service Identity Id",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.XManagedServiceIdentityID"),
			},
			{
				Name:        "scm_ip_security_restrictions_use_main",
				Description: "IP security restrictions for scm to use main",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.ScmIPSecurityRestrictionsUseMain"),
			},
			{
				Name:        "http_20_enabled",
				Description: "Http20Enabled: configures a web site to allow clients to connect over http20",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.HTTP20Enabled"),
			},
			{
				Name:        "min_tls_version",
				Description: "MinTlsVersion: configures the minimum version of TLS required for SSL requests Possible values include: 'OneFullStopZero', 'OneFullStopOne', 'OneFullStopTwo'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.MinTLSVersion"),
			},
			{
				Name:        "scm_min_tls_version",
				Description: "ScmMinTlsVersion: configures the minimum version of TLS required for SSL requests for SCM site Possible values include: 'OneFullStopZero', 'OneFullStopOne', 'OneFullStopTwo'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.ScmMinTLSVersion"),
			},
			{
				Name:        "ftps_state",
				Description: "State of FTP / FTPS service Possible values include: 'AllAllowed', 'FtpsOnly', 'Disabled'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.FtpsState"),
			},
			{
				Name:        "pre_warmed_instance_count",
				Description: "Number of preWarmed instances This setting only applies to the Consumption and Elastic Plans",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.PreWarmedInstanceCount"),
			},
			{
				Name:        "function_app_scale_limit",
				Description: "Maximum number of workers that a site can scale out to This setting only applies to the Consumption and Elastic Premium Plans",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.FunctionAppScaleLimit"),
			},
			{
				Name:        "health_check_path",
				Description: "Health check path",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.HealthCheckPath"),
			},
			{
				Name:        "functions_runtime_scale_monitoring_enabled",
				Description: "Gets or sets a value indicating whether functions runtime scale monitoring is enabled When enabled, the ScaleController will not monitor event sources directly, but will instead call to the runtime to get scale status",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.FunctionsRuntimeScaleMonitoringEnabled"),
			},
			{
				Name:        "website_time_zone",
				Description: "Sets the time zone a site uses for generating timestamps Compatible with Linux and Windows App Service Setting the WEBSITE_TIME_ZONE app setting takes precedence over this config For Linux, expects tz database values https://wwwianaorg/time-zones (for a quick reference see https://enwikipediaorg/wiki/List_of_tz_database_time_zones) For Windows, expects one of the time zones listed under HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\Time Zones",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.WebsiteTimeZone"),
			},
			{
				Name:        "minimum_elastic_instance_count",
				Description: "Number of minimum instance count for a site This setting only applies to the Elastic Plans",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.SiteConfig.MinimumElasticInstanceCount"),
			},
			{
				Name:        "traffic_manager_host_names",
				Description: "Azure Traffic Manager hostnames associated with the app Read-only",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("SiteProperties.TrafficManagerHostNames"),
			},
			{
				Name:        "scm_site_also_stopped",
				Description: "otherwise, <code>false</code> The default is <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.ScmSiteAlsoStopped"),
			},
			{
				Name:        "target_swap_slot",
				Description: "Specifies which deployment slot this app will swap into Read-only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.TargetSwapSlot"),
			},
			{
				Name:        "hosting_environment_profile_id",
				Description: "Resource ID of the App Service Environment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.HostingEnvironmentProfile.ID"),
			},
			{
				Name:        "hosting_environment_profile_name",
				Description: "Name of the App Service Environment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.HostingEnvironmentProfile.Name"),
			},
			{
				Name:        "hosting_environment_profile_type",
				Description: "Resource type of the App Service Environment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.HostingEnvironmentProfile.Type"),
			},
			{
				Name:        "client_affinity_enabled",
				Description: "<code>false</code> to stop sending session affinity cookies, which route client requests in the same session to the same instance Default is <code>true</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.ClientAffinityEnabled"),
			},
			{
				Name:        "client_cert_enabled",
				Description: "otherwise, <code>false</code> Default is <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.ClientCertEnabled"),
			},
			{
				Name:        "client_cert_mode",
				Description: "This composes with ClientCertEnabled setting - ClientCertEnabled: false means ClientCert is ignored - ClientCertEnabled: true and ClientCertMode: Required means ClientCert is required - ClientCertEnabled: true and ClientCertMode: Optional means ClientCert is optional or accepted Possible values include: 'Required', 'Optional', 'OptionalInteractiveUser'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.ClientCertMode"),
			},
			{
				Name:        "client_cert_exclusion_paths",
				Description: "client certificate authentication comma-separated exclusion paths",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.ClientCertExclusionPaths"),
			},
			{
				Name:        "host_names_disabled",
				Description: "otherwise, <code>false</code>  If <code>true</code>, the app is only accessible via API management process",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.HostNamesDisabled"),
			},
			{
				Name:        "custom_domain_verification_id",
				Description: "Unique identifier that verifies the custom domains assigned to the app Customer will add this id to a txt record for verification",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CustomDomainVerificationID"),
			},
			{
				Name:        "outbound_ip_addresses",
				Description: "List of IP addresses that the app uses for outbound connections (eg database access) Includes VIPs from tenants that site can be hosted with current settings Read-only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.OutboundIPAddresses"),
			},
			{
				Name:        "possible_outbound_ip_addresses",
				Description: "List of IP addresses that the app uses for outbound connections (eg database access) Includes VIPs from all tenants except dataComponent Read-only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.PossibleOutboundIPAddresses"),
			},
			{
				Name:        "container_size",
				Description: "Size of the function container",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.ContainerSize"),
			},
			{
				Name:        "daily_memory_time_quota",
				Description: "Maximum allowed daily memory-time quota (applicable on dynamic apps only)",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.DailyMemoryTimeQuota"),
			},
			{
				Name:     "suspended_till_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SiteProperties.SuspendedTill.Time"),
			},
			{
				Name:        "max_number_of_workers",
				Description: "Maximum number of workers This only applies to Functions container",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("SiteProperties.MaxNumberOfWorkers"),
			},
			{
				Name:        "cloning_info_correlation_id",
				Description: "Correlation ID of cloning operation This ID ties multiple cloning operations together to use the same snapshot",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.CorrelationID"),
			},
			{
				Name:        "cloning_info_overwrite",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.Overwrite"),
			},
			{
				Name:        "cloning_info_clone_custom_host_names",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.CloneCustomHostNames"),
			},
			{
				Name:        "cloning_info_clone_source_control",
				Description: "otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.CloneSourceControl"),
			},
			{
				Name:        "cloning_info_source_web_app_id",
				Description: "ARM resource ID of the source app App resource ID is of the form /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftWeb/sites/{siteName} for production slots and /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftWeb/sites/{siteName}/slots/{slotName} for other slots",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.SourceWebAppID"),
			},
			{
				Name:        "cloning_info_source_web_app_location",
				Description: "Location of source app ex: West US or North Europe",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.SourceWebAppLocation"),
			},
			{
				Name:        "cloning_info_hosting_environment",
				Description: "App Service Environment",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.HostingEnvironment"),
			},
			{
				Name:        "cloning_info_app_settings_overrides",
				Description: "Application setting overrides for cloned app If specified, these settings override the settings cloned from source app Otherwise, application settings from source app are retained",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.AppSettingsOverrides"),
			},
			{
				Name:        "cloning_info_configure_load_balancing",
				Description: "<code>true</code> to configure load balancing for source and destination app",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.ConfigureLoadBalancing"),
			},
			{
				Name:        "cloning_info_traffic_manager_profile_id",
				Description: "ARM resource ID of the Traffic Manager profile to use, if it exists Traffic Manager resource ID is of the form /subscriptions/{subId}/resourceGroups/{resourceGroupName}/providers/MicrosoftNetwork/trafficManagerProfiles/{profileName}",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.TrafficManagerProfileID"),
			},
			{
				Name:        "cloning_info_traffic_manager_profile_name",
				Description: "Name of Traffic Manager profile to create This is only needed if Traffic Manager profile does not already exist",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.CloningInfo.TrafficManagerProfileName"),
			},
			{
				Name:        "resource_group",
				Description: "Name of the resource group the app belongs to Read-only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.ResourceGroup"),
			},
			{
				Name:        "is_default_container",
				Description: "<code>true</code> if the app is a default container; otherwise, <code>false</code>",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.IsDefaultContainer"),
			},
			{
				Name:        "default_host_name",
				Description: "Default hostname of the app Read-only",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.DefaultHostName"),
			},
			{
				Name:     "slot_swap_status_timestamp_utc_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("SiteProperties.SlotSwapStatus.TimestampUtc.Time"),
			},
			{
				Name:        "slot_swap_status_source_slot_name",
				Description: "The source slot of the last swap operation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SlotSwapStatus.SourceSlotName"),
			},
			{
				Name:        "slot_swap_status_destination_slot_name",
				Description: "The destination slot of the last swap operation",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.SlotSwapStatus.DestinationSlotName"),
			},
			{
				Name:        "key_vault_reference_identity",
				Description: "Identity to use for Key Vault Reference authentication",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.KeyVaultReferenceIdentity"),
			},
			{
				Name:        "https_only",
				Description: "HttpsOnly: configures a web site to accept only https requests Issues redirect for http requests",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.HTTPSOnly"),
			},
			{
				Name:        "redundancy_mode",
				Description: "Site redundancy mode Possible values include: 'RedundancyModeNone', 'RedundancyModeManual', 'RedundancyModeFailover', 'RedundancyModeActiveActive', 'RedundancyModeGeoRedundant'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SiteProperties.RedundancyMode"),
			},
			{
				Name:        "in_progress_operation_id",
				Description: "Specifies an operation id if this site has a pending operation",
				Type:        schema.TypeUUID,
				Resolver:    schema.PathResolver("SiteProperties.InProgressOperationID"),
			},
			{
				Name:        "storage_account_required",
				Description: "Checks if Customer provided storage account is required",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("SiteProperties.StorageAccountRequired"),
			},
			{
				Name:        "identity_type",
				Description: "Type of managed service identity Possible values include: 'ManagedServiceIdentityTypeSystemAssigned', 'ManagedServiceIdentityTypeUserAssigned', 'ManagedServiceIdentityTypeSystemAssignedUserAssigned', 'ManagedServiceIdentityTypeNone'",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.Type"),
			},
			{
				Name:        "identity_tenant_id",
				Description: "Tenant of managed service identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.TenantID"),
			},
			{
				Name:        "identity_principal_id",
				Description: "Principal Id of managed service identity",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Identity.PrincipalID"),
			},
			{
				Name:        "identity_user_assigned_identities",
				Description: "The list of user assigned identities associated with the resource The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/MicrosoftManagedIdentity/userAssignedIdentities/{identityName}",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Identity.UserAssignedIdentities"),
			},
			{
				Name:        "id",
				Description: "Resource Id",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource Name",
				Type:        schema.TypeString,
			},
			{
				Name:        "kind",
				Description: "Kind of resource",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource Location",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:         "azure_web_app_publishing_profiles",
				Description:  "HostNameSslState SSL-enabled hostname",
				Resolver:     fetchWebAppPublishingProfiles,
				AlwaysDelete: true,
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "ID of web app",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name: "publish_url",
						Type: schema.TypeString,
					},
					{
						Name: "user_name",
						Type: schema.TypeString,
					},
					{
						Name:     "user_pwd",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("UserPWD"),
					},
				},
			},
			{
				Name:        "azure_web_app_host_name_ssl_states",
				Description: "HostNameSslState SSL-enabled hostname",
				Resolver:    fetchWebAppHostNameSslStates,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"app_cq_id", "name"}},
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "ID of web app",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "name",
						Description: "Hostname",
						Type:        schema.TypeString,
					},
					{
						Name:        "ssl_state",
						Description: "SSL type Possible values include: 'SslStateDisabled', 'SslStateSniEnabled', 'SslStateIPBasedEnabled'",
						Type:        schema.TypeString,
					},
					{
						Name:        "virtual_ip",
						Description: "Virtual IP address assigned to the hostname if IP based SSL is enabled",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VirtualIP"),
					},
					{
						Name:        "thumbprint",
						Description: "SSL certificate thumbprint",
						Type:        schema.TypeString,
					},
					{
						Name:        "to_update",
						Description: "Set to <code>true</code> to update existing hostname",
						Type:        schema.TypeBool,
					},
					{
						Name:        "host_type",
						Description: "Indicates whether the hostname is a standard or repository hostname Possible values include: 'HostTypeStandard', 'HostTypeRepository'",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_web_app_connection_strings",
				Description: "ConnStringInfo database connection string information",
				Resolver:    fetchWebAppConnectionStrings,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"app_cq_id", "name"}},
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "ID of web app",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "name",
						Description: "Name of connection string",
						Type:        schema.TypeString,
					},
					{
						Name:        "connection_string",
						Description: "Connection string value",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Type of database Possible values include: 'MySQL', 'SQLServer', 'SQLAzure', 'Custom', 'NotificationHub', 'ServiceBus', 'EventHub', 'APIHub', 'DocDb', 'RedisCache', 'PostgreSQL'",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:         "azure_web_app_handler_mappings",
				Description:  "HandlerMapping the IIS handler mappings used to define which handler processes HTTP requests with certain extension For example, it is used to configure php-cgiexe process to handle all HTTP requests with *php extension",
				Resolver:     fetchWebAppHandlerMappings,
				AlwaysDelete: true,
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "ID of web app",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "extension",
						Description: "Requests with this extension will be handled using the specified FastCGI application",
						Type:        schema.TypeString,
					},
					{
						Name:        "script_processor",
						Description: "The absolute path to the FastCGI application",
						Type:        schema.TypeString,
					},
					{
						Name:        "arguments",
						Description: "Command-line arguments to be passed to the script processor",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:         "azure_web_app_virtual_applications",
				Description:  "VirtualApplication virtual application in an app",
				Resolver:     fetchWebAppVirtualApplications,
				AlwaysDelete: true,
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "ID of web app",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "virtual_path",
						Description: "Virtual path",
						Type:        schema.TypeString,
					},
					{
						Name:        "physical_path",
						Description: "Physical path",
						Type:        schema.TypeString,
					},
					{
						Name:        "preload_enabled",
						Description: "otherwise, <code>false</code>",
						Type:        schema.TypeBool,
					},
				},
				Relations: []*schema.Table{
					{
						Name:         "azure_web_app_virtual_application_virtual_directories",
						Description:  "VirtualDirectory directory for virtual application",
						Resolver:     fetchWebAppVirtualApplicationVirtualDirectories,
						AlwaysDelete: true,
						Columns: []schema.Column{
							{
								Name:        "app_virtual_application_cq_id",
								Description: "Unique ID of azure_web_apps table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "virtual_path",
								Description: "Path to virtual application",
								Type:        schema.TypeString,
							},
							{
								Name:        "physical_path",
								Description: "Physical path",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "azure_web_app_experiments_ramp_up_rules",
				Description: "RampUpRule routing rules for ramp up testing This rule allows to redirect static traffic % to a slot or to gradually change routing % based on performance",
				Resolver:    fetchWebAppExperimentsRampUpRules,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"app_cq_id", "name"}},
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "ID of web app",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "action_host_name",
						Description: "Hostname of a slot to which the traffic will be redirected if decided to Eg myapp-stageazurewebsitesnet",
						Type:        schema.TypeString,
					},
					{
						Name:        "reroute_percentage",
						Description: "Percentage of the traffic which will be redirected to <code>ActionHostName</code>",
						Type:        schema.TypeFloat,
					},
					{
						Name:        "change_step",
						Description: "In auto ramp up scenario this is the step to add/remove from <code>ReroutePercentage</code> until it reaches \\n<code>MinReroutePercentage</code> or <code>MaxReroutePercentage</code> Site metrics are checked every N minutes specified in <code>ChangeIntervalInMinutes</code>\\nCustom decision algorithm can be provided in TiPCallback site extension which URL can be specified in <code>ChangeDecisionCallbackUrl</code>",
						Type:        schema.TypeFloat,
					},
					{
						Name:        "change_interval_in_minutes",
						Description: "Specifies interval in minutes to reevaluate ReroutePercentage",
						Type:        schema.TypeInt,
					},
					{
						Name:        "min_reroute_percentage",
						Description: "Specifies lower boundary above which ReroutePercentage will stay",
						Type:        schema.TypeFloat,
					},
					{
						Name:        "max_reroute_percentage",
						Description: "Specifies upper boundary below which ReroutePercentage will stay",
						Type:        schema.TypeFloat,
					},
					{
						Name:        "change_decision_callback_url",
						Description: "Custom decision algorithm can be provided in TiPCallback site extension which URL can be specified See TiPCallback site extension for the scaffold and contracts https://wwwsiteextensionsnet/packages/TiPCallback/",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ChangeDecisionCallbackURL"),
					},
					{
						Name:        "name",
						Description: "Name of the routing rule The recommended name would be to point to the slot which will receive the traffic in the experiment",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_web_app_auto_heal_rules_triggers_status_codes",
				Description: "StatusCodesBasedTrigger trigger based on status code",
				Resolver:    fetchWebAppAutoHealRulesTriggersStatusCodes,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"app_cq_id", "status"}},
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "ID of web app",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "status",
						Description: "HTTP status code",
						Type:        schema.TypeInt,
					},
					{
						Name:        "sub_status",
						Description: "Request Sub Status",
						Type:        schema.TypeInt,
					},
					{
						Name:        "win32_status",
						Description: "Win32 error code",
						Type:        schema.TypeInt,
					},
					{
						Name:        "path",
						Description: "Request Path",
						Type:        schema.TypeString,
					},
					{
						Name:        "count",
						Description: "Request Count",
						Type:        schema.TypeInt,
					},
					{
						Name:        "time_interval",
						Description: "Time interval",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:         "azure_web_app_auto_heal_rules_triggers_status_codes_ranges",
				Description:  "StatusCodesRangeBasedTrigger trigger based on range of status codes",
				Resolver:     fetchWebAppAutoHealRulesTriggersStatusCodesRanges,
				AlwaysDelete: true,
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "ID of web app",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "status_codes",
						Description: "HTTP status code",
						Type:        schema.TypeString,
					},
					{
						Name: "path",
						Type: schema.TypeString,
					},
					{
						Name:        "count",
						Description: "Request Count",
						Type:        schema.TypeInt,
					},
					{
						Name:        "time_interval",
						Description: "Time interval",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_web_app_auto_heal_rules_triggers_slow_requests_with_paths",
				Description: "SlowRequestsBasedTrigger trigger based on request execution time",
				Resolver:    fetchWebAppAutoHealRulesTriggersSlowRequestsWithPaths,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"app_cq_id", "path"}},
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "ID of web app",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "time_taken",
						Description: "Time taken",
						Type:        schema.TypeString,
					},
					{
						Name:        "path",
						Description: "Request Path",
						Type:        schema.TypeString,
					},
					{
						Name:        "count",
						Description: "Request Count",
						Type:        schema.TypeInt,
					},
					{
						Name:        "time_interval",
						Description: "Time interval",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "azure_web_app_ip_security_restrictions",
				Description: "IPSecurityRestriction IP security restriction on an app",
				Resolver:    fetchWebAppIpSecurityRestrictions,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"app_cq_id", "name"}},
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "ID of web app",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "ip_address",
						Description: "IP address the security restriction is valid for It can be in form of pure ipv4 address (required SubnetMask property) or CIDR notation such as ipv4/mask (leading bit match) For CIDR, SubnetMask property must not be specified",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IPAddress"),
					},
					{
						Name:        "subnet_mask",
						Description: "Subnet mask for the range of IP addresses the restriction is valid for",
						Type:        schema.TypeString,
					},
					{
						Name:        "vnet_subnet_resource_id",
						Description: "Virtual network resource id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VnetSubnetResourceID"),
					},
					{
						Name:        "vnet_traffic_tag",
						Description: "(internal) Vnet traffic tag",
						Type:        schema.TypeInt,
					},
					{
						Name:        "subnet_traffic_tag",
						Description: "(internal) Subnet traffic tag",
						Type:        schema.TypeInt,
					},
					{
						Name:        "action",
						Description: "Allow or Deny access for this IP range",
						Type:        schema.TypeString,
					},
					{
						Name:        "tag",
						Description: "Defines what this IP filter will be used for This is to support IP filtering on proxies Possible values include: 'Default', 'XffProxy', 'ServiceTag'",
						Type:        schema.TypeString,
					},
					{
						Name:        "priority",
						Description: "Priority of IP restriction rule",
						Type:        schema.TypeInt,
					},
					{
						Name:        "name",
						Description: "IP restriction rule name",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "IP restriction rule description",
						Type:        schema.TypeString,
					},
					{
						Name:        "headers",
						Description: "IP restriction rule headers X-Forwarded-Host (https://developermozillaorg/en-US/docs/Web/HTTP/Headers/X-Forwarded-Host#Examples) The matching logic is  - If the property is null or empty (default), all hosts(or lack of) are allowed - A value is compared using ordinal-ignore-case (excluding port number) - Subdomain wildcards are permitted but don't match the root domain For example, *contosocom matches the subdomain foocontosocom  but not the root domain contosocom or multi-level foobarcontosocom - Unicode host names are allowed but are converted to Punycode for matching X-Forwarded-For (https://developermozillaorg/en-US/docs/Web/HTTP/Headers/X-Forwarded-For#Examples) The matching logic is  - If the property is null or empty (default), any forwarded-for chains (or lack of) are allowed - If any address (excluding port number) in the chain (comma separated) matches the CIDR defined by the property X-Azure-FDID and X-FD-HealthProbe The matching logic is exact match",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:        "azure_web_app_scm_ip_security_restrictions",
				Description: "IPSecurityRestriction IP security restriction on an app",
				Resolver:    fetchWebAppScmIpSecurityRestrictions,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"app_cq_id", "name"}},
				Columns: []schema.Column{
					{
						Name:        "app_cq_id",
						Description: "Unique ID of azure_web_apps table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "app_id",
						Description: "ID of web app",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("id"),
					},
					{
						Name:        "ip_address",
						Description: "IP address the security restriction is valid for It can be in form of pure ipv4 address (required SubnetMask property) or CIDR notation such as ipv4/mask (leading bit match) For CIDR, SubnetMask property must not be specified",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("IPAddress"),
					},
					{
						Name:        "subnet_mask",
						Description: "Subnet mask for the range of IP addresses the restriction is valid for",
						Type:        schema.TypeString,
					},
					{
						Name:        "vnet_subnet_resource_id",
						Description: "Virtual network resource id",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("VnetSubnetResourceID"),
					},
					{
						Name:        "vnet_traffic_tag",
						Description: "(internal) Vnet traffic tag",
						Type:        schema.TypeInt,
					},
					{
						Name:        "subnet_traffic_tag",
						Description: "(internal) Subnet traffic tag",
						Type:        schema.TypeInt,
					},
					{
						Name:        "action",
						Description: "Allow or Deny access for this IP range",
						Type:        schema.TypeString,
					},
					{
						Name:        "tag",
						Description: "Defines what this IP filter will be used for This is to support IP filtering on proxies Possible values include: 'Default', 'XffProxy', 'ServiceTag'",
						Type:        schema.TypeString,
					},
					{
						Name:        "priority",
						Description: "Priority of IP restriction rule",
						Type:        schema.TypeInt,
					},
					{
						Name:        "name",
						Description: "IP restriction rule name",
						Type:        schema.TypeString,
					},
					{
						Name:        "description",
						Description: "IP restriction rule description",
						Type:        schema.TypeString,
					},
					{
						Name:        "headers",
						Description: "IP restriction rule headers X-Forwarded-Host (https://developermozillaorg/en-US/docs/Web/HTTP/Headers/X-Forwarded-Host#Examples) The matching logic is  - If the property is null or empty (default), all hosts(or lack of) are allowed - A value is compared using ordinal-ignore-case (excluding port number) - Subdomain wildcards are permitted but don't match the root domain For example, *contosocom matches the subdomain foocontosocom  but not the root domain contosocom or multi-level foobarcontosocom - Unicode host names are allowed but are converted to Punycode for matching X-Forwarded-For (https://developermozillaorg/en-US/docs/Web/HTTP/Headers/X-Forwarded-For#Examples) The matching logic is  - If the property is null or empty (default), any forwarded-for chains (or lack of) are allowed - If any address (excluding port number) in the chain (comma separated) matches the CIDR defined by the property X-Azure-FDID and X-FD-HealthProbe The matching logic is exact match",
						Type:        schema.TypeJSON,
					},
				},
			},
			WebAppAuthSettings(),
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchWebApps(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	svc := meta.(*client.Client).Services().Web.Apps
	response, err := svc.List(ctx)
	if err != nil {
		return err
	}
	for response.NotDone() {
		res <- response.Values()
		if err := response.NextWithContext(ctx); err != nil {
			return err
		}
	}
	return nil
}
func fetchWebAppPublishingProfiles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	svc := meta.(*client.Client).Services().Web.Apps
	response, err := svc.ListPublishingProfileXMLWithSecrets(ctx, *p.ResourceGroup, *p.Name, web.CsmPublishingProfileOptions{})
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	var profiles []PublishingProfile
	xml.Unmarshal(buf.Bytes(), &profiles)
	if err != nil {
		return err
	}

	res <- profiles
	return nil
}
func resolveWebAppAppSettings(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p, ok := resource.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", resource.Item)
	}

	if p.SiteProperties == nil ||
		p.SiteProperties.SiteConfig == nil ||
		p.SiteProperties.SiteConfig.AppSettings == nil {
		return nil
	}

	json := make(map[string]string)
	for _, kp := range *p.SiteProperties.SiteConfig.AppSettings {
		json[*kp.Name] = *kp.Value
	}

	return resource.Set(c.Name, json)
}
func fetchWebAppHostNameSslStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	if p.HostNameSslStates == nil {
		return nil
	}

	res <- *p.HostNameSslStates
	return nil
}
func fetchWebAppConnectionStrings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	if p.SiteProperties == nil ||
		p.SiteProperties.SiteConfig == nil ||
		p.SiteProperties.SiteConfig.ConnectionStrings == nil {
		return nil
	}

	res <- *p.SiteProperties.SiteConfig.ConnectionStrings
	return nil
}
func fetchWebAppHandlerMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	if p.SiteProperties == nil ||
		p.SiteProperties.SiteConfig == nil ||
		p.SiteProperties.SiteConfig.HandlerMappings == nil {
		return nil
	}

	res <- *p.SiteProperties.SiteConfig.HandlerMappings
	return nil
}
func fetchWebAppVirtualApplications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	if p.SiteProperties == nil ||
		p.SiteProperties.SiteConfig == nil ||
		p.SiteProperties.SiteConfig.VirtualApplications == nil {
		return nil
	}

	res <- *p.SiteProperties.SiteConfig.VirtualApplications
	return nil
}
func fetchWebAppVirtualApplicationVirtualDirectories(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.VirtualApplication)
	if !ok {
		return fmt.Errorf("expected web.VirtualApplication but got %T", parent.Item)
	}

	if p.VirtualDirectories == nil {
		return nil
	}

	res <- *p.VirtualDirectories
	return nil
}
func fetchWebAppExperimentsRampUpRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	if p.SiteProperties == nil ||
		p.SiteProperties.SiteConfig == nil ||
		p.SiteProperties.SiteConfig.Experiments == nil ||
		p.SiteProperties.SiteConfig.Experiments.RampUpRules == nil {
		return nil
	}

	res <- *p.SiteProperties.SiteConfig.Experiments.RampUpRules
	return nil
}
func fetchWebAppAutoHealRulesTriggersStatusCodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	if p.SiteProperties == nil ||
		p.SiteProperties.SiteConfig == nil ||
		p.SiteProperties.SiteConfig.AutoHealRules == nil ||
		p.SiteProperties.SiteConfig.AutoHealRules.Triggers == nil ||
		p.SiteProperties.SiteConfig.AutoHealRules.Triggers.StatusCodes == nil {
		return nil
	}

	res <- *p.SiteProperties.SiteConfig.AutoHealRules.Triggers.StatusCodes
	return nil
}
func fetchWebAppAutoHealRulesTriggersStatusCodesRanges(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	if p.SiteProperties == nil ||
		p.SiteProperties.SiteConfig == nil ||
		p.SiteProperties.SiteConfig.AutoHealRules == nil ||
		p.SiteProperties.SiteConfig.AutoHealRules.Triggers == nil ||
		p.SiteProperties.SiteConfig.AutoHealRules.Triggers.StatusCodesRange == nil {
		return nil
	}

	res <- *p.SiteProperties.SiteConfig.AutoHealRules.Triggers.StatusCodesRange
	return nil
}
func fetchWebAppAutoHealRulesTriggersSlowRequestsWithPaths(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	if p.SiteProperties == nil ||
		p.SiteProperties.SiteConfig == nil ||
		p.SiteProperties.SiteConfig.AutoHealRules == nil ||
		p.SiteProperties.SiteConfig.AutoHealRules.Triggers == nil ||
		p.SiteProperties.SiteConfig.AutoHealRules.Triggers.SlowRequestsWithPath == nil {
		return nil
	}

	res <- *p.SiteProperties.SiteConfig.AutoHealRules.Triggers.SlowRequestsWithPath
	return nil
}
func fetchWebAppIpSecurityRestrictions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	if p.SiteProperties == nil ||
		p.SiteProperties.SiteConfig == nil ||
		p.SiteProperties.SiteConfig.IPSecurityRestrictions == nil {
		return nil
	}

	res <- *p.SiteProperties.SiteConfig.IPSecurityRestrictions
	return nil
}
func fetchWebAppScmIpSecurityRestrictions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p, ok := parent.Item.(web.Site)
	if !ok {
		return fmt.Errorf("expected web.Site but got %T", parent.Item)
	}

	if p.SiteProperties == nil ||
		p.SiteProperties.SiteConfig == nil ||
		p.SiteProperties.SiteConfig.ScmIPSecurityRestrictions == nil {
		return nil
	}

	res <- *p.SiteProperties.SiteConfig.ScmIPSecurityRestrictions
	return nil
}

type PublishingProfile struct {
	PublishUrl string `xml:"publishUrl"`
	UserName   string `xml:"userName"`
	UserPWD    string `xml:"userPWD"`
}
