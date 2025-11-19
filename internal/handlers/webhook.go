package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kha7iq/pingme/internal/dispatcher"
	"github.com/kha7iq/pingme/internal/types"
)

// WebhookHandler handles incoming webhook requests
type WebhookHandler struct {
	dispatcher *dispatcher.Dispatcher
}

// NewWebhookHandler creates a new webhook handler
func NewWebhookHandler() *WebhookHandler {
	return &WebhookHandler{
		dispatcher: dispatcher.New(),
	}
}

// WebhookResponse represents the response sent back
type WebhookResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// ServeHTTP implements http.Handler interface
func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests
	if r.Method != http.MethodPost {
		h.sendError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.sendError(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Parse JSON payload
	var req types.WebhookRequest
	if err := json.Unmarshal(body, &req); err != nil {
		h.sendError(w, fmt.Sprintf("Invalid JSON: %v", err), http.StatusBadRequest)
		return
	}

	// Validate request
	if err := h.validateRequest(&req); err != nil {
		h.sendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Log incoming request
	log.Printf("Webhook received: service=%s, message_length=%d", req.Service, len(req.Message))

	// Dispatch message to appropriate service
	if err := h.dispatcher.Dispatch(r.Context(), &req); err != nil {
		log.Printf("Failed to dispatch message to %s: %v", req.Service, err)
		h.sendError(w, fmt.Sprintf("Failed to send message: %v", err), http.StatusInternalServerError)
		return
	}

	// Send success response
	h.sendSuccess(w, fmt.Sprintf("Message sent successfully via %s", req.Service))
}

// validateRequest validates the webhook request
func (h *WebhookHandler) validateRequest(req *types.WebhookRequest) error {
	if req.Service == "" {
		return fmt.Errorf("service field is required")
	}
	if req.Message == "" {
		return fmt.Errorf("message field is required")
	}
	return nil
}

// sendSuccess sends a successful JSON response
func (h *WebhookHandler) sendSuccess(w http.ResponseWriter, message string) {
	resp := WebhookResponse{
		Success: true,
		Message: message,
	}
	h.sendJSON(w, resp, http.StatusOK)
}

// sendError sends an error JSON response
func (h *WebhookHandler) sendError(w http.ResponseWriter, errorMsg string, statusCode int) {
	resp := WebhookResponse{
		Success: false,
		Error:   errorMsg,
	}
	h.sendJSON(w, resp, statusCode)
}

// sendJSON sends a JSON response
func (h *WebhookHandler) sendJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}
