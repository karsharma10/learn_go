package langchain

import (
	"context"
	"fmt"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/textsplitter"
	"log"
	"strings"
	"sync"
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

func GenerateLLMPrompts(ctx context.Context, llm LLM, prompts []string) {
	generator, err := llm.GenerateFromPrompt()
	if err != nil {
		log.Fatal(err)
	}

	promptChannel := make(chan string, 1)
	errChannel := make(chan error, 1)
	wg := sync.WaitGroup{}

	for _, p := range prompts {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			llmText, err := generator(ctx, p)
			if err != nil {
				errChannel <- err
			}
			promptChannel <- llmText
		}(p)
	}

	go func() {
		for p := range promptChannel {
			fmt.Println(p)
		}
	}()
	go func() {
		for err := range errChannel {
			fmt.Println(err)
		}
	}()

	go func() {
		wg.Wait()
		close(promptChannel)
		close(errChannel)
	}()
	wg.Wait()
}

// SummarizationChain is a langchain implementation of using chains for summarizing
func SummarizationChain(ctx context.Context, doc *string, llm llms.Model) {
	llmSummarizationChain := chains.LoadRefineSummarization(llm)
	docs, err := documentloaders.NewText(strings.NewReader(*doc)).LoadAndSplit(ctx,
		textsplitter.NewRecursiveCharacter(),
	)
	outputValues, err := chains.Call(ctx, llmSummarizationChain, map[string]any{"input_documents": docs})

	if err != nil {
		log.Fatal(err)
	}
	out := outputValues["text"].(string)
	fmt.Println(out)
}
