// Theme configuration matching the website's dark aesthetic
export const colors = {
  // Dark backgrounds (matching website)
  background: "#0f0f23", // Main dark background (ink-900)
  surface: "#1a1a2e", // Card/surface background (ink-800)
  elevated: "#16213e", // Elevated surfaces (ink-700)

  // Accent colors (matching website mint palette)
  primary: "#10b981", // Mint-600 (matching website buttons)
  secondary: "#06d6a0", // Mint-500 for secondary actions
  accent: "#34d399", // Mint-400 for hover states

  // Text colors (matching website)
  text: {
    primary: "#ffffff", // White text
    secondary: "#94a3b8", // Slate-400 for secondary text
    muted: "#64748b", // Slate-500 for muted text
    tertiary: "#cbd5e1", // Slate-300 for labels
  },

  // Semantic colors
  error: "#ef4444", // Red-500
  warning: "#f59e0b", // Amber-500
  success: "#10b981", // Mint-600 (matching primary)

  // Borders and separators
  border: "rgba(255, 255, 255, 0.1)", // White with 10% opacity
  borderLight: "rgba(255, 255, 255, 0.2)", // White with 20% opacity

  // Glass/blur effects (matching website)
  glass: "rgba(26, 26, 46, 0.9)", // surface with 90% opacity
  backdrop: "rgba(15, 15, 35, 0.8)", // background with 80% opacity
};

// Typography to match website (Inter font)
export const typography = {
  fontFamily: {
    regular: "Inter-Regular",
    medium: "Inter-Medium",
    semiBold: "Inter-SemiBold",
    bold: "Inter-Bold",
  },
  fontSize: {
    xs: 12,
    sm: 14,
    base: 16,
    lg: 18,
    xl: 20,
    "2xl": 24,
    "3xl": 30,
    "4xl": 36,
  },
  lineHeight: {
    xs: 16,
    sm: 20,
    base: 24,
    lg: 28,
    xl: 28,
    "2xl": 32,
    "3xl": 36,
    "4xl": 40,
  },
};

// Spacing system (0.25rem = 4px base)
export const spacing = {
  xs: 4,
  sm: 8,
  md: 16,
  lg: 24,
  xl: 32,
  "2xl": 48,
  "3xl": 64,
};

// Border radius matching website rounded-xl style
export const borderRadius = {
  sm: 4,
  md: 8,
  lg: 12,
  xl: 16,
  "2xl": 24,
  full: 9999,
};

// Shadow system for elevation
export const shadows = {
  sm: {
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 1 },
    shadowOpacity: 0.2,
    shadowRadius: 2,
    elevation: 2,
  },
  md: {
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.3,
    shadowRadius: 4,
    elevation: 4,
  },
  lg: {
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 4 },
    shadowOpacity: 0.4,
    shadowRadius: 8,
    elevation: 8,
  },
  glass: {
    shadowColor: "#10b981",
    shadowOffset: { width: 0, height: 0 },
    shadowOpacity: 0.3,
    shadowRadius: 20,
    elevation: 12,
  },
};