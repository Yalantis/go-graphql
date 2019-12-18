package graphql

type Response struct {
	Data   interface{} `json:"data"`
	Errors Errors      `json:"errors"`

	StatusCode int `json:"-"`
}
