package models

type FlightAggregation struct {
	Label string `json:"label"` // e.g., "Jan", "Feb", "2022", etc.
	Count int	`json:"count"`  // Number of flights
	Total int	`json:"total"`  // Total flights for the given label (e.g., total flights in January or total flights in 2022)
}