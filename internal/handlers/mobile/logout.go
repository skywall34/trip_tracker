package mobile

import (
	"encoding/json"
	"net/http"
)

// LogoutHandler handles POST requests for mobile logout
type LogoutHandler struct {
}

// LogoutHandlerParams constructor parameters
type LogoutHandlerParams struct {
}

// NewLogoutHandler creates a new LogoutHandler
func NewLogoutHandler(params LogoutHandlerParams) *LogoutHandler {
	return &LogoutHandler{}
}

// ServeHTTP handles POST /api/v1/mobile/auth/logout requests
func (h *LogoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// For mobile logout, we just need to return success
	// The actual token cleanup happens on the client side
	// when the client calls the logout action in Redux

	// In a more complex setup, you might want to:
	// 1. Invalidate the refresh token in a database
	// 2. Add the access token to a blacklist
	// 3. Perform other cleanup operations

	response := ApiResponse{
		Success: true,
		Data:    map[string]string{"message": "Logged out successfully"},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}