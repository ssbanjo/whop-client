package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type RetreiveCompanyResponse types.Company

// Returns the Company that is currently authenticated with the API.
//
// https://dev.whop.com/reference/retrieve_company
func (c Client) RetreiveCompany() (*RetreiveCompanyResponse, error) {

	return internal.ExecuteRequest[RetreiveCompanyResponse](
		"GET",
		retreiveCompanyEndpoint,
		c.headers,
		nil,
		nil,
	)
}
