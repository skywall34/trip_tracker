module.exports = ({ config }) => {
  // Use environment variable or default to localhost
  const apiUrl = process.env.EXPO_PUBLIC_DEV_API_URL || 'http://localhost:3000';
  
  console.log('Mobile API URL:', apiUrl);
  
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
      bundleIdentifier: "com.mias.trips"
    },
    android: {
      adaptiveIcon: {
        foregroundImage: "./assets/adaptive-icon.png",
        backgroundColor: "#ffffff"
      },
      edgeToEdgeEnabled: false,
      package: "com.mias.trips",
      permissions: ["ACCESS_FINE_LOCATION", "ACCESS_COARSE_LOCATION"],
      googleServicesFile: process.env.GOOGLE_SERVICES_JSON || "./google-services.json"
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
      ],
      [
        "@react-native-google-signin/google-signin",
        {
          iosClientId: process.env.EXPO_PUBLIC_GOOGLE_OAUTH_CLIENT_ID_IOS || "",
          androidClientId: process.env.EXPO_PUBLIC_GOOGLE_OAUTH_CLIENT_ID_ANDROID || "731083279268-40ueldcggn03lj76jknhfsfkdso82nds.apps.googleusercontent.com",
          iosUrlScheme: "com.googleusercontent.apps.731083279268-udc3jitbjbn40k281oa1ppourlk9ikuj"
        }
      ]
    ],
    extra: {
      apiUrl: apiUrl,
      googleOAuthClientId: {
        ios: process.env.EXPO_PUBLIC_GOOGLE_OAUTH_CLIENT_ID_IOS || "",
        android: process.env.EXPO_PUBLIC_GOOGLE_OAUTH_CLIENT_ID_ANDROID || "731083279268-40ueldcggn03lj76jknhfsfkdso82nds.apps.googleusercontent.com",
        web: process.env.EXPO_PUBLIC_GOOGLE_OAUTH_CLIENT_ID_WEB || "731083279268-udc3jitbjbn40k281oa1ppourlk9ikuj.apps.googleusercontent.com"
      },
      productionApiUrl: process.env.EXPO_PUBLIC_PRODUCTION_API_URL || "https://fromnto.cloud"
    },
    scheme: "mias-trips"
  };
};