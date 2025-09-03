package mobile

import (
	"encoding/json"
	"net/http"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/models"
)

// PostTripsHandler handles POST requests to create trips
type PostTripsHandler struct {
	tripStore *db.TripStore
}

// PostTripsHandlerParams constructor parameters
type PostTripsHandlerParams struct {
	TripStore *db.TripStore
}

// NewPostTripsHandler creates a new PostTripsHandler
func NewPostTripsHandler(params PostTripsHandlerParams) *PostTripsHandler {
	return &PostTripsHandler{
		tripStore: params.TripStore,
	}
}

// ServeHTTP handles POST /api/v1/trips requests
func (h *PostTripsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get user ID from JWT context
	userID, ok := GetUserIDFromContext(r.Context())
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Parse JSON request body
	var tripData struct {
		Airline       string  `json:"airline"`
		FlightNumber  string  `json:"flight_number"`
		Origin        string  `json:"origin"`
		Destination   string  `json:"destination"`
		DepartureTime uint32  `json:"departure_time"`
		ArrivalTime   uint32  `json:"arrival_time"`
		Reservation   *string `json:"reservation,omitempty"`
		Terminal      *string `json:"terminal,omitempty"`
		Gate          *string `json:"gate,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&tripData); err != nil {
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

	// Create trip model
	trip := m.Trip{
		UserId:        userID,
		Airline:       tripData.Airline,
		FlightNumber:  tripData.FlightNumber,
		Departure:     tripData.Origin,
		Arrival:       tripData.Destination,
		DepartureTime: tripData.DepartureTime,
		ArrivalTime:   tripData.ArrivalTime,
		Reservation:   tripData.Reservation,
		Terminal:      tripData.Terminal,
		Gate:          tripData.Gate,
	}

	// Save to database
	tripID, err := h.tripStore.CreateTrip(trip)
	if err != nil {
		response := ApiResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "CREATE_TRIP_FAILED",
				Message: "Failed to create trip",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	// Get the created trip to return
	createdTrip, err := h.tripStore.GetTripGivenId(int(tripID), userID)
	if err != nil {
		response := ApiResponse{
			Success: false,
			Error: &ErrorResponse{
				Code:    "FETCH_CREATED_TRIP_FAILED",
				Message: "Trip created but failed to fetch details",
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
		Data:    createdTrip,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}