package mobile

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	db "github.com/skywall34/trip-tracker/internal/database"
)

// PutTripsHandler handles PUT requests to update trips
type PutTripsHandler struct {
	tripStore *db.TripStore
}

// PutTripsHandlerParams constructor parameters
type PutTripsHandlerParams struct {
	TripStore *db.TripStore
}

// NewPutTripsHandler creates a new PutTripsHandler
func NewPutTripsHandler(params PutTripsHandlerParams) *PutTripsHandler {
	return &PutTripsHandler{
		tripStore: params.TripStore,
	}
}

// ServeHTTP handles PUT /api/v1/trips/{id} requests
func (h *PutTripsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
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

	// Get existing trip to verify ownership
	existingTrip, err := h.tripStore.GetTripGivenId(tripID, userID)
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

	// Parse JSON request body
	var updateData struct {
		Airline       *string `json:"airline,omitempty"`
		FlightNumber  *string `json:"flight_number,omitempty"`
		Origin        *string `json:"origin,omitempty"`
		Destination   *string `json:"destination,omitempty"`
		DepartureTime *uint32 `json:"departure_time,omitempty"`
		ArrivalTime   *uint32 `json:"arrival_time,omitempty"`
		Reservation   *string `json:"reservation,omitempty"`
		Terminal      *string `json:"terminal,omitempty"`
		Gate          *string `json:"gate,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		response := ApiResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "INVALID_REQUEST",
				Message: "Invalid JSON request body",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Update fields if provided
	if updateData.Airline != nil {
		existingTrip.Airline = *updateData.Airline
	}
	if updateData.FlightNumber != nil {
		existingTrip.FlightNumber = *updateData.FlightNumber
	}
	if updateData.Origin != nil {
		existingTrip.Departure = *updateData.Origin
	}
	if updateData.Destination != nil {
		existingTrip.Arrival = *updateData.Destination
	}
	if updateData.DepartureTime != nil {
		existingTrip.DepartureTime = *updateData.DepartureTime
	}
	if updateData.ArrivalTime != nil {
		existingTrip.ArrivalTime = *updateData.ArrivalTime
	}
	if updateData.Reservation != nil {
		existingTrip.Reservation = updateData.Reservation
	}
	if updateData.Terminal != nil {
		existingTrip.Terminal = updateData.Terminal
	}
	if updateData.Gate != nil {
		existingTrip.Gate = updateData.Gate
	}

	// Save to database
	err = h.tripStore.EditTrip(existingTrip)
	if err != nil {
		response := ApiResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "UPDATE_TRIP_FAILED",
				Message: "Failed to update trip",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Return updated trip
	response := ApiResponse{
		Success: true,
		Data:    existingTrip,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}