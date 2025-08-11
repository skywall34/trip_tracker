document.addEventListener("DOMContentLoaded", function () {
  let mapElement = document.getElementById("map");
  if (!mapElement) {
    return; // Exit if the map element does not exist
  }

  let map = L.map("map").setView([20, 0], 2);
  if (!map) {
    return; // Exit if the map element does not exist
  }

  // Use dark theme map tiles
  L.tileLayer("https://{s}.basemaps.cartocdn.com/dark_all/{z}/{x}/{y}{r}.png", {
    maxZoom: 18,
    attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors &copy; <a href="https://carto.com/attributions">CARTO</a>',
    subdomains: 'abcd'
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

        L.polyline([departure, arrival], { 
          color: "#37f5c0", 
          weight: 3,
          opacity: 0.8,
          dashArray: "5, 10"
        })
          .addTo(map)
          .bindPopup(`<div class="text-white font-semibold">${trip.airline} Flight ${trip.flight_number}</div>`);

        // Add departure marker
        L.marker(departure, { icon: customIcon })
          .addTo(map)
          .bindPopup(`<div class="text-white font-medium">✈️ Departure<br><span class="text-mint-400">${trip.departure}</span></div>`);

        // Add arrival marker
        L.marker(arrival, { icon: customIcon })
          .addTo(map)
          .bindPopup(`<div class="text-white font-medium">🛬 Arrival<br><span class="text-mint-400">${trip.arrival}</span></div>`);
      });
    });
});
