package graphql

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	endpoint   string
	header     http.Header
	httpClient Doer
}

func New(endpoint string, opts ...Option) *Client {
	header := http.Header{}
	header.Set("Content-Type", "application/graphql")

	c := &Client{
		endpoint: endpoint,
		header:   header,
	}

	for _, opt := range opts {
		opt.Apply(c)
	}

	if c.httpClient == nil {
		c.httpClient = http.DefaultClient
	}

	return c
}

func (c *Client) Post(ctx context.Context, request *Request, response *Response) (err error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.endpoint, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	for k, v := range c.header {
		req.Header[k] = v
	}

	for k, v := range request.header {
		req.Header[k] = v
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if e := resp.Body.Close(); e != nil {
			err = e // replace with body close error
		}
	}()

	response.StatusCode = resp.StatusCode

	if http.StatusBadRequest <= resp.StatusCode {
		response.Errors = Errors{ // will be overridden if response includes errors
			Error{Message: http.StatusText(resp.StatusCode)},
		}
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return err
	}

	return nil
}
