package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type AddFreeDaysToMembershipParams struct {
	Expand []string `url:"expand,omitempty" del:"," json:"-"`
	Days   int      `json:"days,omitempty" url:"-"`
}

type AddFreeDaysToMembershipResponse types.Membership

// Extend a Membership's next renewal date or expiration date by a specified number of days.
//
// https://dev.whop.com/reference/add_free_days_membership
func (c Client) AddFreeDaysToMembership(membershipId string, params AddFreeDaysToMembershipParams) (*AddFreeDaysToMembershipResponse, error) {

	return internal.ExecuteRequest[AddFreeDaysToMembershipResponse](
		"POST",
		fmt.Sprintf(addFreeDaysToMembershipEndpoint, membershipId),
		c.headers,
		params,
		params,
	)
}
