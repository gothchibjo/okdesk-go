package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Client represents the Okdesk API client.
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new Okdesk API client.
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// request performs an HTTP request and decodes the response into the given output.
func (c *Client) request(ctx context.Context, method, endpoint string, queryParams map[string]string, output interface{}) error {
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+endpoint, nil)
	if err != nil {
		return err
	}

	// Set API key as query parameter
	q := req.URL.Query()
	q.Add("api_token", c.apiKey)
	for key, value := range queryParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status code: " + resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(output)
}
