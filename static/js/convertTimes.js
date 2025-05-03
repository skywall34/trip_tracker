// convertTimes takes the flight times and transforms input and output to match users' timezones
document.addEventListener("DOMContentLoaded", function () {
  let timezoneInput = document.getElementById("timezone");
  if (timezoneInput) {
    timezoneInput.value = Intl.DateTimeFormat().resolvedOptions().timeZone;
  }
  // Set the timezone when the page loads
  convertTimes(); // Run initially to update timestamps

  if (document.body) {
    document.body.addEventListener("htmx:afterSwap", function () {
      convertTimes(); // Run again after HTMX swaps in new content
    });

    // Ensure the timezone is set before sending the form via HTMX
    document.body.addEventListener("htmx:configRequest", function (event) {
      let form = event.detail.elt.closest("form");
      if (!form) {
        return;
      }
      let timezoneInput = document.getElementById("timezone");
      if (timezoneInput) {
        timezoneInput.value = Intl.DateTimeFormat().resolvedOptions().timeZone;
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
    const utcTime = element.getAttribute("data-utc");
    const timeZone = element.getAttribute("data-tz"); // NEW

    if (utcTime && timeZone) {
      const localDate = new Date(utcTime);

      if (!isNaN(localDate.getTime())) {
        const formattedTime = localDate.toLocaleString(undefined, {
          year: "numeric",
          month: "short",
          day: "numeric",
          hour: "2-digit",
          minute: "2-digit",
          second: "2-digit",
          timeZoneName: "short",
          timeZone: timeZone, // 👈 key fix
        });

        element.innerText = formattedTime;
      } else {
        console.error("Invalid Date Format:", utcTime);
        element.innerText = "Invalid Date";
      }
    }
  });
}
