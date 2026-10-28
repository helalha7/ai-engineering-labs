package classification

import (
	"errors"
	"slices"

	"google.golang.org/genai"
)

var ErrInvalidJSON = errors.New("Invalid JSON")

type Class struct {
	Intent string
}

func (c Class) validate() bool {

	allowedValues := []string{"BOOKING", "SERVICES", "BUSINESS", "PRICES", "UNKNOWN"}
	if !slices.Contains(allowedValues, c.Intent) {
		return false
	}

	return true
}

var schema = genai.Schema{
	Type: genai.TypeObject,
	Properties: map[string]*genai.Schema{
		"intent": &genai.Schema{
			Type: genai.TypeString,
			Enum: []string{
				"BOOKING",
				"SERVICES",
				"BUSINESS",
				"PRICES",
				"UNKNOWN",
			},
			Description: "The Intent of the message sent by a customer",
		},
	},
	Required: []string{"intent"},
}
