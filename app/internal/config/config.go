package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
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

func LoadConfig() (*Config, error) {
	var config Config

	configPath := filepath.Join("config", "config.yaml")

	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	// Разбор YAML в структуру Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}
