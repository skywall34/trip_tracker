package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
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
func GetFlights(client *http.Client, dep string, arr string, limit int, offset int) (flights FlightsAPIResponse){
	accessKey := os.Getenv("API_ACCESS_KEY")
	params := url.Values{}
	params.Add("access_key", accessKey)
	params.Add("limit", strconv.Itoa(limit))
	params.Add("offset", strconv.Itoa(offset))
	// Airpot IATA code
	params.Add("dep_iata", dep)
	params.Add("arr_iata", arr)

	// Construct full URL with query parameters
	fullURL := fmt.Sprintf("%s?%s", FlightsAPIURL, params.Encode())

	// Create the request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check for non-200 response codes
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected response code: %d", resp.StatusCode)
	}

	// Decode JSON response
	var apiResponse FlightsAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		log.Fatalf("Error decoding JSON: %v", err)
	}

	// Print the response
	fmt.Printf("Response: %+v\n", apiResponse)

	return apiResponse
}