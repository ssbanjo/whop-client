package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type UpdateExperienceParams struct {
	Expand []string `url:"expand,omitempty" del:"," json:"-"`
	Name   string   `json:"name,omitempty" url:"-"`
}

type UpdateExperienceResponse types.Experience

// Returns a specific Experience, using its unique ID. Experiences represent what the customer receives or unlocks when they purchase a Membership.
//
// https://dev.whop.com/reference/update_experience
func (c Client) UpdateExperience(experienceId string, params UpdateExperienceParams) (*UpdateExperienceResponse, error) {

	return internal.ExecuteRequest[UpdateExperienceResponse](
		"POST",
		fmt.Sprintf(updateExperienceEndpoint, experienceId),
		c.headers,
		params,
		params,
	)
}
