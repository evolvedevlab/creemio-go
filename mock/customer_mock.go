package mock

import "net/http"

func HandleGetCustomer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetCustomerResponse())
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
