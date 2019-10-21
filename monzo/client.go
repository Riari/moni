package monzo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// Client is the HTTP client used for API requests.
type Client struct {
	Client      *http.Client
	BaseURL     *url.URL
	AccessToken string

	Auth *AuthService
}

func (c *Client) get(endpoint string) *http.Response {
	req, err := c.newRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := c.send(req)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

func (c *Client) newRequest(method, path string, body map[string]interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	return req, nil
}

func (c *Client) send(req *http.Request) (*http.Response, error) {
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
