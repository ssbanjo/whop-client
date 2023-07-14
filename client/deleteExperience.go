package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type DeleteExperienceResponse types.Experience

// Delete an experience.
//
// https://dev.whop.com/reference/delete_experience
func (c Client) DeleteExperience(experienceId string) (*DeleteExperienceResponse, error) {

	return internal.ExecuteRequest[DeleteExperienceResponse](
		"DELETE",
		fmt.Sprintf(deleteExperienceEndpoint, experienceId),
		c.headers,
		nil,
		nil,
	)
}
