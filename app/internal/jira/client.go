package jira

import (
	"encoding/json"
	"fmt"
	"github.com/katheineevse/jira-assistant-bot/internal/config"
	"github.com/katheineevse/jira-assistant-bot/internal/dto"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	BaseURL  string
	Username string
	APIToken string
}

func NewClient(cfg config.JiraConfig) *Client {
	return &Client{
		BaseURL:  cfg.BaseURL,
		Username: cfg.Username,
		APIToken: cfg.APIToken,
	}
}

//TODO перенести в енв данные и поменять публичное апи + считывать данные с конфига

// https://katherinevse.atlassian.net/rest/api/2/search?jql=project=KAN+AND+status%3D%22To+Do%22+AND+assignee+IS+EMPTY
func (c *Client) GetUnassignedIssues(projectKey string, sinceHours int) ([]dto.Issue, error) {
	//TODO брать из конфига
	baseURL := "https://katherinevse.atlassian.net/rest/api/2/search"
	query := "project=KAN AND status='To Do' AND assignee IS EMPTY"

	escapedQuery := url.QueryEscape(query)
	url := fmt.Sprintf("%s?jql=%s", baseURL, escapedQuery)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.SetBasicAuth(c.Username, c.APIToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	fmt.Println("ЭТО БОТ ТОКЕН\n", c.APIToken)
	//TODO убрать такую инициализацию клиента
	client := &http.Client{
		Timeout: 50 * time.Second,
	}

	log.Printf("Sending request to URL: %s", url)
	log.Printf("Request headers: %+v", req.Header)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, body)
	}

	//TODO models
	var result struct {
		Issues []dto.Issue `json:"issues"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Issues, nil
}
