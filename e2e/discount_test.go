package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/evolvedevlab/creemio-go"
	"github.com/stretchr/testify/assert"
)

func TestDiscounts_Get(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	code := "KING100"
	discount, res, err := client.Discounts.Get(context.Background(), &creemio.DiscountRequestQuery{
		DiscountCode: code,
	})
	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(discount)
	a.Equal(code, discount.Code)
}

func TestDiscounts_Create(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	var (
		code   = "TEST100"
		prodID = "prod_5XCDn29MFRGYYkU98u4SI4"
	)
	discount, res, err := client.Discounts.Create(context.Background(), &creemio.CreateDiscountRequest{
		Name:              code,
		Code:              code,
		Type:              creemio.DiscountTypeFixed,
		Duration:          creemio.DiscountDurationOnce,
		Amount:            100,
		Currency:          "USD",
		AppliesToProducts: []string{prodID},
	})

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(discount)
	a.Equal(code, discount.Code)
	a.Contains(discount.AppliesToProducts, prodID)
}

func TestDiscounts_Delete(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	discountID := "dis_5SoHenlTIW9Rv67J3LisO3"
	discount, res, err := client.Discounts.Delete(context.Background(), discountID)

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(discount)
	a.Equal(discountID, discount.ID)
}
