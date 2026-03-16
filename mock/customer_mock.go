package mock

import "net/http"

func HandleGetCustomer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetCustomerResponse())
}

func HandleGetCustomerList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetCustomerListResponse())
}

func HandleGetBillingPortalURL(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"customer_portal_link": "https://creem.io/my-orders/login/xxxxxxxxxx"}`))
}

func GetCustomerResponse() []byte {
	return []byte(`{
  "id": "cus_abc123",
  "mode": "test",
  "object": "customer",
  "email": "user@example.com",
  "name": "John Doe",
  "country": "US",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-01T00:00:00Z"
}`)
}

func GetCustomerListResponse() []byte {
	return []byte(`{
  "items": [
    {
      "id": "cus_abc123<string>",
      "mode": "test",
      "object": "customer",
      "email": "user@example.com",
      "country": "US",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-01T00:00:00Z",
      "name": "John Doe"
    }
  ],
  "pagination": {
    "total_records": 0,
    "total_pages": 0,
    "current_page": 1,
    "next_page": 2,
    "prev_page": null
  }
}`)
}
