package graphql

import "net/http"

type Request struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`

	header http.Header
}

func (r *Request) SetVar(key string, value interface{}) {
	if r.Variables == nil {
		r.Variables = make(map[string]interface{}, 1)
	}
	r.Variables[key] = value
}

func (r *Request) SetHeader(header http.Header) {
	r.header = header
}
