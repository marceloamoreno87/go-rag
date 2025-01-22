package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type OpenAIService struct {
	apiKey string
	client *http.Client
	prompt string
}

func NewOpenAIService(apiKey string) *OpenAIService {
	service := &OpenAIService{
		apiKey: apiKey,
		client: &http.Client{},
	}

	service.prompt = `Você é um especialista em gestão de academias de musculação e bem-estar. 
	Por favor, forneça conselhos detalhados sobre como melhorar a gestão de uma academia. 
	Inclua tópicos como:
	- Gestão de membros e retenção
	- Manutenção e atualização de equipamentos
	- Estratégias de marketing eficazes
	- Desenvolvimento de programas de treinamento e bem-estar
	- Criação de um ambiente motivador e acolhedor para os clientes
	- Tendências atuais no setor de academias
	Responda à seguinte pergunta: `

	return service

}

func (s *OpenAIService) GenerateRAGContext(ctx context.Context, prompt string) (string, error) {
	url := "https://api.openai.com/v1/engines/davinci-codex/completions"
	requestBody, err := json.Marshal(map[string]interface{}{
		"prompt":     prompt,
		"max_tokens": 150,
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.apiKey)

	resp, err := s.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return "", fmt.Errorf("failed to decode response body: %w", err)
	}

	choices, ok := responseBody["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("no choices found in response")
	}

	choice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid choice format")
	}

	text, ok := choice["text"].(string)
	if !ok {
		return "", fmt.Errorf("invalid text format")
	}

	return text, nil
}
