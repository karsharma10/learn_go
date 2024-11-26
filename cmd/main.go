package main

import (
	"context"
	"fmt"
	"github.com/karsharma10/learn_go/models/langchain"
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
}
