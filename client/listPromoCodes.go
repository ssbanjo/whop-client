package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type ListPromoCodesParams struct {
	PageNumber int                   `url:"page,omitempty"`
	PageSize   int                   `url:"per,omitempty"`
	Status     types.PromoCodeStatus `url:"status,omitempty"`
	Expand     []string              `url:"expand,omitempty" del:","`
}

type ListPromoCodesResponse listResponse[types.PromoCode]

// List all promo codes.
//
// https://dev.whop.com/reference/list_promo_codes
func (c Client) ListPromoCodes(params ListPromoCodesParams) (*ListPromoCodesResponse, error) {

	return internal.ExecuteRequest[ListPromoCodesResponse](
		"GET",
		listPromoCodesEndpoint,
		c.headers,
		params,
		nil,
	)
}
