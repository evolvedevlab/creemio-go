package creemio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/evolvedevlab/creemio-go/mock"
	"github.com/stretchr/testify/assert"
)

func TestCustomers_Get(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetCustomer))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	expectedEmail := "user@example.com"
	// For comparing the url request url with search params
	url, err := url.Parse(fmt.Sprintf("/%s/customers", APIVersion))
	if err != nil {
		panic(err)
	}
	q := url.Query()
	q.Set("email", expectedEmail)
	url.RawQuery = q.Encode()

	resp, res, err := c.Customers.Get(context.Background(), &CustomerRequestQuery{
		Email: expectedEmail,
	})

	a.NoError(err)
	a.Equal(url.RequestURI(), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)
	a.Equal(expectedEmail, resp.Email)

	var expected Customer
	err = json.Unmarshal(mock.GetCustomerResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)
}

func TestCustomers_GetWithMissingRequiredField(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetCustomer))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Customers.Get(context.Background(), &CustomerRequestQuery{})

	a.Error(err)
	a.Nil(resp)
	a.Nil(res)
	a.EqualError(err, errCustomerNoQuery.Error())
}

func TestCustomers_GetWithError(t *testing.T) {
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

	resp, res, err := c.Customers.Get(context.Background(), &CustomerRequestQuery{
		Email: "user@example.com",
	})

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}

func TestCustomers_GetBillingPortalURL(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetBillingPortalURL))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Customers.GetBillingPortalURL(context.Background(), "1")

	a.NoError(err)
	a.True(strings.HasPrefix(resp, "https://creem.io/my-orders/login/"))
	a.NotNil(res)
	a.Equal(http.StatusOK, res.Status)
}

func TestCustomers_GetBillingPortalURL_WithError(t *testing.T) {
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

	resp, res, err := c.Customers.GetBillingPortalURL(context.Background(), "1")

	a.Error(err)
	a.Empty(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}
