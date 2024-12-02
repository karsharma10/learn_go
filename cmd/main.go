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
	//source: from https://kdigo.org/wp-content/uploads/2024/03/KDIGO-2024-CKD-Guideline.pdf : (KDIGO Guidelines)
	doc := `The Kidney Disease: Improving Global Outcomes (KDIGO)
			organization was established in 2003 with the mission to
			improve the care and outcomes of people living with kidney
			disease worldwide. The development and implementation of
			global clinical practice guidelines is central to the many activities of KDIGO to fulfill its mission. Twenty years later, we
			are excited to present this update of the KDIGO Clinical
			Practice Guideline for the Evaluation and Management of
			Chronic Kidney Disease (CKD) to complement the existing
			12 guidelines that address various other facets of kidney
			disease management.
	
	
			Our aspiration is that the KDIGO CKD Guideline serves as
			a comprehensive reference for evidence-based practices, offering clear and valuable guidance for the optimal diagnosis
			and treatment of CKD. The updated guideline is the result of
			a rigorous process, extensively detailed in the KDIGO
			Methods Manual. To promote objectivity and transparency,
			we screen Guideline Co-Chairs and Work Group members
			(which include clinicians, researchers, and patients) for conflicts of interest. Over a span of 2–3 years, these individuals
			volunteer their time, starting with the creation of a Scope of
			Work that undergoes an open public review to engage all
			stakeholders. This document is then adapted into a Request
			for Proposal, which is used to enlist an independent Evidence
			Review Team.
	
	
			The Evidence Review Team conducts a systematic review of
			existing literature, extracting studies with appropriate design
			and outcomes deemed important by both people with CKD
			and clinicians. All work is meticulously graded on study
			quality and potential bias, forming the basis for quantifying
			the overall certainty of the evidence using the “Grading of
			Recommendations Assessment, Development, and Evaluation” (GRADE) approach. The penultimate version of the
			guideline also undergoes public review to capture additional
			perspectives. Thus, guidelines are the result of a rigorous and
			objective assessment of available evidence, enriched by the
			collective expertise of healthcare providers, researchers, and
			patients alike. Guideline statements (“We recommend” or
			“We suggest”) reflect clinical questions that were addressed by
			the evidence reviews from the Evidence Review Team. Practice
			points provide guidance on clinical questions that were not,
			and largely could not be, studied by the Evidence Review
			Team.`

	langchain.SummarizationChain(ctx, &doc, newOllama)

	fmt.Println(config.MdHashing("hello"))

	fmt.Println(config.ShaHashing("hello"))
}
