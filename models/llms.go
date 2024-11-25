package models

import (
	"context"
	"github.com/tmc/langchaingo/llms/ollama"
)

type LLM interface {
	GenerateEmbedding() (func(ctx context.Context, texts []string) ([][]float32, error), error)
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
