package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type ListProductsParams struct {
	PageNumber int                     `url:"page,omitempty"`
	PageSize   int                     `url:"per,omitempty"`
	Visibility types.ProductVisibility `url:"visibility,omitempty"`
	Expand     []string                `url:"expand,omitempty" del:","`
}

type ListProductsResponse listResponse[types.Product]

// Returns a collection of Products, based on the supplied filters.
// If no filters are applied, all products (regardless of their currently availability) are accessible from this endpoint.
//
// https://dev.whop.com/reference/list_products
func (c Client) ListProducts(params ListProductsParams) (*ListProductsResponse, error) {

	return internal.ExecuteRequest[ListProductsResponse](
		"GET",
		listProductsEndpoint,
		c.headers,
		params,
		nil,
	)
}
