package openAI

import (
	"context"
	"fmt"
	"github.com/openai/openai-go"
	"io"
	"os"
	"time"
)

type uploadFineTuningDataset interface {
	uploadDocument(file string) error
}

type openAIFineTuning struct {
	Client *openai.Client
}

func (o *openAIFineTuning) uploadDocument(file string) error {
	ctx := context.Background()

	if o.Client == nil {
		return fmt.Errorf("need to intialize openAI client")
	}

	fmt.Println("==> Uploading file...")
	data, err := os.Open(file)
	uploadFile, err := o.Client.Files.New(ctx, openai.FileNewParams{
		File:    openai.F[io.Reader](data),
		Purpose: openai.F(openai.FilePurposeFineTune),
	})
	if err != nil {
		return fmt.Errorf("error with openAI client")
	}
	fmt.Printf("Uploaded file with ID: %s\n", uploadFile.ID)
	fmt.Println("Waiting for file to be processed")

	for {
		uploadFile, err = o.Client.Files.Get(ctx, uploadFile.ID)
		if err != nil {
			panic(err)
		}
		fmt.Printf("File status: %s\n", uploadFile.Status)
		if uploadFile.Status == "processed" {
			break
		}
		time.Sleep(time.Second)
	}

	return nil
}
