package graphql

import "strings"

type Errors []Error

func (errs Errors) Error() string {
	msgs := make([]string, len(errs))
	for i, err := range errs {
		msgs[i] = err.Message
	}
	return strings.Join(msgs, "; ")
}

type Error struct {
	ErrorType string `json:"errorType,omitempty"`
	Message   string `json:"message"`

	Path       []interface{}          `json:"path,omitempty"`
	Locations  []Location             `json:"locations,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

func (err Error) Error() string {
	return err.Message
}

type Location struct {
	Line   int `json:"line,omitempty"`
	Column int `json:"column,omitempty"`
}
