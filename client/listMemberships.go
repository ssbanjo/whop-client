package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type ListMembershipsParams struct {
	PageNumber    int                    `url:"page,omitempty"`
	PageSize      int                    `url:"per,omitempty"`
	Expand        []string               `url:"expand,omitempty" del:","`
	Status        types.MembershipStatus `url:"status,omitempty"`
	PlanId        string                 `url:"plan_id,omitempty"`
	ProductId     string                 `url:"product_id,omitempty"`
	UserId        string                 `url:"user_id,omitempty"`
	DiscordId     string                 `url:"discord_id,omitempty"`
	WalletAddress string                 `url:"wallet_address,omitempty"`
	Valid         bool                   `url:"valid,omitempty"`
	HideMetadata  bool                   `url:"hide_metadata,omitempty"`
}

type ListMembershipsResponse listResponse[types.Membership]

// Returns a collection of Memberships, based on the supplied filters.
//
// https://dev.whop.com/reference/list_memberships
func (c Client) ListMemberships(params ListMembershipsParams) (*ListMembershipsResponse, error) {

	return internal.ExecuteRequest[ListMembershipsResponse](
		"GET",
		listMembershipsEndpoint,
		c.headers,
		params,
		nil,
	)
}
