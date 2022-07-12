service          = "azure"
output_directory = "."
add_generate     = true

resource "azure" "subscription" "tenants" {
  path = "github.com/Azure/azure-sdk-for-go/services/subscription/mgmt/2020-09-01/subscription.TenantIDDescription"
  description = "Azure tenant information"

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
      "subscription_id",
      "id"
    ]
  }

  column "id" {
    type = "string"
    description = "The fully qualified ID of the tenant"
  }

  column "tenant_id" {
    type = "string"
    description = "The tenant ID"
  }

  multiplex "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.SubscriptionMultiplex"
  }

  deleteFilter "AzureSubscription" {
    path = "github.com/cloudquery/cq-provider-azure/client.DeleteSubscriptionFilter"
  }
}
