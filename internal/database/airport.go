package database

import (
	"database/sql"
)

type AirportStore struct {
	db *sql.DB
}

type NewAirportStoreParams struct {
	DB *sql.DB
}

func NewAirportStore(params NewAirportStoreParams) *TripStore {
	return &TripStore{db: params.DB}
}

// Criteria for airport search
// This file should only contain methods to query Airport focused queries
// For example, trip.go focuses on Trip related queries even though it joins with airports
// This helps keep the code organized and maintainable