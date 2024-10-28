package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{
		"Authorization": []string{"ApiKey 123"},
	}
	APIKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("Failed to get API key from valid params: %v", err)
	}
	if APIKey != "123" {
		t.Errorf("GetAPIKey returned wrong token. Got %v, want %v", APIKey, "123")
	}

	invalidHeaders := http.Header{}
	_, err = GetAPIKey(invalidHeaders)
	if err == nil {
		t.Error("GetAPIKey should have failed for invalid headers")
	}
}
