package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type CreatePlanParams struct {
	PlanType                 string               `json:"plan_type,omitempty" url:"-"`
	BaseCurrency             string               `json:"base_currency,omitempty" url:"-"`
	CustomFields             []string             `json:"custom_fields,omitempty" url:"-"`
	PeggedCurrencies         []string             `json:"pegged_currencies,omitempty" url:"-"`
	ReleaseMethodSettings    any                  `json:"release_method_settings,omitempty" url:"-"`
	Requirements             any                  `json:"requirements,omitempty" url:"-"`
	Metadata                 any                  `json:"metadata,omitempty" url:"-"`
	ProductID                string               `json:"product_id,omitempty" url:"-"`
	AccessPassID             string               `json:"access_pass_id,omitempty" url:"-"`
	AllowMultipleQuantity    bool                 `json:"allow_multiple_quantity,omitempty" url:"-"`
	BillingPeriod            int                  `json:"billing_period,omitempty" url:"-"`
	CardPayments             bool                 `json:"card_payments,omitempty" url:"-"`
	CoinbaseCommerceAccepted bool                 `json:"coinbase_commerce_accepted,omitempty" url:"-"`
	ExpirationDays           int                  `json:"expiration_days,omitempty" url:"-"`
	GracePeriodDays          int                  `json:"grace_period_days,omitempty" url:"-"`
	InitialPrice             float64              `json:"initial_price,omitempty" url:"-"`
	InternalNotes            string               `json:"internal_notes,omitempty" url:"-"`
	OnePerUser               bool                 `json:"one_per_user,omitempty" url:"-"`
	Refillable               bool                 `json:"refillable,omitempty" url:"-"`
	ReleaseMethod            string               `json:"release_method,omitempty" url:"-"`
	RenewalPrice             float64              `json:"renewal_price,omitempty" url:"-"`
	ShortLink                string               `json:"short_link,omitempty" url:"-"`
	Stock                    int                  `json:"stock,omitempty" url:"-"`
	TrialPeriodDays          int                  `json:"trial_period_days,omitempty" url:"-"`
	UnlimitedStock           bool                 `json:"unlimited_stock,omitempty" url:"-"`
	Visibility               types.PlanVisibility `json:"visibility,omitempty" url:"-"`
	Expand                   []string             `url:"expand,omitempty" del:"," json:"-"`
}

type CreatePlanResponse types.Plan

// Create a plan.
//
// https://dev.whop.com/reference/create_plan
func (c Client) CreatePlan(params CreatePlanParams) (*CreatePlanResponse, error) {

	return internal.ExecuteRequest[CreatePlanResponse](
		"POST",
		createPlanEndpoint,
		c.headers,
		params,
		params,
	)
}
