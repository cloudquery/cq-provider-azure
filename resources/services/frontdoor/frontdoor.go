package frontdoor

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource frontdoor --config gen.hcl --output .
func Frontdoors() *schema.Table {
	return &schema.Table{
		Name:         "azure_frontdoors_frontdoor",
		Description:  "FrontDoor front Door represents a collection of backend endpoints to route traffic to along with rules that specify how traffic is sent there.",
		Resolver:     fetchFrontdoorsFrontdoors,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "properties_resource_state",
				Description: "ResourceState - Resource status of the Front Door",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.ResourceState"),
			},
			{
				Name:        "properties_provisioning_state",
				Description: "ProvisioningState - READ-ONLY; Provisioning state of the Front Door.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:        "properties_cname",
				Description: "Cname - READ-ONLY; The host that each frontendEndpoint must CNAME to.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Cname"),
			},
			{
				Name:        "properties_frontdoor_id",
				Description: "FrontdoorID - READ-ONLY; The Id of the frontdoor.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.FrontdoorID"),
			},
			{
				Name:        "properties_friendly_name",
				Description: "FriendlyName - A friendly name for the frontDoor",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.FriendlyName"),
			},
			{
				Name:        "backend_pools_settings_enforce_certificate_name_checks",
				Description: "EnforceCertificateNameCheck - Whether to enforce certificate name check on HTTPS requests to all backend pools",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.BackendPoolsSettings.EnforceCertificateNameCheck"),
			},
			{
				Name:        "properties_backend_pools_settings_send_recv_timeout_seconds",
				Description: "SendRecvTimeoutSeconds - Send and receive timeout on forwarding request to the backend",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Properties.BackendPoolsSettings.SendRecvTimeoutSeconds"),
			},
			{
				Name:        "properties_enabled_state",
				Description: "EnabledState - Operational status of the Front Door load balancer",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.EnabledState"),
			},
			{
				Name:        "id",
				Description: "ID - READ-ONLY; Resource ID.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Name - READ-ONLY; Resource name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Type - READ-ONLY; Resource type.",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Location - Resource location.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Tags - Resource tags.",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "azure_frontdoors_frontdoor_rules_engines",
				Description: "RulesEngine a rules engine configuration containing a list of rules that will run to modify the runtime behavior of the request and response.",
				Resolver:    fetchFrontdoorFrontdoorRulesEngines,
				Columns: []schema.Column{
					{
						Name:        "frontdoor_cq_id",
						Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "rules_engine_properties_resource_state",
						Description: "ResourceState - Resource status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RulesEngineProperties.ResourceState"),
					},
					{
						Name:        "name",
						Description: "Name - READ-ONLY; Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Type - READ-ONLY; Resource type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "ID - READ-ONLY; Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_frontdoors_frontdoor_rules_engine_rules",
						Description: "RulesEngineRule contains a list of match conditions, and an action on how to modify the request/response",
						Resolver:    fetchFrontdoorFrontdoorRulesEngineRules,
						Columns: []schema.Column{
							{
								Name:        "frontdoor_rules_engine_cq_id",
								Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor_rules_engines table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "Name - A name to refer to this specific rule.",
								Type:        schema.TypeString,
							},
							{
								Name:        "priority",
								Description: "Priority - A priority assigned to this rule.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "action_route_configuration_override",
								Description: "RouteConfigurationOverride - Override the route configuration.",
								Type:        schema.TypeJSON,
								Resolver:    resolveFrontdoorRulesEngineRulesActionRouteConfigurationOverride,
							},
							{
								Name:        "match_processing_behavior",
								Description: "MatchProcessingBehavior - If this rule is a match should the rules engine continue running the remaining rules or stop",
								Type:        schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "azure_frontdoors_frontdoor_rules_engine_rule_action_request_header_actions",
								Description: "HeaderAction an action that can manipulate an http header.",
								Resolver:    fetchFrontdoorFrontdoorRulesEngineRuleActionRequestHeaderActions,
								Columns: []schema.Column{
									{
										Name:        "frontdoor_rules_engine_rule_cq_id",
										Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor_rules_engine_rules table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "header_action_type",
										Description: "HeaderActionType - Which type of manipulation to apply to the header",
										Type:        schema.TypeString,
									},
									{
										Name:        "header_name",
										Description: "HeaderName - The name of the header this action will apply to.",
										Type:        schema.TypeString,
									},
									{
										Name:        "value",
										Description: "Value - The value to update the given header name with",
										Type:        schema.TypeString,
									},
								},
							},
							{
								Name:        "azure_frontdoors_frontdoor_rules_engine_rule_action_response_header_actions",
								Description: "HeaderAction an action that can manipulate an http header.",
								Resolver:    fetchFrontdoorFrontdoorRulesEngineRuleActionResponseHeaderActions,
								Columns: []schema.Column{
									{
										Name:        "frontdoor_rules_engine_rule_cq_id",
										Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor_rules_engine_rules table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "header_action_type",
										Description: "HeaderActionType - Which type of manipulation to apply to the header",
										Type:        schema.TypeString,
									},
									{
										Name:        "header_name",
										Description: "HeaderName - The name of the header this action will apply to.",
										Type:        schema.TypeString,
									},
									{
										Name:        "value",
										Description: "Value - The value to update the given header name with",
										Type:        schema.TypeString,
									},
								},
							},
							{
								Name:        "azure_frontdoors_frontdoor_rules_engine_rule_match_conditions",
								Description: "RulesEngineMatchCondition define a match condition",
								Resolver:    fetchFrontdoorFrontdoorRulesEngineRuleMatchConditions,
								Columns: []schema.Column{
									{
										Name:        "frontdoor_rules_engine_rule_cq_id",
										Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor_rules_engine_rules table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "rules_engine_match_variable",
										Description: "RulesEngineMatchVariable - Match Variable",
										Type:        schema.TypeString,
									},
									{
										Name:        "selector",
										Description: "Selector - Name of selector in RequestHeader or RequestBody to be matched",
										Type:        schema.TypeString,
									},
									{
										Name:        "rules_engine_operator",
										Description: "RulesEngineOperator - Describes operator to apply to the match condition",
										Type:        schema.TypeString,
									},
									{
										Name:        "negate_condition",
										Description: "NegateCondition - Describes if this is negate condition or not",
										Type:        schema.TypeBool,
									},
									{
										Name:        "rules_engine_match_value",
										Description: "RulesEngineMatchValue - Match values to match against",
										Type:        schema.TypeStringArray,
									},
									{
										Name:        "transforms",
										Description: "Transforms - List of transforms",
										Type:        schema.TypeStringArray,
									},
								},
							},
						},
					},
				},
			},
			{
				Name:        "azure_frontdoors_frontdoor_routing_rules",
				Description: "RoutingRule a routing rule represents a specification for traffic to treat and where to send it, along with health probe information.",
				Resolver:    fetchFrontdoorFrontdoorRoutingRules,
				Columns: []schema.Column{
					{
						Name:        "frontdoor_cq_id",
						Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "routing_rule_properties_resource_state",
						Description: "ResourceState - Resource status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoutingRuleProperties.ResourceState"),
					},
					{
						Name:        "routing_rule_properties_accepted_protocols",
						Description: "AcceptedProtocols - Protocol schemes to match for this rule",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("RoutingRuleProperties.AcceptedProtocols"),
					},
					{
						Name:        "routing_rule_properties_patterns_to_match",
						Description: "PatternsToMatch - The route patterns of the rule.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("RoutingRuleProperties.PatternsToMatch"),
					},
					{
						Name:        "routing_rule_properties_enabled_state",
						Description: "EnabledState - Whether to enable use of this rule",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoutingRuleProperties.EnabledState"),
					},
					{
						Name:        "routing_rule_properties_route_configuration",
						Description: "RouteConfiguration - A reference to the routing configuration.",
						Type:        schema.TypeJSON,
						Resolver:    resolveFrontdoorRoutingRulesRoutingRulePropertiesRouteConfiguration,
					},
					{
						Name:        "routing_rule_properties_rules_engine_id",
						Description: "ID - Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoutingRuleProperties.RulesEngine.ID"),
					},
					{
						Name:        "routing_rule_properties_web_application_firewall_policy_link_id",
						Description: "ID - Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoutingRuleProperties.WebApplicationFirewallPolicyLink.ID"),
					},
					{
						Name:        "name",
						Description: "Name - Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Type - READ-ONLY; Resource type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "ID - Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_frontdoors_frontdoor_routing_rule_routing_rule_properties_frontend_endpoints",
						Description: "SubResource reference to another subresource.",
						Resolver:    fetchFrontdoorFrontdoorRoutingRuleRoutingRulePropertiesFrontendEndpoints,
						Columns: []schema.Column{
							{
								Name:        "frontdoor_routing_rule_cq_id",
								Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor_routing_rules table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "id",
								Description: "ID - Resource ID.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("ID"),
							},
						},
					},
				},
			},
			{
				Name:        "azure_frontdoors_frontdoor_load_balancing_settings",
				Description: "LoadBalancingSettingsModel load balancing settings for a backend pool",
				Resolver:    fetchFrontdoorFrontdoorLoadBalancingSettings,
				Columns: []schema.Column{
					{
						Name:        "frontdoor_cq_id",
						Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "load_balancing_settings_properties_resource_state",
						Description: "ResourceState - Resource status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LoadBalancingSettingsProperties.ResourceState"),
					},
					{
						Name:        "load_balancing_settings_properties_sample_size",
						Description: "SampleSize - The number of samples to consider for load balancing decisions",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LoadBalancingSettingsProperties.SampleSize"),
					},
					{
						Name:        "load_balancing_settings_properties_successful_samples_required",
						Description: "SuccessfulSamplesRequired - The number of samples within the sample period that must succeed",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LoadBalancingSettingsProperties.SuccessfulSamplesRequired"),
					},
					{
						Name:        "additional_latency_milliseconds",
						Description: "AdditionalLatencyMilliseconds - The additional latency in milliseconds for probes to fall into the lowest latency bucket",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("LoadBalancingSettingsProperties.AdditionalLatencyMilliseconds"),
					},
					{
						Name:        "name",
						Description: "Name - Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Type - READ-ONLY; Resource type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "ID - Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:        "azure_frontdoors_frontdoor_health_probe_settings",
				Description: "HealthProbeSettingsModel load balancing settings for a backend pool",
				Resolver:    fetchFrontdoorFrontdoorHealthProbeSettings,
				Columns: []schema.Column{
					{
						Name:        "frontdoor_cq_id",
						Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "health_probe_settings_properties_resource_state",
						Description: "ResourceState - Resource status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.ResourceState"),
					},
					{
						Name:        "health_probe_settings_properties_path",
						Description: "Path - The path to use for the health probe",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.Path"),
					},
					{
						Name:        "health_probe_settings_properties_protocol",
						Description: "Protocol - Protocol scheme to use for this probe",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.Protocol"),
					},
					{
						Name:        "health_probe_settings_properties_interval_in_seconds",
						Description: "IntervalInSeconds - The number of seconds between health probes.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.IntervalInSeconds"),
					},
					{
						Name:        "health_probe_settings_properties_health_probe_method",
						Description: "HealthProbeMethod - Configures which HTTP method to use to probe the backends defined under backendPools",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.HealthProbeMethod"),
					},
					{
						Name:        "health_probe_settings_properties_enabled_state",
						Description: "EnabledState - Whether to enable health probes to be made against backends defined under backendPools",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("HealthProbeSettingsProperties.EnabledState"),
					},
					{
						Name:        "name",
						Description: "Name - Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Type - READ-ONLY; Resource type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "ID - Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
			{
				Name:        "azure_frontdoors_frontdoor_backend_pools",
				Description: "BackendPool a backend pool is a collection of backends that can be routed to.",
				Resolver:    fetchFrontdoorFrontdoorBackendPools,
				Columns: []schema.Column{
					{
						Name:        "frontdoor_cq_id",
						Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "backend_pool_properties_resource_state",
						Description: "ResourceState - Resource status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("BackendPoolProperties.ResourceState"),
					},
					{
						Name:        "backend_pool_properties_load_balancing_settings_id",
						Description: "ID - Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("BackendPoolProperties.LoadBalancingSettings.ID"),
					},
					{
						Name:        "backend_pool_properties_health_probe_settings_id",
						Description: "ID - Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("BackendPoolProperties.HealthProbeSettings.ID"),
					},
					{
						Name:        "name",
						Description: "Name - Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Type - READ-ONLY; Resource type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "ID - Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_frontdoors_frontdoor_backend_pool_backends",
						Description: "Backend backend address of a frontDoor load balancer.",
						Resolver:    fetchFrontdoorFrontdoorBackendPoolBackends,
						Columns: []schema.Column{
							{
								Name:        "frontdoor_backend_pool_cq_id",
								Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor_backend_pools table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "address",
								Description: "Address - Location of the backend (IP address or FQDN)",
								Type:        schema.TypeString,
							},
							{
								Name:        "private_link_alias",
								Description: "PrivateLinkAlias - The Alias of the Private Link resource",
								Type:        schema.TypeString,
							},
							{
								Name:        "private_link_resource_id",
								Description: "PrivateLinkResourceID - The Resource Id of the Private Link resource",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PrivateLinkResourceID"),
							},
							{
								Name:        "private_link_location",
								Description: "PrivateLinkLocation - The location of the Private Link resource",
								Type:        schema.TypeString,
							},
							{
								Name:        "private_endpoint_status",
								Description: "PrivateEndpointStatus - READ-ONLY; The Approval status for the connection to the Private Link",
								Type:        schema.TypeString,
							},
							{
								Name:        "private_link_approval_message",
								Description: "PrivateLinkApprovalMessage - A custom message to be included in the approval request to connect to the Private Link",
								Type:        schema.TypeString,
							},
							{
								Name:        "http_port",
								Description: "HTTPPort - The HTTP TCP port number",
								Type:        schema.TypeInt,
								Resolver:    schema.PathResolver("HTTPPort"),
							},
							{
								Name:        "https_port",
								Description: "HTTPSPort - The HTTPS TCP port number",
								Type:        schema.TypeInt,
								Resolver:    schema.PathResolver("HTTPSPort"),
							},
							{
								Name:        "enabled_state",
								Description: "EnabledState - Whether to enable use of this backend",
								Type:        schema.TypeString,
							},
							{
								Name:        "priority",
								Description: "Priority - Priority to use for load balancing",
								Type:        schema.TypeInt,
							},
							{
								Name:        "weight",
								Description: "Weight - Weight of this endpoint for load balancing purposes.",
								Type:        schema.TypeInt,
							},
							{
								Name:        "backend_host_header",
								Description: "BackendHostHeader - The value to use as the host header sent to the backend",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "azure_frontdoors_frontdoor_frontend_endpoints",
				Description: "FrontendEndpoint a frontend endpoint used for routing.",
				Resolver:    fetchFrontdoorFrontdoorFrontendEndpoints,
				Columns: []schema.Column{
					{
						Name:        "frontdoor_cq_id",
						Description: "Unique CloudQuery ID of azure_frontdoors_frontdoor table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "frontend_endpoint_properties_resource_state",
						Description: "ResourceState - Resource status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.ResourceState"),
					},
					{
						Name:        "frontend_endpoint_properties_custom_https_provisioning_state",
						Description: "CustomHTTPSProvisioningState - READ-ONLY; Provisioning status of Custom Https of the frontendEndpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSProvisioningState"),
					},
					{
						Name:        "frontend_endpoint_properties_custom_https_provisioning_substate",
						Description: "CustomHTTPSProvisioningSubstate - READ-ONLY; Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSProvisioningSubstate"),
					},
					{
						Name:        "certificate_source",
						Description: "CertificateSource - Defines the source of the SSL certificate",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.CertificateSource"),
					},
					{
						Name:        "protocol_type",
						Description: "ProtocolType - Defines the TLS extension protocol that is used for secure delivery",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.ProtocolType"),
					},
					{
						Name:        "minimum_tls_version",
						Description: "MinimumTLSVersion - The minimum TLS version required from the clients to establish an SSL handshake with Front Door",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.MinimumTLSVersion"),
					},
					{
						Name:        "key_vault_certificate_source_vault_id",
						Description: "ID - Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.KeyVaultCertificateSourceParameters.Vault.ID"),
					},
					{
						Name:        "key_vault_certificate_source_secret_name",
						Description: "SecretName - The name of the Key Vault secret representing the full certificate PFX",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.KeyVaultCertificateSourceParameters.SecretName"),
					},
					{
						Name:        "key_vault_certificate_source_secret_version",
						Description: "SecretVersion - The version of the Key Vault secret representing the full certificate PFX",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.KeyVaultCertificateSourceParameters.SecretVersion"),
					},
					{
						Name:        "certificate_type",
						Description: "CertificateType - Defines the type of the certificate used for secure connections to a frontendEndpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.CustomHTTPSConfiguration.CertificateSourceParameters.CertificateType"),
					},
					{
						Name:        "frontend_endpoint_properties_host_name",
						Description: "HostName - The host name of the frontendEndpoint",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.HostName"),
					},
					{
						Name:        "frontend_endpoint_properties_session_affinity_enabled_state",
						Description: "SessionAffinityEnabledState - Whether to allow session affinity on this host",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.SessionAffinityEnabledState"),
					},
					{
						Name:        "frontend_endpoint_properties_session_affinity_ttl_seconds",
						Description: "SessionAffinityTTLSeconds - UNUSED",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.SessionAffinityTTLSeconds"),
					},
					{
						Name:        "frontend_endpoint_properties_id",
						Description: "ID - Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("FrontendEndpointProperties.WebApplicationFirewallPolicyLink.ID"),
					},
					{
						Name:        "name",
						Description: "Name - Resource name.",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Type - READ-ONLY; Resource type.",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "ID - Resource ID.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchFrontdoorsFrontdoors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorRulesEngines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorRulesEngineRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func resolveFrontdoorRulesEngineRulesActionRouteConfigurationOverride(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorRulesEngineRuleActionRequestHeaderActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorRulesEngineRuleActionResponseHeaderActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorRulesEngineRuleMatchConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorRoutingRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func resolveFrontdoorRoutingRulesRoutingRulePropertiesRouteConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorRoutingRuleRoutingRulePropertiesFrontendEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorLoadBalancingSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorHealthProbeSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorBackendPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorBackendPoolBackends(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontdoorFrontendEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
