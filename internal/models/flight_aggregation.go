package models

type FlightAggregation struct {
	Label string `json:"label"` // e.g., "Jan", "Feb", "2022", etc.
	Count int	`json:"count"`  // Number of flights
}