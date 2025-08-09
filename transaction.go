package creemio

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type Transaction struct {
	ID             string `json:"id"`
	Mode           Mode   `json:"mode"`
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

type TransactionList struct {
	Items      []Transaction `json:"items"`
	Pagination Pagination    `json:"pagination"`
}

type TransactionListQuery struct {
	CustomerID string
	OrderID    string
	ProductID  string
	PageNumber int
	PageSize   int
}

type TransactionService struct {
	client *Client
}

func (s *TransactionService) List(ctx context.Context, query *TransactionListQuery) (*TransactionList, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/transactions", "search")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	if query != nil {
		q := req.URL.Query()
		if len(query.CustomerID) > 0 {
			q.Set("customer_id", query.CustomerID)
		}
		if len(query.OrderID) > 0 {
			q.Set("order_id", query.OrderID)
		}
		if len(query.ProductID) > 0 {
			q.Set("product_id", query.ProductID)
		}
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

	var result TransactionList
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, newResponse(res, body), err
	}

	return &result, newResponse(res, body), nil
}
