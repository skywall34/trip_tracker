package models

/**

CREATE TABLE airports (
    iata_code TEXT PRIMARY KEY,
    name TEXT,
    country TEXT,
    latitude REAL,
    longitude REAL,
    region TEXT
);
**/

type Airport struct {
	IataCode  string `json:"iata_code"`
	Name	  string `json:"name"`
	Country   string `json:"country"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Region    string `json:"region"`
}
