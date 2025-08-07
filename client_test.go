package creemio

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewClient_Defaults(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	c := New()

	a.NotNil(c)
	a.Equal(APIURL, c.baseURL)
	a.NotNil(c.httpClient)

	// Ensure services are initialized
	a.NotNil(c.Checkouts)
	a.NotNil(c.Customers)
	a.NotNil(c.Subscriptions)
	a.NotNil(c.Transactions)
	a.NotNil(c.Products)
	a.NotNil(c.Discounts)
	a.NotNil(c.Licenses)
}

func TestNewClient_WithAPIKey(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	apiKey := "test-key"
	c := New(WithAPIKey(apiKey))

	a.Equal(apiKey, c.apiKey)
}

func TestNewClient_WithBaseURL(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	customURL := "https://custom.api"
	c := New(WithBaseURL(customURL))

	a.Equal(customURL, c.baseURL)
}

func TestNewClient_WithHTTPClient(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	customClient := &http.Client{Timeout: 5 * time.Second}
	c := New(WithHTTPClient(customClient))

	a.Equal(customClient, c.httpClient)
	a.Equal(5*time.Second, c.httpClient.Timeout)
}

func TestNewClient_CombinesOptions(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	apiKey := "combo-key"
	customURL := "https://combo.api"
	customHTTP := &http.Client{Timeout: 2 * time.Second}

	c := New(
		WithAPIKey(apiKey),
		WithBaseURL(customURL),
		WithHTTPClient(customHTTP),
	)

	a.Equal(apiKey, c.apiKey)
	a.Equal(customURL, c.baseURL)
	a.Equal(customHTTP, c.httpClient)
}
