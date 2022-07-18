service          = "azure"
output_directory = "."
add_generate     = true

resource "azure" "" "front_door" {
  path        = "github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor.FrontDoor"
  description = "Front Door represents a collection of backend endpoints to route traffic to along with rules that specify how traffic is sent there."

  userDefinedColumn "subscription_id" {
    type        = "string"
    description = "Azure subscription ID"
    resolver "resolveAzureSubscription" {
      path = "github.com/cloudquery/cq-provider-azure/client.ResolveAzureSubscription"
    }
  }

  multiplex "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.SubscriptionMultiplex"
  }

  deleteFilter "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.DeleteSubscriptionFilter"
  }

  column "id" {
    description = "Resource ID"
  }

  column "name" {
    description = "Resource name"
  }

  column "type" {
    description = "Resource type"
  }

  column "location" {
    description = "Resource location"
  }

  column "tags" {
    description = "Resource tags"
  }

  column "properties_resource_state" {
    rename      = "resource_state"
    description = "Resource status of the Front Door"
  }

  column "properties_provisioning_state" {
    rename      = "provisioning_state"
    description = "Provisioning state of the Front Door"
  }

  column "properties_cname" {
    rename      = "cname"
    description = "The host that each frontend endpoint must CNAME to"
  }

  column "properties_frontdoor_id" {
    rename      = "frontdoor_id"
    description = "The ID of the Front Door"
  }

  column "properties_friendly_name" {
    rename      = "friendly_name"
    description = "A friendly name for the Front Door"
  }

  column "properties_backend_pools_settings_enforce_certificate_name_check" {
    rename      = "enforce_certificate_name_check"
    description = "Whether to enforce certificate name check on HTTPS requests to all backend pools"
  }

  column "properties_backend_pools_settings_send_recv_timeout_seconds" {
    rename      = "send_recv_timeout_seconds"
    description = "Send and receive timeout on forwarding request to the backend"
  }

  column "properties_enabled_state" {
    rename      = "enabled_state"
    description = "Operational status of the Front Door load balancer"
  }

  relation "azure" "front_door" "properties_rules_engines" {
    rename      = "rules_engines"
    description = "Rules engine configuration containing a list of rules that will run to modify the runtime behavior of the request and response."

    column "rules_engine_properties_resource_state" {
      rename      = "resource_state"
      description = "Resource status"
    }

    column "id" {
      description = "Resource ID"
    }

    column "name" {
      description = "Resource name"
    }

    column "type" {
      description = "Resource type"
    }

    relation "azure" "front_door" "rules_engine_properties_rules" {
      rename      = "rules"
      description = "A list of rules that define a particular Rules Engine Configuration."

      column "name" {
        description = "A name to refer to the rule"
      }

      column "priority" {
        description = "A priority assigned to the rule"
      }

      column "match_processing_behavior" {
        description = "If the rule is a match should the rules engine continue running the remaining rules or stop"
      }

      column "action_route_configuration_override" {
        rename            = "route_configuration_override"
        description       = "Override the route configuration"
        type              = "json"
        generate_resolver = true
      }

      relation "azure" "front_door" "action_request_header_actions" {
        rename      = "request_header_actions"
        description = "A list of header actions to apply from the request from AFD to the origin."

        column "header_action_type" {
          rename      = "action_type"
          description = "Which type of manipulation to apply to the header"
        }

        column "header_name" {
          rename      = "name"
          description = "The name of the header the action will apply to"
        }

        column "value" {
          description = "The value to update the given header name with"
        }
      }

      relation "azure" "front_door" "action_response_header_actions" {
        rename      = "response_header_actions"
        description = "A list of header actions to apply from the response from AFD to the client."

        column "header_action_type" {
          rename      = "action_type"
          description = "Which type of manipulation to apply to the header"
        }

        column "header_name" {
          rename      = "name"
          description = "The name of the header the action will apply to"
        }

        column "value" {
          description = "The value to update the given header name with"
        }
      }

      relation "azure" "front_door" "match_conditions" {
        description = "A list of match conditions that must meet in order for the actions of the rule to run. Having no match conditions means the actions will always run."

        column "rules_engine_match_variable" {
          rename      = "match_variable"
          description = "Match variable"
        }

        column "selector" {
          description = "Name of selector in request header or request body to be matched"
        }

        column "rules_engine_operator" {
          rename      = "operator"
          description = "Describes operator to apply to the match condition"
        }

        column "negate_condition" {
          description = "Describes if this is negate condition or not"
        }

        column "rules_engine_match_value" {
          rename      = "match_value"
          description = "Match values to match against"
        }

        column "transforms" {
          description = "List of transforms"
        }
      }
    }
  }

  relation "azure" "frontdoor" "properties_routing_rules" {
    rename      = "routing_rules"
    description = "Routing rules represent specifications for traffic to treat and where to send it, along with health probe information."

    column "routing_rule_properties_resource_state" {
      rename      = "resource_state"
      description = "Resource status"
    }

    column "routing_rule_properties_accepted_protocols" {
      rename      = "accepted_protocols"
      description = "Protocol schemes to match for the rule"
    }

    column "routing_rule_properties_patterns_to_match" {
      rename      = "patterns_to_match"
      description = "The route patterns of the rule"
    }

    column "routing_rule_properties_enabled_state" {
      rename      = "enabled_state"
      description = "Whether the rule is enabled"
    }

    column "routing_rule_properties_route_configuration" {
      rename            = "route_configuration"
      description       = "A reference to the routing configuration"
      type              = "json"
      generate_resolver = true
    }

    column "routing_rule_properties_rules_engine_id" {
      rename      = "rules_engine_id"
      description = "ID of a specific Rules Engine Configuration to apply to the route"
    }

    column "routing_rule_properties_web_application_firewall_policy_link_id" {
      rename      = "web_application_firewall_policy_link_id"
      description = "ID of the Web Application Firewall policy for each routing rule (if applicable)"
    }

    column "id" {
      description = "Resource ID"
    }

    column "name" {
      description = "Resource name"
    }

    column "type" {
      description = "Resource type"
    }

    relation "azure" "frontdoor" "routing_rule_properties_frontend_endpoints" {
      rename      = "frontend_endpoints"
      description = "Frontend endpoints associated with the rule."

      column "id" {
        description = "Resource ID"
      }
    }
  }

  relation "azure" "frontdoor" "properties_load_balancing_settings" {
    rename      = "load_balancing_settings"
    description = "Load balancing settings for a backend pool associated with the Front Door instance"

    column "load_balancing_settings_properties_resource_state" {
      rename      = "resource_state"
      description = "Resource status"
    }

    column "load_balancing_settings_properties_sample_size" {
      rename      = "sample_size"
      description = "The number of samples to consider for load balancing decisions"
    }

    column "load_balancing_settings_properties_successful_samples_required" {
      rename      = "successful_samples_required"
      description = "The number of samples within the sample period that must succeed"
    }

    column "load_balancing_settings_properties_additional_latency_milliseconds" {
      rename      = "additional_latency_milliseconds"
      description = "The additional latency in milliseconds for probes to fall into the lowest latency bucket"
    }

    column "id" {
      description = "Resource ID"
    }

    column "name" {
      description = "Resource name"
    }

    column "type" {
      description = "Resource type"
    }
  }

  relation "azure" "frontdoor" "properties_health_probe_settings" {
    rename      = "health_probe_settings"
    description = "Health probe settings for a backend pool associated with this Front Door instance"

    column "health_probe_settings_properties_resource_state" {
      rename      = "resource_state"
      description = "Resource status"
    }

    column "health_probe_settings_properties_path" {
      rename      = "path"
      description = "The path to use for the health probe"
    }

    column "protocol" {
      rename      = "protocol"
      description = "Protocol scheme to use for the health probe"
    }

    column "health_probe_settings_properties_interval_in_seconds" {
      rename      = "interval_in_seconds"
      description = "The number of seconds between health probes"
    }

    column "health_probe_settings_properties_health_probe_method" {
      rename      = "health_probe_method"
      description = "Which HTTP method is used to perform the health probe"
    }

    column "health_probe_settings_properties_enabled_state" {
      rename      = "enabled_state"
      description = "Whether the health probe is enabled"
    }

    column "id" {
      description = "Resource ID"
    }

    column "name" {
      description = "Resource name"
    }

    column "type" {
      description = "Resource type"
    }
  }

  relation "azure" "frontdoor" "properties_backend_pools" {
    rename      = "backend_pools"
    skip_prefix = true
    relation "azure" "frontdoor" "backend_pool_properties_backends" {
      rename      = "backends"
      skip_prefix = true
    }
  }

  relation "azure" "frontdoor" "properties_frontend_endpoints" {
    rename      = "frontend_endpoints"
    skip_prefix = true
    column "frontend_endpoint_properties_custom_https_configuration_certificate_source" {
      skip_prefix = true
      rename      = "certificate_source"
    }
    column "frontend_endpoint_properties_custom_https_configuration_protocol_type" {
      skip_prefix = true
      rename      = "protocol_type"
    }
    column "frontend_endpoint_properties_custom_https_configuration_minimum_tls_version" {
      skip_prefix = true
      rename      = "minimum_tls_version"
    }
    column "frontend_endpoint_properties_custom_https_configuration_key_vault_certificate_source_parameters" {
      skip_prefix = true
      rename      = "key_vault_certificate_source_parameters"
    }
    column "frontend_endpoint_properties_custom_https_configuration_vault_id" {
      skip_prefix = true
      rename      = "key_vault_certificate_source_vault_id"
    }
    column "frontend_endpoint_properties_custom_https_configuration_secret_name" {
      skip_prefix = true
      rename      = "key_vault_certificate_source_secret_name"
    }
    column "frontend_endpoint_properties_custom_https_configuration_secret_version" {
      skip_prefix = true
      rename      = "key_vault_certificate_source_secret_version"
    }
    column "frontend_endpoint_properties_custom_https_configuration_certificate_source_parameters" {
      skip_prefix = true
      rename      = "certificate_source_parameters"
    }
    column "frontend_endpoint_properties_custom_https_configuration_certificate_type" {
      skip_prefix = true
      rename      = "certificate_type"
    }
    column "frontend_endpoint_properties_web_application_firewall_policy_link" {
      skip_prefix = true
      rename      = "web_application_firewall_policy_link"
    }
  }
}