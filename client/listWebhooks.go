package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type ListWebhooksParams struct {
	PageNumber int  `url:"page,omitempty"`
	PageSize   int  `url:"per,omitempty"`
	Enabled    bool `url:"enabled,omitempty"`
}

type ListWebhooksResponse listResponse[types.Webhook]

// List all Webhook URLs.
//
// https://dev.whop.com/reference/list_webhooks
func (c Client) ListWebhooks(params ListWebhooksParams) (*ListWebhooksResponse, error) {

	return internal.ExecuteRequest[ListWebhooksResponse](
		"GET",
		listWebhooksEndpoint,
		c.headers,
		params,
		nil,
	)
}
