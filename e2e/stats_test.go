package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/evolvedevlab/creemio-go"
	"github.com/stretchr/testify/assert"
)

func TestStats_GetMetricsSummary(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	summary, res, err := client.Stats.GetMetricsSummary(context.Background(), &creemio.MetricsSummaryQuery{
		Currency: creemio.CurrencyUSD,
	})
	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	// collected from dashborad manually
	// WARN: these values needs to be set manually for testing
	var (
		totalProducts = 7
		totalSubs     = 13
		totalCust     = 9
		totalPayments = 20
		activeSubs    = 1
	)

	a.Equal(totalProducts, summary.Totals.TotalProducts)
	a.Equal(totalSubs, summary.Totals.TotalSubscriptions)
	a.Equal(totalCust, summary.Totals.TotalCustomers)
	a.Equal(totalPayments, summary.Totals.TotalPayments)
	a.Equal(activeSubs, summary.Totals.ActiveSubscriptions)
}
