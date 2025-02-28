package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gothchibjo/okdesk-go/okdesk/models"
)

// GetIssue retrieves a single issue by ID.
func (c *Client) GetIssue(ctx context.Context, issueID int) (*models.Issue, error) {
	var issue models.Issue
	err := c.request(ctx, http.MethodGet, "/api/v1/issues/"+strconv.Itoa(issueID), nil, &issue)
	if err != nil {
		return nil, err
	}
	return &issue, nil
}

// GetIssuesList retrieves a list of issues with optional filters.
func (c *Client) GetIssuesList(ctx context.Context, filters map[string]string) ([]models.Issue, error) {
	var issues []models.Issue
	err := c.request(ctx, http.MethodGet, "/api/v1/issues/list", filters, &issues)
	if err != nil {
		return nil, err
	}
	return issues, nil
}
