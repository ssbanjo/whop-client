package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type CancelMembershipParams struct {
	Expand []string `url:"expand,omitempty" del:"," json:"-"`
}

type CancelMembershipResponse types.Membership

// Cancels a specific Membership, using its unique ID. It will set the Membership to cancel at period end.
//
// https://dev.whop.com/reference/cancel_membership
func (c Client) CancelMembership(membershipId string, params CancelMembershipParams) (*CancelMembershipResponse, error) {

	return internal.ExecuteRequest[CancelMembershipResponse](
		"POST",
		fmt.Sprintf(cancelMembershipEndpoint, membershipId),
		c.headers,
		params,
		nil,
	)

}
