package types

type PlanVisibility string

const (
	PlanVisibilityVisible   PlanVisibility = "visible"
	PlanVisibilityHidden    PlanVisibility = "hidden"
	PlanVisibilityArchived  PlanVisibility = "archived"
	PlanVisibilityQuickLink PlanVisibility = "quick_link"
)

type Plan struct {
	ID                     string         `json:"id"`
	Product                string         `json:"product"`
	PlanType               string         `json:"plan_type"`
	ReleaseMethod          string         `json:"release_method"`
	Visibility             PlanVisibility `json:"visibility"`
	BillingPeriod          int            `json:"billing_period"`
	InternalNotes          string         `json:"internal_notes"`
	Metadata               any            `json:"metadata"`
	DirectLink             string         `json:"direct_link"`
	RenewalPrice           string         `json:"renewal_price"`
	InitialPrice           string         `json:"initial_price"`
	BaseCurrency           string         `json:"base_currency"`
	Requirements           any            `json:"requirements"`
	ReleaseMethodSettings  any            `json:"release_method_settings"`
	AcceptedPaymentMethods []string       `json:"accepted_payment_methods"`
	Stock                  int            `json:"stock"`
	UnlimitedStock         bool           `json:"unlimited_stock"`
	CreatedAt              int            `json:"created_at"`
	AccessPass             string         `json:"access_pass"`
	CardPayments           bool           `json:"card_payments"`
}
