# creemio-go

Go SDK for [creem.io](https://creem.io)

[![Go Reference](https://pkg.go.dev/badge/github.com/evolvedevlab/creemio-go.svg)](https://pkg.go.dev/github.com/evolvedevlab/creemio-go)
[![Test Status](https://github.com/evolvedevlab/creemio-go/actions/workflows/test.yml/badge.svg)](https://github.com/evolvedevlab/creemio-go/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

## Installation

```sh
go get -u github.com/evolvedevlab/creemio-go@latest
```

## Initialize the Client

Optional configs can be added as shown below.

```go
package main

import (
    "github.com/evolvedevlab/creemio-go"
    "net/http"
    "os"
)

func main() {
    client := creemio.New(creemio.WithAPIKey(os.Getenv("API_KEY")))
}
```

**With other options:**

```go
client := creemio.New(
    creemio.WithAPIKey(os.Getenv("API_KEY")),
    creemio.WithBaseURL(creemio.TestAPIURL),    // Set a custom base URL
    creemio.WithHTTPClient(http.DefaultClient), // Provide a custom HTTP client
)
```

## Error Handling

```go
subscription, response, err := client.Subscriptions.Get(context.Background(), "sub_xxxxx")
if err != nil {
    var apiErr *creemio.APIError
    if errors.As(err, &apiErr) {
        // Handle API-specific error
    }
    // Handle other errors
}
```

## WebHooks

### Handling Events

```go
func WebHookHandler(w http.ResponseWriter, r *http.Request) {
    sig := r.Header.Get("creem-signature")
    // Authenticate the webhook signature

    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Fatal(err)
    }
    defer r.Body.Close()

    var payload creemio.WebHookRequest
    if err := json.Unmarshal(body, &payload); err != nil {
        log.Fatal(err)
    }

    // Handle events
    switch payload.EventType {
    case creemio.WebHookEventSubscriptionActive:
        var sub creemio.WebHookSubscriptionRequest
        if err := json.Unmarshal(body, &sub); err != nil {
            log.Fatal(err)
        }
        // Handle subscription active event
    case creemio.WebHookEventCheckoutCompleted:
        var checkout creemio.WebHookCheckoutRequest
        if err := json.Unmarshal(body, &checkout); err != nil {
            log.Fatal(err)
        }
        // Handle checkout completed event
    default:
        panic("invalid event")
    }
}
```

For further information, check the [offical docs](https://docs.creem.io/learn/webhooks).

### Creem Signature Verification

```go
// secret is the webhook token generated from dashboard
func VerifyCreemSignature(r *http.Request, secret string) bool {
	sig := r.Header.Get("creem-signature")
	if len(sig) == 0 {
		return false
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return false
	}

	// Reconstructing the request body for later usage
	r.Body = io.NopCloser(strings.NewReader(string(body)))

	// Generate HMAC-SHA256
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(body)
	expectedMAC := mac.Sum(nil)
	expectedSig := hex.EncodeToString(expectedMAC)

	// Constant-time compare
	return hmac.Equal([]byte(sig), []byte(expectedSig))
}
```

This implementation is same as the [JS version in the official docs](https://docs.creem.io/learn/webhooks/verify-webhook-requests#how-to-verify-creem-signature).

## Implemented

- **Checkouts**
  - `GET /v1/checkouts` - Get Checkout Session
  - `POST /v1/checkouts` - Create Checkout Session
- **Products**
  - `POST /v1/products` - Create Product
  - `GET /v1/products` - Get Product
  - `GET /v1/products` - Get Products List
- **Customers**
  - `GET /v1/customers` - Get Customer
  - `GET /v1/customers/billing` - Get Customer Billing Portal Link
- **Transactions**
  - `GET /v1/transactions` - Get Transactions List
- **Licenses**
  - `POST /v1/licenses` - Validate License Key
  - `POST /v1/licenses` - Activate License Key
  - `POST /v1/licenses` - Deactivate License Key
- **Discounts**
  - `GET /v1/discounts` - Create Discount Code
  - `POST /v1/discounts` - Create Discount Code
  - `DELETE /v1/discounts` - Delete Discount Code
- **Subscriptions**
  - `GET /v1/subscriptions` - Get Subscription
  - `POST /v1/subscriptions/{id}` - Update Subscription
  - `POST /v1/subscriptions/{id}` - Upgrade Subscription
  - `POST /v1/subscriptions/{id}/cancel` - Cancel Subscription

## Issues

Some apis still don't work properly. Please check the [issues](https://github.com/evolvedevlab/creemio-go/issues).

## Testing

```sh
make test
```
