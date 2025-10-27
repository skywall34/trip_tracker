package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/skywall34/trip-tracker/internal/models"
	"github.com/skywall34/trip-tracker/templates"
)

type GetPlaceSearchHandler struct{}

type GetPlaceSearchHandlerParams struct{}

func NewGetPlaceSearchHandler(params GetPlaceSearchHandlerParams) *GetPlaceSearchHandler {
	return &GetPlaceSearchHandler{}
}

func (h *GetPlaceSearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		// Return empty results
		templates.PlaceSearchResults([]models.GooglePlaceSuggestion{}).Render(r.Context(), w)
		return
	}

	// Call Google Places Autocomplete API (New)
	apiURL := "https://places.googleapis.com/v1/places:autocomplete"

	// Create request body
	requestBody := map[string]interface{}{
		"input": query,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Goog-Api-Key", os.Getenv("GOOGLE_PLACES_API_KEY"))

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

	var result models.GooglePlaceAutocomplete
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		http.Error(w, "Error parsing API response", http.StatusInternalServerError)
		return
	}

	// Return results as HTML using templ
	err = templates.PlaceSearchResults(result.Suggestions).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
