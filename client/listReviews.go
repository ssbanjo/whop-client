package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type ListReviewsParams struct {
	PageNumber int      `url:"page,omitempty"`
	PageSize   int      `url:"per,omitempty"`
	UserId     string   `url:"user_id,omitempty"`
	ProductId  string   `url:"product_id,omitempty"`
	Expand     []string `url:"expand,omitempty" del:","`
}

type ListReviewsResponse listResponse[types.Review]

// Returns a list of all your reviews. Can be searched by User or by Product.
//
// https://dev.whop.com/reference/list_reviews
func (c Client) ListReviews(params ListReviewsParams) (*ListReviewsResponse, error) {

	return internal.ExecuteRequest[ListReviewsResponse](
		"GET",
		listReviewsEndpoint,
		c.headers,
		params,
		nil,
	)
}
