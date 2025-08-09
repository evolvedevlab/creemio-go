package creemio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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

	reqData := &CheckoutCreateRequest{
		ProductID:    expectedProductID,
		RequestID:    expectedReqID,
		Units:        expectedUnit,
		DiscountCode: "SUMMER20",
		Customer: &CheckoutCustomer{
			ID:    expectedCustomerID,
			Email: "paul@gmail.com",
		},
		CustomField: []CustomField{
			{
				Type:     "text",
				Key:      "key",
				Label:    "Enter Key",
				Optional: true,
				Text: &TextSpec{
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

	resp, res, err := c.Checkouts.Create(context.Background(), reqData)

	a.NoError(err)
	a.NotNil(resp)
	a.Equal(http.StatusOK, res.Status)
	a.Equal(expectedProductID, resp.Product.ID)
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

	resp, res, err := c.Checkouts.Create(context.Background(), &CheckoutCreateRequest{
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

	resp, res, err := c.Checkouts.Create(context.Background(), &CheckoutCreateRequest{})

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
	url, err := url.Parse(fmt.Sprintf("/%s/checkouts", APIVersion))
	if err != nil {
		panic(err)
	}

	q := url.Query()
	q.Set("checkout_id", checkoutID)
	url.RawQuery = q.Encode()

	resp, res, err := c.Checkouts.Get(context.Background(), checkoutID)

	a.NoError(err)
	a.Equal(url.RequestURI(), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)
	a.Equal(checkoutID, resp.ID)

	// Have to do it like this, we cannot compare json here as some field(s)
	// can either be string or object in json.
	var expectedCheckout Checkout
	err = json.Unmarshal(mock.GetCheckoutResponse(), &expectedCheckout)

	a.NoError(err)
	a.Equal(expectedCheckout, *resp)
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
