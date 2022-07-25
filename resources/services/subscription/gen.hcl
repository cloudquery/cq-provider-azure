service          = "azure"
output_directory = "."
add_generate     = true

description_modifier "remove_read_only" {
  words = ["READ-ONLY; "]
}

description_modifier "remove_field_name" {
  regex = ".+- "
}

// /zz_generated_models.go
resource "azure" "subscription" "subscriptions" {
  path = "github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription.Model"
  description = "Azure subscription information"

  userDefinedColumn "subscription_id" {
    type        = "string"
    description = "Azure subscription id"
    resolver "resolveAzureSubscription" {
      path = "github.com/cloudquery/cq-provider-azure/client.ResolveAzureSubscription"
    }
  }
  column "subscription_id" {
    skip = true
  }
  options {
    primary_keys = [
      "id"
    ]
  }

  column "subscription_policies_location_placement_id" {
    rename = "location_placement_id"
  }

  column "subscription_policies_quota_id" {
    rename = "quota_id"
  }

  column "subscription_policies_spending_limit" {
    rename = "spending_limit"
  }

  multiplex "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.SubscriptionMultiplex"
  }

  deleteFilter "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.DeleteSubscriptionFilter"
  }
}

resource "azure" "subscription" "tenants" {
  path = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions.TenantIDDescription"
  description = "Azure tenant information"

  options {
    primary_keys = [
      "id"
    ]
  }

  column "id" {
    type = "string"
    description = "The fully qualified ID of the tenant"
  }

  multiplex "SingleSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.SingleSubscriptionMultiplex"
  }
}