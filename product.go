package creemio

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

type BillingType string

const (
	BillingTypeRecurring BillingType = "recurring"
	BillingTypeOneTime   BillingType = "onetime"
)

// Product is either a full object or just an ID string depending on the API context.
type Product struct {
	ID                string      `json:"id"`
	Mode              Mode        `json:"mode"`
	Object            string      `json:"object"`
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	ImageURL          string      `json:"image_url"`
	Features          []Feature   `json:"features"`
	Price             int         `json:"price"`
	Currency          string      `json:"currency"`
	BillingType       BillingType `json:"billing_type"`
	BillingPeriod     string      `json:"billing_period"`
	Status            string      `json:"status"`
	TaxMode           string      `json:"tax_mode"`
	TaxCategory       string      `json:"tax_category"`
	ProductURL        string      `json:"product_url"`
	DefaultSuccessURL string      `json:"default_success_url"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
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

type CreateProductRequest struct {
	Name              string        `json:"name"`
	Description       string        `json:"description"`
	ImageURL          string        `json:"image_url,omitempty"`
	Price             int           `json:"price"`
	Currency          string        `json:"currency"`
	BillingType       BillingType   `json:"billing_type"`
	BillingPeriod     string        `json:"billing_period,omitempty"`
	TaxMode           string        `json:"tax_mode,omitempty"`
	TaxCategory       string        `json:"tax_category,omitempty"`
	DefaultSuccessURL string        `json:"default_success_url,omitempty"`
	CustomField       []CustomField `json:"custom_field,omitempty"`
}

type ProductList struct {
	Items      []Product  `json:"items"`
	Pagination Pagination `json:"pagination"`
}

type ProductListQuery struct {
	PageNumber int
	PageSize   int
}

type ProductService struct {
	client *Client
}

func (s *ProductService) Create(ctx context.Context, data *CreateProductRequest) (*Product, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/products")

	// Validation for required fields
	if data == nil {
		return nil, nil, errors.New("required fields are missing")
	}
	if len(data.Name) == 0 {
		return nil, nil, errors.New("name is required")
	}
	if data.Price <= 0 {
		return nil, nil, errors.New("price must be greater than 0")
	}
	if len(data.Currency) == 0 {
		return nil, nil, errors.New("currency is required")
	}
	if len(data.BillingType) == 0 {
		return nil, nil, errors.New("billing_type is required")
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

	var product Product
	if err := json.Unmarshal(body, &product); err != nil {
		return nil, newResponse(res, body), err
	}

	return &product, newResponse(res, body), nil
}

func (s *ProductService) Get(ctx context.Context, id string) (*Product, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/products")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	q := req.URL.Query()
	q.Set("product_id", id)
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

	var product Product
	if err := json.Unmarshal(body, &product); err != nil {
		return nil, newResponse(res, body), err
	}

	return &product, newResponse(res, body), nil
}

func (s *ProductService) List(ctx context.Context, query *ProductListQuery) (*ProductList, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/products", "search")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	if query != nil {
		q := req.URL.Query()
		if query.PageNumber > 0 {
			q.Set("page_number", strconv.Itoa(query.PageNumber))
		}
		if query.PageSize > 0 {
			q.Set("page_size", strconv.Itoa(query.PageSize))
		}
		req.URL.RawQuery = q.Encode()
	}

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

	var result ProductList
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, newResponse(res, body), err
	}

	return &result, newResponse(res, body), nil
}
