// PWA Feature Enhancements
class PWAFeatures {
  constructor() {
    this.initializeFeatures();
  }

  initializeFeatures() {
    this.setupGeolocation();
    this.setupCamera();
    this.setupPullToRefresh();
    this.setupOfflineSync();
    this.setupShareAPI();
  }

  // Geolocation for automatic location detection
  setupGeolocation() {
    const locationBtns = document.querySelectorAll("[data-get-location]");
    locationBtns.forEach((btn) => {
      btn.addEventListener("click", this.getCurrentLocation.bind(this));
    });
  }

  async getCurrentLocation() {
    if (!("geolocation" in navigator)) {
      alert("Geolocation is not supported");
      return;
    }

    try {
      const position = await new Promise((resolve, reject) => {
        navigator.geolocation.getCurrentPosition(resolve, reject, {
          enableHighAccuracy: true,
          timeout: 10000,
          maximumAge: 600000,
        });
      });

      const { latitude, longitude } = position.coords;

      // Reverse geocoding to get airport/city
      const response = await fetch(
        `/api/location?lat=${latitude}&lng=${longitude}`
      );
      const location = await response.json();

      // Fill in form fields
      const originField = document.querySelector("#origin");
      if (originField && location.airport_code) {
        originField.value = location.airport_code;
        htmx.trigger(originField, "change");
      }
    } catch (error) {
      console.error("Error getting location:", error);
      alert("Unable to get your location. Please check permissions.");
    }
  }

  // Camera access for trip photos
  setupCamera() {
    const cameraInputs = document.querySelectorAll(
      'input[type="file"][accept*="image"]'
    );
    cameraInputs.forEach((input) => {
      // Add camera capture attribute for mobile
      input.setAttribute("capture", "camera");

      input.addEventListener("change", this.handleImageCapture.bind(this));
    });
  }

  handleImageCapture(event) {
    const file = event.target.files[0];
    if (file) {
      // Show preview
      const reader = new FileReader();
      reader.onload = (e) => {
        const preview = document.querySelector("#image-preview");
        if (preview) {
          preview.src = e.target.result;
          preview.style.display = "block";
        }
      };
      reader.readAsDataURL(file);

      // Compress image for upload
      this.compressImage(file).then((compressedFile) => {
        // Upload compressed image
        this.uploadImage(compressedFile);
      });
    }
  }

  async compressImage(file, quality = 0.8) {
    const canvas = document.createElement("canvas");
    const ctx = canvas.getContext("2d");
    const img = new Image();

    return new Promise((resolve) => {
      img.onload = () => {
        const maxWidth = 1200;
        const maxHeight = 800;
        let { width, height } = img;

        if (width > height) {
          if (width > maxWidth) {
            height = (height * maxWidth) / width;
            width = maxWidth;
          }
        } else {
          if (height > maxHeight) {
            width = (width * maxHeight) / height;
            height = maxHeight;
          }
        }

        canvas.width = width;
        canvas.height = height;

        ctx.drawImage(img, 0, 0, width, height);

        canvas.toBlob(resolve, "image/jpeg", quality);
      };

      img.src = URL.createObjectURL(file);
    });
  }

  // Pull-to-refresh functionality
  setupPullToRefresh() {
    let startY = 0;
    let pullDistance = 0;
    let isPulling = false;

    const container = document.querySelector(".main-content");
    if (!container) return;

    container.addEventListener("touchstart", (e) => {
      if (window.scrollY === 0) {
        startY = e.touches[0].clientY;
        isPulling = true;
      }
    });

    container.addEventListener("touchmove", (e) => {
      if (!isPulling) return;

      const currentY = e.touches[0].clientY;
      pullDistance = Math.max(0, currentY - startY);

      if (pullDistance > 0) {
        e.preventDefault();
        container.style.setProperty(
          "--pull-distance",
          `${pullDistance * 0.5}px`
        );

        if (pullDistance > 100) {
          container.classList.add("pull-to-refresh-ready");
        }
      }
    });

    container.addEventListener("touchend", () => {
      if (isPulling && pullDistance > 100) {
        this.refreshContent();
      }

      isPulling = false;
      pullDistance = 0;
      container.style.setProperty("--pull-distance", "0");
      container.classList.remove("pull-to-refresh-ready");
    });
  }

  refreshContent() {
    // Trigger HTMX refresh of current page content
    const refreshTarget = document.querySelector("[data-refresh-target]");
    if (refreshTarget) {
      htmx.trigger(refreshTarget, "refresh");
    } else {
      window.location.reload();
    }
  }

  // Offline sync functionality
  setupOfflineSync() {
    // Store offline actions in IndexedDB
    window.addEventListener("online", this.syncOfflineActions.bind(this));

    // Intercept form submissions when offline
    document.addEventListener("htmx:beforeRequest", (event) => {
      if (!navigator.onLine) {
        event.preventDefault();
        this.storeOfflineAction(event.detail);
        this.showOfflineMessage();
      }
    });
  }

  storeOfflineAction(action) {
    const offlineActions = JSON.parse(
      localStorage.getItem("offlineActions") || "[]"
    );
    offlineActions.push({
      ...action,
      timestamp: Date.now(),
    });
    localStorage.setItem("offlineActions", JSON.stringify(offlineActions));
  }

  async syncOfflineActions() {
    const offlineActions = JSON.parse(
      localStorage.getItem("offlineActions") || "[]"
    );

    for (const action of offlineActions) {
      try {
        await fetch(action.url, {
          method: action.method || "POST",
          body: action.body,
          headers: action.headers,
        });
      } catch (error) {
        console.error("Failed to sync action:", error);
      }
    }

    localStorage.removeItem("offlineActions");
    this.showSyncMessage();
  }

  showOfflineMessage() {
    this.showToast("Action saved. Will sync when online.", "info");
  }

  showSyncMessage() {
    this.showToast("Data synced successfully!", "success");
  }

  // Web Share API
  setupShareAPI() {
    const shareButtons = document.querySelectorAll("[data-share]");
    shareButtons.forEach((btn) => {
      btn.addEventListener("click", this.shareContent.bind(this));
    });
  }

  async shareContent(event) {
    const shareData = JSON.parse(event.target.dataset.share);

    if (navigator.share) {
      try {
        await navigator.share(shareData);
      } catch (error) {
        console.error("Error sharing:", error);
      }
    } else {
      // Fallback to clipboard
      await navigator.clipboard.writeText(shareData.url);
      this.showToast("Link copied to clipboard!", "success");
    }
  }

  // Toast notifications
  showToast(message, type = "info") {
    const toast = document.createElement("div");
    toast.className = `fixed top-4 right-4 z-50 px-4 py-2 rounded-lg text-white ${
      type === "success"
        ? "bg-green-600"
        : type === "error"
        ? "bg-red-600"
        : "bg-blue-600"
    }`;
    toast.textContent = message;

    document.body.appendChild(toast);

    setTimeout(() => {
      toast.remove();
    }, 3000);
  }
}

// Initialize PWA features when DOM is loaded
document.addEventListener("DOMContentLoaded", () => {
  new PWAFeatures();
});