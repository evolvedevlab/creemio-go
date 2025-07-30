package creemio

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

var errCustomerNoQuery = errors.New("either customer_id or email is required")

type Customer struct {
	ID        string    `json:"id"`
	Mode      string    `json:"mode"`
	Object    string    `json:"object"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Either ID or Email should be present
type CustomerRequestQuery struct {
	ID    string
	Email string
}

type CustomerService struct {
	client *Client
}

func (s *CustomerService) Get(ctx context.Context, query *CustomerRequestQuery) (*Customer, *Response, error) {
	if len(query.ID) == 0 && len(query.Email) == 0 {
		return nil, nil, errCustomerNoQuery
	}

	targetUrl := makeUrl(s.client.baseURL, "/customers")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)

	q := req.URL.Query()
	if len(query.ID) > 0 {
		q.Set("customer_id", query.ID)
	} else {
		q.Set("email", query.Email)
	}
	req.URL.RawQuery = q.Encode()

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, newResponse(res), err
	}
	defer res.Body.Close()

	var customer Customer
	if err := json.NewDecoder(res.Body).Decode(&customer); err != nil {
		return nil, newResponse(res), err
	}

	return &customer, newResponse(res), nil
}
