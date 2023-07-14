package types

type Payment struct {
	ID                 string `json:"id"`
	Product            string `json:"product"`
	Membership         string `json:"membership"`
	Plan               string `json:"plan"`
	User               string `json:"user"`
	FinalAmount        int    `json:"final_amount"`
	Subtotal           int    `json:"subtotal"`
	Currency           string `json:"currency"`
	Last4              string `json:"last4"`
	LastPaymentAttempt int    `json:"last_payment_attempt"`
	NextPaymentAttempt int    `json:"next_payment_attempt"`
	PaymentsFailed     int    `json:"payments_failed"`
	PaymentProcessor   string `json:"payment_processor"`
	RefundedAmount     int    `json:"refunded_amount"`
	RefundedAt         int    `json:"refunded_at"`
	Status             string `json:"status"`
	WalletAddress      string `json:"wallet_address"`
	CryptoTxHash       string `json:"crypto_tx_hash"`
}
