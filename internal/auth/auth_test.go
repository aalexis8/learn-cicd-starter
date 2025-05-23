package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name           string
		headers        http.Header
		expectedAPIKey string
		expectError    bool
	}{
		{
			name: "valid API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-api-key"},
			},
			expectedAPIKey: "my-secret-api-key",
			expectError:    true,
		},
		{
			name: "missing API key",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expectedAPIKey: "",
			expectError:    false,
		},
		{
			name: "invalid header format",
			headers: http.Header{
				"Authorization": []string{"Bearer my-secret-api-key"},
			},
			expectedAPIKey: "",
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.expectError {
				t.Errorf("GetAPIKey() error = %v, expectError %v", err, tt.expectError)
				return
			}
			if apiKey != tt.expectedAPIKey {
				t.Errorf("GetAPIKey() = %v, expected %v", apiKey, tt.expectedAPIKey)
			}
		})
	}
}
