package creemio

import (
	"encoding/json"
	"errors"
)

var (
	errRequiredMissingField        = errors.New("missing required fields")
	errRequiredFieldProductID      = errors.New("product_id is required")
	errRequiredFieldSubscriptionID = errors.New("subscription_id is required")
)

// API Error response from creemio
type APIError struct {
	TraceID string `json:"trace_id"`
	Status  int    `json:"status"`
	Err     string `json:"error"`
	Message any    `json:"message"`
}

func newAPIError(data []byte) error {
	var apiErr APIError
	if err := json.Unmarshal(data, &apiErr); err != nil {
		return err
	}

	return &apiErr
}

func (e *APIError) Error() string {
	return e.Err
}
