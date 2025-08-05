package mock

import "net/http"

func HandleGetTransaction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetTransactionResponse())
}

func GetTransactionResponse() []byte {
	return []byte(`{
  "items": [
    {
      "id": "tran_75a91IxFrn1VRZKdhFkzQ0",
      "object": "transaction",
      "amount": 4900,
      "amount_paid": 0,
      "currency": "USD",
      "type": "invoice",
      "tax_country": "IN",
      "tax_amount": 0,
      "discount_amount": 4900,
      "status": "paid",
      "refunded_amount": null,
      "order": "ord_3tfjICgjCz9EDPd06vvCbj",
      "subscription": "sub_1dBhaI9z4LYBWexl7J0RVE",
      "customer": "cust_4tyb89MPnfXrumvx0kwXdG",
      "description": "Subscription payment",
      "period_start": 1754323156000,
      "period_end": 1785859156000,
      "created_at": 1754323160037,
      "mode": "test"
    },
    {
      "id": "tran_5LywaKvSflFkyHChpDv7W5",
      "object": "transaction",
      "amount": 4900,
      "amount_paid": 0,
      "currency": "USD",
      "type": "invoice",
      "tax_country": "IN",
      "tax_amount": 0,
      "discount_amount": 4900,
      "status": "paid",
      "refunded_amount": null,
      "order": "ord_1x53viuAO3xtV45ZypVOAz",
      "subscription": "sub_1oGmra1vt3G0vwJ40LkNOd",
      "customer": "cust_4tyb89MPnfXrumvx0kwXdG",
      "description": "Subscription payment",
      "period_start": 1754322396000,
      "period_end": 1785858396000,
      "created_at": 1754322400167,
      "mode": "test"
    }
  ],
  "pagination": {
    "total_records": 2,
    "total_pages": 1,
    "current_page": 1,
    "next_page": null,
    "prev_page": null
  }
}`)
}
