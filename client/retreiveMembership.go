package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type RetreiveMembershipParams struct {
	Expand []string `url:"expand,omitempty" del:","`
}

type RetreiveMembershipResponse types.Membership

// Returns a specific Membership, using its unique ID.
//
// https://dev.whop.com/reference/retrieve_membership
func (c Client) RetriveMembership(membershipId string, params RetreiveMembershipParams) (*RetreiveMembershipResponse, error) {

	return internal.ExecuteRequest[RetreiveMembershipResponse](
		"GET",
		fmt.Sprintf(retreiveMembershipEndpoint, membershipId),
		c.headers,
		params,
		nil,
	)
}
