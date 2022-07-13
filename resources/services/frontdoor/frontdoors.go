package frontdoor

import (
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource cdns --config gen.hcl --output .
func FrontDoors() *schema.Table {
	return nil
}
