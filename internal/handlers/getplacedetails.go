package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/skywall34/trip-tracker/internal/models"
	"github.com/skywall34/trip-tracker/templates"
)

type GetPlaceDetailsHandler struct{}

type GetPlaceDetailsHandlerParams struct{}

func NewGetPlaceDetailsHandler(params GetPlaceDetailsHandlerParams) *GetPlaceDetailsHandler {
	return &GetPlaceDetailsHandler{}
}

func (h *GetPlaceDetailsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	placeID := r.URL.Query().Get("place_id")
	if placeID == "" {
		http.Error(w, "place_id parameter required", http.StatusBadRequest)
		return
	}

	// Call Google Place Details API (New)
	apiURL := fmt.Sprintf("https://places.googleapis.com/v1/places/%s", placeID)

	// Create HTTP request
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	// Set headers
	req.Header.Set("X-Goog-Api-Key", os.Getenv("GOOGLE_PLACES_API_KEY"))
	req.Header.Set("X-Goog-FieldMask", "id,displayName,formattedAddress,location,types")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error calling Google API", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Check for API errors
	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Google API error: %d", resp.StatusCode), http.StatusInternalServerError)
		return
	}

	var result models.GooglePlaceDetails
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, "Error parsing API response", http.StatusInternalServerError)
		return
	}

	// Auto-categorize based on types
	category := ""
	for _, t := range result.Types {
		switch t {
		case "restaurant", "food":
			category = "Restaurant"
		case "lodging":
			category = "Hotel"
		case "museum":
			category = "Museum"
		case "park":
			category = "Park"
		case "tourist_attraction":
			category = "Landmark"
		}
		if category != "" {
			break
		}
	}

	// Render the modal form with populated data
	err = templates.AddPlaceModalForm(
		result.ID,
		result.DisplayName.Text,
		result.FormattedAddress,
		result.Location.Latitude,
		result.Location.Longitude,
		category,
	).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
