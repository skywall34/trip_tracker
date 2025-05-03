package models

type Trip struct {
    ID                   int     `json:"id"`
    UserId               int     `json:"user_id"`
    Departure            string  `json:"departure"`
    Arrival              string  `json:"arrival"`
    DepartureTime        uint32  `json:"departure_time"`
    ArrivalTime          uint32  `json:"arrival_time"`
    Airline              string  `json:"airline"`
    FlightNumber         string  `json:"flight_number"`
    Reservation          *string `json:"reservation,omitempty"`
    Terminal             *string `json:"terminal,omitempty"`
    Gate                 *string `json:"gate,omitempty"`
    DepartureLat         float64 `json:"departure_lat"`
    DepartureLon         float64 `json:"departure_lon"`
    ArrivalLat           float64 `json:"arrival_lat"`
    ArrivalLon           float64 `json:"arrival_lon"`
    ArrivalTimezone      *string `json:"arrival_timezone,omitempty"` // Not part of DB, we add this later
    DepartureTimezone    *string `json:"departure_timezone,omitempty"` // Not part of DB, we add this later via timezone reference map
}
