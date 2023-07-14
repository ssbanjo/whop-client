package client

import (
	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type CreateWebhookParams struct {
	Url     string               `json:"url,omitempty" url:"-"`
	Enabled bool                 `json:"enabled,omitempty" url:"-"`
	Events  []types.WebhookEvent `json:"events,omitempty" url:"-"`
}

type CreateWebhookResponse types.Webhook

// Create a new Webhook URL.
//
// https://dev.whop.com/reference/create_webhook
func (c Client) CreateWebhook(params CreateWebhookParams) (*CreateWebhookResponse, error) {

	return internal.ExecuteRequest[CreateWebhookResponse](
		"POST",
		createWebhookEndpoint,
		c.headers,
		nil,
		params,
	)
}
