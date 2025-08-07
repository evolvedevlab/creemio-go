package mock

import "net/http"

func HandleGetDiscount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetDiscountResponse())
}

func HandlePostCreateDiscount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetDiscountResponse())
}

func HandlePostDeleteDiscount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetDiscountResponse())
}

func HandleDeleteDiscount(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetDiscountResponse())
}

func GetDiscountResponse() []byte {
	return []byte(`{  
	"id": "dis_6SIeFZHfREPdOsopCVegdD",
	"object": "discount",
	"store_id": "sto_34Tj56CUJ9XoKfEkFFXs0s",
	"status": "active",
	"name": "KING100",
	"code": "KING100",
	"type": "percentage",
	"percentage": 100,
	"duration": "once",
	"applies_to_products": [
		"prod_10rxEcEUn5fYkSBpXWhdjR",
		"prod_2c3Gjarm2xbSOAj8bgepT5"
	],
	"created_at": "2025-08-04T15:45:30.725Z",
	"mode": "test"
  }`)
}
