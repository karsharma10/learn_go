package openAI

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
)

func runOpenAI() {
	client := openai.NewClient()
	ctx := context.Background()
	question := "How big is the earth?"
	completion, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
			openai.SystemMessage(question),
		}),
		Seed:  openai.Int(1),
		Model: openai.F(openai.ChatModelGPT4o),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(completion.Choices[0].Message.Content)
}
