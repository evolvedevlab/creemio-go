package creemio

import (
	"net/http"
)

const (
	APIURL     = "https://api.creem.io"
	TestAPIURL = "https://test-api.creem.io"
	APIVersion = "v1"
)

type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string

	Checkouts     *CheckoutService
	Customers     *CustomerService
	Subscriptions *SubscriptionService
	Transactions  *TransactionService
	Products      *ProductService
	Discounts     *DiscountService
}

type Option func(*Client)

func New(opts ...Option) *Client {
	c := &Client{
		baseURL:    APIURL,
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.Checkouts = &CheckoutService{client: c}
	c.Customers = &CustomerService{client: c}
	c.Subscriptions = &SubscriptionService{client: c}
	c.Transactions = &TransactionService{client: c}
	c.Products = &ProductService{client: c}
	c.Discounts = &DiscountService{client: c}

	return c
}
