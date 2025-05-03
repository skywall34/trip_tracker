package models

import (
	"encoding/json"
	"os"
)

// Country struct represents the ISO country code and name
// Used to determine if a user visited a specific country and to populate on a world map
type AirportTimezone struct {
	Code        string `json:"code"`
    CountryCode string `json:"countryCode"`
    Timezone    string `json:"timezone"`
	Offset      Offset `json:"offset"`
}

type Offset struct {
	Gmt int8 `json:"gmt"`
	Dst int8 `json:"dst"`
}

var AirportTimezoneMap []AirportTimezone

func LoadAirportTimezonesFromFile(path string) error {
    file, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    return json.Unmarshal(file, &AirportTimezoneMap)
}