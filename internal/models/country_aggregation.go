package models

type CountryAggregation struct {
	Name  string `json:"name"`  // e.g., "Jan", "Feb", "2022", etc.
	Count int    `json:"count"` // Number of countries per aggregation (year, month)
}