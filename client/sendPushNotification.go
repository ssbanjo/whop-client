package client

import (
	"github.com/ssbanjo/whop-client/internal"
)

type SendPushNotificationParams struct {
	Body      string `json:"body,omitempty" url:"-"`
	Title     string `json:"title,omitempty" url:"-"`
	Link      string `json:"link,omitempty" url:"-"`
	ProductId string `json:"product_ids,omitempty" url:"-"`
}

// Through the Whop consumer app for iOS and Android, a company can send custom tailored, time sensitive notifications to their members.
//
// https://dev.whop.com/reference/create_push_notification
func (c Client) SendPushNotification(params SendPushNotificationParams) error {

	_, err := internal.ExecuteRequest[any](
		"POST",
		sendPushNotificationEndpoint,
		c.headers,
		params,
		params,
	)

	return err
}
