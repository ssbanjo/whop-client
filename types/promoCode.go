package types

import "time"

type PromoCodeType string

const (
	PromoCodeTypePercentage PromoCodeType = "percentage"
	PromoCodeTypeFlatAmount PromoCodeType = "flat_amount"
)

type PromoCodeStatus string

const (
	PromoCodeStatusActive   PromoCodeStatus = "active"
	PromoCodeStatusInactive PromoCodeStatus = "inactive"
	PromoCodeStatusArchived PromoCodeStatus = "archived"
)

type PromoCode struct {
	ID                 string          `json:"id"`
	CreatedAt          time.Time       `json:"created_at"`
	AmountOff          string          `json:"amount_off"`
	BaseCurrency       string          `json:"base_currency"`
	Code               string          `json:"code"`
	ExpirationDatetime int             `json:"expiration_datetime"`
	NewUsersOnly       bool            `json:"new_users_only"`
	NumberOfIntervals  int             `json:"number_of_intervals"`
	PlanIds            []string        `json:"plan_ids"`
	PromoType          PromoCodeType   `json:"promo_type"`
	Status             PromoCodeStatus `json:"status"`
	Stock              int             `json:"stock"`
	Uses               int             `json:"uses"`
}
