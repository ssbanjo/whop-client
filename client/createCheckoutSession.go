package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type CreateCheckoutSessionsParams struct {
	PlanId        string `json:"plan_id,omitempty"`
	AffiliateCode string `json:"affiliate_code,omitempty"`
	RedirectUrl   string `json:"redirect_url,omitempty"`
	Metadata      any    `json:"metadata,omitempty"`
}

type CreateCheckoutSessionsResponse types.CheckoutSession

// Creates a new Checkout Session.
//
// https://dev.whop.com/reference/create_checkout_session
func (c Client) CreateCheckoutSessions(params CreateCheckoutSessionsParams) (*CreateCheckoutSessionsResponse, error) {

	return internal.ExecuteRequest[CreateCheckoutSessionsResponse](
		"POST",
		createCheckoutSessionEndpoint,
		c.headers,
		nil,
		params,
	)
}
