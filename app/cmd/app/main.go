package main

import (
	"github.com/katheineevse/jira-assistant-bot/internal/config"
	"github.com/katheineevse/jira-assistant-bot/internal/jira"
	"github.com/katheineevse/jira-assistant-bot/internal/scheduler"
	"github.com/katheineevse/jira-assistant-bot/internal/telegram"
	"github.com/katheineevse/jira-assistant-bot/internal/usecase"
	"log"
	"time"
)

//TODO добавить обработчик только рабочее время !

func main() {

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	log.Printf("Loaded configuration: %+v", cfg)

	jiraClient := jira.NewClient(cfg.JiraCfg)
	tgClient := telegram.NewClient(cfg.TgCfg)

	//TODO вынести логику отправки сообщений
	//gocron.NewScheduler

	notifyUseCase := usecase.NewNotifyUnassignedIssues(jiraClient, tgClient, cfg.ProgCfg)

	scheduler := scheduler.New(notifyUseCase, cfg.SchedulerCfg)
	scheduler.Start()

	time.Sleep(1 * time.Minute) //TODO del me
	//err = notify.Execute(50)
	//if err != nil {
	//	log.Fatalf("Error executing notification: %v", err)
	//}
}
