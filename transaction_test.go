package creemio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/evolvedevlab/creemio-go/mock"
	"github.com/stretchr/testify/assert"
)

func TestTransactions_List(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetTransaction))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	// For comparing the url request url with search params
	url, err := url.Parse(fmt.Sprintf("/%s/transactions/search", APIVersion))
	if err != nil {
		panic(err)
	}

	var (
		custID   = "cus_123"
		prodID   = "prod_123"
		orderID  = "1"
		pageNum  = 1
		pageSize = 2
	)

	q := url.Query()
	q.Set("customer_id", custID)
	q.Set("order_id", orderID)
	q.Set("product_id", prodID)
	q.Set("page_number", strconv.Itoa(pageNum))
	q.Set("page_size", strconv.Itoa(pageSize))
	url.RawQuery = q.Encode()

	resp, res, err := c.Transactions.List(context.Background(), &TransactionListQuery{
		CustomerID: custID,
		OrderID:    orderID,
		ProductID:  prodID,
		PageNumber: pageNum,
		PageSize:   pageSize,
	})

	a.NoError(err)
	a.Equal(url.RequestURI(), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)

	var expected TransactionList
	err = json.Unmarshal(mock.GetTransactionResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)
}

func TestTransactions_ListWithError(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Transactions.List(context.Background(), nil)

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}
