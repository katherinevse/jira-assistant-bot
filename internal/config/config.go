package config

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

func LoadConfig() *Config {
	return &Config{
		JiraCfg:      JiraConfig{},
		TgCfg:        TgConfig{},
		ProgCfg:      ProjectConfig{},
		SchedulerCfg: SchedulerConfig{},
	}
}
