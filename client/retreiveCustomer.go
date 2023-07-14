package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type RetreiveCustomerResponse types.Customer

// Retrieves a customer.
//
// https://dev.whop.com/reference/retrieve_customer
func (c Client) RetriveCustomer(search string) (*RetreiveCustomerResponse, error) {

	return internal.ExecuteRequest[RetreiveCustomerResponse](
		"GET",
		fmt.Sprintf(retreiveCustomerEndpoint, search),
		c.headers,
		nil,
		nil,
	)
}
