package models

type AirlineAggregation struct {
	Name string `json:"name"` // e.g., "Jan", "Feb", "2022", etc.
	Count int	`json:"count"`  // Number of airlines per aggregation (year, month)
}