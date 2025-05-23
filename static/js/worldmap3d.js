// https://github.com/bobbyroe/vertex-earth

import * as THREE from "three";
import { OrbitControls } from "three/addons/controls/OrbitControls.js";

//
// ─── HELPERS ───────────────────────────────────────────────────────────────────
//

function randomSpherePoint(minR = 25, maxR = 50) {
  const r = Math.random() * (maxR - minR) + minR;
  const u = Math.random();
  const v = Math.random();
  const theta = 2 * Math.PI * u; // full rotation around
  const phi = Math.acos(2 * v - 1); // distribution from pole to pole
  return new THREE.Vector3(
    r * Math.sin(phi) * Math.cos(theta),
    r * Math.sin(phi) * Math.sin(theta),
    r * Math.cos(phi)
  );
}

function getStarfield({ numStars = 500, sprite } = {}) {
  const positions = [],
    colors = [],
    color = new THREE.Color();
  for (let i = 0; i < numStars; i++) {
    const p = randomSpherePoint();
    positions.push(p.x, p.y, p.z);
    color.setHSL(0.6, 0.2, Math.random());
    colors.push(color.r, color.g, color.b);
  }
  const geo = new THREE.BufferGeometry();
  geo.setAttribute("position", new THREE.Float32BufferAttribute(positions, 3));
  geo.setAttribute("color", new THREE.Float32BufferAttribute(colors, 3));
  const mat = new THREE.PointsMaterial({
    size: 0.2,
    vertexColors: true,
    map: sprite,
    transparent: true,
  });
  return new THREE.Points(geo, mat);
}

function latLongToVector3(latitude, longitude, radius = 1, height = 0.01) {
  // Convert degrees → radians
  const latRad = THREE.MathUtils.degToRad(90 - latitude);
  const lonRad = THREE.MathUtils.degToRad(longitude);

  // Total radius with optional height offset
  const r = radius + height;

  // Note the minus on X!
  const x = -r * Math.sin(latRad) * Math.cos(lonRad);
  const y = r * Math.cos(latRad);
  const z = r * Math.sin(latRad) * Math.sin(lonRad);

  return new THREE.Vector3(x, y, z);
}

function createFlightPath(from, to, globeGroup) {
  const start = latLongToVector3(from.lat, from.lon, 1, 0);
  const end = latLongToVector3(to.lat, to.lon, 1, 0);
  const distance = start.distanceTo(end);

  const mid = start
    .clone()
    .add(end)
    .multiplyScalar(0.5)
    .normalize()
    .multiplyScalar(1 + distance * 0.25);

  const curve = new THREE.QuadraticBezierCurve3(start, mid, end);
  const points = curve.getPoints(100);
  const geometry = new THREE.BufferGeometry().setFromPoints(points);
  const material = new THREE.LineBasicMaterial({ color: 0xffff00 });
  const line = new THREE.Line(geometry, material);

  const marker = new THREE.Mesh(
    new THREE.SphereGeometry(0.01, 8, 6),
    new THREE.MeshBasicMaterial({ color: 0xff0000 })
  );
  globeGroup.add(marker);

  return { line, curve, marker, progress: Math.random() };
}

async function loadTrips() {
  const response = await fetch("/api/trips");
  const data = await response.json();

  const trips = data.map((trip) => ({
    from: {
      lat: trip.departure_lat,
      lon: trip.departure_lon,
      name: trip.departure,
    },
    to: {
      lat: trip.arrival_lat,
      lon: trip.arrival_lon,
      name: trip.arrival,
    },
  }));

  return trips;
}

//
// ─── SHADERS ───────────────────────────────────────────────────────────────────
//

const vertexShader = `
  uniform float size;
  uniform sampler2D elevTexture;
  varying vec2 vUv;
  varying float vVisible;

  void main() {
    vUv = uv;
    vec4 mvPosition = modelViewMatrix * vec4(position, 1.0);
    float elv = texture2D(elevTexture, vUv).r;
    vec3 vNormal = normalMatrix * normal;
    vVisible = step(0.0, dot(-normalize(mvPosition.xyz), normalize(vNormal)));
    mvPosition.z += 0.35 * elv;
    gl_PointSize = size;
    gl_Position = projectionMatrix * mvPosition;
  }
`;

const fragmentShader = `
  uniform sampler2D colorTexture;
  uniform sampler2D alphaTexture;
  varying vec2 vUv;
  varying float vVisible;

  void main() {
    if (floor(vVisible + 0.1) == 0.0) discard;
    float alpha = 1.0 - texture2D(alphaTexture, vUv).r;
    vec3 color = texture2D(colorTexture, vUv).rgb;
    gl_FragColor = vec4(color, alpha);
  }
`;

//
// ─── ENTRYPOINT ────────────────────────────────────────────────────────────────
//

