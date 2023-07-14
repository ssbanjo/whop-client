package types

type MembershipStatus string

const (
	MembershipStatusDrafted    MembershipStatus = "drafted"
	MembershipStatusTrailing   MembershipStatus = "trailing"
	MembershipStatusActive     MembershipStatus = "active"
	MembershipStatusPastDue    MembershipStatus = "past_due"
	MembershipStatusCompleted  MembershipStatus = "completed"
	MembershipStatusCanceled   MembershipStatus = "canceled"
	MembershipStatusExpired    MembershipStatus = "expired"
	MembershipStatusUnresolved MembershipStatus = "unresolved"
)

type MembershipPayment string

const (
	MembershipPaymentFree     MembershipPayment = "free"
	MembershipPaymentStripe   MembershipPayment = "stripe"
	MembershipPaymentCoinbase MembershipPayment = "coinbase"
	MembershipPaymentCrypto   MembershipPayment = "crypto"
)

type Membership struct {
	ID                    string            `json:"id"`
	Product               string            `json:"product"`
	User                  string            `json:"user"`
	Plan                  string            `json:"plan"`
	Email                 string            `json:"email"`
	Status                MembershipStatus  `json:"status"`
	Valid                 bool              `json:"valid"`
	CancelAtPeriodEnd     bool              `json:"cancel_at_period_end"`
	PaymentProcessor      MembershipPayment `json:"payment_processor"`
	LicenseKey            string            `json:"license_key"`
	Metadata              any               `json:"metadata"`
	Quantity              int               `json:"quantity"`
	WalletAddress         string            `json:"wallet_address"`
	CustomFieldsResponses any               `json:"custom_fields_responses"`
	Discord               struct {
		ID       string `json:"id"`
		Username string `json:"username"`
	} `json:"discord"`
	NftTokens          []NftToken `json:"nft_tokens"`
	ExpiresAt          int        `json:"expires_at"`
	RenewalPeriodStart int        `json:"renewal_period_start"`
	RenewalPeriodEnd   int        `json:"renewal_period_end"`
	CreatedAt          int        `json:"created_at"`
	ManageURL          string     `json:"manage_url"`
	AccessPass         string     `json:"access_pass"`
}
