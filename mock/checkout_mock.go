package mock

import "net/http"

func HandleGetCheckout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetCheckoutResponse())
}

func HandlePostCheckout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write(GetCheckoutResponse())
}

func GetCheckoutResponse() []byte {
	return []byte(`{
  "id": "1",
  "mode": "test",
  "object": "checkout",
  "status": "pending",
  "request_id": "req_1234567890",
  "product": "prod_987654321",
  "units": 1,
  "order": {
    "id": "order_1234567890",
    "mode": "test",
    "object": "order",
    "customer": "cust_1234567890",
    "product": "prod_987654321",
    "transaction": "tx_1234567890",
    "discount": "dis_1234567890",
    "amount": 2000,
    "sub_total": 1800,
    "tax_amount": 200,
    "discount_amount": 100,
    "amount_due": 1900,
    "amount_paid": 1900,
    "currency": "EUR",
    "fx_amount": 15,
    "fx_currency": "EUR",
    "fx_rate": 1.2,
    "status": "pending",
    "type": "recurring",
    "affiliate": "aff_123",
    "created_at": "2023-09-13T00:00:00Z",
    "updated_at": "2023-09-13T00:00:00Z"
  },
  "subscription": "sub_1234567890",
  "customer": "cust_1234567890",
  "custom_fields": [
    {
      "type": "text",
      "key": "custom_key",
      "label": "Custom Label",
      "optional": true,
      "text": {
        "max_length": 123,
        "min_length": 123
      }
    }
  ],
  "checkout_url": "https://creem.io/checkout/some-id",
  "success_url": "https://example.com/return",
  "feature": [
    {
      "license": {
        "id": "lic_1234567890",
        "mode": "test",
        "object": "license",
        "status": "active",
        "key": "ABC123-XYZ456-XYZ456-XYZ456",
        "activation": 5,
        "activation_limit": 1,
        "expires_at": "2023-09-13T00:00:00Z",
        "created_at": "2023-09-13T00:00:00Z",
        "instance": {
          "id": "inst_1234567890",
          "mode": "test",
          "object": "license-instance",
          "name": "My Customer License Instance",
          "status": "active",
          "created_at": "2023-09-13T00:00:00Z"
        }
      }
    }
  ],
  "metadata": {
    "userId": "user_123",
    "visitCount": 42,
    "lastVisit": "2023-04-01"
  }
}`)
}
