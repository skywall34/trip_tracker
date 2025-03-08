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
}
