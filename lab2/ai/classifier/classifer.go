package classifier

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/genai"
)

type Classifier struct {
	Client *genai.Client
}

func (classifier Classifier) Classify(ctx context.Context, message string) (Class, error) {
	gcc := genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(
			`
			Role:You are customer messages intent classifier for a business,
			Task: Classify the customer message,
			Fallback: If you can't decide certainly choose UNKNOWN
			`,
			genai.RoleUser,
		),
		ResponseMIMEType: "application/json",
		ResponseSchema:   &schema,
		MaxOutputTokens:  1000,
	}

	resp, err := classifier.Client.Models.GenerateContent(
		ctx,
		"gemini-3.5-flash",
		genai.Text(message),
		&gcc,
	)
	if err != nil {
		return Class{}, err
	}

	fmt.Println(resp.Text())

	var c Class
	err = json.Unmarshal([]byte(resp.Text()), &c)
	if err != nil {
		return Class{}, err
	}

	return c, nil
}
