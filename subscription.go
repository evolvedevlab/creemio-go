package creemio

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Feature struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
}

type Product struct {
	ID                string    `json:"id"`
	Mode              string    `json:"mode"`
	Object            string    `json:"object"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	ImageURL          string    `json:"image_url"`
	Features          []Feature `json:"features"`
	Price             int       `json:"price"`
	Currency          string    `json:"currency"`
	BillingType       string    `json:"billing_type"`
	BillingPeriod     string    `json:"billing_period"`
	Status            string    `json:"status"`
	TaxMode           string    `json:"tax_mode"`
	TaxCategory       string    `json:"tax_category"`
	ProductURL        string    `json:"product_url"`
	DefaultSuccessURL string    `json:"default_success_url"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (p *Product) UnmarshalJSON(data []byte) error {
	// If it's a string, treat it as product ID
	if len(data) > 0 && data[0] == '"' {
		return json.Unmarshal(data, &p.ID)
	}

	// Otherwise, try to parse it as a full Product object
	type alias Product // avoid recursion
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = Product(tmp)
	return nil
}

type SubscriptionItem struct {
	ID        string `json:"id"`
	Mode      string `json:"mode,omitempty"`
	Object    string `json:"object,omitempty"`
	ProductID string `json:"product_id"`
	PriceID   string `json:"price_id"`
	Units     int    `json:"units"`
}

type Transaction struct {
	ID             string `json:"id"`
	Mode           string `json:"mode"`
	Object         string `json:"object"`
	Amount         int    `json:"amount"`
	AmountPaid     int    `json:"amount_paid"`
	DiscountAmount int    `json:"discount_amount"`
	Currency       string `json:"currency"`
	Type           string `json:"type"`
	TaxCountry     string `json:"tax_country"`
	TaxAmount      int    `json:"tax_amount"`
	Status         string `json:"status"`
	RefundedAmount int    `json:"refunded_amount"`
	Order          string `json:"order"`
	Subscription   string `json:"subscription"`
	Customer       string `json:"customer"`
	Description    string `json:"description"`
	PeriodStart    int    `json:"period_start"`
	PeriodEnd      int    `json:"period_end"`
	CreatedAt      int    `json:"created_at"`
}

type SubscriptionStatus string

const (
	SubscriptionStatusActive   SubscriptionStatus = "active"
	SubscriptionStatusCanceled SubscriptionStatus = "canceled"
	SubscriptionStatusUnpaid   SubscriptionStatus = "unpaid"
	SubscriptionStatusPaused   SubscriptionStatus = "paused"
	SubscriptionStatusTrialing SubscriptionStatus = "trialing"
)

type Subscription struct {
	ID                     string             `json:"id"`
	Mode                   string             `json:"mode"`
	Object                 string             `json:"object"`
	Product                Product            `json:"product"`
	Customer               Customer           `json:"customer"`
	Items                  []SubscriptionItem `json:"items"`
	CollectionMethod       string             `json:"collection_method"`
	Status                 SubscriptionStatus `json:"status"`
	LastTransactionID      string             `json:"last_transaction_id"`
	LastTransaction        Transaction        `json:"last_transaction"`
	LastTransactionDate    time.Time          `json:"last_transaction_date"`
	NextTransactionDate    time.Time          `json:"next_transaction_date"`
	CurrentPeriodStartDate time.Time          `json:"current_period_start_date"`
	CurrentPeriodEndDate   time.Time          `json:"current_period_end_date"`
	CanceledAt             *time.Time         `json:"canceled_at"`
	CreatedAt              time.Time          `json:"created_at"`
	UpdatedAt              time.Time          `json:"updated_at"`
}

type SubscriptionUpdateBehavior string

const (
	ProrationChargeImmediately SubscriptionUpdateBehavior = "proration-charge-immediately"
	ProrationCharge            SubscriptionUpdateBehavior = "proration-charge"
	ProrationNone              SubscriptionUpdateBehavior = "proration-none"
)

type UpdateSubscriptionRequest struct {
	SubscriptionID string             `json:"-"`
	Items          []SubscriptionItem `json:"items,omitempty"`
	// proration-charge-immediately | proration-charge | proration-none
	UpdateBehavior SubscriptionUpdateBehavior `json:"update_behavior,omitempty"`
}

type SubscriptionService struct {
	client *Client
}

func (s *SubscriptionService) Get(ctx context.Context, id string) (*Subscription, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/subscriptions")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)

	q := req.URL.Query()
	q.Set("subscription_id", id)
	req.URL.RawQuery = q.Encode()

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, newResponse(res), err
	}
	if res.StatusCode >= 400 {
		return nil, newResponse(res), newAPIError(res.Body)
	}
	defer res.Body.Close()

	var sub Subscription
	if err := json.NewDecoder(res.Body).Decode(&sub); err != nil {
		return nil, newResponse(res), err
	}

	return &sub, newResponse(res), nil
}

func (s *SubscriptionService) Update(ctx context.Context, data *UpdateSubscriptionRequest) (*Subscription, *Response, error) {
	if len(data.SubscriptionID) == 0 {
		return nil, nil, errRequiredFieldSubscriptionID
	}

	targetUrl := makeUrl(s.client.baseURL, "/subscriptions", data.SubscriptionID)

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
		return nil, newResponse(res), err
	}
	if res.StatusCode >= 400 {
		return nil, newResponse(res), newAPIError(res.Body)
	}
	defer res.Body.Close()

	var sub Subscription
	if err := json.NewDecoder(res.Body).Decode(&sub); err != nil {
		return nil, newResponse(res), err
	}

	return &sub, newResponse(res), nil
}

func (s *SubscriptionService) Cancel(ctx context.Context, id string) (*Subscription, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/subscriptions", id, "cancel")

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", s.client.apiKey)

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, newResponse(res), err
	}
	if res.StatusCode >= 400 {
		return nil, newResponse(res), newAPIError(res.Body)
	}
	defer res.Body.Close()

	var sub Subscription
	if err := json.NewDecoder(res.Body).Decode(&sub); err != nil {
		return nil, newResponse(res), err
	}

	return &sub, newResponse(res), nil
}

// TODO: Upgrade()
