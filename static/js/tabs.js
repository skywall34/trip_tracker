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
      slider.classList.remove("translate-x-full");
      slider.classList.add("translate-x-0");
      setActiveTab(tabUpcoming, tabPast);
    });
  
    tabPast.addEventListener("click", () => {
      slider.classList.remove("translate-x-0");
      slider.classList.add("translate-x-full");
      setActiveTab(tabPast, tabUpcoming);
    });
  });