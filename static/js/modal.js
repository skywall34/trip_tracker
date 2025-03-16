document.addEventListener("DOMContentLoaded", function () {
  const addTripBtn = document.getElementById("add-trip-btn");
  const createTripForm = document.getElementById("create-trip-form");
  const closeTripForm = document.getElementById("close-trip-form");

  // Show modal when clicking the add button
  if (addTripBtn) {
    addTripBtn.addEventListener("click", function () {
      createTripForm.classList.remove("hidden");
    });
  }

  // Hide modal when clicking cancel button
  if (closeTripForm) {
    closeTripForm.addEventListener("click", function () {
      createTripForm.classList.add("hidden");
    });
  }

  // Automatically hide modal after form submission via HTMX
  document.body.addEventListener("htmx:afterRequest", function () {
    createTripForm.classList.add("hidden");
  });
});
