package creemio

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

var errDiscountNoQuery = errors.New("either discount_id or discount_code must be present")

type DiscountStatus string

const (
	DiscountStatusActive    = "active"
	DiscountStatusDraft     = "draft"
	DiscountStatusExpired   = "expired"
	DiscountStatusScheduled = "scheduled"
)

type DiscountType string

const (
	DiscountTypePercentage = "percentage"
	DiscountTypeFixed      = "fixed"
)

type DiscountDuration string

const (
	DiscountDurationForever   = "forever"
	DiscountDurationOnce      = "once"
	DiscountDurationRepeating = "repeating"
)

type Discount struct {
	ID                string           `json:"id"`
	Mode              Mode             `json:"mode"`
	Object            string           `json:"object"`
	Status            DiscountStatus   `json:"status"`
	Name              string           `json:"name"`
	Code              string           `json:"code"`
	Type              DiscountType     `json:"type"`
	Amount            int              `json:"amount,omitempty"`
	Currency          string           `json:"currency,omitempty"`
	Percentage        int              `json:"percentage,omitempty"`
	ExpiryDate        *time.Time       `json:"expiry_date,omitempty"`
	MaxRedemptions    int              `json:"max_redemptions,omitempty"`
	Duration          DiscountDuration `json:"duration"`
	DurationInMonths  int              `json:"duration_in_months,omitempty"`
	AppliesToProducts []string         `json:"applies_to_products,omitempty"`
}

type CreateDiscountRequest struct {
	Name              string           `json:"name"`
	Type              DiscountType     `json:"type"`
	Duration          DiscountDuration `json:"duration"`
	AppliesToProducts []string         `json:"applies_to_products"`
	Code              string           `json:"code,omitempty"`
	Amount            int              `json:"amount,omitempty"`
	Currency          string           `json:"currency,omitempty"`
	Percentage        int              `json:"percentage,omitempty"`
	ExpiryDate        *time.Time       `json:"expiry_date,omitempty"`
	MaxRedemptions    int              `json:"max_redemptions,omitempty"`
	DurationInMonths  int              `json:"duration_in_months,omitempty"`
}

type DiscountRequestQuery struct {
	DiscountID   string
	DiscountCode string
}

type DiscountService struct {
	client *Client
}

func (s *DiscountService) Get(ctx context.Context, query *DiscountRequestQuery) (*Discount, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/discounts")

	if query == nil || (len(query.DiscountID) == 0 && len(query.DiscountCode) == 0) {
		return nil, nil, errDiscountNoQuery
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	if len(query.DiscountID) > 0 {
		q.Set("discount_id", query.DiscountID)
	}
	if len(query.DiscountCode) > 0 {
		q.Set("discount_code", query.DiscountCode)
	}
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

	var discount Discount
	if err := json.Unmarshal(body, &discount); err != nil {
		return nil, newResponse(res, body), err
	}

	return &discount, newResponse(res, body), nil
}

func (s *DiscountService) Create(ctx context.Context, data *CreateDiscountRequest) (*Discount, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/discounts")

	if data == nil {
		return nil, nil, errRequiredMissingField
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

	var discount Discount
	if err := json.Unmarshal(body, &discount); err != nil {
		return nil, newResponse(res, body), err
	}

	return &discount, newResponse(res, body), nil
}

func (s *DiscountService) Delete(ctx context.Context, id string) (*Discount, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/discounts", id, "delete")

	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, targetUrl, nil)
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

	var discount Discount
	if err := json.Unmarshal(body, &discount); err != nil {
		return nil, newResponse(res, body), err
	}

	return &discount, newResponse(res, body), nil
}
