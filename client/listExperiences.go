package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type ListExperiencesParams struct {
	PageNumber     int                  `url:"page,omitempty"`
	PageSize       int                  `url:"per,omitempty"`
	ExperienceType types.ExperienceType `url:"experience_type,omitempty"`
	Expand         []string             `url:"expand,omitempty" del:","`
}

type ListExperiencesResponse listResponse[types.Experience]

// Returns a collection of Experiences, based on the supplied filters. Experiences represent what the customer receives or unlocks when they purchase an Product.
//
// https://dev.whop.com/reference/list_experiences
func (c Client) ListExperiences(params ListExperiencesParams) (*ListExperiencesResponse, error) {

	return internal.ExecuteRequest[ListExperiencesResponse](
		"GET",
		listExperiencesEndpoint,
		c.headers,
		params,
		nil,
	)
}
