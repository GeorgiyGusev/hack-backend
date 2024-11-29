package gateway

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Gateway представляет структуру для работы с API нейросети.
type Gateway struct {
	client http.Client
}

// NewGateway создает новый экземпляр Gateway.
func NewGateway() *Gateway {
	return &Gateway{client: http.Client{}}
}

// Message представляет структуру запроса для API нейросети.
type Message struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// Response представляет структуру ответа от API нейросети.
type Response struct {
	Model    string `json:"model"`
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

// SendMessage отправляет сообщение нейросети и возвращает ответ.
func (g *Gateway) SendMessage(ctx context.Context, message string) (string, error) {
	// Создаем тело запроса
	body := Message{
		Model:  "llama3",
		Prompt: fmt.Sprintf("INSTRUCTIONS FOR MODEL: You are a smart assistant for the MininCode education system in the Nizhny Novgorod region. You must help and advise people and children to find IT objects and companies in Nizhny Novgorod and provide advice on the IT sphere. | MESSAGE FROM USER -> %s", message),
		Stream: false,
	}

	// Сериализуем тело запроса в JSON
	marshaled, err := json.Marshal(body)
	if err != nil {
		return "", fmt.Errorf("failed to marshal message: %w", err)
	}

	// Создаем HTTP-запрос с контекстом
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://kowlad123321456654.tplinkdns.com:11434/api/generate", bytes.NewReader(marshaled))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Отправляем запрос
	resp, err := g.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected response status: %d, body: %s", resp.StatusCode, string(body))
	}

	// Читаем тело ответа
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	// Парсим JSON-ответ
	var response Response
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	// Проверяем флаг завершения и возвращаем ответ
	if !response.Done {
		return "", fmt.Errorf("response not marked as done")
	}

	return response.Response, nil
}
