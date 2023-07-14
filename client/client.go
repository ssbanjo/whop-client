package client

import "net/http"

type Client struct {
	headers http.Header
}

func NewClient(apiKey string) *Client {

	headers := http.Header{
		"Authorization": {"Bearer " + apiKey},
		"accept":        {"application/json"},
		"content-type":  {"application/json"},
	}

	return &Client{
		headers: headers,
	}

}
