package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type DeleteCheckoutSessionsResponse types.CheckoutSession

// Deletes a specific Checkout Session, using its unique ID.
//
// https://dev.whop.com/reference/delete_checkout_session
func (c Client) DeleteCheckoutSessions(sessionId string) (*DeleteCheckoutSessionsResponse, error) {

	return internal.ExecuteRequest[DeleteCheckoutSessionsResponse](
		"DELETE",
		fmt.Sprintf(deleteCheckoutSessionEndpoint, sessionId),
		c.headers,
		nil,
		nil,
	)
}
