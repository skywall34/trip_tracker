package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/internal/models"
	"github.com/skywall34/trip-tracker/templates"
)

// TODO: Secondary Index
//  userIndex[trip.UserId] = append(userIndex[trip.UserId], trip)
var trips = map[int]models.Trip{
	1:  {UserId: 1, Departure: "JFK", Arrival: "LHR", DepartureTime: 1672531200, ArrivalTime: 1672560000, Airline: "British Airways", FlightNumber: "BA117", Reservation: "ABC123", Terminal: "T7", Gate: "B22"},
	2:  {UserId: 1, Departure: "SFO", Arrival: "NRT", DepartureTime: 1740481594, ArrivalTime: 1740567994, Airline: "ANA", FlightNumber: "NH107", Reservation: "XYZ456", Terminal: "T3", Gate: "C15"},
	3:  {UserId: 3, Departure: "LAX", Arrival: "SYD", DepartureTime: 1672693200, ArrivalTime: 1672756800, Airline: "Qantas", FlightNumber: "QF12", Reservation: "QF789", Terminal: "T4", Gate: "D5"},
	4:  {UserId: 2, Departure: "ORD", Arrival: "CDG", DepartureTime: 1672779600, ArrivalTime: 1672813200, Airline: "Air France", FlightNumber: "AF65", Reservation: "AF001", Terminal: "T5", Gate: "E8"},
	5:  {UserId: 5, Departure: "MIA", Arrival: "YYZ", DepartureTime: 1672866000, ArrivalTime: 1672876800, Airline: "Air Canada", FlightNumber: "AC129", Reservation: "AC567", Terminal: "T2", Gate: "F3"},
	6:  {UserId: 2, Departure: "DFW", Arrival: "DXB", DepartureTime: 1672952400, ArrivalTime: 1673016000, Airline: "Emirates", FlightNumber: "EK222", Reservation: "EK999", Terminal: "T1", Gate: "G12"},
	7:  {UserId: 4, Departure: "SEA", Arrival: "PEK", DepartureTime: 1673038800, ArrivalTime: 1673098800, Airline: "Hainan Airlines", FlightNumber: "HU7962", Reservation: "HU123", Terminal: "T6", Gate: "H9"},
	8:  {UserId: 8, Departure: "BOS", Arrival: "KEF", DepartureTime: 1673125200, ArrivalTime: 1673143200, Airline: "Icelandair", FlightNumber: "FI632", Reservation: "FI888", Terminal: "T7", Gate: "I7"},
	9:  {UserId: 9, Departure: "ATL", Arrival: "JNB", DepartureTime: 1673211600, ArrivalTime: 1673275200, Airline: "Delta", FlightNumber: "DL200", Reservation: "DL777", Terminal: "T3", Gate: "J4"},
	10: {UserId: 6, Departure: "DEN", Arrival: "MEX", DepartureTime: 1673298000, ArrivalTime: 1673312400, Airline: "Aeromexico", FlightNumber: "AM33", Reservation: "AM222", Terminal: "T4", Gate: "K6"},
	11: {UserId: 1, Departure: "DEN", Arrival: "MEX", DepartureTime: 1740308794, ArrivalTime: 1740395194, Airline: "Aeromexico", FlightNumber: "AM33", Reservation: "AM333", Terminal: "T4", Gate: "L2"},
}


func HtmxTripsHandler(w http.ResponseWriter, r *http.Request) {

    filterPast := r.URL.Query().Get("past")
    headerVal := r.Header.Get("HX-Request")

    ctx := r.Context()
    userId, ok := ctx.Value(m.UserKey).(int)
    fmt.Printf("User ID: %d, Filter Past: %s HeaderVal: %s, OK: %t \n", userId, filterPast, headerVal, ok)
    if !ok {
        // redirect to home
        http.Redirect(w, r, "/login", http.StatusSeeOther)
        return
    }

    // if htmx is requesting, return only the fragment
    if headerVal == "true" {
        var userTrips []models.Trip
        for i, trip := range trips {
            // if trip.UserId == userId {
            //     userTrips = append(userTrips, trips[i])
            // }
            if trip.UserId == userId {
                userTrips = append(userTrips, trips[i])
            }
        }
        if len(userTrips) == 0 {
            http.Error(w, "No trips found for this user", http.StatusNotFound)
            return
        }
        // TODO: Add filter for past trips
        if filterPast == "true" {
            renderErr := templates.RenderPastTrips(userTrips).Render(r.Context(), w)
            if renderErr != nil {
                http.Error(w, "Error rendering template", http.StatusInternalServerError)
                return
            }
        } else {
            // We want upcoming, which are trips that are coming in the future (up to 1 week)
            var filteredTrips []models.Trip
            for _, trip := range userTrips {
                // get the current unix time
                currentUnixTime := time.Now().Unix()
                // get the unix time for 1 week from now
                oneWeekFromNow := currentUnixTime + (7 * 24 * 60 * 60)
                if trip.DepartureTime > uint32(currentUnixTime) && trip.DepartureTime < uint32(oneWeekFromNow) {
                    filteredTrips = append(filteredTrips, trip)
                }
            }
            renderErr := templates.RenderTrips(filteredTrips).Render(r.Context(), w)
            if renderErr != nil {
                http.Error(w, "Error rendering template", http.StatusInternalServerError)
                return
            }
        }
    } else {
        // Otherwise, return the full page
        c := templates.TripsPage()
        templates.Layout(c, "Trips").Render(r.Context(), w)
    }
}

func GetTripsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // get the id of the trip to be deleted if it exists
    tripId := r.URL.Query().Get("id")
    if tripId != "" { 
        tripNum, err := strconv.Atoi(tripId)
        if err != nil {
            http.Error(w, "Invalid trip ID, trip ID must be number!", http.StatusBadRequest)
            return
        }
        trip, exists := trips[tripNum]
        if !exists {
            http.Error(w, "Trip not found", http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(trip)    
    } else {
        json.NewEncoder(w).Encode(trips)
    }
}

func GetTripsForUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // get the id of the trip to be deleted if it exists
    var userTrips []models.Trip
    userId := r.URL.Query().Get("user")
    if userId != "" { 
        userIdNum, err := strconv.Atoi(userId)
        if err != nil {
            http.Error(w, "Invalid user ID, user ID must be number!", http.StatusBadRequest)
            return
        }
        for i, trip := range trips {
            if trip.UserId == userIdNum {
                userTrips = append(userTrips, trips[i])
            }
        }
        if len(userTrips) == 0 {
            http.Error(w, "No trips found for this user", http.StatusNotFound)
            return
        }
        json.NewEncoder(w).Encode(userTrips)    
    } else {
        http.Error(w, "Missing user ID", http.StatusBadRequest)
        return
    }
}

func PostTripsHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement POST handler
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse and validate the request body
    var newTrip models.Trip
    err := json.NewDecoder(r.Body).Decode(&newTrip)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Check if UserId exists for this trip
    _, exists := Users[newTrip.UserId]
    if !exists {
        http.Error(w, "User ID does not exist", http.StatusBadRequest)
        return
    }


    // Validation logic 
    if newTrip.Departure == "" || newTrip.Arrival == "" || newTrip.Airline == "" || newTrip.FlightNumber == "" { 
        http.Error(w, "Missing required fields", http.StatusBadRequest)
        marshaled, err := json.MarshalIndent(newTrip, "", "   ")
        if err != nil {
            log.Fatalf("marshaling error: %s", err)
        }
        fmt.Println(string(marshaled))
        return
    }

    if newTrip.ArrivalTime <= newTrip.DepartureTime {
		http.Error(w, "ArrivalTime must be after DepartureTime", http.StatusBadRequest)
		return
	}

    // Validate the request body and add the new trip to the trips slice
    tripId := len(trips) + 1
    trips[tripId] = newTrip

    w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTrip)
}

func EditTripsHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement check to make sure user has access to edit this trip
    // The middleware should therefore not only authenticate the user but also authorize them to edit the specific trip.
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // get the id of the trip to be deleted
    tripId := r.URL.Query().Get("id")
    if tripId == "" {
        http.Error(w, "Missing trip ID", http.StatusBadRequest)
        return
    }

    tripNum, err := strconv.Atoi(tripId)
    if err != nil {
        http.Error(w, "Invalid trip ID, trip ID must be number!", http.StatusBadRequest)
        return
    }

    // Check if the trip exists
	if _, exists := trips[tripNum]; !exists {
		http.Error(w, "Trip not found", http.StatusNotFound)
		return
	}

    // Parse and validate the incoming JSON
	var updatedTrip models.Trip
	err = json.NewDecoder(r.Body).Decode(&updatedTrip)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}


    // Validation logic
	if updatedTrip.Departure == "" || updatedTrip.Arrival == "" {
		http.Error(w, "Departure and Arrival are required fields", http.StatusBadRequest)
		return
	}

	if updatedTrip.DepartureTime == 0 || updatedTrip.ArrivalTime == 0 {
		http.Error(w, "DepartureTime and ArrivalTime are required fields", http.StatusBadRequest)
		return
	}

	if updatedTrip.ArrivalTime <= updatedTrip.DepartureTime {
		http.Error(w, "ArrivalTime must be after DepartureTime", http.StatusBadRequest)
		return
	}

	if updatedTrip.Airline == "" || updatedTrip.FlightNumber == "" {
		http.Error(w, "Airline and FlightNumber are required fields", http.StatusBadRequest)
		return
	}

	// Update the trip
	trips[tripNum] = updatedTrip

	// Respond with the updated trip
	json.NewEncoder(w).Encode(updatedTrip)

}

func DeleteTripsHandler(w http.ResponseWriter, r *http.Request) {
    // TODO: Middleware to make sure user has access to delete this trip
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // get the id of the trip to be deleted
    tripId := r.URL.Query().Get("id")
    if tripId == "" {
        http.Error(w, "Missing trip ID", http.StatusBadRequest)
        return
    }

    tripNum, err := strconv.Atoi(tripId)
    if err != nil {
        http.Error(w, "Invalid trip ID, trip ID must be number!", http.StatusBadRequest)
        return
    }

    // Check if the trip exists
	if _, exists := trips[tripNum]; !exists {
		http.Error(w, "Trip not found", http.StatusNotFound)
		return
	}

	// Delete the trip
	delete(trips, tripNum)

	// Respond with success
	w.WriteHeader(http.StatusNoContent)
}