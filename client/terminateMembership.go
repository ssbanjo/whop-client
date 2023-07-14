package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type TerminateMembershipParams struct {
	Expand []string `url:"expand,omitempty" del:"," json:"-"`
}

type TerminateMembershipResponse types.Membership

// Terminates a specific Membership, using its unique ID. It will immediately invalidate the Membership, and the User will/should have their Product's Experiences unfulfilled.
// Termination is irreversible, so a User will have to re-purchase if they wish to continue their access.
//
// https://dev.whop.com/reference/terminate_membership
func (c Client) TerminateMembership(membershipId string, params TerminateMembershipParams) (*TerminateMembershipResponse, error) {

	return internal.ExecuteRequest[TerminateMembershipResponse](
		"POST",
		fmt.Sprintf(terminateMembershipEndpoint, membershipId),
		c.headers,
		params,
		nil,
	)
}
