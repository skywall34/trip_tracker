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

  const customIcon = new L.Icon({
    iconUrl: "/static/images/marker-icon.png",
    iconRetinaUrl: "/static/images/marker-icon-2x.png",
    shadowUrl: "/static/images/marker-shadow.png",
    iconSize: [25, 41],
    iconAnchor: [12, 41],
    popupAnchor: [1, -34],
    tooltipAnchor: [16, -28],
    shadowSize: [41, 41],
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
        L.marker(departure, { icon: customIcon })
          .addTo(map)
          .bindPopup(`Departure: ${trip.departure}`);

        // Add arrival marker
        L.marker(arrival, { icon: customIcon })
          .addTo(map)
          .bindPopup(`Arrival: ${trip.arrival}`);
      });
    });
});
