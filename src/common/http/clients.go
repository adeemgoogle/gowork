package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	baseUrl    string
	httpClient *http.Client

	headers map[string]Header
}

type Header struct {
	Value     string
	IsDefault bool
}

type Response struct {
	res  *http.Response
	Body []byte
}

func (r *Response) Unmarshal(v any) error {
	//return json.NewDecoder(r.res.Body).Decode(&v)
	return json.Unmarshal(r.Body, &v)
}

func NewClient(baseUrl string) *Client {
	return &Client{
		baseUrl: baseUrl,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: 10 * time.Second,
	}
}

func (c *Client) Get(ctx context.Context, endpoint string) (*Response, error) {
	url := c.baseUrl + endpoint
	fmt.Println("Request url: ", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	reqOp := c.setOptions(req)
	return c.sendRequest(ctx, reqOp)
}

func (c *Client) Post(ctx context.Context, endpoint string, body []byte) (*Response, error) {
	url := c.baseUrl + endpoint
	fmt.Println("Request url: ", url)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	reqOp := c.setOptions(req)
	return c.sendRequest(ctx, reqOp)
}

func (c *Client) setOptions(req *http.Request) *http.Request {
	for key, header := range c.headers {
		req.Header.Set(key, header.Value)
	}
	return req
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request) (*Response, error) {
	reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	res, err := c.httpClient.Do(req.WithContext(reqCtx))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return &Response{res, body}, nil
}
