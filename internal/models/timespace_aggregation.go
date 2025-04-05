package models


type TimeSpaceAggregation struct {
	TotalHours float32 `json:"total_hours"`
	TotalKm int `json:"total_km"` // Total kilometers traveled
}