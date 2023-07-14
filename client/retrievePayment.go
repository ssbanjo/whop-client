package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type RetreivePaymentsParams struct {
	Expand []string `url:"expand,omitempty" del:"," json:"-"`
}

type RetreivePaymentsResponse types.Payment

// Returns a specific Payment, using its unique ID. Payments represent a monetary transaction between a User and a Company.
// Only Payments with a status of paid have been successfully processed.
//
// https://dev.whop.com/reference/retrieve_payment
func (c Client) RetrivePayment(paymentId string, params RetreivePaymentsParams) (*RetreivePaymentsResponse, error) {

	return internal.ExecuteRequest[RetreivePaymentsResponse](
		"GET",
		listPaymentsEndpoint,
		c.headers,
		params,
		nil,
	)
}
