document.addEventListener("DOMContentLoaded", function () {
  // Handle the close button for forms loaded via HTMX
  document.body.addEventListener("click", function(event) {
    if (event.target && (event.target.id === "close-trip-form" || event.target.id === "close-login-form")) {
      // Find the parent container and clear it
      const manualCreateDiv = document.getElementById("manual-create");
      if (manualCreateDiv) {
        manualCreateDiv.innerHTML = "";
      }
      
      // Also handle the old form if it exists (on the dedicated create trip page)
      const createTripFormDiv = document.getElementById("create-trip-form");
      if (createTripFormDiv) {
        createTripFormDiv.classList.add("hidden");
      }
    }
  });

  // Legacy support for the dedicated create trip page
  let createTripElement = document.getElementById("create-trip");
  if (createTripElement) {
    const addTripBtn = document.getElementById("add-trip-btn");
    const createTripForm = document.getElementById("create-trip-form");

    // Show modal when clicking the add button
    if (addTripBtn) {
      addTripBtn.addEventListener("click", function () {
        if (createTripForm) {
          createTripForm.classList.remove("hidden");
        }
      });
    }
  }

  // Clear search results when a place suggestion is clicked
  document.body.addEventListener("click", function(event) {
    const target = event.target.closest('[data-place-suggestion="true"]');
    if (target) {
      const searchResults = document.getElementById("search-results");
      const searchInput = document.getElementById("place-search-input");
      if (searchResults) {
        searchResults.innerHTML = "";
      }
      if (searchInput) {
        searchInput.value = "";
      }
    }
  });

  // Close place modal when cancel button is clicked
  document.body.addEventListener("click", function(event) {
    if (event.target && event.target.id === "close-place-modal") {
      const modal = document.getElementById("add-place-modal");
      if (modal) {
        modal.classList.add("hidden");
        modal.classList.remove("flex");
      }
    }
  });

  // Automatically hide forms after form submission via HTMX
  document.body.addEventListener("htmx:afterRequest", function (event) {
    // Clear the manual create div on the home page
    const manualCreateDiv = document.getElementById("manual-create");
    if (manualCreateDiv && event.detail.xhr.responseURL && event.detail.xhr.responseURL.includes("/trips")) {
      manualCreateDiv.innerHTML = "";
    }

    // Hide the modal on the dedicated create trip page
    const createTripFormDiv = document.getElementById("create-trip-form");
    if (createTripFormDiv) {
      createTripFormDiv.classList.add("hidden");
    }

    // Hide the add place modal after successful place creation
    const addPlaceModal = document.getElementById("add-place-modal");
    if (addPlaceModal && event.detail.xhr.responseURL && event.detail.xhr.responseURL.includes("/places")) {
      // Check if the request was successful (status 2xx)
      if (event.detail.successful) {
        addPlaceModal.classList.add("hidden");
        addPlaceModal.classList.remove("flex");
      }
    }
  });
});
