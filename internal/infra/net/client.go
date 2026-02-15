package net

import (
	"context"
	"io"
	"net/http"
	"strings"
)

// Client implements domain.HTTPClient using net/http.
type Client struct {
	client *http.Client
}

// New creates a new HTTP client.
func New() *Client {
	return &Client{client: &http.Client{}}
}

func (c *Client) Get(ctx context.Context, url string, headers map[string]string) (int, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, "", err
	}
	setHeaders(req, headers)
	return c.do(req)
}

func (c *Client) Post(ctx context.Context, url string, body string, headers map[string]string) (int, string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return 0, "", err
	}
	if _, ok := headers["Content-Type"]; !ok {
		req.Header.Set("Content-Type", "application/json")
	}
	setHeaders(req, headers)
	return c.do(req)
}

func (c *Client) do(req *http.Request) (int, string, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return 0, "", err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", err
	}
	return resp.StatusCode, string(data), nil
}

func setHeaders(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}
}
