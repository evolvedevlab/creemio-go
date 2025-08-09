package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/evolvedevlab/creemio-go"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptions_Get(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	subID := "sub_6syj1nIKE9fpJ5LbJ9PuQa"
	sub, res, err := client.Subscriptions.Get(context.Background(), subID)

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(sub)
	a.Equal(subID, sub.ID)
}

func TestSubscriptions_Update(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	var (
		subID   = "sub_6syj1nIKE9fpJ5LbJ9PuQa"
		subItem = creemio.SubscriptionItem{
			ID:        "sitem_5eRoo0COAMIUYJCMPogZyl",
			ProductID: "prod_10rxEcEUn5fYkSBpXWhdjR",
			PriceID:   "pprice_7SzOVL4yhHPsfkTgMIMOhl",
		}
	)
	sub, res, err := client.Subscriptions.Update(context.Background(), &creemio.UpdateSubscriptionRequest{
		SubscriptionID: subID,
		UpdateBehavior: creemio.ProrationCharge,
		Items:          []creemio.SubscriptionItem{subItem},
	})

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(sub)
	a.Equal(subID, sub.ID)
	a.Equal(subItem.ID, sub.Items[0].ID)
}

func TestSubscriptions_Upgrade(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	var (
		subID  = "sub_6syj1nIKE9fpJ5LbJ9PuQa"
		prodID = "prod_10rxEcEUn5fYkSBpXWhdjR"
	)
	sub, res, err := client.Subscriptions.Upgrade(context.Background(), &creemio.UpgradeSubscriptionRequest{
		SubscriptionID: subID,
		ProductID:      prodID,
		UpdateBehavior: creemio.ProrationCharge,
	})

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(sub)
	a.Equal(subID, sub.ID)
	a.Equal(prodID, sub.Product.ID)
}

func TestSubscriptions_Cancel(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	subID := "sub_6syj1nIKE9fpJ5LbJ9PuQa"
	sub, res, err := client.Subscriptions.Cancel(context.Background(), subID)

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(sub)
	a.Equal(subID, sub.ID)
}
