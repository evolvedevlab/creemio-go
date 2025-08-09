package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/evolvedevlab/creemio-go"
	"github.com/stretchr/testify/assert"
)

func TestCustomers_Get(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	email := "nn@gmail.com"
	customer, res, err := client.Customers.Get(context.Background(), &creemio.CustomerRequestQuery{
		Email: email,
	})

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotEmpty(*customer)
	a.Equal(email, customer.Email)
}

func TestCustomers_GetBillingPortalURL(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	custID := "cust_35R82UpDVmQgC5gsKu4R3a"
	billingURL, res, err := client.Customers.GetBillingPortalURL(context.Background(), custID)

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotEmpty(billingURL)
}
