package e2e

import (
	"context"
	"net/http"
	"testing"

	"github.com/evolvedevlab/creemio-go"
	"github.com/stretchr/testify/assert"
)

func TestTransactions_List(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	results, res, err := client.Transactions.List(context.Background(), &creemio.TransactionListQuery{
		PageSize: 100,
	})
	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.GreaterOrEqual(len(results.Items), 1)
	a.Equal(len(results.Items), results.Pagination.TotalRecords)
	a.NotEmpty(results.Pagination)
}
