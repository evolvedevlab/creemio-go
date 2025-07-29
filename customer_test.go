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
	resp, res, err := c.Customers.Get(context.Background(), CustomerRequestQuery{
		Email: expectedEmail,
	})

	a.NoError(err)
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)
	a.Equal(expectedEmail, resp.Email)

	// Marshalling again for comparing expected json data with the received json data
	rawJson, err := json.Marshal(resp)

	a.NoError(err)
	a.JSONEq(string(mock.GetCustomerResponse()), string(rawJson))
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

	resp, res, err := c.Customers.Get(context.Background(), CustomerRequestQuery{})

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

	resp, res, err := c.Customers.Get(context.Background(), CustomerRequestQuery{
		Email: "user@example.com",
	})

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}
