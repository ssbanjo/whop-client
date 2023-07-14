package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type ListCustomersParams struct {
	PageNumber int `url:"page,omitempty"`
	PageSize   int `url:"per,omitempty"`
}

type ListCustomersResponse listResponse[types.Customer]

// Returns a list of all your members. Members are users who have a currently active membership.
//
// https://dev.whop.com/reference/list_customers
func (c Client) ListCustomers(params ListCustomersParams) (*ListCustomersResponse, error) {

	return internal.ExecuteRequest[ListCustomersResponse](
		"GET",
		listCustomersEndpoint,
		c.headers,
		params,
		nil,
	)
}
