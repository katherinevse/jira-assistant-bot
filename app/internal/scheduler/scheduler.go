package scheduler

import (
	"github.com/go-co-op/gocron"
	"github.com/katheineevse/jira-assistant-bot/internal/config"
	"github.com/katheineevse/jira-assistant-bot/internal/usecase"
	"log"
	"time"
)

type Scheduler struct {
	UseCase      *usecase.NotifyUnassignedIssues
	SchedulerCfg config.SchedulerConfig
}

func New(notifyUseCase *usecase.NotifyUnassignedIssues, schedulerCfg config.SchedulerConfig) *Scheduler {
	return &Scheduler{
		UseCase:      notifyUseCase,
		SchedulerCfg: schedulerCfg,
	}
}

func (s *Scheduler) Start() {
	scheduler := gocron.NewScheduler(time.UTC)

	interval, err := time.ParseDuration(s.SchedulerCfg.Interval)
	if err != nil {
		log.Fatalf("Ошибка преобразования интервала: %v", err)
	}

	_, err = scheduler.Every(interval).Do(s.executeTask)
	if err != nil {
		log.Fatalf("Error scheduling task: %v", err)
	}

	scheduler.StartAsync()
}

func (s *Scheduler) executeTask() {
	log.Println("Checking for unassigned issues...")
	err := s.UseCase.Execute(50)
	if err != nil {
		log.Printf("Error executing notification: %v", err)
	}
}
