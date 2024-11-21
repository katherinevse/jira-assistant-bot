package main

import (
	"github.com/katheineevse/jira-assistant-bot/internal/config"
	"github.com/katheineevse/jira-assistant-bot/internal/jira"
	"github.com/katheineevse/jira-assistant-bot/internal/telegram"
	"github.com/katheineevse/jira-assistant-bot/internal/usecase"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	log.Printf("Loaded configuration: %+v", cfg)

	jiraClient := jira.NewClient(cfg.JiraCfg)
	tgClient := telegram.NewClient(cfg.TgCfg)

	//TODO вынести логику отправки сообщений

	notify := usecase.NewNotifyUnassignedIssues(jiraClient, tgClient)

	err = notify.Execute("KAN", 50)
	if err != nil {
		log.Fatalf("Error executing notification: %v", err)
	}
}
