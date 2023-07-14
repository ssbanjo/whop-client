package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type RetrivePlanParams struct {
	Expand []string `url:"expand,omitempty" del:","`
}

type RetreivePlanResponse types.Plan

// Returns a specific Plan, using its unique ID. Plans represent a pricing and release method configuration used to purchase a Product.
//
// https://dev.whop.com/reference/retrieve_plan
func (c Client) RetreivePlan(planId string, params RetrivePlanParams) (*RetreivePlanResponse, error) {

	return internal.ExecuteRequest[RetreivePlanResponse](
		"GET",
		fmt.Sprintf(retrivePlanEndpoint, planId),
		c.headers,
		params,
		nil,
	)
}
