package models

type Place struct {
	ID          int     `json:"id"`
	UserID      int     `json:"user_id"`
	PlaceID     string  `json:"place_id"`       // Google Place ID
	Name        string  `json:"name"`
	Address     *string `json:"address,omitempty"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	VisitDate   uint32  `json:"visit_date"`     // Unix timestamp
	Category    *string `json:"category,omitempty"`
	Notes       *string `json:"notes,omitempty"`
	MarkerColor string  `json:"marker_color"`
	CreatedAt   uint32  `json:"created_at"`
	UpdatedAt   uint32  `json:"updated_at"`
}

// For combining places and trips in the timeline
type TimelineItem struct {
	Type      string `json:"type"` // "place" or "trip"
	Place     *Place `json:"place,omitempty"`
	Trip      *Trip  `json:"trip,omitempty"`
	Timestamp uint32 `json:"timestamp"` // For sorting
}

// Google Places API (New) response structures

// Autocomplete (New) API response
type GooglePlaceAutocomplete struct {
	Suggestions []GooglePlaceSuggestion `json:"suggestions"`
}

type GooglePlaceSuggestion struct {
	PlacePrediction *GooglePlacePrediction `json:"placePrediction,omitempty"`
	QueryPrediction *GoogleQueryPrediction `json:"queryPrediction,omitempty"`
}

type GooglePlacePrediction struct {
	Place          string                    `json:"place"`
	PlaceID        string                    `json:"placeId"`
	Text           GooglePlaceText           `json:"text"`
	StructuredFormat *GoogleStructuredFormat `json:"structuredFormat,omitempty"`
	Types          []string                  `json:"types,omitempty"`
}

type GoogleQueryPrediction struct {
	Text GooglePlaceText `json:"text"`
}

type GooglePlaceText struct {
	Text string `json:"text"`
}

type GoogleStructuredFormat struct {
	MainText      GooglePlaceText `json:"mainText"`
	SecondaryText GooglePlaceText `json:"secondaryText"`
}

// Place Details (New) API response
type GooglePlaceDetails struct {
	ID               string              `json:"id"`
	DisplayName      GooglePlaceText     `json:"displayName"`
	FormattedAddress string              `json:"formattedAddress"`
	Location         GooglePlaceLocation `json:"location"`
	Types            []string            `json:"types"`
}

type GooglePlaceLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
