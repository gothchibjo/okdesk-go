# Okdesk Go Client

This is a Go client for the Okdesk API.

## Installation

```sh
go get github.com/gothchibjo/okdesk-go
```

## Usage

```go
package main

import (
    "fmt"
    "log"
    "time"

    "github.com/gothchibjo/okdesk-go/okdesk"
)

func main() {
    // Create client with custom timeout
    client := okdesk.NewClient("https://your-company.okdesk.ru", "your-api-key", &okdesk.ClientOptions{
        Timeout: 10 * time.Second,
    })

    // Filters for searching issues (example: only open)
    filters := map[string]string{
        "status": "open",
    }

    // Get list of issues
    issues, err := client.GetIssuesList(filters)
    if (err != nil) {
        log.Fatalf("Failed to get issues list: %v", err)
    }

    // Print list of issues
    for _, issue := range issues {
        fmt.Printf("ID: %d, Title: %s, Status: %s\n", issue.ID, issue.Title, issue.Status)
    }

    // Example call to GetIssue
    issueID := 13
    issue, err := client.GetIssue(issueID)
    if (err != nil) {
        log.Fatalf("Failed to get issue: %v", err)
    }
    fmt.Printf("Issue: %+v\n", issue)
}
```
