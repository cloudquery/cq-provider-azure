package frontdoor

import (
	"context"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource front_door --config gen.hcl --output .
func FrontDoors() *schema.Table {
	return &schema.Table{
		Name:         "azure_front_door",
		Description:  "Front Door represents a collection of backend endpoints to route traffic to along with rules that specify how traffic is sent there.",
		Resolver:     fetchFrontDoors,
		Multiplex:    client.SubscriptionMultiplex,
		DeleteFilter: client.DeleteSubscriptionFilter,
		Columns: []schema.Column{
			{
				Name:        "subscription_id",
				Description: "Azure subscription ID",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAzureSubscription,
			},
			{
				Name:        "resource_state",
				Description: "Resource status of the Front Door",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.ResourceState"),
			},
			{
				Name:        "provisioning_state",
				Description: "Provisioning state of the Front Door",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.ProvisioningState"),
			},
			{
				Name:        "cname",
				Description: "The host that each frontend endpoint must CNAME to",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.Cname"),
			},
			{
				Name:        "frontdoor_id",
				Description: "The ID of the Front Door",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.FrontdoorID"),
			},
			{
				Name:        "friendly_name",
				Description: "A friendly name for the Front Door",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.FriendlyName"),
			},
			{
				Name:        "enforce_certificate_name_check",
				Description: "Whether to enforce certificate name check on HTTPS requests to all backend pools",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.BackendPoolsSettings.EnforceCertificateNameCheck"),
			},
			{
				Name:        "send_recv_timeout_seconds",
				Description: "Send and receive timeout on forwarding request to the backend",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Properties.BackendPoolsSettings.SendRecvTimeoutSeconds"),
			},
			{
				Name:        "enabled_state",
				Description: "Operational status of the Front Door load balancer",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Properties.EnabledState"),
			},
			{
				Name:        "id",
				Description: "Resource ID",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ID"),
			},
			{
				Name:        "name",
				Description: "Resource name",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "Resource type",
				Type:        schema.TypeString,
			},
			{
				Name:        "location",
				Description: "Resource location",
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
				Name:        "azure_front_door_rules_engines",
				Description: "Rules engine configuration containing a list of rules that will run to modify the runtime behavior of the request and response.",
				Resolver:    fetchFrontDoorFrontDoorRulesEngines,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_door table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "resource_state",
						Description: "Resource status",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RulesEngineProperties.ResourceState"),
					},
					{
						Name:        "name",
						Description: "Resource name",
						Type:        schema.TypeString,
					},
					{
						Name:        "type",
						Description: "Resource type",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "Resource ID",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ID"),
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "azure_front_door_rules_engine_rules",
						Description: "A list of rules that define a particular Rules Engine Configuration.",
						Resolver:    fetchFrontDoorFrontDoorRulesEngineRules,
						Columns: []schema.Column{
							{
								Name:        "front_door_rules_engine_cq_id",
								Description: "Unique CloudQuery ID of azure_front_door_rules_engines table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "A name to refer to this specific rule",
								Type:        schema.TypeString,
							},
							{
								Name:        "priority",
								Description: "A priority assigned to this rule",
								Type:        schema.TypeInt,
							},
							{
								Name:        "route_configuration_override",
								Description: "Override the route configuration",
								Type:        schema.TypeJSON,
								Resolver:    resolveFrontDoorRulesEngineRulesRouteConfigurationOverride,
							},
							{
								Name:        "match_processing_behavior",
								Description: "If this rule is a match should the rules engine continue running the remaining rules or stop",
								Type:        schema.TypeString,
							},
						},
						Relations: []*schema.Table{
							{
								Name:        "azure_front_door_rules_engine_rule_request_header_actions",
								Description: "A list of header actions to apply from the request from AFD to the origin.",
								Resolver:    fetchFrontDoorFrontDoorRulesEngineRuleRequestHeaderActions,
								Columns: []schema.Column{
									{
										Name:        "front_door_rules_engine_rule_cq_id",
										Description: "Unique CloudQuery ID of azure_front_door_rules_engine_rules table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "action_type",
										Description: "Which type of manipulation to apply to the header",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("HeaderActionType"),
									},
									{
										Name:        "name",
										Description: "The name of the header this action will apply to",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("HeaderName"),
									},
									{
										Name:        "value",
										Description: "The value to update the given header name with",
										Type:        schema.TypeString,
									},
								},
							},
							{
								Name:        "azure_front_door_rules_engine_rule_response_header_actions",
								Description: "A list of header actions to apply from the response from AFD to the client.",
								Resolver:    fetchFrontDoorFrontDoorRulesEngineRuleResponseHeaderActions,
								Columns: []schema.Column{
									{
										Name:        "front_door_rules_engine_rule_cq_id",
										Description: "Unique CloudQuery ID of azure_front_door_rules_engine_rules table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "action_type",
										Description: "Which type of manipulation to apply to the header",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("HeaderActionType"),
									},
									{
										Name:        "name",
										Description: "The name of the header this action will apply to",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("HeaderName"),
									},
									{
										Name:        "value",
										Description: "The value to update the given header name with",
										Type:        schema.TypeString,
									},
								},
							},
							{
								Name:        "azure_front_door_rules_engine_rule_match_conditions",
								Description: "A list of header actions to apply from the response from AFD to the client.",
								Resolver:    fetchFrontDoorFrontDoorRulesEngineRuleMatchConditions,
								Columns: []schema.Column{
									{
										Name:        "front_door_rules_engine_rule_cq_id",
										Description: "Unique CloudQuery ID of azure_front_door_rules_engine_rules table (FK)",
										Type:        schema.TypeUUID,
										Resolver:    schema.ParentIdResolver,
									},
									{
										Name:        "match_variable",
										Description: "Match variable",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("RulesEngineMatchVariable"),
									},
									{
										Name:        "selector",
										Description: "Name of selector in request header or request body to be matched",
										Type:        schema.TypeString,
									},
									{
										Name:        "operator",
										Description: "Describes operator to apply to the match condition",
										Type:        schema.TypeString,
										Resolver:    schema.PathResolver("RulesEngineOperator"),
									},
									{
										Name:        "negate_condition",
										Description: "Describes if this is negate condition or not",
										Type:        schema.TypeBool,
									},
									{
										Name:        "match_value",
										Description: "Match values to match against",
										Type:        schema.TypeStringArray,
										Resolver:    schema.PathResolver("RulesEngineMatchValue"),
									},
									{
										Name:        "transforms",
										Description: "List of transforms",
										Type:        schema.TypeStringArray,
									},
								},
							},
						},
					},
				},
			},
			{
				Name:        "azure_front_door_routing_rules",
				Description: "RoutingRule a routing rule represents a specification for traffic to treat and where to send it, along with health probe information.",
				Resolver:    fetchFrontdoorFrontDoorRoutingRules,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_door table (FK)",
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
						Resolver:    resolveFrontDoorRoutingRulesRoutingRulePropertiesRouteConfiguration,
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
						Name:        "azure_front_door_routing_rule_routing_rule_properties_frontend_endpoints",
						Description: "SubResource reference to another subresource.",
						Resolver:    fetchFrontdoorFrontDoorRoutingRuleRoutingRulePropertiesFrontendEndpoints,
						Columns: []schema.Column{
							{
								Name:        "front_door_routing_rule_cq_id",
								Description: "Unique CloudQuery ID of azure_front_door_routing_rules table (FK)",
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
				Name:        "azure_front_door_load_balancing_settings",
				Description: "LoadBalancingSettingsModel load balancing settings for a backend pool",
				Resolver:    fetchFrontdoorFrontDoorLoadBalancingSettings,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_door table (FK)",
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
				Name:        "azure_front_door_health_probe_settings",
				Description: "HealthProbeSettingsModel load balancing settings for a backend pool",
				Resolver:    fetchFrontdoorFrontDoorHealthProbeSettings,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_door table (FK)",
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
				Name:        "azure_front_door_backend_pools",
				Description: "BackendPool a backend pool is a collection of backends that can be routed to.",
				Resolver:    fetchFrontdoorFrontDoorBackendPools,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_door table (FK)",
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
						Name:        "azure_front_door_backend_pool_backends",
						Description: "Backend backend address of a frontDoor load balancer.",
						Resolver:    fetchFrontdoorFrontDoorBackendPoolBackends,
						Columns: []schema.Column{
							{
								Name:        "front_door_backend_pool_cq_id",
								Description: "Unique CloudQuery ID of azure_front_door_backend_pools table (FK)",
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
				Name:        "azure_front_door_frontend_endpoints",
				Description: "FrontendEndpoint a frontend endpoint used for routing.",
				Resolver:    fetchFrontdoorFrontDoorFrontendEndpoints,
				Columns: []schema.Column{
					{
						Name:        "front_door_cq_id",
						Description: "Unique CloudQuery ID of azure_front_door table (FK)",
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

func fetchFrontDoors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontDoorFrontDoorRulesEngines(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontDoorFrontDoorRulesEngineRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func resolveFrontDoorRulesEngineRulesRouteConfigurationOverride(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func fetchFrontDoorFrontDoorRulesEngineRuleRequestHeaderActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontDoorFrontDoorRulesEngineRuleResponseHeaderActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontDoorFrontDoorRulesEngineRuleMatchConditions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontDoorRoutingRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func resolveFrontDoorRoutingRulesRoutingRulePropertiesRouteConfiguration(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	panic("not implemented")
}
func fetchFrontdoorFrontDoorRoutingRuleRoutingRulePropertiesFrontendEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontDoorLoadBalancingSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontDoorHealthProbeSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontDoorBackendPools(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontDoorBackendPoolBackends(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
func fetchFrontdoorFrontDoorFrontendEndpoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	panic("not implemented")
}
