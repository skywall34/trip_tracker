import * as THREE from "three"; // → https://cdn.jsdelivr.net/...
import { OrbitControls } from "three/addons/controls/OrbitControls.js";

// Check if the current page is /worldmap3d before rendering
if (window.location.pathname === "/worldmap3d") {
  console.log("In worldmap3d");

  // ...existing code to render the 3D world map...
} else {
  console.log("NOt in worldmap3d");
}
