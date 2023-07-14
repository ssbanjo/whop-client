package types

type CheckoutSession struct {
	Id            string `json:"id"`
	RedirectUrl   string `json:"redirect_url"`
	AffiliateCode string `json:"affiliate_code"`
	Metadata      any    `json:"metadata"`
	PlanId        string `json:"plan_id"`
	PurchaseUrl   string `json:"purchase_url"`
}
