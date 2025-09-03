package mobile

import (
	"encoding/json"
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
)

// GetTripsHandler handles GET requests for trips list
type GetTripsHandler struct {
	tripStore *db.TripStore
}

// GetTripsHandlerParams constructor parameters
type GetTripsHandlerParams struct {
	TripStore *db.TripStore
}

// NewGetTripsHandler creates a new GetTripsHandler
func NewGetTripsHandler(params GetTripsHandlerParams) *GetTripsHandler {
	return &GetTripsHandler{
		tripStore: params.TripStore,
	}
}

// ServeHTTP handles GET /api/v1/trips requests
func (h *GetTripsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from JWT context
	userID, ok := GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get trips for user
	trips, err := h.tripStore.GetTripsGivenUser(userID)
	if err != nil {
		response := ApiResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "FETCH_TRIPS_FAILED",
				Message: "Failed to fetch trips",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return successful response
	response := ApiResponse{
		Success: true,
		Data:    trips,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}