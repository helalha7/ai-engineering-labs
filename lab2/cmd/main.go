package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

var tool = genai.Tool{
	FunctionDeclarations: []*genai.FunctionDeclaration{
		&genai.FunctionDeclaration{
			Name:        "business_informations",
			Description: "this function used to retrieve the business informations",
		},
		&genai.FunctionDeclaration{
			Name:        "services",
			Description: "this function used to retrieve informations about the services that the business provides",
		},
	},
}

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := genai.NewClient(
		ctx,
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}

	gcc := genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(
			`
			you are customer support for a business
			you should generate response messages for the customer which is messaging through whatsapp
			Don't just spit informations in the customer face only answer his specific question
			`,
			genai.RoleUser,
		),
		Tools: []*genai.Tool{
			&tool,
		},
	}

	resp, err := client.Models.GenerateContent(
		ctx,
		"gemini-3.5-flash",
		genai.Text("Hey who are you and what services do you offfer ?"),
		&gcc,
	)
	if err != nil {
		log.Fatal(err)
	}

	calls := resp.FunctionCalls()

	if len(calls) == 0 {
		fmt.Println(resp.Text())
		return
	}

	var results []*genai.Part

	for _, call := range calls {
		var data map[string]any

		switch call.Name {
		case "business_informations":
			data = map[string]any{
				"name":     "Example Barbershop",
				"location": "Haifa",
				"hours":    "09:00-18:00",
			}

		case "services":
			data = map[string]any{
				"services": []string{
					"Haircut",
					"Beard trimming",
				},
			}

		default:
			log.Fatalf("unknown function: %s", call.Name)
		}

		results = append(results, &genai.Part{
			FunctionResponse: &genai.FunctionResponse{
				ID:       call.ID,
				Name:     call.Name,
				Response: data,
			},
		})
	}

	// Include the original user message and Gemini's complete
	// response, including its function calls.
	history := []*genai.Content{
		genai.NewContentFromText(
			"how much a haricut costs?",
			genai.RoleUser,
		),
		resp.Candidates[0].Content,
		genai.NewContentFromParts(results, genai.RoleUser),
	}

	finalResp, err := client.Models.GenerateContent(
		ctx,
		"gemini-3.5-flash",
		history,
		&gcc,
	)
	if err != nil {
		log.Fatal(err)
	}

	if len(finalResp.FunctionCalls()) > 0 {
		log.Fatal("model requested additional function calls")
	}

	fmt.Println(finalResp.Text())

}
