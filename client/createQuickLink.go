package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type CreateQuickLinkParams struct {
	Stock           int      `json:"stock,omitempty" url:"-"`
	Metadata        any      `json:"metadata,omitempty" url:"-"`
	Requirements    any      `json:"requirements,omitempty" url:"-"`
	InternalNotes   string   `json:"internal_notes,omitempty" url:"-"`
	ShortLink       string   `json:"short_link,omitempty" url:"-"`
	CustomPassword  string   `json:"custom_password,omitempty" url:"-"`
	TrialPeriodDays int      `json:"trial_period_days,omitempty" url:"-"`
	Expand          []string `url:"expand,omitempty" del:"," json:"-"`
}

type CreateQuickLinkResponse types.Plan

// Create a quick link (or release link) for a plan. This will copy over all data from an existing plan and make a new plan with the same data that can be used as a one-time release link.
// You can change the stock and metadata on the new plan.
//
// https://dev.whop.com/reference/create_quick_link
func (c Client) CreateQuickLink(planId string, params CreateQuickLinkParams) (*CreateQuickLinkResponse, error) {

	return internal.ExecuteRequest[CreateQuickLinkResponse](
		"POST",
		fmt.Sprintf(createQuickLinkEndpoint, planId),
		c.headers,
		params,
		params,
	)
}
