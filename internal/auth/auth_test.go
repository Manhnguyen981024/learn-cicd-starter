package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header := http.Header{}
	if _, err := GetAPIKey(header); err == nil {
		t.Fatalf("Expected: ErrNoAuthHeaderIncluded but got %v", err)
	}
	header.Add("Authorization", "ABC123")

	if _, err := GetAPIKey(header); err == nil {
		t.Fatalf("Expected: Emalformed authorizatio but got %v", err)
	}

	expectedApiKey := "abc123"
	header.Set("Authorization", "ApiKey "+expectedApiKey)
	actualApiKey, err := GetAPIKey(header)
	if err != nil {
		t.Fatalf("Expected: get API key success with key abc123 but got an error %v", err)
	}

	if expectedApiKey != actualApiKey {
		t.Fatalf("Expected: get API key with value %v but got an %v", expectedApiKey, actualApiKey)
	}
	t.Log("Test success")
}
