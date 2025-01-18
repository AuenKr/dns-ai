package llm

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func GenerateContent(query string) (string, error) {
	ctx := context.Background()

	apiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if !ok {
		return "", errors.New("API key required")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.5-flash-8b")

	model.SetTemperature(1)
	model.SetTopK(40)
	model.SetTopP(0.95)
	model.SetMaxOutputTokens(8192)
	model.ResponseMIMEType = "text/plain"
	model.SystemInstruction = &genai.Content{
		Parts: []genai.Part{genai.Text("You are DNS Assistant. Respond to DNS TXT queries with concise answers, ensuring responses do not exceed 255 characters. For math queries, provide only the result without additional details or context.")},
	}

	resp, err := model.GenerateContent(ctx, genai.Text(query))
	if err != nil {
		return "", err
	}

	var response string = ""
	for _, part := range resp.Candidates[0].Content.Parts {
		switch p := part.(type) {
		case genai.Text:
			response += string(p)
		}
	}

	response = strings.TrimSpace(response)

	return response, nil
}
