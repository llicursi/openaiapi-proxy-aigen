package handlers

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jsonResponse(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		expected string
	}{
		{
			name:     "string response",
			data:     map[string]string{"message": "test message"},
			expected: `{"message":"test message"}` + "\n",
		},
		{
			name:     "number response",
			data:     map[string]int{"count": 42},
			expected: `{"count":42}` + "\n",
		},
		{
			name: "complex response",
			data: map[string]interface{}{
				"string": "value",
				"number": 123,
				"bool":   true,
				"nested": map[string]string{
					"key": "value",
				},
			},
			expected: `{"bool":true,"nested":{"key":"value"},"number":123,"string":"value"}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			jsonResponse(w, tt.data)

			// Check Content-Type header
			assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

			// Check response body
			assert.Equal(t, tt.expected, w.Body.String())

			// Verify the response is valid JSON
			var result interface{}
			err := json.Unmarshal(w.Body.Bytes(), &result)
			assert.NoError(t, err, "Response should be valid JSON")
		})
	}
}
