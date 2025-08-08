package tests

import (
	"testing"
	"time"

	"github.com/skywall34/trip-tracker/internal/handlers"
	"github.com/skywall34/trip-tracker/internal/models"
)

func TestParseLocalToUTC(t *testing.T) {
	orig := models.AirportTimezoneLookup
	models.AirportTimezoneLookup = map[string]string{"JFK": "America/New_York"}
	t.Cleanup(func() { models.AirportTimezoneLookup = orig })

	got, err := handlers.ParseLocalToUTC("2023-01-02T15:04", "JFK", "America/Los_Angeles")
	if err != nil {
		t.Fatalf("ParseLocalToUTC returned error: %v", err)
	}

	expected := time.Date(2023, time.January, 2, 20, 4, 0, 0, time.UTC)
	if !got.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestParseLocalToUTCFallbackToProvidedTimezone(t *testing.T) {
	orig := models.AirportTimezoneLookup
	models.AirportTimezoneLookup = map[string]string{}
	t.Cleanup(func() { models.AirportTimezoneLookup = orig })

	got, err := handlers.ParseLocalToUTC("2023-01-02T15:04", "XXX", "America/Los_Angeles")
	if err != nil {
		t.Fatalf("ParseLocalToUTC returned error: %v", err)
	}

	expected := time.Date(2023, time.January, 2, 23, 4, 0, 0, time.UTC)
	if !got.Equal(expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}
