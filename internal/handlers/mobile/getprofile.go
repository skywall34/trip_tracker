package mobile

import (
	"encoding/json"
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
)

// GetProfileHandler handles GET requests for user profile
type GetProfileHandler struct {
	userStore *db.UserStore
}

// GetProfileHandlerParams constructor parameters
type GetProfileHandlerParams struct {
	UserStore *db.UserStore
}

// NewGetProfileHandler creates a new GetProfileHandler
func NewGetProfileHandler(params GetProfileHandlerParams) *GetProfileHandler {
	return &GetProfileHandler{
		userStore: params.UserStore,
	}
}

// ProfileResponse represents the user profile data returned to mobile
type ProfileResponse struct {
	ID           int    `json:"id"`
	Username     *string `json:"username,omitempty"`
	FirstName    *string `json:"first_name,omitempty"`
	LastName     *string `json:"last_name,omitempty"`
	Email        string  `json:"email"`
	AuthProvider string  `json:"auth_provider"`
	CreatedAt    string  `json:"created_at"` // formatted as string for mobile
}

// ServeHTTP handles GET /api/v1/profile requests
func (h *GetProfileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from JWT context
	userID, ok := GetUserIDFromContext(r.Context())
	if !ok {
		response := ApiResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "UNAUTHORIZED",
				Message: "User not authenticated",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get user details from database
	user, err := h.userStore.GetUserGivenID(userID)
	if err != nil {
		response := ApiResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "FETCH_PROFILE_FAILED",
				Message: "Failed to fetch user profile",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Convert to response format
	profile := ProfileResponse{
		ID:           user.ID,
		Email:        user.Email,
		AuthProvider: user.AuthProvider,
		CreatedAt:    user.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}

	// Add optional fields if they exist
	if user.Username != "" {
		profile.Username = &user.Username
	}
	if user.FirstName != "" {
		profile.FirstName = &user.FirstName
	}
	if user.LastName != "" {
		profile.LastName = &user.LastName
	}

	// Return successful response
	response := ApiResponse{
		Success: true,
		Data:    profile,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}