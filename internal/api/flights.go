package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

// APIResponse represents the overall response structure
type FlightsAPIResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       []Flight   `json:"data"`
}

// Pagination represents pagination details
type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
	Total  int `json:"total"`
}

// Flight represents the flight details
type Flight struct {
	FlightDate   string     `json:"flight_date"`
	FlightStatus string     `json:"flight_status"`
	Departure    Airport    `json:"departure"`
	Arrival      Airport    `json:"arrival"`
	Airline      Airline    `json:"airline"`
	FlightInfo   FlightInfo `json:"flight"`
	Aircraft     *Aircraft  `json:"aircraft,omitempty"` // omitempty will remove the key from json if value is null
	Live         *LiveData  `json:"live,omitempty"` // * allows distinguishing between nil and an actual value even if empty string or zero
}

// Airport represents airport details for departure and arrival
type Airport struct {
	Airport        string     `json:"airport"`
	Timezone       string     `json:"timezone"`
	IATA           string     `json:"iata"`
	ICAO           string     `json:"icao"`
	Terminal       *string    `json:"terminal,omitempty"`
	Gate           *string    `json:"gate,omitempty"`
	Baggage        *string    `json:"baggage,omitempty"`
	Delay          *int       `json:"delay,omitempty"`
	Scheduled      time.Time  `json:"scheduled"`
	Estimated      time.Time  `json:"estimated"`
	Actual         *time.Time `json:"actual,omitempty"`
	EstimatedRunway *time.Time `json:"estimated_runway,omitempty"`
	ActualRunway   *time.Time `json:"actual_runway,omitempty"`
}

// Airline represents the airline details
type Airline struct {
	Name string `json:"name"`
	IATA string `json:"iata"`
	ICAO string `json:"icao"`
}

// FlightInfo represents the flight details
type FlightInfo struct {
	Number     string      `json:"number"`
	IATA       string      `json:"iata"`
	ICAO       string      `json:"icao"`
	CodeShared interface{} `json:"codeshared,omitempty"` // Change type if structured data is expected
}

// Aircraft represents aircraft details
type Aircraft struct {
	Registration string `json:"registration"`
	IATA         string `json:"iata"`
	ICAO         string `json:"icao"`
	ICAO24       string `json:"icao24"`
}

// LiveData represents live flight data
type LiveData struct {
	Updated        time.Time `json:"updated"`
	Latitude       float64   `json:"latitude"`
	Longitude      float64   `json:"longitude"`
	Altitude       float64   `json:"altitude"`
	Direction      float64   `json:"direction"`
	SpeedHorizontal float64  `json:"speed_horizontal"`
	SpeedVertical  float64   `json:"speed_vertical"`
	IsGround       bool      `json:"is_ground"`
}

const FlightsAPIURL = "https://api.aviationstack.com/v1/flights"



// We'll use this function to get current status of flights
// flight_iata example: "DL171"
// Limit will always be one
func GetFlight(flightIATA string) ( *FlightsAPIResponse, error ){
	accessKey := os.Getenv("API_ACCESS_KEY")
	params := url.Values{}
	params.Add("access_key", accessKey)
	params.Add("flight_iata", flightIATA) // flight IATA code to get the status of a specific flight
	params.Add("limit", "1") // Limit to 1 to get the most recent flight status

	// Construct full URL with query parameters
	fullURL := fmt.Sprintf("%s?%s", FlightsAPIURL, params.Encode())

	// Send the request
	resp, err := http.Get(fullURL) // Using http.Get for simplicity since we don't need custom headers here
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 response codes
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}

	// Decode JSON response
	var apiResponse FlightsAPIResponse

	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}

	if len(apiResponse.Data) == 0 {
		return nil, fmt.Errorf("no flights found for IATA %s", flightIATA)
	}

	// Print the response
	fmt.Printf("Response: %+v\n", apiResponse)

	return &apiResponse, nil
}