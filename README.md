# creemio-go

Go SDK for creem.io

## Installation

```sh
go get -u github.com/evolvedevlab/creemio-go@latest
```

## Init the Client

Optional configs can be added as shown below.

```go
package main

import (
    "github.com/evolvedevlab/creemio-go"
)

func main() {
    client := creemio.New(creemio.WithAPIKey(os.Getenv("API_KEY")))
}
```

**With other options:**

```go
client := creemio.New(
    creemio.WithAPIKey(os.Getenv("API_KEY")),
    creemio.WithBaseURL(creemio.BaseTestAPIUrl), // Or another url
    creemio.WithHTTPClient(http.DefaultClient) // Custom http client
)
```

## Error Handling

```go
subscription, response, err := client.Subscriptions.Get(context.Background(), "1")
if err != nil {
    var apiErr *creemio.APIError
	if errors.As(err, &apiErr) {
		// handle api error
	}
    // handle other error
}
```

## WebHooks

```go
// These helper structs will be added later in the sdk.
type WebHookData struct {
    ID    string `json:"id"`
    Event string `json:"eventType"`
}

type WebHookSubscription struct {
    // ...
}
type WebHookCheckout struct {
    // ...
}
func WebHookHandler(w http.ResponseWriter, r *http.Request) {
    sig := r.Header.Get("creem-signature")
    // Authenticate the webhook signature

    body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	var payload WebHookData
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Fatal(err)
	}

    switch payload.Event {
	    case creemio.WebHookEventSubscriptionActive:
            var sub WebHookSubscription
            if err := json.Unmarshal(body, &sub); err != nil {
                log.Fatal(err)
            }
	    	// handle subscription active
    	case creemio.WebHookEventCheckoutCompleted:
            var checkout WebHookCheckout
            if err := json.Unmarshal(body, &checkout); err != nil {
                log.Fatal(err)
            }
		    // handle checkout completed
	    default:
	    	panic("invalid event type")
	}

}
```

For further information, check the [offical docs](https://docs.creem.io/learn/webhooks).

## Issues

Some apis still don't work properly. Please check the [issues](https://github.com/evolvedevlab/creemio-go/issues).

## Testing

```sh
make test
```
