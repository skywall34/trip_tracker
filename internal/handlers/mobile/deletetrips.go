package mobile

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	db "github.com/skywall34/trip-tracker/internal/database"
)

// DeleteTripsHandler handles DELETE requests to remove trips
type DeleteTripsHandler struct {
	tripStore *db.TripStore
}

// DeleteTripsHandlerParams constructor parameters
type DeleteTripsHandlerParams struct {
	TripStore *db.TripStore
}

// NewDeleteTripsHandler creates a new DeleteTripsHandler
func NewDeleteTripsHandler(params DeleteTripsHandlerParams) *DeleteTripsHandler {
	return &DeleteTripsHandler{
		tripStore: params.TripStore,
	}
}

// ServeHTTP handles DELETE /api/v1/trips/{id} requests
func (h *DeleteTripsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from JWT context
	userID, ok := GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Extract trip ID from URL path
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 5 {
		http.Error(w, "Trip ID required", http.StatusBadRequest)
		return
	}
	
	tripIDStr := pathParts[len(pathParts)-1]
	tripID, err := strconv.Atoi(tripIDStr)
	if err != nil {
		response := ApiResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "INVALID_TRIP_ID",
				Message: "Invalid trip ID",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Verify trip exists and belongs to user
	_, err = h.tripStore.GetTripGivenId(tripID, userID)
	if err != nil {
		response := ApiResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "TRIP_NOT_FOUND",
				Message: "Trip not found",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Delete the trip
	err = h.tripStore.DeleteTrip(tripID)
	if err != nil {
		response := ApiResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "DELETE_TRIP_FAILED",
				Message: "Failed to delete trip",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return success response
	response := ApiResponse{
		Success: true,
		Data: map[string]interface{}{
			"message": "Trip deleted successfully",
			"trip_id": tripID,
		},
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}