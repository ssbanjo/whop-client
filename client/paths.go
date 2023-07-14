package client

const whopApiPath string = "https://api.whop.com/api/v2"

const (
	retreiveCompanyEndpoint         string = whopApiPath + "/company"
	listMembershipsEndpoint         string = whopApiPath + "/memberships"
	retreiveMembershipEndpoint      string = whopApiPath + "/memberships/%s"
	updateMembershipEndpoint        string = whopApiPath + "/memberships/%s"
	addFreeDaysToMembershipEndpoint string = whopApiPath + "/memberships/%s/add_free_days"
	terminateMembershipEndpoint     string = whopApiPath + "/memberships/%s/terminate"
	cancelMembershipEndpoint        string = whopApiPath + "/memberships/%s/cancel"
	validateLicenseKeyEndpoint      string = whopApiPath + "/memberships/%s/validate_license"
	listPaymentsEndpoint            string = whopApiPath + "/payments"
	retrivePaymentsEndpoint         string = whopApiPath + "/payments/%s"
	listPlansEndpoint               string = whopApiPath + "/plans"
	createPlanEndpoint              string = whopApiPath + "/plans"
	retrivePlanEndpoint             string = whopApiPath + "/plans/%s"
	createQuickLinkEndpoint         string = whopApiPath + "/plans/%s/create_quick_link"
	listProductsEndpoint            string = whopApiPath + "/products"
	retreiveProductEndpoint         string = whopApiPath + "/products/%s"
	updateProductEndpoint           string = whopApiPath + "/products/%s"
	listCheckoutSessionsEndpoint    string = whopApiPath + "/checkout_sessions"
	createCheckoutSessionEndpoint   string = whopApiPath + "/checkout_sessions"
	retreiveCheckoutSessionEndpoint string = whopApiPath + "/checkout_sessions/%s"
	deleteCheckoutSessionEndpoint   string = whopApiPath + "/checkout_sessions/%s"
	listCustomersEndpoint           string = whopApiPath + "/customers"
	retreiveCustomerEndpoint        string = whopApiPath + "/customers/%s"
	listPromoCodesEndpoint          string = whopApiPath + "/promo_codes"
	createPromoCodeEndpoint         string = whopApiPath + "/promo_codes"
	retrivePromoCodeEndpoint        string = whopApiPath + "/promo_codes/%s"
	updatePromoCodeEndpoint         string = whopApiPath + "/promo_codes/%s"
	deletePromoCodeEndpoint         string = whopApiPath + "/promo_codes/%s"
	listExperiencesEndpoint         string = whopApiPath + "/experiences"
	retreiveExperienceEndpoint      string = whopApiPath + "/experiences/%s"
	updateExperienceEndpoint        string = whopApiPath + "/experiences/%s"
	deleteExperienceEndpoint        string = whopApiPath + "/experiences/%s"
	sendPushNotificationEndpoint    string = whopApiPath + "/push_notifications"
	listReviewsEndpoint             string = whopApiPath + "/reviews"
	retreiveReviewEndpoint          string = whopApiPath + "/reviews/%s"
	listWebhooksEndpoint            string = whopApiPath + "/webhooks"
	createWebhookEndpoint           string = whopApiPath + "/webhooks"
	retreiveWebhookEndpoint         string = whopApiPath + "/webhooks/%s"
	updateWebhookEndpoint           string = whopApiPath + "/webhooks/%s"
	deleteWebhookEndpoint           string = whopApiPath + "/webhooks/%s"
)