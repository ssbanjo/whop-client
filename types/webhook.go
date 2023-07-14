package types

type WebhookEvent string

const (
	WebhookEventMembershipWentValid                WebhookEvent = "membership_went_valid"
	WebhookEventMembershipWentInvalid              WebhookEvent = "membership_went_invalid"
	WebhookEventMembershipMetadataUpdated          WebhookEvent = "membership_metadata_updated"
	WebhookEventMembershipExperienceClaimed        WebhookEvent = "membership_experience_claimed"
	WebhookEventMembershipCancelAtPeriodEndChanged WebhookEvent = "membership_cancel_at_period_end_changed"
	WebhookEventPaymentSucceeded                   WebhookEvent = "payment_succeeded"
	WebhookEventPaymentFailed                      WebhookEvent = "payment_failed"
	WebhookEventPaymentAffiliateRewardCreated      WebhookEvent = "payment_affiliate_reward_created"
)

type Webhook struct {
	Id        string `json:"id"`
	Enabled   bool   `json:"enabled"`
	Url       string `json:"url"`
	CreatedAt string `json:"created_at"`
}
