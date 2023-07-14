package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type UpdateMembershipParams struct {
	Expand   []string `url:"expand,omitempty" del:"," json:"-"`
	Metadata any      `json:"metadata,omitempty" url:"-"`
}

type UpdateMembershipResponse types.Membership

// Use the update license endpoint to add metadata like hwid to an instance of the License.
// This can help you track things like user preferences and the device a user is using your software from.
//
// https://dev.whop.com/reference/update_membership
func (c Client) UpdateMembership(membershipId string, params UpdateMembershipParams) (*UpdateMembershipResponse, error) {

	return internal.ExecuteRequest[UpdateMembershipResponse](
		"POST",
		fmt.Sprintf(updateMembershipEndpoint, membershipId),
		c.headers,
		params,
		params,
	)
}
