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

  relation "azure" "frontdoor" "properties_rules_engines" {
    rename      = "rules_engines"
    skip_prefix = true
    relation "azure" "frontdoor" "rules_engine_properties_rules" {
      rename      = "rules"
      skip_prefix = true
      column "action_route_configuration_override" {
        type              = "json"
        generate_resolver = true
      }
    }
  }

  relation "azure" "frontdoor" "properties_routing_rules" {
    rename      = "routing_rules"
    skip_prefix = true
    column "routing_rule_properties_route_configuration" {
      type              = "json"
      generate_resolver = true
    }
  }

  relation "azure" "frontdoor" "properties_load_balancing_settings" {
    rename      = "load_balancing_settings"
    skip_prefix = true
    column "load_balancing_settings_properties_additional_latency_milliseconds" {
      rename = "additional_latency_milliseconds"
    }
  }

  relation "azure" "frontdoor" "properties_health_probe_settings" {
    rename      = "health_probe_settings"
    skip_prefix = true
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