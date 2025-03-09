document.addEventListener("DOMContentLoaded", function () {
  console.log("DOM fully loaded, running convertTimes()");
  convertTimes(); // Run initially to update timestamps

  if (document.body) {
    document.body.addEventListener("htmx:afterSwap", function () {
      console.log("HTMX Content Updated! Running convertTimes()");
      convertTimes(); // Run again after HTMX swaps in new content
    });

    document.body.addEventListener("htmx:configRequest", function (event) {
      let form = event.detail.elt.closest("form");
      if (!form) return;

      // Find datetime-local inputs
      let departureInput = form.querySelector('input[name="departuretime"]');
      let arrivalInput = form.querySelector('input[name="arrivaltime"]');

      if (departureInput && arrivalInput) {
        let departureDate = new Date(departureInput.value);
        let arrivalDate = new Date(arrivalInput.value);

        if (!isNaN(departureDate) && !isNaN(arrivalDate)) {
          // Convert to UTC and manually format for datetime-local input
          departureInput.value = formatDatetimeLocal(departureDate);
          arrivalInput.value = formatDatetimeLocal(arrivalDate);
        }
      }
    });
  } else {
    console.error("Error: document.body is null");
  }
});

// Function to format date correctly for datetime-local input (YYYY-MM-DDTHH:MM:SS)
function formatDatetimeLocal(date) {
  let year = date.getUTCFullYear();
  let month = String(date.getUTCMonth() + 1).padStart(2, "0");
  let day = String(date.getUTCDate()).padStart(2, "0");
  let hours = String(date.getUTCHours()).padStart(2, "0");
  let minutes = String(date.getUTCMinutes()).padStart(2, "0");
  let seconds = String(date.getUTCSeconds()).padStart(2, "0");

  return `${year}-${month}-${day}T${hours}:${minutes}:${seconds}`;
}

function convertTimes() {
  document.querySelectorAll(".time-convert").forEach((element) => {
    let utcTime = element.getAttribute("data-utc");

    if (utcTime) {
      let localDate = new Date(utcTime);

      if (!isNaN(localDate.getTime())) {
        let formattedTime = localDate.toLocaleString(undefined, {
          year: "numeric",
          month: "short",
          day: "numeric",
          hour: "2-digit",
          minute: "2-digit",
          second: "2-digit",
          timeZoneName: "short",
        });
        element.innerText = formattedTime;
      } else {
        console.error("Invalid Date Format:", utcTime);
        element.innerText = "Invalid Date";
      }
    }
  });
}
