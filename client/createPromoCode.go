package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type CreatePromoCodeParams struct {
	AmountOff          float64             `json:"amount_off,omitempty" url:"-"`
	BaseCurrency       string              `json:"base_currency,omitempty" url:"-"`
	Code               string              `json:"code,omitempty" url:"-"`
	PromoType          types.PromoCodeType `json:"promo_type,omitempty" url:"-"`
	ExpirationDatetime int                 `json:"expiration_datetime,omitempty" url:"-"`
	NewUsersOnly       bool                `json:"new_users_only,omitempty" url:"-"`
	NumberOfIntervals  int                 `json:"number_of_intervals,omitempty" url:"-"`
	PlanIds            []string            `json:"plan_ids,omitempty" url:"-"`
	Stock              int                 `json:"stock,omitempty" url:"-"`
	UnlimitedStock     bool                `json:"unlimited_stock,omitempty" url:"-"`
	Metadata           any                 `json:"metadata,omitempty" url:"-"`
	Expand             []string            `url:"expand,omitempty" del:"," json:"-"`
}

type CreatePromoCodeResponse types.PromoCode

// Creates a new promo code with the given parameters. A promo code can offer a discount or a free period for a specific set of plans or for all plans in your company.
//
// https://dev.whop.com/reference/create_promo_code
func (c Client) CreatePromoCode(params CreatePromoCodeParams) (*CreatePromoCodeResponse, error) {

	return internal.ExecuteRequest[CreatePromoCodeResponse](
		"POST",
		createPromoCodeEndpoint,
		c.headers,
		params,
		params,
	)
}
