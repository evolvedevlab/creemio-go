package creemio

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type CheckoutService struct {
	client *Client
}

type CheckoutCreateRequest struct {
	RequestID    string                `json:"request_id"`
	ProductID    string                `json:"product_id"`
	Units        int                   `json:"units"`
	DiscountCode string                `json:"discount_code,omitempty"`
	Customer     *CheckoutCustomer     `json:"customer,omitempty"`
	CustomField  []CheckoutCustomField `json:"custom_field,omitempty"`
	SuccessURL   string                `json:"success_url"`
	Metadata     map[string]any        `json:"metadata,omitempty"`
}

type CheckoutCustomer struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type CheckoutCustomField struct {
	Type     string            `json:"type"`
	Key      string            `json:"key"`
	Label    string            `json:"label"`
	Optional bool              `json:"optional"`
	Text     *CheckoutTextSpec `json:"text,omitempty"`
}

type CheckoutTextSpec struct {
	MaxLength int `json:"max_length"`
	MinLength int `json:"min_length"`
}

type CheckoutResponse struct {
	ID           string                `json:"id"`
	Mode         string                `json:"mode"`
	Object       string                `json:"object"`
	Status       string                `json:"status"`
	RequestID    string                `json:"request_id"`
	Product      string                `json:"product"`
	Units        int                   `json:"units"`
	Order        *CheckoutOrder        `json:"order"`
	Subscription string                `json:"subscription"`
	Customer     string                `json:"customer"`
	CustomFields []CheckoutCustomField `json:"custom_fields"`
	CheckoutURL  string                `json:"checkout_url"`
	SuccessURL   string                `json:"success_url"`
	Feature      []CheckoutFeature     `json:"feature"`
	Metadata     map[string]any        `json:"metadata"`
}

type CheckoutOrder struct {
	ID             string    `json:"id"`
	Mode           string    `json:"mode"`
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

type License struct {
	ID              string           `json:"id"`
	Mode            string           `json:"mode"`
	Object          string           `json:"object"`
	Status          string           `json:"status"`
	Key             string           `json:"key"`
	Activation      int              `json:"activation"`
	ActivationLimit int              `json:"activation_limit"`
	ExpiresAt       time.Time        `json:"expires_at"`
	CreatedAt       time.Time        `json:"created_at"`
	Instance        *LicenseInstance `json:"instance"`
}

type LicenseInstance struct {
	ID        string    `json:"id"`
	Mode      string    `json:"mode"`
	Object    string    `json:"object"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (c *CheckoutService) Get(ctx context.Context, id string) (*CheckoutResponse, *Response, error) {
	targetUrl := makeUrl(c.client.baseURL, "/checkouts", id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", c.client.apiKey)

	res, err := c.client.httpClient.Do(req)
	if err != nil {
		return nil, newResponse(res), err
	}
	defer res.Body.Close()

	var checkout CheckoutResponse
	if err := json.NewDecoder(res.Body).Decode(&checkout); err != nil {
		return nil, newResponse(res), err
	}

	return &checkout, newResponse(res), nil
}

func (c *CheckoutService) Create(ctx context.Context, data CheckoutCreateRequest) (*CheckoutResponse, *Response, error) {
	targetUrl := makeUrl(c.client.baseURL, "/checkouts")

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
	req.Header.Set("x-api-key", c.client.apiKey)

	res, err := c.client.httpClient.Do(req)
	if err != nil {
		return nil, newResponse(res), err
	}
	defer res.Body.Close()

	var checkout CheckoutResponse
	if err := json.NewDecoder(res.Body).Decode(&checkout); err != nil {
		return nil, newResponse(res), err
	}

	return &checkout, newResponse(res), nil
}
