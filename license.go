package creemio

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type LicenseStatus string

const (
	LicenseStatusInactive LicenseStatus = "inactive"
	LicenseStatusActive   LicenseStatus = "active"
	LicenseStatusExpired  LicenseStatus = "expired"
	LicenseStatusDisabled LicenseStatus = "disabled"
)

type License struct {
	ID              string           `json:"id"`
	Mode            Mode             `json:"mode"`
	Object          string           `json:"object"`
	Status          string           `json:"status"`
	Key             string           `json:"key"`
	Activation      int              `json:"activation"`
	ActivationLimit int              `json:"activation_limit"`
	ExpiresAt       *time.Time       `json:"expires_at"`
	CreatedAt       time.Time        `json:"created_at"`
	Instance        *LicenseInstance `json:"instance"`
}

type LicenseInstance struct {
	ID        string    `json:"id"`
	Mode      Mode      `json:"mode"`
	Object    string    `json:"object"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type LicenseRequest struct {
	Key        string `json:"key"`
	InstanceID string `json:"instance_id"`
}

type LicenseActivateRequest struct {
	Key          string `json:"key"`
	InstanceName string `json:"instance_name"`
}

type LicenseValidateRequest = LicenseRequest
type LicenseDeactivateRequest = LicenseRequest

type LicenseService struct {
	client *Client
}

func (s *LicenseService) Activate(ctx context.Context, data *LicenseActivateRequest) (*License, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/licenses", "activate")

	if data == nil {
		return nil, nil, errRequiredMissingField
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetUrl, bytes.NewReader(payload))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, newResponse(res), err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, newResponse(res), newAPIError(res.Body)
	}

	var license License
	if err := json.NewDecoder(res.Body).Decode(&license); err != nil {
		return nil, newResponse(res), err
	}

	return &license, newResponse(res), nil
}

func (s *LicenseService) Deactivate(ctx context.Context, data *LicenseDeactivateRequest) (*License, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/licenses", "deactivate")

	if data == nil {
		return nil, nil, errRequiredMissingField
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetUrl, bytes.NewReader(payload))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, newResponse(res), err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, newResponse(res), newAPIError(res.Body)
	}

	var license License
	if err := json.NewDecoder(res.Body).Decode(&license); err != nil {
		return nil, newResponse(res), err
	}

	return &license, newResponse(res), nil
}

func (s *LicenseService) Validate(ctx context.Context, data *LicenseValidateRequest) (*License, *Response, error) {
	targetUrl := makeUrl(s.client.baseURL, "/licenses", "validate")

	if data == nil {
		return nil, nil, errRequiredMissingField
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, targetUrl, bytes.NewReader(payload))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("x-api-key", s.client.apiKey)
	req.Header.Set("Content-Type", "application/json")

	res, err := s.client.httpClient.Do(req)
	if err != nil {
		return nil, newResponse(res), err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return nil, newResponse(res), newAPIError(res.Body)
	}

	var license License
	if err := json.NewDecoder(res.Body).Decode(&license); err != nil {
		return nil, newResponse(res), err
	}

	return &license, newResponse(res), nil
}
