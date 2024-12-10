package main

import (
	"context"
	"fmt"
	"log"
	"os"

	aai "github.com/AssemblyAI/assemblyai-go-sdk"
)

func main() {
	apiKey := "6e179305765c4d56a199c45676f7cc65"
	// audioURL := "https://assembly.ai/news.mp4"

	file, err := os.Open("./audioFiles/harvard.wav")
	if err != nil {
		log.Fatal("Error recievd in reading audio file - ", err.Error())
	}
	defer file.Close()

	client := aai.NewClient(apiKey)

	transcript, err := client.Transcripts.TranscribeFromReader(context.Background(), file, &aai.TranscriptOptionalParams{
		SentimentAnalysis: aai.Bool(true),
	})

	if err != nil {
		fmt.Println("Something bad happened:", err)
		os.Exit(1)
	}

	for _, res := range transcript.SentimentAnalysisResults {
		fmt.Println(aai.ToString(res.Text))
		fmt.Println(res.Sentiment)
		fmt.Println(aai.ToFloat64(res.Confidence))
		fmt.Println("Error recieved - ", err)
	}

	fmt.Println(*transcript.Text)
}
