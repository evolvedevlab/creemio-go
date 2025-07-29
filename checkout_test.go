package creemio

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evolvedevlab/creemio-go/mock"
	"github.com/stretchr/testify/assert"
)

func TestCheckouts_Create(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostCheckout))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	var (
		expectedProductID  = "prod_987654321"
		expectedReqID      = "req_1234567890"
		expectedUnit       = 1
		expectedCustomerID = "cust_1234567890"
		expectedSuccessUrl = "https://example.com/return"
	)

	req := CheckoutCreateRequest{
		ProductID:    expectedProductID,
		RequestID:    expectedReqID,
		Units:        expectedUnit,
		DiscountCode: "SUMMER20",
		Customer: &CheckoutCustomer{
			ID:    expectedCustomerID,
			Email: "paul@gmail.com",
		},
		CustomField: []CheckoutCustomField{
			{
				Type:     "text",
				Key:      "key",
				Label:    "Enter Key",
				Optional: true,
				Text: &CheckoutTextSpec{
					MaxLength: 10,
					MinLength: 12,
				},
			},
		},
		SuccessURL: expectedSuccessUrl,
		Metadata: map[string]any{
			"userID": expectedCustomerID,
		},
	}

	resp, res, err := c.Checkouts.Create(context.Background(), req)

	a.NoError(err)
	a.NotNil(resp)
	a.Equal(http.StatusCreated, res.Status)
	a.Equal(expectedProductID, resp.Product)
	a.Equal(expectedReqID, resp.RequestID)
	a.Equal(expectedSuccessUrl, resp.SuccessURL)
}

func TestCheckouts_CreateWithError(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Checkouts.Create(context.Background(), CheckoutCreateRequest{
		ProductID: "1",
	})

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}

func TestCheckouts_CreateWithMissingProductID(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostCheckout))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Checkouts.Create(context.Background(), CheckoutCreateRequest{})

	a.Error(err)
	a.EqualError(err, errRequiredFieldProductID.Error())
	a.Nil(resp)
	a.Nil(res)
}

func TestCheckouts_Get(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetCheckout))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	checkoutID := "1"
	resp, res, err := c.Checkouts.Get(context.Background(), checkoutID)

	a.NoError(err)
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)
	a.Equal(checkoutID, resp.ID)

	// Marshalling again for comparing expected json data with the received json data
	rawJson, err := json.Marshal(resp)

	a.NoError(err)
	a.JSONEq(string(mock.GetCheckoutResponse()), string(rawJson))
}

func TestCheckouts_GetWithError(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Checkouts.Get(context.Background(), "1")

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}
