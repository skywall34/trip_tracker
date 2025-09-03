package mobile

import (
	"encoding/json"
	"net/http"
)

// Refresh token handler
type RefreshTokenHandler struct {
	*MobileAuthHandler
}

func NewRefreshTokenHandler(authHandler *MobileAuthHandler) *RefreshTokenHandler {
	return &RefreshTokenHandler{
		MobileAuthHandler: authHandler,
	}
}

func (h *RefreshTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response := LoginResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "INVALID_REQUEST",
				Message: "Invalid request body",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Verify refresh token
	claims, err := h.verifyToken(req.RefreshToken)
	if err != nil {
		response := LoginResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "INVALID_REFRESH_TOKEN",
				Message: "Invalid or expired refresh token",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Generate new tokens
	tokens, err := h.generateTokens(claims.UserID, claims.Email)
	if err != nil {
		response := LoginResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "TOKEN_GENERATION_FAILED",
				Message: "Failed to generate new tokens",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return new tokens
	response := struct {
		Success bool           `json:"success"`
		Data    *TokenResponse `json:"data,omitempty"`
		Error   *ErrorResponse `json:"error,omitempty"`
	}{
		Success: true,
		Data:    tokens,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}