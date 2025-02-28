package okdesk

import (
	"context"
	"time"

	"github.com/gothchibjo/okdesk-go/internal/api"
	"github.com/gothchibjo/okdesk-go/okdesk/models"
)

// Client is the main structure for interacting with the Okdesk API.
type Client struct {
	apiClient *api.Client
	timeout   time.Duration
}

// ClientOptions holds optional parameters for creating a Client.
type ClientOptions struct {
	Timeout time.Duration
}

// NewClient creates a new Okdesk API client with optional parameters.
func NewClient(baseURL, apiKey string, opts *ClientOptions) *Client {
	timeout := 30 * time.Second // default timeout
	if opts != nil && opts.Timeout != 0 {
		timeout = opts.Timeout
	}
	return &Client{
		apiClient: api.NewClient(baseURL, apiKey),
		timeout:   timeout,
	}
}

// GetIssue retrieves a single issue by ID.
func (c *Client) GetIssue(issueID int) (*models.Issue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	return c.apiClient.GetIssue(ctx, issueID)
}

// GetIssuesList retrieves a list of issues with optional filters.
func (c *Client) GetIssuesList(filters map[string]string) ([]models.Issue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()
	return c.apiClient.GetIssuesList(ctx, filters)
}
