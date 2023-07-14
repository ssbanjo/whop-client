package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type UpdateWebhookParams struct {
	Enabled bool   `json:"enabled,omitempty" url:"-"`
	Url     string `json:"url,omitempty" url:"-"`
}

type UpdateWebhookResponse types.Webhook

// Update a webhook URL.
//
// https://dev.whop.com/reference/update_webhook
func (c Client) UpdateWebhook(webhookId string, params UpdateWebhookParams) (*UpdateWebhookResponse, error) {

	return internal.ExecuteRequest[UpdateWebhookResponse](
		"POST",
		fmt.Sprintf(updateWebhookEndpoint, webhookId),
		c.headers,
		nil,
		params,
	)
}
