package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/katheineevse/jira-assistant-bot/internal/config"
	"log"
	"net/http"
)

type Client struct {
	BotToken string
	ChatID   string
}

func NewClient(cfg config.TgConfig) *Client {
	return &Client{
		BotToken: cfg.BotToken,
		ChatID:   cfg.ChatID,
	}
}

func (c *Client) SendMessage(message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", c.BotToken)

	data := map[string]string{
		"chat_id": c.ChatID,
		"text":    message,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Telegram API returned status: %d", resp.StatusCode)
		return fmt.Errorf("failed to send message to Telegram, status code: %d", resp.StatusCode)
	}

	log.Println("Message sent to Telegram successfully")
	return nil

}
