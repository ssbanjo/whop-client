package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type UpdateProductParams struct {
	PageNumber  int                     `json:"page,omitempty" url:"-"`
	PageSize    int                     `json:"per,omitempty" url:"-"`
	Name        string                  `json:"name,omitempty" url:"-"`
	OnePerUser  bool                    `json:"one_per_user,omitempty" url:"-"`
	Visibility  types.ProductVisibility `json:"visibility,omitempty" url:"-"`
	Shuffleable bool                    `json:"shuffleable,omitempty" url:"-"`
	Expand      []string                `url:"expand,omitempty" del:"," json:"-"`
}

type UpdateProductResponse types.Product

// Update a Product's attributes, using its unique ID. Keep in mind, any changes made to the pass will be reflected on the User's account.
//
// https://dev.whop.com/reference/update_product
func (c Client) UpdateProduct(productId string, params UpdateProductParams) (*UpdateProductResponse, error) {

	return internal.ExecuteRequest[UpdateProductResponse](
		"POST",
		fmt.Sprintf(updateProductEndpoint, productId),
		c.headers,
		params,
		params,
	)
}
