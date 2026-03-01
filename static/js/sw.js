const CACHE_NAME = "mias-trips-v1.1.0";
const urlsToCache = [
  "/fromnto/",
  "/fromnto/statistics",
  "/fromnto/worldmap",
  "/fromnto/static/css/output.css",
  "/fromnto/static/js/htmx.min.js",
  "/fromnto/static/js/convertTimes.js",
  "/fromnto/static/js/modal.js",
  "/fromnto/static/js/response-targets.js",
  "/fromnto/static/js/tabs.js",
  "/fromnto/static/js/leaflet.js",
  "/fromnto/static/js/map.js",
  "/fromnto/static/js/pwa-features.js",
  "/fromnto/static/css/mobile.css",
  "/fromnto/static/icons/icon-192x192.png",
  "/fromnto/static/icons/icon-512x512.png",
];

// Install Service Worker
self.addEventListener("install", (event) => {
  event.waitUntil(
    caches
      .open(CACHE_NAME)
      .then((cache) => cache.addAll(urlsToCache))
      .then(() => self.skipWaiting())
  );
});

// Activate Service Worker
self.addEventListener("activate", (event) => {
  event.waitUntil(
    caches
      .keys()
      .then((cacheNames) => {
        return Promise.all(
          cacheNames.map((cacheName) => {
            if (cacheName !== CACHE_NAME) {
              return caches.delete(cacheName);
            }
          })
        );
      })
      .then(() => self.clients.claim())
  );
});

// Fetch Strategy: Cache First for static assets, Network First for data
self.addEventListener("fetch", (event) => {
  if (event.request.method !== "GET") return;

  const url = new URL(event.request.url);

  // Cache first for static assets
  if (url.pathname.startsWith("/fromnto/static/")) {
    event.respondWith(
      caches.match(event.request).then((response) => {
        return response || fetch(event.request);
      })
    );
    return;
  }

  // Network first for dynamic content
  event.respondWith(
    fetch(event.request)
      .then((response) => {
        // Cache successful responses
        if (response.status === 200) {
          const responseClone = response.clone();
          caches.open(CACHE_NAME).then((cache) => {
            cache.put(event.request, responseClone);
          });
        }
        return response;
      })
      .catch(() => {
        // Fallback to cache when offline
        return caches.match(event.request).then((response) => {
          if (response) {
            return response;
          }
          // Return offline page for navigation requests
          if (event.request.mode === "navigate") {
            return caches.match("/fromnto/offline");
          }
        });
      })
  );
});

// Background Sync for offline actions
self.addEventListener("sync", (event) => {
  if (event.tag === "background-sync") {
    event.waitUntil(syncOfflineActions());
  }
});

async function syncOfflineActions() {
  // Implement offline action sync logic
  console.log("Syncing offline actions...");
}
