package main

import (
	"context"
	"fmt"
	"github.com/karsharma10/learn_go/config"
	"github.com/karsharma10/learn_go/models/langchain"
	langchainOllama "github.com/tmc/langchaingo/llms/ollama"
	"log"
)

func main() {
	ctx := context.Background()
	ollama := langchain.NewOllama("llama3.2")
	ollamaEmbedding, err := ollama.GenerateEmbedding()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	text := []string{"hello"}
	embedding, err := ollamaEmbedding(ctx, text)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Println("Embedding: ", embedding)

	ollamaPrompt, err := ollama.GenerateFromPrompt()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	message, err := ollamaPrompt(ctx, "Hello, How are you?")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println(message)

	langchain.GenerateLLMPrompts(ctx, ollama, []string{"Hello, How are you?", "What is the capital of India?"})

	fun := config.WithOpenAI()

	set := config.Configs{}

	fun(&set)

	fmt.Println(set.OpenAIKey)
	newOllama, _ := langchainOllama.New(langchainOllama.WithModel("llama3.2"))
	doc := `AI applications are summarizing articles, writing stories and 
	engaging in long conversations — and large language models are doing 
	the heavy lifting.
	
	A large language model, or LLM, is a deep learning model that can 
	understand, learn, summarize, translate, predict, and generate text and other 
	content based on knowledge gained from massive datasets.
	
	Large language models - successful applications of 
	transformer models. They aren’t just for teaching AIs human languages, 
	but for understanding proteins, writing software code, and much, much more.
	
	In addition to accelerating natural language processing applications — 
	like translation, chatbots, and AI assistants — large language models are 
	used in healthcare, software development, and use cases in many other fields.`

	langchain.SummarizationChain(ctx, &doc, newOllama)
}
