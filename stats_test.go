package creemio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/evolvedevlab/creemio-go/mock"
	"github.com/stretchr/testify/assert"
)

func TestStats_GetMetricsSummary(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetMetricsSummary))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	currency := CurrencyUSD
	// For comparing the url request url with search params
	url, err := url.Parse(fmt.Sprintf("/%s/stats/summary", APIVersion))
	if err != nil {
		panic(err)
	}
	q := url.Query()
	q.Set("currency", string(currency))
	url.RawQuery = q.Encode()

	resp, res, err := c.Stats.GetMetricsSummary(context.Background(), &MetricsSummaryQuery{
		Currency: currency,
	})

	a.NoError(err)
	a.Equal(url.RequestURI(), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)

	var expected MetricsSummary
	err = json.Unmarshal(mock.GetMetricsSummaryResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)
}

func TestStats_GetMetricsSummaryWithMissingCurrency(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetMetricsSummary))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Stats.GetMetricsSummary(context.Background(), &MetricsSummaryQuery{})

	a.Error(err)
	a.EqualError(err, errRequiredFieldCurrency.Error())
	a.Nil(resp)
	a.Nil(res)
}

func TestStats_GetMetricsSummaryWithError(t *testing.T) {
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

	resp, res, err := c.Stats.GetMetricsSummary(context.Background(), &MetricsSummaryQuery{
		Currency: CurrencyUSD,
	})

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}
