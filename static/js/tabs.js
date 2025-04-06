document.addEventListener("DOMContentLoaded", function () {
    const slider = document.getElementById("tab-slider");
    const tabUpcoming = document.getElementById("tab-upcoming");
    const tabPast = document.getElementById("tab-past");
  
    if (!slider || !tabUpcoming || !tabPast) return; // Exit if not on Trips page
  
    const setActiveTab = (activeBtn, inactiveBtn) => {
      activeBtn.classList.add("text-green-700", "font-bold");
      inactiveBtn.classList.remove("text-green-700", "font-bold");
    };
  
    tabUpcoming.addEventListener("click", () => {
      slider.style.transform = "translateX(0%)";
      setActiveTab(tabUpcoming, tabPast);
    });
  
    tabPast.addEventListener("click", () => {
      slider.style.transform = "translateX(100%)";
      setActiveTab(tabPast, tabUpcoming);
    });
  });