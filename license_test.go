package creemio

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evolvedevlab/creemio-go/mock"
	"github.com/stretchr/testify/assert"
)

func TestLicenses_Activate(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostActivateLicense))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	var (
		licenseKey  = "key-1"
		licenseName = "license-1"
	)

	resp, res, err := c.Licenses.Activate(context.Background(), &LicenseActivateRequest{
		Key:          licenseKey,
		InstanceName: licenseName,
	})

	a.NoError(err)
	a.NotNil(resp)
	a.NotNil(res)
	a.Equal(fmt.Sprintf("/%s/licenses/activate", APIVersion), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)

	var expected License
	err = json.Unmarshal(mock.GetLicenseResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)

}

func TestLicenses_ActivateWithMissingRequiredField(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostActivateLicense))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Licenses.Activate(context.Background(), nil)

	a.Error(err)
	a.EqualError(err, errRequiredMissingField.Error())
	a.Nil(resp)
	a.Nil(res)
}

func TestLicenses_ActivateWithError(t *testing.T) {
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

	resp, res, err := c.Licenses.Activate(context.Background(), &LicenseActivateRequest{
		Key:          "key-1",
		InstanceName: "license 1",
	})

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}

func TestLicenses_Deactivate(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostDeactivateLicense))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	var (
		licenseID  = "1"
		licenseKey = "key-1"
	)

	resp, res, err := c.Licenses.Deactivate(context.Background(), &LicenseDeactivateRequest{
		InstanceID: licenseID,
		Key:        licenseKey,
	})

	a.NoError(err)
	a.NotNil(resp)
	a.NotNil(res)
	a.Equal(fmt.Sprintf("/%s/licenses/deactivate", APIVersion), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)

	var expected License
	err = json.Unmarshal(mock.GetLicenseResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)

}

func TestLicenses_DeactivateWithMissingRequiredField(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostDeactivateLicense))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Licenses.Deactivate(context.Background(), nil)

	a.Error(err)
	a.EqualError(err, errRequiredMissingField.Error())
	a.Nil(resp)
	a.Nil(res)
}

func TestLicenses_DeactivateWithError(t *testing.T) {
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

	resp, res, err := c.Licenses.Deactivate(context.Background(), &LicenseDeactivateRequest{
		InstanceID: "1",
		Key:        "key-1",
	})

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}

func TestLicenses_Validate(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostValidateLicense))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	var (
		licenseKey = "key-1"
		licenseID  = "1"
	)

	resp, res, err := c.Licenses.Validate(context.Background(), &LicenseValidateRequest{
		InstanceID: licenseID,
		Key:        licenseKey,
	})

	a.NoError(err)
	a.NotNil(resp)
	a.NotNil(res)
	a.Equal(fmt.Sprintf("/%s/licenses/validate", APIVersion), res.RequestURL.RequestURI())
	a.Equal(http.StatusOK, res.Status)

	var expected License
	err = json.Unmarshal(mock.GetLicenseResponse(), &expected)

	a.NoError(err)
	a.Equal(expected, *resp)
}

func TestLicenses_ValidateWithMissingRequiredField(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	s := httptest.NewServer(http.HandlerFunc(mock.HandlePostValidateLicense))
	defer s.Close()

	c := New(
		WithBaseURL(s.URL),
		WithAPIKey(""),
	)

	resp, res, err := c.Licenses.Validate(context.Background(), nil)

	a.Error(err)
	a.EqualError(err, errRequiredMissingField.Error())
	a.Nil(resp)
	a.Nil(res)
}

func TestLicenses_ValidateWithError(t *testing.T) {
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

	resp, res, err := c.Licenses.Validate(context.Background(), &LicenseValidateRequest{
		InstanceID: "1",
		Key:        "key-1",
	})

	a.Error(err)
	a.Nil(resp)
	a.NotNil(res)
	a.Equal(http.StatusInternalServerError, res.Status)
}
