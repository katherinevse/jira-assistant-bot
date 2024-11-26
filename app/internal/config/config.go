package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	JiraCfg      JiraConfig      `yaml:"jira"`
	TgCfg        TgConfig        `yaml:"telegram"`
	ProgCfg      ProjectConfig   `yaml:"project"`
	SchedulerCfg SchedulerConfig `yaml:"scheduler"`
}

type JiraConfig struct {
	BaseURL  string `yaml:"base_url"`
	Username string `yaml:"username"`
	APIToken string `yaml:"api_token"`
}

type TgConfig struct {
	BotToken string `yaml:"bot_token"`
	ChatID   string `yaml:"chat_id"`
}

type ProjectConfig struct {
	Key string `yaml:"key"`
}

type SchedulerConfig struct {
	Interval string `yaml:"interval"`
}

func New() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки .env файла: %v", err)
	} else {
		fmt.Println("Файл .env успешно загружен")
	}

	return &Config{
		JiraCfg: JiraConfig{
			BaseURL:  os.Getenv("JIRA_BASE_URL"),
			Username: os.Getenv("JIRA_USERNAME"),
			APIToken: os.Getenv("JIRA_API_TOKEN"),
		},
		TgCfg: TgConfig{
			BotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
			ChatID:   os.Getenv("TELEGRAM_CHAT_ID"),
		},
		ProgCfg: ProjectConfig{
			Key: os.Getenv("PROJECT_KEY"),
		},
		SchedulerCfg: SchedulerConfig{
			Interval: os.Getenv("SCHEDULER_INTERVAL"),
		},
	}, nil
}
