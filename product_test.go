package creemio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/evolvedevlab/creemio-go/mock"
	"github.com/stretchr/testify/assert"
)

func TestProducts_Create(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostCreateProduct))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Products.Create(context.Background(), &CreateProductRequest{
		Name:        "product 1",
		Price:       100,
		Currency:    "USD",
		BillingType: "every-month",
	})

	a.NoError(err)
	a.NotNil(resp)
	a.NotNil(res)
	a.Equal(fmt.Sprintf("/%s/products", APIVersion), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)

	var expected Product
	err = json.Unmarshal(mock.GetProductResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)
}

func TestProducts_CreateWithMissingRequiredField(t *testing.T) {
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

	resp, res, err := c.Products.Create(context.Background(), &CreateProductRequest{})

	a.Error(err)
	a.Nil(resp)
	a.Nil(res)
}

func TestProducts_CreateWithError(t *testing.T) {
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

	resp, res, err := c.Products.Create(context.Background(), &CreateProductRequest{
		Name:        "product 1",
		Price:       100,
		Currency:    "USD",
		BillingType: "every-month",
	})

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}

func TestProducts_List(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetProductList))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	// For comparing the url request url with search params
	url, err := url.Parse(fmt.Sprintf("/%s/products/search", APIVersion))
	if err != nil {
		panic(err)
	}

	var (
		pageNum  = 1
		pageSize = 2
	)

	q := url.Query()
	q.Set("page_number", strconv.Itoa(pageNum))
	q.Set("page_size", strconv.Itoa(pageSize))
	url.RawQuery = q.Encode()

	resp, res, err := c.Products.List(context.Background(), &ProductListQuery{
		PageNumber: pageNum,
		PageSize:   pageSize,
	})

	a.NoError(err)
	a.Equal(url.RequestURI(), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)

	var expected ProductList
	err = json.Unmarshal(mock.GetProductListResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)
}

func TestProducts_ListWithError(t *testing.T) {
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

	resp, res, err := c.Products.List(context.Background(), nil)

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}

func TestProducts_Get(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetProduct))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	prodID := "1"
	// For comparing the url request url with search params
	url, err := url.Parse(fmt.Sprintf("/%s/products", APIVersion))
	if err != nil {
		panic(err)
	}
	q := url.Query()
	q.Set("product_id", prodID)
	url.RawQuery = q.Encode()

	resp, res, err := c.Products.Get(context.Background(), prodID)

	a.NoError(err)
	a.Equal(url.RequestURI(), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)

	var expected Product
	err = json.Unmarshal(mock.GetProductResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)
}

func TestProducts_GetWithError(t *testing.T) {
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

	resp, res, err := c.Products.Get(context.Background(), "1")

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}
