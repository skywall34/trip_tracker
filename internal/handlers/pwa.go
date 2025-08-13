package handlers

import (
	"net/http"
)

// PWAManifestHandler serves the web app manifest
type PWAManifestHandler struct{}

func NewPWAManifestHandler() *PWAManifestHandler {
	return &PWAManifestHandler{}
}

func (h *PWAManifestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "public, max-age=604800")

	http.ServeFile(w, r, "static/manifest.json")
}

// ServiceWorkerHandler serves the service worker
type ServiceWorkerHandler struct{}

func NewServiceWorkerHandler() *ServiceWorkerHandler {
	return &ServiceWorkerHandler{}
}

func (h *ServiceWorkerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/javascript")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	http.ServeFile(w, r, "static/js/sw.js")
}

// OfflineHandler serves offline page
type OfflineHandler struct{}

func NewOfflineHandler() *OfflineHandler {
	return &OfflineHandler{}
}

func (h *OfflineHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	offlineHTML := `
	<!DOCTYPE html>
	<html lang="en" class="dark">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Offline - Mia's Trips</title>
		<link rel="stylesheet" href="/static/css/output.css">
	</head>
	<body class="bg-ink-900 text-slate-300 min-h-screen flex items-center justify-center">
		<div class="text-center max-w-md mx-auto px-4">
			<div class="w-24 h-24 bg-emerald-500 rounded-full flex items-center justify-center mx-auto mb-6">
				<span class="text-4xl">✈️</span>
			</div>
			<h1 class="text-3xl font-bold text-emerald-400 mb-4">You're Offline</h1>
			<p class="text-slate-400 mb-6">Your trip data is still available!</p>
			<button onclick="window.location.reload()"
				class="bg-emerald-600 text-white px-6 py-3 rounded-lg hover:bg-emerald-700">
				Try Again
			</button>
		</div>
	</body>
	</html>`

	w.Write([]byte(offlineHTML))
}