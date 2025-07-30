package creemio

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeUrl(t *testing.T) {
	tests := []struct {
		name     string
		baseURL  string
		path     string
		params   []string
		expected string
	}{
		{
			name:     "basic URL with no params",
			baseURL:  "https://api.example.com",
			path:     "/customers",
			expected: "https://api.example.com/v1/customers",
		},
		{
			name:     "baseURL with trailing slash",
			baseURL:  "https://api.example.com/",
			path:     "/customers",
			expected: "https://api.example.com/v1/customers",
		},
		{
			name:     "path without leading slash",
			baseURL:  "https://api.example.com",
			path:     "customers",
			expected: "https://api.example.com/v1/customers",
		},
		{
			name:     "URL with one param",
			baseURL:  "https://api.example.com",
			path:     "/customers",
			params:   []string{"abc123"},
			expected: "https://api.example.com/v1/customers/abc123",
		},
		{
			name:     "URL with multiple params",
			baseURL:  "https://api.example.com",
			path:     "/products",
			params:   []string{"123", "licenses"},
			expected: "https://api.example.com/v1/products/123/licenses",
		},
		{
			name:     "path with trailing slash",
			baseURL:  "https://api.example.com",
			path:     "/customers/",
			params:   []string{"abc123"},
			expected: "https://api.example.com/v1/customers/abc123",
		},
		{
			name:     "baseURL and path both without slashes",
			baseURL:  "https://api.example.com",
			path:     "customers",
			params:   []string{"123"},
			expected: "https://api.example.com/v1/customers/123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := makeUrl(tt.baseURL, tt.path, tt.params...)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
