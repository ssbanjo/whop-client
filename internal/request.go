package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	net_url "net/url"

	"github.com/google/go-querystring/query"
	"golang.org/x/exp/slices"
)

func prepareRequest(method string, url string, headers http.Header, params any, body any) (*http.Request, error) {

	u, _ := net_url.Parse(url)
	var rawBody []byte = nil

	if params != nil {

		newQuery, _ := query.Values(params)
		u.RawQuery = newQuery.Encode()
	}

	if body != nil {

		rawBody, _ = json.Marshal(body)
	}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(rawBody))

	if err != nil {
		return nil, err
	}

	req.Header = headers

	return req, nil

}

func sendRequest(req *http.Request, resp any) error {

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	if !slices.Contains([]int{200, 201, 204}, res.StatusCode) {
		return fmt.Errorf("request failed with status code %d: %s", res.StatusCode, string(body))
	}

	return json.Unmarshal(body, resp)
}

func ExecuteRequest[T any](method string, url string, headers http.Header, query any, body any) (*T, error) {

	var res *T

	req, err := prepareRequest(method, url, headers, query, body)

	if err != nil {
		return nil, fmt.Errorf("failed to prepare request: %v", err)
	}

	err = sendRequest(req, &res)

	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	return res, nil

}
