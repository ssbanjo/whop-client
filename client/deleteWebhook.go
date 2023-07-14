package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type DeleteWebhookResponse types.Webhook

// Delete a specific Webhook URL.
//
// https://dev.whop.com/reference/delete_webhook
func (c Client) DeleteWebhook(webhookId string) (*DeleteWebhookResponse, error) {

	return internal.ExecuteRequest[DeleteWebhookResponse](
		"DELETE",
		fmt.Sprintf(deleteWebhookEndpoint, webhookId),
		c.headers,
		nil,
		nil,
	)
}
