package classification

import (
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/genai"
)

type Classifier struct {
	Client *genai.Client
}

func (c *Classifier) Classify(message string) (Class, error) {

	gcc := genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(
			"you classify customer messages intent for a business",
			genai.RoleUser,
		),
		ResponseMIMEType: "application/json",
		ResponseSchema:   &schema,
		MaxOutputTokens:  400,
	}

	resp, err := c.Client.Models.GenerateContent(
		context.Background(),
		"gemini-3.5-flash",
		genai.Text(message),
		&gcc,
	)
	if err != nil {
		return Class{}, err
	}

	jsIntent := resp.Text()
	var cl Class
	err = json.Unmarshal([]byte(jsIntent), &cl)
	if err != nil {
		return Class{}, ErrInvalidJSON
	}

	if !cl.validate() {
		return Class{}, errors.New("errs")
	}

	return cl, nil
}
