package handlers

import (
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

// NewHandler creates a new Handler with the given configuration
func NewHandler(config Config) *Handler {
	// Setup rotating logger using config
	logWriter := &lumberjack.Logger{
		Filename:   config.Logging.Filename,
		MaxSize:    config.Logging.MaxSize,
		MaxBackups: config.Logging.MaxBackups,
		MaxAge:     config.Logging.MaxAge,
		Compress:   config.Logging.Compress,
	}

	// Create multi-writer for both file and console
	multiWriter := io.MultiWriter(os.Stdout, logWriter)

	// Create logger
	logger := log.New(multiWriter, "", log.LstdFlags)

	return &Handler{
		config: config,
		logger: logger,
	}
}

// forwardRequest handles the proxying of HTTP requests to the configured forward URL.
// It copies all headers, method, and body from the original request to the forwarded request,
// then copies the response back to the original client.
//
// Parameters:
//   - w: http.ResponseWriter to write the response back to the client
//   - r: *http.Request containing the original request to be forwarded
//
// Returns:
//   - error: nil if successful, otherwise contains the error encountered
func (h *Handler) forwardRequest(w http.ResponseWriter, r *http.Request) error {
	h.logger.Printf("Forwarding request to %s", h.config.Server.ForwardURL)

	client := &http.Client{}

	url := h.config.Server.ForwardURL + r.URL.Path
	req, err := http.NewRequest(r.Method, url, r.Body)
	if err != nil {
		return err
	}

	req.Header = r.Header
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	return err
}
