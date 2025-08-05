package mock

import "net/http"

func HandlePostCreateProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetProductResponse())
}

func HandleGetProductList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetProductListResponse())
}

func GetProductListResponse() []byte {
	return []byte(`{
  "items": [
    {
      "id": "prod_abc123",
      "mode": "test",
      "object": "product",
      "name": "Premium Membership",
      "description": "This is a sample product description.",
      "image_url": "https://example.com/image.jpg",
      "features": [
        {
          "id": "feat_001",
          "type": "discord-access",
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
      "product_url": "https://creem.io/product/prod_abc123",
      "default_success_url": "https://example.com/?status=successful",
      "created_at": "2023-01-01T00:00:00Z",
      "updated_at": "2023-01-15T00:00:00Z"
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

func HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetProductResponse())
}

func GetProductResponse() []byte {
	return []byte(`{
  "id": "prod_abc123",
  "mode": "test",
  "object": "product",
  "name": "Premium Membership",
  "description": "This is a sample product description.",
  "image_url": "https://example.com/image.jpg",
  "features": [
    {
      "id": "feat_001",
      "type": "discord-access",
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
  "product_url": "https://creem.io/product/prod_abc123",
  "default_success_url": "https://example.com/?status=successful",
  "created_at": "2023-01-01T00:00:00Z",
  "updated_at": "2023-01-15T00:00:00Z"
}`)
}
