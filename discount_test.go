package creemio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/evolvedevlab/creemio-go/mock"
	"github.com/stretchr/testify/assert"
)

func TestDiscounts_Create(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostCreateDiscount))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	// For comparing the url request url with search params
	url, err := url.Parse(fmt.Sprintf("/%s/discounts", APIVersion))
	if err != nil {
		panic(err)
	}

	resp, res, err := c.Discounts.Create(context.Background(), &CreateDiscountRequest{
		Name:              "KING100",
		Type:              DiscountTypePercentage,
		Duration:          DiscountDurationOnce,
		AppliesToProducts: []string{"1", "2"},
	})

	a.NoError(err)
	a.Equal(url.RequestURI(), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)

	var expected Discount
	err = json.Unmarshal(mock.GetDiscountResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)
}

func TestDiscounts_CreateWithMissingRequiredField(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostCreateDiscount))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Discounts.Create(context.Background(), nil)

	a.Error(err)
	a.Nil(resp)
	a.Nil(res)
	a.EqualError(err, errRequiredMissingField.Error())
}

func TestDiscounts_CreateWithError(t *testing.T) {
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

	resp, res, err := c.Discounts.Create(context.Background(), &CreateDiscountRequest{
		Name:              "KING100",
		Type:              DiscountTypePercentage,
		Duration:          DiscountDurationOnce,
		AppliesToProducts: []string{"1", "2"},
	})

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}

func TestDiscounts_Delete(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostDeleteDiscount))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	discountID := "1"
	// For comparing the url request url with search params
	url, err := url.Parse(fmt.Sprintf("/%s/discounts/%s/delete", APIVersion, discountID))
	if err != nil {
		panic(err)
	}

	resp, res, err := c.Discounts.Delete(context.Background(), discountID)

	a.NoError(err)
	a.Equal(url.RequestURI(), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)

	var expected Discount
	err = json.Unmarshal(mock.GetDiscountResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)
}

func TestDiscounts_DeleteWithError(t *testing.T) {
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

	resp, res, err := c.Discounts.Delete(context.Background(), "1")

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}

func TestDiscounts_Get(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetDiscount))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	var (
		DiscountID   = "1"
		DiscountCode = "KING100"
	)

	// For comparing the url request url with search params
	url, err := url.Parse(fmt.Sprintf("/%s/discounts", APIVersion))
	if err != nil {
		panic(err)
	}
	q := url.Query()
	q.Set("discount_id", DiscountID)
	q.Set("discount_code", DiscountCode)
	url.RawQuery = q.Encode()

	resp, res, err := c.Discounts.Get(context.Background(), &DiscountRequestQuery{
		DiscountID:   DiscountID,
		DiscountCode: DiscountCode,
	})

	a.NoError(err)
	a.Equal(url.RequestURI(), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)

	var expected Discount
	err = json.Unmarshal(mock.GetDiscountResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)
}

func TestDiscounts_GetWithMissingRequiredField(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetDiscount))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Discounts.Get(context.Background(), &DiscountRequestQuery{})

	a.Error(err)
	a.Nil(resp)
	a.Nil(res)
	a.EqualError(err, errDiscountNoQuery.Error())
}

func TestDiscounts_GetWithError(t *testing.T) {
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

	resp, res, err := c.Discounts.Get(context.Background(), &DiscountRequestQuery{
		DiscountCode: "KING100",
	})

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}
