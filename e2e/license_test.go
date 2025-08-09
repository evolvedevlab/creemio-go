package e2e

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/evolvedevlab/creemio-go"
	"github.com/stretchr/testify/assert"
)

func TestLicense_Activate(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	var (
		key  = "XTRJ9-0NZTW-UDN8E-L0EWM-V3NHQ"
		name = "key 1"
	)
	license, res, err := client.Licenses.Activate(context.Background(), &creemio.LicenseActivateRequest{
		Key:          key,
		InstanceName: name,
	})

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(license)
	a.Equal(key, license.Key)
	a.Equal(name, license.Instance.Name)
	fmt.Println(*license.Instance)
}

func TestLicense_Deactivate(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	var (
		key        = "XTRJ9-0NZTW-UDN8E-L0EWM-V3NHQ"
		instanceID = "lki_2dyyxID4FyvaDDFbqothEF"
	)
	license, res, err := client.Licenses.Deactivate(context.Background(), &creemio.LicenseDeactivateRequest{
		Key:        key,
		InstanceID: instanceID,
	})

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(license)
	a.Equal(key, license.Key)
	a.Equal(instanceID, license.Instance.ID)
}

func TestLicense_Validate(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	var (
		key        = "XTRJ9-0NZTW-UDN8E-L0EWM-V3NHQ"
		instanceID = "lki_2dyyxID4FyvaDDFbqothEF"
	)
	license, res, err := client.Licenses.Validate(context.Background(), &creemio.LicenseValidateRequest{
		Key:        key,
		InstanceID: instanceID,
	})

	a.NoError(err)

	// http response
	a.Equal(http.StatusOK, res.Status)

	a.NotNil(license)
	a.Equal(key, license.Key)
	a.Equal(instanceID, license.Instance.ID)
}
