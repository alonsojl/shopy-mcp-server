package shopy

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const defaultTimeout = 3 * time.Second

type HTTPClient struct {
	ctx    context.Context
	method string
	url    string
	body   io.Reader
	client *http.Client
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

func (c *HTTPClient) NewRequest(ctx context.Context, method, url string, body io.Reader) *HTTPClient {
	c.ctx = ctx
	c.method = method
	c.url = url
	c.body = body
	return c
}

func (c *HTTPClient) Decode(v any) ([]byte, error) {
	req, err := http.NewRequestWithContext(c.ctx, c.method, c.url, c.body)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, fmt.Errorf("decode body: %w", err)
	}

	raw, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("json marshal: %w", err)
	}

	return raw, nil
}
