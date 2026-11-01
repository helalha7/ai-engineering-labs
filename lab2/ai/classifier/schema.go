package classifier

import (
	"slices"

	"google.golang.org/genai"
)

type Class struct {
	Intent string `json:"intent"`
}

var AllowedValues = []string{"BOOKING", "SERVICES", "PRICES", "UNKNOWN", "GENERAL_QA"}

func (c Class) Validate() bool {
	return slices.Contains(AllowedValues, c.Intent)
}

var schema = genai.Schema{
	Type: genai.TypeObject,
	Properties: map[string]*genai.Schema{
		"intent": &genai.Schema{
			Type:        genai.TypeString,
			Enum:        AllowedValues,
			Description: "The intent of the customer message you should specify",
		},
	},
	Required: []string{"intent"},
}
