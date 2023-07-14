package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type RetrivePromoCodeParams struct {
	Expand []string `url:"expand,omitempty" del:","`
}

type RetrivePromoCodeResponse types.PromoCode

// Retrieve a promo code.
//
// https://dev.whop.com/reference/retrieve_promo_code
func (c Client) RetrivePromoCode(promoCodeId string, params RetrivePromoCodeParams) (*RetrivePromoCodeResponse, error) {

	return internal.ExecuteRequest[RetrivePromoCodeResponse](
		"GET",
		fmt.Sprintf(retrivePromoCodeEndpoint, promoCodeId),
		c.headers,
		params,
		nil,
	)
}
