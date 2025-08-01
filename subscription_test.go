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

func TestSubscriptions_Cancel(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostCancelSubscription))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	subID := "1"
	resp, res, err := c.Subscriptions.Cancel(context.Background(), "1")

	a.NoError(err)
	a.NotNil(resp)
	a.NotNil(res)
	a.Equal(fmt.Sprintf("/%s/subscriptions/%s/cancel", APIVersion, subID), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)

	var expectedSub Subscription
	err = json.Unmarshal(mock.GetSubscriptionResponse(), &expectedSub)

	a.NoError(err)
	a.Equal(expectedSub, *resp)

}

func TestSubscriptions_CancelWithError(t *testing.T) {
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

	resp, res, err := c.Subscriptions.Cancel(context.Background(), "1")

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}

func TestSubscriptions_Update(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostUpdateSubscription))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	subID := "1"
	resp, res, err := c.Subscriptions.Update(context.Background(), &UpdateSubscriptionRequest{
		SubscriptionID: subID,
		UpdateBehavior: ProrationCharge,
		Items: []SubscriptionItem{
			{
				ID:        subID,
				Mode:      "test",
				Object:    "subscription",
				ProductID: "1",
				PriceID:   "1",
				Units:     1,
			},
		},
	})

	a.NoError(err)
	a.NotNil(resp)
	a.NotNil(res)
	a.Equal(fmt.Sprintf("/%s/subscriptions/%s", APIVersion, subID), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)

	var expectedSub Subscription
	err = json.Unmarshal(mock.GetSubscriptionResponse(), &expectedSub)

	a.NoError(err)
	a.Equal(expectedSub, *resp)
}

func TestSubscriptions_UpdateWithMissingSubscriptionID(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostUpdateSubscription))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Subscriptions.Update(context.Background(), &UpdateSubscriptionRequest{})

	a.Error(err)
	a.EqualError(err, errRequiredFieldSubscriptionID.Error())
	a.Nil(resp)
	a.Nil(res)
}

func TestSubscriptions_UpdateWithError(t *testing.T) {
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

	resp, res, err := c.Subscriptions.Update(context.Background(), &UpdateSubscriptionRequest{
		SubscriptionID: "1",
	})

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}

func TestSubscriptions_Get(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandleGetSubscription))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	subID := "sub_abc123"
	url, err := url.Parse(fmt.Sprintf("/%s/subscriptions", APIVersion))
	if err != nil {
		panic(err)
	}
	q := url.Query()
	q.Set("subscription_id", subID)
	url.RawQuery = q.Encode()

	resp, res, err := c.Subscriptions.Get(context.Background(), subID)

	a.NoError(err)
	a.Equal(url.RequestURI(), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)
	a.NotNil(resp)

	var expectedSub Subscription
	err = json.Unmarshal(mock.GetSubscriptionResponse(), &expectedSub)

	a.NoError(err)
	a.Equal(expectedSub, *resp)
}

func TestSubscriptions_GetWithError(t *testing.T) {
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

	resp, res, err := c.Subscriptions.Get(context.Background(), "1")

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}
