package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/evolvedevlab/creemio-go"
	"github.com/stretchr/testify/assert"
)

func TestCheckouts_Create(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	var (
		prodID     = "prod_10rxEcEUn5fYkSBpXWhdjR"
		email      = "npaul@gmail.com"
		successURL = "https://creem.io"
	)

	checkout, res, err := client.Checkouts.Create(context.Background(), &creemio.CheckoutCreateRequest{
		RequestID: "1",
		ProductID: prodID,
		Units:     1,
		Customer: &creemio.CheckoutCustomer{
			Email: email,
		},
		CustomField: []creemio.CustomField{
			{
				Type:  "text",
				Key:   "test",
				Label: "Enter test",
			},
		},
		SuccessURL: successURL,
		Metadata:   map[string]any{"email": email},
	})

	a.NoError(err)

	// http status code
	a.Equal(http.StatusOK, res.Status)

	a.Equal(prodID, checkout.Product.ID)
	a.Equal(email, checkout.Metadata["email"])
	a.Equal(successURL, checkout.SuccessURL)
}

func TestCheckouts_Get(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	var (
		checkoutID = "ch_3ryTHZ4Tz7qTZ2rMSUMBOS"
		email      = "npaul@gmail.com"
	)
	checkout, res, err := client.Checkouts.Get(context.Background(), checkoutID)

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.Equal(checkoutID, checkout.ID)
	a.Equal(email, checkout.Metadata["email"])
}
