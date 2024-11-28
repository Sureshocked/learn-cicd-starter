package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// Test case 1: Valid API key
	t.Run("valid api key", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("Authorization", "ApiKey testkey123")
		key, err := GetAPIKey(headers)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if key != "testkey123" {
			t.Errorf("Expected key 'testkey123', got '%s'", key)
		}
	})

	// Test case 2: Missing header
	t.Run("missing auth header", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("Expected ErrNoAuthHeaderIncluded, got %v", err)
		}
	})

	// Test case 3: Malformed header
	t.Run("malformed header", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("Authorization", "InvalidFormat")
		_, err := GetAPIKey(headers)
		if err == nil || err.Error() != "malformed authorization header" {
			t.Errorf("Expected malformed authorization header error, got %v", err)
		}
	})
}
