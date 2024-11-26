package usecase

import (
	"fmt"
	"github.com/katheineevse/jira-assistant-bot/internal/config"
	"github.com/katheineevse/jira-assistant-bot/internal/dto"
	"github.com/katheineevse/jira-assistant-bot/internal/jira"
	"github.com/katheineevse/jira-assistant-bot/internal/telegram"

	"log"
)

type NotifyUnassignedIssues struct {
	JiraClient *jira.Client
	TgClient   *telegram.Client
	ProjConfig config.ProjectConfig
}

func NewNotifyUnassignedIssues(jiraClient *jira.Client, tgClient *telegram.Client, projConfig config.ProjectConfig) *NotifyUnassignedIssues {
	return &NotifyUnassignedIssues{
		JiraClient: jiraClient,
		TgClient:   tgClient,
		ProjConfig: projConfig,
	}
}

func (n *NotifyUnassignedIssues) Execute(sinceHours int) error {
	unassignIssues, err := n.JiraClient.GetUnassignedIssues(n.ProjConfig.Key, sinceHours)
	if err != nil {
		log.Fatalf("Failed to get issues: %v", err)

	}

	message := n.formatMessage(unassignIssues)

	err = n.TgClient.SendMessage(message)
	if err != nil {
		return err
	}

	return nil

}

// TODO fix summary printing
func (n *NotifyUnassignedIssues) formatMessage(issues []dto.Issue) string {
	var message string
	message = "⚠️" + "JIRA TASK WARNING\n" + "Проблема: длительное время нет исполнителя по задаче \n" + "Решение: пожалуйста, назначьте исполнителя:\n"
	for _, issue := range issues {
		message += fmt.Sprintf(" %s + %s)\n", issue.Summary, issue.Link)
	}
	return message
}
