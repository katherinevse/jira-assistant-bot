package main

import (
	"fmt"
	"github.com/katheineevse/jira-assistant-bot/internal/config"
	"github.com/katheineevse/jira-assistant-bot/internal/jira"
	"log"
)

var cfg *config.Config

func init() {
	cfg = config.LoadConfig()
}
func main() {
	jiraClient := jira.NewClient(cfg.JiraCfg)

	//Unassign
	unassignIssues, err := jiraClient.GetUnassignedIssues("KAN", 50)
	if err != nil {
		log.Fatalf("Failed to get issues: %v", err)
	}

	for _, issue := range unassignIssues {
		fmt.Printf("Issue: %s - %s\n", issue.Key, issue.URL)
	}

}
