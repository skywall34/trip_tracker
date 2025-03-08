package handlers

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/skywall34/trip-tracker/internal/database"
	m "github.com/skywall34/trip-tracker/internal/middleware"
	"github.com/skywall34/trip-tracker/internal/models"
	"github.com/skywall34/trip-tracker/templates"
)

type GetTripHandler struct {
    tripStore *db.TripStore
}

type GetTripHandlerParams struct {
    TripStore *db.TripStore
}

func NewGetTripHandler(params GetTripHandlerParams) *GetTripHandler {
    return &GetTripHandler{
        tripStore: params.TripStore,
    }
}

/**
INSERT INTO trips (user_id, departure, arrival, departure_time, arrival_time, airline, flight_number, reservation, terminal, gate) VALUES
(1, 'JFK', 'LHR', 1672531200, 1672560000, 'British Airways', 'BA117', 'ABC123', 'T7', 'B22'),
(1, 'SFO', 'NRT', 1740481594, 1740567994, 'ANA', 'NH107', 'XYZ456', 'T3', 'C15'),
(3, 'LAX', 'SYD', 1672693200, 1672756800, 'Qantas', 'QF12', 'QF789', 'T4', 'D5'),
(2, 'ORD', 'CDG', 1672779600, 1672813200, 'Air France', 'AF65', 'AF001', 'T5', 'E8'),
(5, 'MIA', 'YYZ', 1672866000, 1672876800, 'Air Canada', 'AC129', 'AC567', 'T2', 'F3'),
(2, 'DFW', 'DXB', 1672952400, 1673016000, 'Emirates', 'EK222', 'EK999', 'T1', 'G12'),
(4, 'SEA', 'PEK', 1673038800, 1673098800, 'Hainan Airlines', 'HU7962', 'HU123', 'T6', 'H9'),
(8, 'BOS', 'KEF', 1673125200, 1673143200, 'Icelandair', 'FI632', 'FI888', 'T7', 'I7'),
(9, 'ATL', 'JNB', 1673211600, 1673275200, 'Delta', 'DL200', 'DL777', 'T3', 'J4'),
(6, 'DEN', 'MEX', 1673298000, 1673312400, 'Aeromexico', 'AM33', 'AM222', 'T4', 'K6'),
(1, 'DEN', 'MEX', 1740308794, 1740395194, 'Aeromexico', 'AM33', 'AM333', 'T4', 'L2');
**/


func (t *GetTripHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

        userTrips, err := t.tripStore.GetTripsGivenUser(userId)
        if err != nil {
            http.Error(w, "Error getting trips", http.StatusInternalServerError)
            return
        }
        if len(userTrips) == 0 {
            http.Error(w, "No trips found for this user", http.StatusNotFound)
            return
        }

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

