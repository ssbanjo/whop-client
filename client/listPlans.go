package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type ListPlansParams struct {
	PageNumber int                  `url:"page,omitempty"`
	PageSize   int                  `url:"per,omitempty"`
	Visibility types.PlanVisibility `url:"visibility,omitempty"`
	ProductId  string               `url:"product_id,omitempty"`
	Expand     []string             `url:"expand,omitempty" del:","`
}

type ListPlansResponse listResponse[types.Plan]

// Returns a collection of Plans, based on the supplied filters. If no filters are applied, all plans (regardless of their currently availability) are accessible from this endpoint.
//
// https://dev.whop.com/reference/list_plans
func (c Client) ListPlans(params ListPlansParams) (*ListPlansResponse, error) {

	return internal.ExecuteRequest[ListPlansResponse](
		"GET",
		listPlansEndpoint,
		c.headers,
		params,
		nil,
	)
}
