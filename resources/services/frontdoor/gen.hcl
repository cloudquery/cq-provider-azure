service          = "azure"
output_directory = "."
add_generate     = true

resource "azure" "frontdoors" "cdns" {
  path = "github.com/Azure/azure-sdk-for-go/services/frontdoor/mgmt/2020-11-01/frontdoor.FrontDoor"

  relation "azure" "frontdoors" "action_response_header_actions" {
    column "cdn_properties_rules_engine_rules_engine_properties_rules" {
      skip = true
    }
  }
}