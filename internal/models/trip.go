package models

type Trip struct {
    ID int64 `json:"id"`
    UserId int64 `json:"user_id"`
    Departure string `json:"departure"`
    Arrival string `json:"arrival"`
    DepartureTime uint32 `json:"departure_time"`
    ArrivalTime uint32 `json:"arrival_time"`
    Airline string `json:"airline"`
    FlightNumber string `json:"flight_number"`
    Reservation string `json:"reservation"`
    Terminal string `json:"terminal"`
    Gate string `json:"gate"`
    DepartureLat  float64 `json:"departure_lat"`
    DepartureLon  float64 `json:"departure_lon"`
    ArrivalLat float64 `json:"arrival_lat"`
    ArrivalLon float64 `json:"arrival_lon"`
}
