package models

import (
	"context"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
)

type LLM interface {
	GenerateEmbedding() (func(ctx context.Context, texts []string) ([][]float32, error), error)
	GenerateFromPrompt() (func(ctx context.Context, text string) (string, error), error)
}

type OllamaModel struct {
	model string
}

var _ LLM = (*OllamaModel)(nil)

func NewOllama(model string) *OllamaModel {
	return &OllamaModel{
		model: model,
	}
}

// GenerateEmbedding returns a function that takes a context and a slice of strings and returns a slice of slices of float32 and an error
func (o *OllamaModel) GenerateEmbedding() (func(ctx context.Context, texts []string) ([][]float32, error), error) {
	llm, err := ollama.New(ollama.WithModel(o.model), ollama.WithRunnerEmbeddingOnly(true))
	if err != nil {
		return nil, err
	}
	embedding := func(ctx context.Context, texts []string) ([][]float32, error) {
		return llm.CreateEmbedding(ctx, texts)
	}
	return embedding, nil
}

func (o *OllamaModel) GenerateFromPrompt() (func(ctx context.Context, text string) (string, error), error) {
	llm, err := ollama.New(ollama.WithModel(o.model))
	if err != nil {
		return nil, err
	}
	generatedText := func(ctx context.Context, text string) (string, error) {
		return llm.Call(ctx, text)
	}
	return generatedText, nil
}

type OpenAIModel struct {
	model          string
	embeddingModel string
}

var _ LLM = (*OpenAIModel)(nil)

func NewOpenAI(model string, embeddingModel string) *OpenAIModel {
	return &OpenAIModel{
		model:          model,
		embeddingModel: embeddingModel,
	}
}

func (o *OpenAIModel) GenerateEmbedding() (func(ctx context.Context, texts []string) ([][]float32, error), error) {
	llm, err := openai.New(openai.WithEmbeddingModel(o.embeddingModel))
	if err != nil {
		return nil, err
	}
	embedding := func(ctx context.Context, texts []string) ([][]float32, error) {
		return llm.CreateEmbedding(ctx, texts)
	}
	return embedding, nil
}

func (o *OpenAIModel) GenerateFromPrompt() (func(ctx context.Context, text string) (string, error), error) {
	llm, err := openai.New(openai.WithModel(o.model))
	if err != nil {
		return nil, err
	}
	generatedText := func(ctx context.Context, text string) (string, error) {
		return llm.Call(ctx, text)
	}
	return generatedText, nil
}
