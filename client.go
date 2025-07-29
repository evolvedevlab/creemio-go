package creemio

import (
	"net/http"
	"time"
)

const (
	BaseAPIURL           = "https://api.creem.io"
	APIVersion           = "v1"
	DefaultClientTimeout = 10 * time.Second
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string

	Checkouts *CheckoutService
}

type Option func(*Client)

func New(opts ...Option) *Client {
	c := &Client{
		baseURL: BaseAPIURL,
		httpClient: &http.Client{
			Timeout: DefaultClientTimeout,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Checkouts = &CheckoutService{client: c}

	return c
}
