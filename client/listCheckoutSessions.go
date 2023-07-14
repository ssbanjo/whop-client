package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type ListCheckoutSessionsParams struct {
	PlanId     string `url:"plan_id,omitempty"`
	PageNumber int    `url:"page,omitempty"`
	PageSize   int    `url:"per,omitempty"`
}

type ListCheckoutSessionsResponse listResponse[types.CheckoutSession]

// Returns a collection of Checkout Sessions, based on the supplied filters.
// A Checkout Session allows specific data to be attached to a Membership upon purchase and customization of the checkout experience.
//
// https://dev.whop.com/reference/list_checkout_sessions
func (c Client) ListCheckoutSessions(params ListCheckoutSessionsParams) (*ListCheckoutSessionsResponse, error) {

	return internal.ExecuteRequest[ListCheckoutSessionsResponse](
		"GET",
		listCheckoutSessionsEndpoint,
		c.headers,
		params,
		nil,
	)
}
