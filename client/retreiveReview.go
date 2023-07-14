package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type RetriveReviewParams struct {
	Expand []string `url:"expand,omitempty" del:","`
}

type RetriveReviewResponse types.Review

// Retrieves a specific Review, using its unique ID.
//
// https://dev.whop.com/reference/retrieve_review
func (c Client) RetriveReview(reviewId string, params RetriveReviewParams) (*RetriveReviewResponse, error) {

	return internal.ExecuteRequest[RetriveReviewResponse](
		"GET",
		fmt.Sprintf(retreiveReviewEndpoint, reviewId),
		c.headers,
		params,
		nil,
	)
}
