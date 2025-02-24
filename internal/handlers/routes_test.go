package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandler_MethodNotAllowed(t *testing.T) {
	// Test cases for different HTTP methods
	methods := []string{http.MethodGet, http.MethodPut, http.MethodDelete, http.MethodPatch}
	endpoints := []string{
		"/v1/chat/completions",
		"/v1/completions",
		"/v1/embeddings",
		"/v1/moderations",
	}

	h := NewHandler(Config{})

	for _, endpoint := range endpoints {
		for _, method := range methods {
			t.Run(method+"_"+endpoint, func(t *testing.T) {
				// Create request
				req := httptest.NewRequest(method, endpoint, nil)
				w := httptest.NewRecorder()

				// Route request to appropriate handler
				switch endpoint {
				case "/v1/chat/completions":
					h.ChatCompletions(w, req)
				case "/v1/completions":
					h.Completions(w, req)
				case "/v1/embeddings":
					h.Embeddings(w, req)
				case "/v1/moderations":
					h.Moderations(w, req)
				}

				// Assert response
				assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
				assert.Contains(t, w.Body.String(), "Only POST requests are allowed")
			})
		}
	}
}

func TestHandler_ValidPostRequests(t *testing.T) {
	endpoints := []string{
		"/v1/chat/completions",
		"/v1/completions",
		"/v1/embeddings",
		"/v1/moderations",
	}

	h := NewHandler(Config{})

	for _, endpoint := range endpoints {
		t.Run("POST_"+endpoint, func(t *testing.T) {
			// Create a POST request with empty JSON body
			req := httptest.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer([]byte("{}")))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Route request to appropriate handler
			switch endpoint {
			case "/v1/chat/completions":
				h.ChatCompletions(w, req)
			case "/v1/completions":
				h.Completions(w, req)
			case "/v1/embeddings":
				h.Embeddings(w, req)
			case "/v1/moderations":
				h.Moderations(w, req)
			}

			// Since forwardRequest is not implemented in this test,
			// we expect an internal server error
			assert.Equal(t, http.StatusInternalServerError, w.Code)
			assert.Contains(t, w.Body.String(), "Error forwarding request")
		})
	}
}

func TestSetupRoutes(t *testing.T) {
	h := NewHandler(Config{})

	// This test ensures that routes are properly registered
	SetupRoutes(h)

	// Create a test server using the default ServeMux
	server := httptest.NewServer(http.DefaultServeMux)
	defer server.Close()

	// Test each endpoint
	endpoints := []string{
		"/v1/chat/completions",
		"/v1/completions",
		"/v1/embeddings",
		"/v1/moderations",
	}

	client := &http.Client{}

	for _, endpoint := range endpoints {
		t.Run("RouteExists_"+endpoint, func(t *testing.T) {
			// Test that POST requests are accepted
			req, err := http.NewRequest(http.MethodPost, server.URL+endpoint, bytes.NewBuffer([]byte("{}")))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			// We expect an error because the forward request will fail,
			// but the route should exist and accept the request
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)

			// Test that other methods are rejected
			req, err = http.NewRequest(http.MethodGet, server.URL+endpoint, nil)
			require.NoError(t, err)

			resp, err = client.Do(req)
			require.NoError(t, err)
			defer resp.Body.Close()

			assert.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
		})
	}
}
