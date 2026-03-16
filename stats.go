package creemio

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
)

var (
	errRequiredFieldCurrency = errors.New("currency is required")
)

type Interval string

const (
	IntervalDay   Interval = "day"
	IntervalWeek  Interval = "week"
	IntervalMonth Interval = "month"
)

type MetricsSummary struct {
	Totals  Totals   `json:"totals"`
	Periods []Period `json:"periods"`
}

type Totals struct {
	TotalProducts              int     `json:"totalProducts"`
	TotalSubscriptions         int     `json:"totalSubscriptions"`
	TotalCustomers             int     `json:"totalCustomers"`
	TotalPayments              int     `json:"totalPayments"`
	ActiveSubscriptions        int     `json:"activeSubscriptions"`
	TotalRevenue               float64 `json:"totalRevenue"`
	TotalNetRevenue            float64 `json:"totalNetRevenue"`
	NetMonthlyRecurringRevenue float64 `json:"netMonthlyRecurringRevenue"`
	MonthlyRecurringRevenue    float64 `json:"monthlyRecurringRevenue"`
}

type Period struct {
	Timestamp    int64   `json:"timestamp"`
	GrossRevenue float64 `json:"grossRevenue"`
	NetRevenue   float64 `json:"netRevenue"`
}

type MetricsSummaryQuery struct {
	Currency  Currency
	Interval  Interval
	StartDate int64
	EndDate   int64
}

type StatsService struct {
	client *Client
}

// currency is required
func (s *StatsService) GetMetricsSummary(ctx context.Context, query *MetricsSummaryQuery) (*MetricsSummary, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/stats", "summary")

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetUrl, nil)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	if query == nil || len(query.Currency) == 0 {
		return nil, nil, errRequiredFieldCurrency
	}

	q := req.URL.Query()
	q.Set("currency", string(query.Currency))
	if len(query.Interval) > 0 {
		q.Set("interval", string(query.Interval))
	}
	if query.StartDate > 0 {
		q.Set("start_date", strconv.FormatInt(query.StartDate, 10))
	}
	if query.EndDate > 0 {
		q.Set("end_date", strconv.FormatInt(query.EndDate, 10))
	}
	req.URL.RawQuery = q.Encode()

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, newResponse(res, body), err
	}
	if res.StatusCode >= 400 {
		return nil, newResponse(res, body), newAPIError(body)
	}

	var summary MetricsSummary
	if err := json.Unmarshal(body, &summary); err != nil {
		return nil, newResponse(res, body), err
	}

	return &summary, newResponse(res, body), nil
}
