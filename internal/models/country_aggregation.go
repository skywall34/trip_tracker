package models

type CountryAggregation struct {
	Label  string `json:"label"`  // e.g., "Jan", "Feb", "2022", etc.
	Country string `json:"country"` // Country name
	Count int    `json:"count"` // Number of countries per aggregation (year, month)
}