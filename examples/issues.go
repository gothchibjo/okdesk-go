package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gothchibjo/okdesk-go/okdesk"
)

func main() {
	// Create client with custom timeout
	client := okdesk.NewClient("https://comfort.okdesk.ru", "0e19daa3eeb3e35f11e4464b71cb851da7317f88", &okdesk.ClientOptions{
		Timeout: 10 * time.Second,
	})

	// Filters for searching issues (example: only open)
	//
	// All filter parameters see in the [Okdeks API documentation](https://apidocs.okdesk.ru/apidoc/#!poluchenie-detalizirovannogo-spiska-po-parametram-poluchenie-detalizirovannogo-spiska-po-parametram)
	filters := map[string]string{
		"status_codes[]": "opened",
	}

	// Get list of issues
	issues, err := client.GetIssuesList(filters)
	if err != nil {
		log.Fatalf("Failed to get issues list: %v", err)
	}

	// Print list of issues
	for _, issue := range issues {
		fmt.Printf("ID: %d, Title: %s, Status: %s\n", issue.ID, issue.Title, issue.Status)
	}

	// Example call to GetIssue
	issueID := 13
	issue, err := client.GetIssue(issueID)
	if err != nil {
		log.Fatalf("Failed to get issue: %v", err)
	}
	fmt.Printf("Issue: %+v\n", issue)
}
