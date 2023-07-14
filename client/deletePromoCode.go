package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type DeletePromoCodeResponse types.PromoCode

// Delete a promo code. You can only delete a promo code that has not been used yet.
//
// We recommend using the update endpoint and setting its status to 'archived' over this.
//
// https://dev.whop.com/reference/delete_promo_code
func (c Client) DeletePromoCode(promoCodeId string) (*DeletePromoCodeResponse, error) {

	return internal.ExecuteRequest[DeletePromoCodeResponse](
		"DELETE",
		fmt.Sprintf(deletePromoCodeEndpoint, promoCodeId),
		c.headers,
		nil,
		nil,
	)
}
