package client

import (
	"fmt"

	"github.com/ssbanjo/whop-client/internal"
	"github.com/ssbanjo/whop-client/types"
)

type RetriveWebhookParams struct {
	Expand []string `url:"expand,omitempty" del:","`
}

type RetriveWebhookResponse types.Webhook

// Retrieves a specific Review, using its unique ID.
//
// https://dev.whop.com/reference/retrieve_review
func (c Client) RetriveWebhook(webhookId string, params RetriveWebhookParams) (*RetriveWebhookResponse, error) {

	return internal.ExecuteRequest[RetriveWebhookResponse](
		"GET",
		fmt.Sprintf(retreiveWebhookEndpoint, webhookId),
		c.headers,
		params,
		nil,
	)
}
