package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/v41/github"
)

var (
	slackWebhookURL = os.Getenv("SLACK_WEBHOOK_URL")
	repoOwner       = os.Getenv("REPO_OWNER")
	repoName        = os.Getenv("REPO_NAME")
	threshold       = os.Getenv("THRESHOLD_DAYS")
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

func HandleRequest(ctx context.Context, event EventBridgeEvent) {
	client := github.NewClient(nil) // No authentication

	now := time.Now()

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

	days, err := strconv.Atoi(threshold)
	if err != nil {
		fmt.Println("Error parsing DAYS_DURATION:", err)
		return
	}

	thresholdDuration := time.Duration(days) * 24 * time.Hour

	for _, pr := range prs {
		prCreatedAt := pr.GetCreatedAt()
		if prCreatedAt.After(now.Add(-thresholdDuration)) {
			sendSlackNotification(pr)
		}
	}

	eventJSON, _ := json.MarshalIndent(event, "", "  ")
	fmt.Println(string(eventJSON))
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
