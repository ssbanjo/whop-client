package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type RetrieveProductParams struct {
	Expand []string `url:"expand,omitempty" del:","`
}

type RetrieveProductResponse types.Product

// Returns a specific Product, using its unique ID. Products represent what the customer has purchased from your company.
// Therefore, the Product can be used to provide context to the user (and your application) on what they have purchased.
//
// https://dev.whop.com/reference/retrieve_product
func (c Client) RetrieveProduct(productId string, params RetrieveProductParams) (*RetrieveProductResponse, error) {

	return internal.ExecuteRequest[RetrieveProductResponse](
		"GET",
		fmt.Sprintf(retreiveProductEndpoint, productId),
		c.headers,
		params,
		nil,
	)
}
