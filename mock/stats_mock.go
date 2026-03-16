package mock

import "net/http"

func HandleGetMetricsSummary(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetMetricsSummaryResponse())
}

func GetMetricsSummaryResponse() []byte {
	return []byte(`{
  "totals": {
    "totalProducts": 12,
    "totalSubscriptions": 48,
    "totalCustomers": 35,
    "totalPayments": 62,
    "activeSubscriptions": 21,
    "totalRevenue": 553939,
    "totalNetRevenue": 478094,
    "netMonthlyRecurringRevenue": 89500,
    "monthlyRecurringRevenue": 94200
  },
  "periods": [
    {
      "timestamp": 1763337600000,
      "grossRevenue": 2999,
      "netRevenue": 2909
    },
    {
      "timestamp": 1763942400000,
      "grossRevenue": 32989,
      "netRevenue": 31998
    },
    {
      "timestamp": 1764547200000,
      "grossRevenue": 47984,
      "netRevenue": 46542
    },
    {
      "timestamp": 1765152000000,
      "grossRevenue": 125958,
      "netRevenue": 122173
    },
    {
      "timestamp": 1765756800000,
      "grossRevenue": 343968,
      "netRevenue": 278372
    },
    {
      "timestamp": 1766361600000,
      "grossRevenue": 0,
      "netRevenue": 0
    },
    {
      "timestamp": 1766966400000,
      "grossRevenue": 0,
      "netRevenue": 0
    },
    {
      "timestamp": 1767571200000,
      "grossRevenue": 225240,
      "netRevenue": 192096
    }
  ]
}`)
}
