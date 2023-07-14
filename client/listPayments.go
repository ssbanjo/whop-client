package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type ListPaymentsParams struct {
	PageNumber   int      `url:"page,omitempty"`
	PageSize     int      `url:"per,omitempty"`
	Status       string   `url:"status,omitempty"`
	MembershipId string   `url:"membership_id,omitempty"`
	Expand       []string `url:"expand,omitempty" del:","`
}

type ListPaymentsResponse listResponse[types.Payment]

// Returns a collection of Payments, based on the supplied filters. Payments represent a monetary transaction between a User and a Company.
// Only Payments with a status of paid have been successfully processed.
//
// https://dev.whop.com/reference/list_payments
func (c Client) ListPayments(params ListPaymentsParams) (*ListPaymentsResponse, error) {

	return internal.ExecuteRequest[ListPaymentsResponse](
		"GET",
		listPaymentsEndpoint,
		c.headers,
		params,
		nil,
	)
}
