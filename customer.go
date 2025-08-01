package creemio

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
)

var (
	errCustomerNoQuery           = errors.New("either customer_id or email is required")
	errInvalidCustomerPortalLink = errors.New("invalid customer portal link")
)

// Subscription is either a full object or just an ID string depending on the API context.
type Customer struct {
	ID        string    `json:"id"`
	Mode      Mode      `json:"mode"`
	Object    string    `json:"object"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Customer) UnmarshalJSON(data []byte) error {
	// If it's a string, treat it as customer ID
	if len(data) > 0 && data[0] == '"' {
		return json.Unmarshal(data, &c.ID)
	}

	// Otherwise, try to parse it as a full Customer object
	type alias Customer // avoid recursion
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = Customer(tmp)
	return nil
}

type customerPortalResponse struct {
	CustomerPortalLink string `json:"customer_portal_link"`
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
	req.Header.Set("Content-Type", "application/json")

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

	if res.StatusCode >= 400 {
		return nil, newResponse(res), newAPIError(res.Body)
	}

	var customer Customer
	if err := json.NewDecoder(res.Body).Decode(&customer); err != nil {
		return nil, newResponse(res), err
	}

	return &customer, newResponse(res), nil
}

// https://docs.creem.io/learn/customers/customer-portal#response
func (s *CustomerService) GetBillingPortalURL(ctx context.Context, customerID string) (string, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/customers", "billing")

	reqBody := map[string]string{"customer_id": customerID}
	payload, err := json.Marshal(reqBody)
	if err != nil {
		return "", nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetUrl, bytes.NewReader(payload))
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return "", newResponse(res), err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", newResponse(res), newAPIError(res.Body)
	}

	var resp customerPortalResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return "", newResponse(res), err
	}

	link := resp.CustomerPortalLink
	if len(link) == 0 || !strings.HasPrefix(link, "https://creem.io/my-orders/login/") {
		return "", newResponse(res), errInvalidCustomerPortalLink
	}

	return resp.CustomerPortalLink, newResponse(res), nil
}
