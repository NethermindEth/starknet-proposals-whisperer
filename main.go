package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/v41/github"
)

var (
	slackWebhookURL = os.Getenv("SLACK_WEBHOOK_URL")
	repoOwner       = os.Getenv("REPO_OWNER")
	repoName        = os.Getenv("REPO_NAME")
	threshold       = os.Getenv("THRESHOLD_DAYS") // Not used anymore
)

type EventBridgeEvent struct {
	Version    string                 `json:"version"`
	ID         string                 `json:"id"`
	DetailType string                 `json:"detail-type"`
	Source     string                 `json:"source"`
	Account    string                 `json:"account"`
	Time       string                 `json:"time"`
	Region     string                 `json:"region"`
	Resources  []string               `json:"resources"`
	Detail     map[string]interface{} `json:"detail"`
}

func HandleRequest(ctx context.Context, in EventBridgeEvent) {
	client := github.NewClient(nil) // No authentication
	fmt.Println(in)

	opt := &github.PullRequestListOptions{
		State: "open",
		ListOptions: github.ListOptions{
			PerPage: 10,
		},
	}

	prs, _, err := client.PullRequests.List(ctx, repoOwner, repoName, opt)
	if err != nil {
		fmt.Println("Error fetching PRs:", err)
		return
	}

	// Parse EventsBridge time
	now, err := time.Parse(time.RFC3339, in.Time)

	// The Lambda is triggered every day at 10am UTC and 5pm UTC
	// PRs should be checked between the time the Lambda is triggered
	// and the time it was triggered the last time.
	type threshold struct {
		time              time.Time
		thresholdDuration time.Duration
	}
	times := []threshold{
		{
			time.Date(now.Year(), now.Month(), now.Day(), 10, 0, 0, 0, time.UTC),
			17 * time.Hour, // 17 hours after 5pm UTC
		},
		{
			time.Date(now.Year(), now.Month(), now.Day(), 17, 0, 0, 0, time.UTC),
			7 * time.Hour, // 7 hours after 10am UTC
		},
	}
	var thresholdDuration time.Duration
	//thresholdDuration := time.Duration(days) * 24 * time.Hour
	for _, t := range times {
		if now.Before(t.time) || now.Equal(t.time) {
			thresholdDuration = t.thresholdDuration
			break
		}
	}

	for _, pr := range prs {
		prCreatedAt := pr.GetCreatedAt()
		if prCreatedAt.After(now.Add(-thresholdDuration)) {
			sendSlackNotification(pr)
		}
	}

}

func sendSlackNotification(pr *github.PullRequest) {
	msg := fmt.Sprintf(
		":github: *New Pull Request Submitted!* :tada:\n"+
			"*Title:* %s\n"+
			"*Author:* %s\n"+
			"*URL:* <%s|View on GitHub>\n"+
			"*Created At:* %s",
		pr.GetTitle(),
		pr.GetUser().GetLogin(),
		pr.GetHTMLURL(),
		pr.GetCreatedAt().Format(time.RFC822),
	)

	payload := map[string]string{
		"text": msg,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	resp, err := http.Post(slackWebhookURL, "application/json", strings.NewReader(string(jsonData)))
	if err != nil {
		fmt.Println("Error sending notification:", err)
		return
	}
	defer resp.Body.Close()
}

func main() {
	lambda.Start(HandleRequest)
	//HandleRequest()
	fmt.Println("Done")
}
