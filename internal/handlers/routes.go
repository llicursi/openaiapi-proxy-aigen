package handlers

import "net/http"

// ChatCompletions godoc
// @Summary      Forward chat completions requests
// @Description  Forwards chat completion requests to the configured endpoint
// @Tags         openai
// @Accept       json
// @Produce      json
// @Success      200  {object}  ApiResponse
// @Failure      500  {object}  ApiResponse
// @Router       /chat/completions [post]
func (h *Handler) ChatCompletions(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("ChatCompletions request from %s: %s %s", r.RemoteAddr, r.Method, r.URL.Path)
	if err := h.forwardRequest(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse(w, ApiResponse{
			Message: "Error forwarding request: " + err.Error(),
			Object:  "error",
		})
	}
}

// Completions godoc
// @Summary      Forward completions requests
// @Description  Forwards completion requests to the configured endpoint
// @Tags         openai
// @Accept       json
// @Produce      json
// @Success      200  {object}  ApiResponse
// @Failure      500  {object}  ApiResponse
// @Router       /completions [post]
func (h *Handler) Completions(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("Completions request from %s: %s %s", r.RemoteAddr, r.Method, r.URL.Path)
	if err := h.forwardRequest(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse(w, ApiResponse{
			Message: "Error forwarding request: " + err.Error(),
			Object:  "error",
		})
	}
}

// Embeddings godoc
// @Summary      Forward embeddings requests
// @Description  Forwards embedding requests to the configured endpoint
// @Tags         openai
// @Accept       json
// @Produce      json
// @Success      200  {object}  ApiResponse
// @Failure      500  {object}  ApiResponse
// @Router       /embeddings [post]
func (h *Handler) Embeddings(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("Embeddings request from %s: %s %s", r.RemoteAddr, r.Method, r.URL.Path)
	if err := h.forwardRequest(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse(w, ApiResponse{
			Message: "Error forwarding request: " + err.Error(),
			Object:  "error",
		})
	}
}

// Moderations godoc
// @Summary      Forward moderations requests
// @Description  Forwards moderation requests to the configured endpoint
// @Tags         openai
// @Accept       json
// @Produce      json
// @Success      200  {object}  ApiResponse
// @Failure      500  {object}  ApiResponse
// @Router       /moderations [post]
func (h *Handler) Moderations(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("Moderations request from %s: %s %s", r.RemoteAddr, r.Method, r.URL.Path)
	if err := h.forwardRequest(w, r); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonResponse(w, ApiResponse{
			Message: "Error forwarding request: " + err.Error(),
			Object:  "error",
		})
	}
}

func SetupRoutes(h *Handler) {
	http.HandleFunc("/v1/chat/completions", h.ChatCompletions)
	http.HandleFunc("/v1/completions", h.Completions)
	http.HandleFunc("/v1/embeddings", h.Embeddings)
	http.HandleFunc("/v1/moderations", h.Moderations)
}
