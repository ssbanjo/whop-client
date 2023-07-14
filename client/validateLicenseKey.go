package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type ValidateLicenseKeyParams struct {
	Expand   []string `url:"expand,omitempty" del:"," json:"-"`
	Metadata any      `json:"metadata,omitempty" url:"-"`
}

type ValidateLicenseKeyResponse types.Membership

// Validate an instance of your software by sending metadata about the user. For example, by sending metadata containing the HWID of the device most recently logged in,
// you can ensure an instance of your software is only being used on one device at a time.
//
// https://dev.whop.com/reference/validate_license
func (c Client) ValidateLicenseKey(licenseKey string, params ValidateLicenseKeyParams) (*ValidateLicenseKeyResponse, error) {

	return internal.ExecuteRequest[ValidateLicenseKeyResponse](
		"POST",
		fmt.Sprintf(validateLicenseKeyEndpoint, licenseKey),
		c.headers,
		params,
		params,
	)
}
