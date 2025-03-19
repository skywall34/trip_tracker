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

  fetch("/api/trips")
    .then((response) => response.json())
    .then((data) => {
      data.forEach((trip) => {
        let latlngs = [
          [trip.departure_lat, trip.departure_lon],
          [trip.arrival_lat, trip.arrival_lon],
        ];
        L.polyline(latlngs, { color: "blue", weight: 2 })
          .addTo(map)
          .bindPopup(`${trip.airline} Flight ${trip.flight_number}`);
      });
    });
});
