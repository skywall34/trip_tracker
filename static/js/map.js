document.addEventListener("DOMContentLoaded", function () {
  let mapElement = document.getElementById("map");
  if (!mapElement) {
    return; // Exit if the map element does not exist
  }

  let map = L.map("map").setView([20, 0], 2);
  if (!map) {
    return; // Exit if the map element does not exist
  }
  L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
    maxZoom: 10,
  }).addTo(map);

  L.Icon.Default.mergeOptions({
    iconUrl: "../../images/marker-icon.png",
    iconRetinaUrl: "../../images/marker-icon.png",
    shadowUrl: "../../images/marker-shadow.png",
  });

  fetch("/api/trips")
    .then((response) => response.json())
    .then((data) => {
      data.forEach((trip) => {
        let departure = [trip.departure_lat, trip.departure_lon];
        let arrival = [trip.arrival_lat, trip.arrival_lon];

        L.polyline([departure, arrival], { color: "blue", weight: 2 })
          .addTo(map)
          .bindPopup(`${trip.airline} Flight ${trip.flight_number}`);

        // Add departure marker
        L.marker(departure)
          .addTo(map)
          .bindPopup(`Departure: ${trip.departure_airport}`);

        // Add arrival marker
        L.marker(arrival)
          .addTo(map)
          .bindPopup(`Arrival: ${trip.arrival_airport}`);
      });
    });
});
