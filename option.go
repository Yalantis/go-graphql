package graphql

type Option interface {
	Apply(*Client)
}

func WithHttpClient(client Doer) Option {
	return withHttpClient{httpClient: client}
}

type withHttpClient struct {
	httpClient Doer
}

func (w withHttpClient) Apply(c *Client) {
	c.httpClient = w.httpClient
}

func WithAuth(value string) Option {
	return withAuth{value: value}
}

type withAuth struct {
	value string
}

func (w withAuth) Apply(c *Client) {
	c.header.Set("Authorization", w.value)
}
