service          = "azure"
output_directory = "."
add_generate     = true

resource "azure" "frontdoors" "cdns" {
  path = "github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor.FrontDoor"

  userDefinedColumn "subscription_id" {
    type        = "string"
    description = "Azure subscription id"
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

  relation "azure" "frontdoor" "properties_rules_engines" {
    relation "azure" "frontdoor" "rules_engine_properties_rules" {
      column "action_route_configuration_override" {
        type              = "json"
        generate_resolver = true
      }
    }
  }
  relation "azure" "frontdoor" "properties_routing_rules" {
    column "routing_rule_properties_route_configuration" {
      type              = "json"
      generate_resolver = true
    }
  }
}