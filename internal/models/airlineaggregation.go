package models

type AirlineAggregation struct {
	Label string `json:"label"` // e.g., "Jan", "Feb", "2022", etc.
	Count int	`json:"count"`  // Number of airlines per aggregation (year, month)
}