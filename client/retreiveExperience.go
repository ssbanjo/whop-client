package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type RetreiveExperienceParams struct {
	Expand []string `url:"expand,omitempty" del:","`
}

type RetreiveExperienceResponse types.Experience

// Returns a specific Experience, using its unique ID. Experiences represent what the customer receives or unlocks when they purchase a Membership.
//
// https://dev.whop.com/reference/retrieve_experience
func (c Client) RetriveExperience(experienceId string, params RetreiveExperienceParams) (*RetreiveExperienceResponse, error) {

	return internal.ExecuteRequest[RetreiveExperienceResponse](
		"GET",
		fmt.Sprintf(retreiveExperienceEndpoint, experienceId),
		c.headers,
		params,
		nil,
	)
}
