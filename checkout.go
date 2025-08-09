package creemio

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type CheckoutCustomer struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

type Checkout struct {
	ID           string            `json:"id"`
	Mode         Mode              `json:"mode"`
	Object       string            `json:"object"`
	Status       string            `json:"status"`
	RequestID    string            `json:"request_id"`
	Product      *Product          `json:"product"`
	Units        int               `json:"units"`
	Order        *CheckoutOrder    `json:"order"`
	Subscription *Subscription     `json:"subscription"`
	Customer     *Customer         `json:"customer"`
	CustomFields []CustomField     `json:"custom_fields"`
	CheckoutURL  string            `json:"checkout_url"`
	SuccessURL   string            `json:"success_url"`
	Feature      []CheckoutFeature `json:"feature"`
	Metadata     map[string]any    `json:"metadata"`
}

type CheckoutOrder struct {
	ID             string    `json:"id"`
	Mode           Mode      `json:"mode"`
	Object         string    `json:"object"`
	Customer       string    `json:"customer"`
	Product        string    `json:"product"`
	Transaction    string    `json:"transaction"`
	Discount       string    `json:"discount"`
	Amount         int       `json:"amount"`
	SubTotal       int       `json:"sub_total"`
	TaxAmount      int       `json:"tax_amount"`
	DiscountAmount int       `json:"discount_amount"`
	AmountDue      int       `json:"amount_due"`
	AmountPaid     int       `json:"amount_paid"`
	Currency       string    `json:"currency"`
	FxAmount       int       `json:"fx_amount"`
	FxCurrency     string    `json:"fx_currency"`
	FxRate         float64   `json:"fx_rate"`
	Status         string    `json:"status"`
	Type           string    `json:"type"`
	Affiliate      string    `json:"affiliate"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CheckoutFeature struct {
	License *License `json:"license"`
}

// product_id is required
type CheckoutCreateRequest struct {
	RequestID    string            `json:"request_id,omitempty"`
	ProductID    string            `json:"product_id"`
	Units        int               `json:"units,omitempty"`
	DiscountCode string            `json:"discount_code,omitempty"`
	Customer     *CheckoutCustomer `json:"customer,omitempty"`
	CustomField  []CustomField     `json:"custom_field,omitempty"`
	SuccessURL   string            `json:"success_url,omitempty"`
	Metadata     map[string]any    `json:"metadata,omitempty"`
}

type CheckoutService struct {
	client *Client
}

func (s *CheckoutService) Get(ctx context.Context, id string) (*Checkout, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/checkouts")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Set("checkout_id", id)
	req.URL.RawQuery = q.Encode()

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, newResponse(res, body), err
	}
	if res.StatusCode >= 400 {
		return nil, newResponse(res, body), newAPIError(body)
	}

	var checkout Checkout
	if err := json.Unmarshal(body, &checkout); err != nil {
		return nil, newResponse(res, body), err
	}

	return &checkout, newResponse(res, body), nil
}

func (s *CheckoutService) Create(ctx context.Context, data *CheckoutCreateRequest) (*Checkout, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/checkouts")

	if len(data.ProductID) == 0 {
		return nil, nil, errRequiredFieldProductID
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetUrl, bytes.NewReader(payload))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.client.apiKey)

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, newResponse(res, body), err
	}
	if res.StatusCode >= 400 {
		return nil, newResponse(res, body), newAPIError(body)
	}

	var checkout Checkout
	if err := json.Unmarshal(body, &checkout); err != nil {
		return nil, newResponse(res, body), err
	}

	return &checkout, newResponse(res, body), nil
}
