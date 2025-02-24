package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	config := Config{
		Logging: LoggingConfig{
			Filename:   "test.log",
			MaxSize:    10,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		},
	}

	h := NewHandler(config)
	assert.NotNil(t, h)
	assert.NotNil(t, h.logger)
	assert.Equal(t, config, h.config)
}

func TestHandler_ForwardRequest(t *testing.T) {
	// Create a test server to forward requests to
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Test", "test-value")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test response"))
	}))
	defer testServer.Close()

	config := Config{
		Server: ServerConfig{
			ForwardURL: testServer.URL,
		},
	}

	h := NewHandler(config)

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "successful forward",
			method:         "POST",
			path:           "/test",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "forward with different method",
			method:         "GET",
			path:           "/test",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			err := h.forwardRequest(w, req)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Equal(t, "test-value", w.Header().Get("X-Test"))
			assert.Equal(t, "test response", w.Body.String())
		})
	}
}
