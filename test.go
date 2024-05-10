package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	godotenv.Load()
	var result string
	ctx := context.Background()
	// Set your Gemini API Key as a new variable in System Environment Variables
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-1.0-pro")
	resp, err := model.GenerateContent(ctx, genai.Text("Tell me the Story of Aladdin"))
	if err != nil {
		log.Fatal(err)
	}

	result = printResponse(resp)
	fmt.Println(result)

	byteArray := []byte(result)
	fmt.Println(byteArray)
	// return output, nil
	// fmt.Println(resp.Candidates.Content)
}

func printResponse(resp *genai.GenerateContentResponse) string {
	var output string
	for _, cand := range resp.Candidates {

		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				output += fmt.Sprintln(part)
				// result += part + genai.Part("\n") // Append each part with a newline
			}
			// output := []byte(cand.Content)
		}
	}
	fmt.Println("---")
	// fmt.Println(output)
	return output
}
