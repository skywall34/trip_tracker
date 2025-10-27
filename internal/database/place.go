package database

import (
	"database/sql"
	"sort"
	"time"

	m "github.com/skywall34/trip-tracker/internal/models"
)

type PlaceStore struct {
	db *sql.DB
}

func NewPlaceStore(db *sql.DB) *PlaceStore {
	return &PlaceStore{db: db}
}

// CreatePlace inserts a new place
func (p *PlaceStore) CreatePlace(place m.Place) (int, error) {
	now := uint32(time.Now().Unix())

	query := `
        INSERT INTO places (
            user_id, place_id, name, address, latitude, longitude,
            visit_date, category, notes, marker_color, created_at, updated_at
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `

	result, err := p.db.Exec(
		query,
		place.UserID,
		place.PlaceID,
		place.Name,
		place.Address,
		place.Latitude,
		place.Longitude,
		place.VisitDate,
		place.Category,
		place.Notes,
		place.MarkerColor,
		now,
		now,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// GetPlacesForUser retrieves all places for a user
func (p *PlaceStore) GetPlacesForUser(userID int) ([]m.Place, error) {
	query := `
        SELECT
            id, user_id, place_id, name, address, latitude, longitude,
            visit_date, category, notes, marker_color, created_at, updated_at
        FROM places
        WHERE user_id = ?
        ORDER BY visit_date DESC
    `

	rows, err := p.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var places []m.Place
	for rows.Next() {
		var place m.Place
		err := rows.Scan(
			&place.ID,
			&place.UserID,
			&place.PlaceID,
			&place.Name,
			&place.Address,
			&place.Latitude,
			&place.Longitude,
			&place.VisitDate,
			&place.Category,
			&place.Notes,
			&place.MarkerColor,
			&place.CreatedAt,
			&place.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}

	return places, nil
}

// GetPlaceByID retrieves a single place by ID
func (p *PlaceStore) GetPlaceByID(placeID, userID int) (m.Place, error) {
	query := `
        SELECT
            id, user_id, place_id, name, address, latitude, longitude,
            visit_date, category, notes, marker_color, created_at, updated_at
        FROM places
        WHERE id = ? AND user_id = ?
    `

	var place m.Place
	err := p.db.QueryRow(query, placeID, userID).Scan(
		&place.ID,
		&place.UserID,
		&place.PlaceID,
		&place.Name,
		&place.Address,
		&place.Latitude,
		&place.Longitude,
		&place.VisitDate,
		&place.Category,
		&place.Notes,
		&place.MarkerColor,
		&place.CreatedAt,
		&place.UpdatedAt,
	)

	return place, err
}

// UpdatePlace updates an existing place
func (p *PlaceStore) UpdatePlace(place m.Place) error {
	query := `
        UPDATE places
        SET name = ?, address = ?, visit_date = ?, category = ?,
            notes = ?, marker_color = ?, updated_at = ?
        WHERE id = ? AND user_id = ?
    `

	now := uint32(time.Now().Unix())

	_, err := p.db.Exec(
		query,
		place.Name,
		place.Address,
		place.VisitDate,
		place.Category,
		place.Notes,
		place.MarkerColor,
		now,
		place.ID,
		place.UserID,
	)

	return err
}

// DeletePlace deletes a place
func (p *PlaceStore) DeletePlace(placeID, userID int) error {
	query := `DELETE FROM places WHERE id = ? AND user_id = ?`
	_, err := p.db.Exec(query, placeID, userID)
	return err
}

// GetPlacesFilteredByYear gets places for a specific year
func (p *PlaceStore) GetPlacesFilteredByYear(userID int, year string) ([]m.Place, error) {
	query := `
        SELECT
            id, user_id, place_id, name, address, latitude, longitude,
            visit_date, category, notes, marker_color, created_at, updated_at
        FROM places
        WHERE user_id = ?
        AND strftime('%Y', datetime(visit_date, 'unixepoch')) = ?
        ORDER BY visit_date DESC
    `

	rows, err := p.db.Query(query, userID, year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var places []m.Place
	for rows.Next() {
		var place m.Place
		err := rows.Scan(
			&place.ID,
			&place.UserID,
			&place.PlaceID,
			&place.Name,
			&place.Address,
			&place.Latitude,
			&place.Longitude,
			&place.VisitDate,
			&place.Category,
			&place.Notes,
			&place.MarkerColor,
			&place.CreatedAt,
			&place.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}

	return places, nil
}

// GetPlacesFilteredByCategory gets places for a specific category
func (p *PlaceStore) GetPlacesFilteredByCategory(userID int, category string) ([]m.Place, error) {
	query := `
        SELECT
            id, user_id, place_id, name, address, latitude, longitude,
            visit_date, category, notes, marker_color, created_at, updated_at
        FROM places
        WHERE user_id = ? AND category = ?
        ORDER BY visit_date DESC
    `

	rows, err := p.db.Query(query, userID, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var places []m.Place
	for rows.Next() {
		var place m.Place
		err := rows.Scan(
			&place.ID,
			&place.UserID,
			&place.PlaceID,
			&place.Name,
			&place.Address,
			&place.Latitude,
			&place.Longitude,
			&place.VisitDate,
			&place.Category,
			&place.Notes,
			&place.MarkerColor,
			&place.CreatedAt,
			&place.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}

	return places, nil
}

// GetPlaceStats returns statistics about places
func (p *PlaceStore) GetPlaceStats(userID int) (map[string]int, error) {
	stats := make(map[string]int)

	// Total places
	var totalPlaces int
	err := p.db.QueryRow(`SELECT COUNT(*) FROM places WHERE user_id = ?`, userID).Scan(&totalPlaces)
	if err != nil {
		return nil, err
	}
	stats["total_places"] = totalPlaces

	// Places by category
	rows, err := p.db.Query(`
        SELECT category, COUNT(*)
        FROM places
        WHERE user_id = ? AND category IS NOT NULL
        GROUP BY category
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		var count int
		if err := rows.Scan(&category, &count); err != nil {
			return nil, err
		}
		stats[category] = count
	}

	return stats, nil
}

// GetCombinedTimeline returns both trips and places sorted by date
func (p *PlaceStore) GetCombinedTimeline(userID int, tripStore *TripStore) ([]m.TimelineItem, error) {
	var timeline []m.TimelineItem

	// Get all places
	places, err := p.GetPlacesForUser(userID)
	if err != nil {
		return nil, err
	}

	// Get all trips
	trips, err := tripStore.GetTripsGivenUser(userID)
	if err != nil {
		return nil, err
	}

	// Combine into timeline items
	for i := range places {
		timeline = append(timeline, m.TimelineItem{
			Type:      "place",
			Place:     &places[i],
			Timestamp: places[i].VisitDate,
		})
	}

	for i := range trips {
		timeline = append(timeline, m.TimelineItem{
			Type:      "trip",
			Trip:      &trips[i],
			Timestamp: trips[i].DepartureTime,
		})
	}

	// Sort by timestamp (descending - newest first)
	sort.Slice(timeline, func(i, j int) bool {
		return timeline[i].Timestamp > timeline[j].Timestamp
	})

	return timeline, nil
}