if (window.location.pathname !== "/worldmap3d") {
  console.log("Not on /worldmap3d → skipping ThreeJS init.");
} else {
  // scene, camera, renderer
  const scene = new THREE.Scene();
  const camera = new THREE.PerspectiveCamera(
    45,
    innerWidth / innerHeight,
    0.1,
    1000
  );
  const renderer = new THREE.WebGLRenderer({ antialias: true });
  camera.position.set(0, 0, 3.5);
  renderer.setSize(innerWidth, innerHeight);
  renderer.setPixelRatio(window.devicePixelRatio);
  document.body.appendChild(renderer.domElement);

  // controls
  const controls = new OrbitControls(camera, renderer.domElement);
  controls.enableDamping = true;

  // raycaster + tooltip
  const raycaster = new THREE.Raycaster();
  const mouse = new THREE.Vector2();
  const tooltip = document.createElement("div");
  tooltip.style.cssText = `
    position: absolute;
    background: #000000cc;
    color: #fff;
    padding: 6px 10px;
    border-radius: 6px;
    pointer-events: none;
    font-size: 12px;
    display: none;
  `;
  document.body.appendChild(tooltip);

  // textures
  const loader = new THREE.TextureLoader();
  const starSprite = loader.load("/static/images/circle.png");
  const colorMap = loader.load("/static/images/03_earthlights1k.jpg");
  const elevMap = loader.load("/static/images/01_earthbump1k.jpg");
  const alphaMap = loader.load("/static/images/02_earthspec1k.jpg");
  // globe group
  const globeGroup = new THREE.Group();
  scene.add(globeGroup);

  // wireframe base
  globeGroup.add(
    new THREE.Mesh(
      new THREE.IcosahedronGeometry(1, 10),
      new THREE.MeshBasicMaterial({ color: 0x202020, wireframe: true })
    )
  );

  // shader‐driven earth points
  const pointsGeo = new THREE.IcosahedronGeometry(1, 120);
  const pointsMat = new THREE.ShaderMaterial({
    uniforms: {
      size: { value: 4.0 },
      colorTexture: { value: colorMap },
      elevTexture: { value: elevMap },
      alphaTexture: { value: alphaMap },
    },
    vertexShader,
    fragmentShader,
    transparent: true,
  });
  globeGroup.add(new THREE.Points(pointsGeo, pointsMat));

  // starfield + light
  scene.add(getStarfield({ numStars: 4500, sprite: starSprite }));
  scene.add(new THREE.HemisphereLight(0xffffff, 0x080820, 3));

  const trips = await loadTrips();

  const airportMarkers = [];
  const flightPaths = [];
  const markerGeo = new THREE.SphereGeometry(0.015, 8, 6);

  trips.forEach(({ from, to }) => {
    [from, to].forEach((pt) => {
      const m = new THREE.Mesh(
        markerGeo,
        new THREE.MeshBasicMaterial({ color: 0x00ffcc })
      );
      m.position.copy(latLongToVector3(pt.lat, pt.lon));
      m.userData = pt;
      globeGroup.add(m);
      airportMarkers.push(m);
    });
    const fp = createFlightPath(from, to, globeGroup);
    globeGroup.add(fp.line);
    flightPaths.push(fp);
  });

  // hover tooltip
  renderer.domElement.addEventListener("pointermove", (e) => {
    const rect = renderer.domElement.getBoundingClientRect();
    mouse.x = ((e.clientX - rect.left) / rect.width) * 2 - 1;
    mouse.y = -((e.clientY - rect.top) / rect.height) * 2 + 1;

    raycaster.setFromCamera(mouse, camera);
    const hits = raycaster.intersectObjects(airportMarkers);
    if (hits.length) {
      const pt = hits[0].object.userData;
      tooltip.style.left = `${e.clientX + 10}px`;
      tooltip.style.top = `${e.clientY + 10}px`;
      tooltip.innerHTML = `<strong>${pt.name}</strong><br>Lat: ${pt.lat.toFixed(
        2
      )}<br>Lon: ${pt.lon.toFixed(2)}`;
      tooltip.style.display = "block";
    } else {
      tooltip.style.display = "none";
    }
  });

  // animate
  (function animate() {
    requestAnimationFrame(animate);
    controls.update();
    // globeGroup.rotation.y += 0.002;
    flightPaths.forEach((fp) => {
      fp.progress = (fp.progress + 0.002) % 1;
      fp.marker.position.copy(fp.curve.getPoint(fp.progress));
    });
    renderer.render(scene, camera);
  })();

  // resize
  window.addEventListener("resize", () => {
    camera.aspect = innerWidth / innerHeight;
    camera.updateProjectionMatrix();
    renderer.setSize(innerWidth, innerHeight);
  });
}
