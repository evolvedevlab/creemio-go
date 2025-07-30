# creemio-go

Go SDK for creem.io

# Installation

```sh
go get github.com/evolvedevlab/creemio-go
```

# Usage 

## Init the Client

Optional configs can be added as shown.
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
    // handle error
}
```

## WebHooks

TODO: Will be documented later.

# Testing
```sh
make test
```