package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type UpdatePromoCodeParams struct {
	PlanIds        []string              `json:"plan_ids,omitempty" url:"-"`
	Stock          int                   `json:"stock,omitempty" url:"-"`
	UnlimitedStock bool                  `json:"unlimited_stock,omitempty" url:"-"`
	Metadata       any                   `json:"metadata,omitempty" url:"-"`
	Status         types.PromoCodeStatus `json:"status,omitempty" url:"-"`
	Expand         []string              `url:"expand,omitempty" del:"," json:"-"`
}

type UpdatePromoCodeResponse types.PromoCode

// Update a promo code's applicable plans.
//
// https://dev.whop.com/reference/update_promo_code
func (c Client) UpdatePromoCode(promoCodeId string, params UpdatePromoCodeParams) (*UpdatePromoCodeResponse, error) {

	return internal.ExecuteRequest[UpdatePromoCodeResponse](
		"POST",
		fmt.Sprintf(updatePromoCodeEndpoint, promoCodeId),
		c.headers,
		params,
		params,
	)
}
