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

  // Create custom icons for different airport types
  const layoverIcon = new L.Icon({
    iconUrl: "/static/images/marker-icon.png",
    iconRetinaUrl: "/static/images/marker-icon-2x.png",
    shadowUrl: "/static/images/marker-shadow.png",
    iconSize: [30, 49],
    iconAnchor: [15, 49],
    popupAnchor: [1, -34],
    tooltipAnchor: [16, -28],
    shadowSize: [41, 41],
  });

  fetch("/api/trips")
    .then((response) => response.json())
    .then((data) => {
      // Check if data has the expected structure - if not, assume old API format
      if (!data.hasOwnProperty('standalone_trips') || !data.hasOwnProperty('connecting_trips')) {
        // Convert old format to new format
        const oldData = data;
        data = {
          standalone_trips: Array.isArray(oldData) ? oldData : [],
          connecting_trips: []
        };
      }

      // Ensure connecting_trips is never null
      if (data.connecting_trips === null) {
        data.connecting_trips = [];
      }

      const airportMarkers = new Map(); // Track unique airports

      // Process standalone trips
      if (data.standalone_trips && Array.isArray(data.standalone_trips)) {
        data.standalone_trips.forEach((trip) => {
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

        // Track airports for unique markers
        if (!airportMarkers.has(trip.departure)) {
          airportMarkers.set(trip.departure, {
            position: departure,
            code: trip.departure,
            type: 'departure'
          });
        }
        if (!airportMarkers.has(trip.arrival)) {
          airportMarkers.set(trip.arrival, {
            position: arrival,
            code: trip.arrival,
            type: 'arrival'
          });
        }
      });
      }

      // Process connecting trips
      if (data.connecting_trips && Array.isArray(data.connecting_trips)) {
        data.connecting_trips.forEach((conn) => {
        let departure = [conn.FromTrip.departure_lat, conn.FromTrip.departure_lon];
        let layover = [conn.FromTrip.arrival_lat, conn.FromTrip.arrival_lon];
        let arrival = [conn.ToTrip.arrival_lat, conn.ToTrip.arrival_lon];

        // First leg - solid line
        L.polyline([departure, layover], {
          color: "#37f5c0",
          weight: 4,
          opacity: 0.9
        })
          .addTo(map)
          .bindPopup(`<div class="text-white font-semibold">Leg 1: ${conn.FromTrip.airline} ${conn.FromTrip.flight_number}<br><span class="text-mint-400">${conn.FromTrip.departure} → ${conn.FromTrip.arrival}</span></div>`);

        // Second leg - solid line
        L.polyline([layover, arrival], {
          color: "#37f5c0",
          weight: 4,
          opacity: 0.9
        })
          .addTo(map)
          .bindPopup(`<div class="text-white font-semibold">Leg 2: ${conn.ToTrip.airline} ${conn.ToTrip.flight_number}<br><span class="text-mint-400">${conn.ToTrip.departure} → ${conn.ToTrip.arrival}</span></div>`);

        // Track airports
        if (!airportMarkers.has(conn.FromTrip.departure)) {
          airportMarkers.set(conn.FromTrip.departure, {
            position: departure,
            code: conn.FromTrip.departure,
            type: 'departure'
          });
        }
        if (!airportMarkers.has(conn.FromTrip.arrival)) {
          airportMarkers.set(conn.FromTrip.arrival, {
            position: layover,
            code: conn.FromTrip.arrival,
            type: 'layover'
          });
        }
        if (!airportMarkers.has(conn.ToTrip.arrival)) {
          airportMarkers.set(conn.ToTrip.arrival, {
            position: arrival,
            code: conn.ToTrip.arrival,
            type: 'arrival'
          });
        }
        });
      }

      // Add unique airport markers
      airportMarkers.forEach((airport) => {
        const icon = airport.type === 'layover' ? layoverIcon : customIcon;
        const emoji = airport.type === 'layover' ? '🔄' :
                     airport.type === 'departure' ? '✈️' : '🛬';
        const label = airport.type === 'layover' ? 'Layover' :
                     airport.type === 'departure' ? 'Departure' : 'Arrival';

        L.marker(airport.position, { icon: icon })
          .addTo(map)
          .bindPopup(`<div class="text-white font-medium">${emoji} ${label}<br><span class="text-mint-400">${airport.code}</span></div>`);
      });
    });
});
