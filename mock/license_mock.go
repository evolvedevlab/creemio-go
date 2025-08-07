package mock

import "net/http"

func HandlePostValidateLicense(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetLicenseResponse())
}

func HandlePostActivateLicense(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetLicenseResponse())
}

func HandlePostDeactivateLicense(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(GetLicenseResponse())
}

func GetLicenseResponse() []byte {
	return []byte(`{
  "id": "lic_abc123",
  "mode": "test",
  "object": "license",
  "status": "active",
  "key": "ABC123-XYZ456-XYZ456-XYZ456",
  "activation": 5,
  "activation_limit": 1,
  "expires_at": "2023-09-13T00:00:00Z",
  "created_at": "2023-09-13T00:00:00Z",
  "instance": {
    "id": "inst_456xyz",
    "mode": "test",
    "object": "license-instance",
    "name": "My Customer License Instance",
    "status": "active",
    "created_at": "2023-09-13T00:00:00Z"
  }
}`)
}
