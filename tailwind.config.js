/** tailwind.config.js */
module.exports = {
  content: ["./templates/**/*.templ", "./internal/**/*.go", "./static/**/*.js"],
  theme: {
    extend: {
      colors: {
        ink: { 900: "#0a0b10", 800: "#0f1118", 700: "#141828" },
        mint: { 400: "#37f5c0", 500: "#26e0b0", 600: "#13c39a" },
        grape: { 400: "#a78bff", 500: "#8f6bff", 600: "#6f4dff" },
      },
      fontFamily: {
        sans: [
          "Inter",
          "ui-sans-serif",
          "system-ui",
          "Segoe UI",
          "Helvetica",
          "Arial",
        ],
        mono: [
          "JetBrains Mono",
          "ui-monospace",
          "SFMono-Regular",
          "Menlo",
          "monospace",
        ],
      },
      boxShadow: {
        glow: "0 0 0 1px rgba(55,245,192,.35), 0 0 24px rgba(55,245,192,.25)",
        grape:
          "0 0 0 1px rgba(167,139,255,.25), 0 0 18px rgba(167,139,255,.18)",
      },
      backgroundImage: {
        mesh: "radial-gradient(1200px 600px at 10% 0%, rgba(167,139,255,.18), transparent 40%), radial-gradient(900px 500px at 90% 10%, rgba(55,245,192,.12), transparent 50%), radial-gradient(700px 400px at 50% 100%, rgba(56,189,248,.10), transparent 50%)",
        grain:
          "url(\"data:image/svg+xml;utf8,<svg xmlns='http://www.w3.org/2000/svg' width='160' height='160' viewBox='0 0 160 160'><filter id='n'><feTurbulence type='fractalNoise' baseFrequency='.8' numOctaves='2'/><feColorMatrix type='saturate' values='0'/><feComponentTransfer><feFuncA type='table' tableValues='0 0 .02 .04 .02 0'/></feComponentTransfer></filter><rect width='100%' height='100%' filter='url(%23n)'/></svg>\")",
      },
      keyframes: {
        pulseGlow: {
          "0%,100%": {
            boxShadow: "0 0 0 0 rgba(55,245,192,0),0 0 18px rgba(55,245,192,0)",
          },
          "50%": {
            boxShadow:
              "0 0 0 1px rgba(55,245,192,.25),0 0 18px rgba(55,245,192,.25)",
          },
        },
      },
      animation: { pulseGlow: "pulseGlow 3s ease-in-out infinite" },
    },
  },
  darkMode: "class",
  plugins: [],
};
