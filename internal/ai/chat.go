package ai

import (
	"context"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func GetAIResponse(question string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are Dotlight, an AI-powered strategy assistant designed to help teams align around product goals, improve decision-making, and clarify next steps. Keep responses concise, clear, and helpful.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: question,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
