package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/evolvedevlab/creemio-go"
	"github.com/stretchr/testify/assert"
)

func TestProducts_List(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	results, res, err := client.Products.List(context.Background(), &creemio.ProductListQuery{
		PageSize: 100,
	})
	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.GreaterOrEqual(len(results.Items), 1)
	a.Equal(len(results.Items), results.Pagination.TotalRecords)
	a.NotEmpty(results.Pagination)
}

func TestProducts_Get(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	prodID := "prod_5XCDn29MFRGYYkU98u4SI4"
	product, res, err := client.Products.Get(context.Background(), prodID)
	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(product)
	a.Equal(prodID, product.ID)
}

func TestProducts_Create(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	product, res, err := client.Products.Create(context.Background(), &creemio.CreateProductRequest{
		Name:        "product 1",
		Description: "this is a test desc used in test env",
		Price:       100,
		Currency:    "USD",
		BillingType: creemio.BillingTypeOneTime,
	})

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(product)
}
