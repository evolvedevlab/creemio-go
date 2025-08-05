package mock

import "net/http"

func HandlePostUpgradeSubscription(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetSubscriptionResponse())
}

func HandlePostCancelSubscription(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetSubscriptionResponse())
}

func HandlePostUpdateSubscription(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetSubscriptionResponse())
}

func HandleGetSubscription(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetSubscriptionResponse())
}

func GetSubscriptionResponse() []byte {
	return []byte(`{
  "id": "sub_abc123",
  "mode": "test",
  "object": "subscription",
  "product": {
    "id": "prod_123",
    "mode": "test",
    "object": "product",
    "name": "Pro Plan",
    "description": "This is a sample product description.",
    "image_url": "https://example.com/image.jpg",
    "features": [
      {
        "id": "feat_001",
        "type": "feature",
        "description": "Get access to discord server."
      }
    ],
    "price": 400,
    "currency": "EUR",
    "billing_type": "recurring",
    "billing_period": "every-month",
    "status": "active",
    "tax_mode": "inclusive",
    "tax_category": "saas",
    "product_url": "https://creem.io/product/prod_123123123123",
    "default_success_url": "https://example.com/?status=successful",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  },
  "customer": {
    "id": "cus_456",
    "mode": "test",
    "object": "customer",
    "email": "user@example.com",
    "name": "John Doe",
    "country": "US",
    "created_at": "2023-01-01T00:00:00Z",
    "updated_at": "2023-01-01T00:00:00Z"
  },
  "items": [
    {
      "id": "item_789",
      "mode": "test",
      "object": "subscription_item",
      "product_id": "prod_123",
      "price_id": "price_123",
      "units": 123
    }
  ],
  "collection_method": "charge_automatically",
  "status": "active",
  "last_transaction_id": "tran_3e6Z6TzvHKdsjEgXnGDEp0",
  "last_transaction": {
    "id": "tran_3e6Z6TzvHKdsjEgXnGDEp0",
    "mode": "test",
    "object": "transaction",
    "amount": 2000,
    "amount_paid": 2000,
    "discount_amount": 2000,
    "currency": "EUR",
    "type": "recurring",
    "tax_country": "US",
    "tax_amount": 2000,
    "status": "paid",
    "refunded_amount": 0,
    "order": "order_123",
    "subscription": "sub_abc123",
    "customer": "cus_456",
    "description": "Monthly subscription payment",
    "period_start": 1720000000,
    "period_end": 1722592000,
    "created_at": 1720000000
  },
  "last_transaction_date": "2024-09-12T12:34:56Z",
  "next_transaction_date": "2024-10-12T12:34:56Z",
  "current_period_start_date": "2024-09-12T12:34:56Z",
  "current_period_end_date": "2024-10-12T12:34:56Z",
  "canceled_at": null,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-09-12T12:34:56Z"
}`)
}
