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
  });
});
