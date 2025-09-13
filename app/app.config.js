const axios = require('axios');

const getNgrokUrl = async () => {
  try {
    // ngrok exposes a local API at http://localhost:4040
    const response = await axios.get('http://localhost:4040/api/tunnels', {
      timeout: 2000
    });
    
    const tunnel = response.data.tunnels.find(t => 
      t.config.addr === 'http://localhost:3000' && t.proto === 'https'
    );
    
    if (tunnel) {
      console.log('🚀 Found ngrok tunnel:', tunnel.public_url);
      return tunnel.public_url;
    }
    
    console.log('⚠️  No ngrok tunnel found for port 3000');
    return null;
  } catch (error) {
    console.log('⚠️  ngrok not running, using localhost');
    return null;
  }
};

module.exports = async ({ config }) => {
  const apiUrl = await getNgrokUrl() || 'http://localhost:3000';
  
  console.log('📱 Mobile API URL:', apiUrl);
  
  return {
    name: "Mia's Trips",
    slug: "mias-trips",
    version: "1.0.0",
    orientation: "portrait",
    icon: "./assets/icon.png",
    userInterfaceStyle: "light",
    newArchEnabled: true,
    splash: {
      image: "./assets/splash-icon.png",
      resizeMode: "contain",
      backgroundColor: "#ffffff"
    },
    ios: {
      supportsTablet: true,
      bundleIdentifier: "com.miasname.trips"
    },
    android: {
      adaptiveIcon: {
        foregroundImage: "./assets/adaptive-icon.png",
        backgroundColor: "#ffffff"
      },
      edgeToEdgeEnabled: true,
      package: "com.miasname.trips",
      permissions: ["ACCESS_FINE_LOCATION", "ACCESS_COARSE_LOCATION"]
    },
    web: {
      favicon: "./assets/favicon.png"
    },
    plugins: [
      "expo-secure-store",
      [
        "expo-location",
        {
          locationAlwaysAndWhenInUsePermission: "Allow Mia's Trips to use your location to automatically detect airports and enhance trip planning."
        }
      ]
    ],
    extra: {
      apiUrl: apiUrl
    },
    scheme: "mias-trips"
  };
};