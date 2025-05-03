package models

import (
	"encoding/json"
	"os"
)

// Country struct represents the ISO country code and name
// Used to determine if a user visited a specific country and to populate on a world map
var AirportTimezoneLookup = map[string]string{}

func LoadAirportTimezonesFromFile(path string) error {
    file, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    return json.Unmarshal(file, &AirportTimezoneLookup)
}