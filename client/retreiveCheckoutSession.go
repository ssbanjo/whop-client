package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type RetreiveCheckoutSessionResponse types.CheckoutSession

// Retrieve a Checkout Session.
//
// https://dev.whop.com/reference/retrieve_checkout_session
func (c Client) RetreiveCheckoutSession(sessionId string) (*RetreiveCheckoutSessionResponse, error) {

	return internal.ExecuteRequest[RetreiveCheckoutSessionResponse](
		"GET",
		fmt.Sprintf(retreiveCheckoutSessionEndpoint, sessionId),
		c.headers,
		nil,
		nil,
	)
}
