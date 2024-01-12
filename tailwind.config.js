/** @type {import('tailwindcss').Config} */
module.exports = {
  daisyui: {
    themes: [
      {
        main: {
          primary: "#66402C",
          secondary: "#A67B5B",
          accent: "#D4A15F",
          neutral: "#DFDFDF",
          "base-100": "#ffffff",
          background: "#E0D7C6",
          info: "#ffffff",
          success: "#00ffff",
          warning: "#ffffff",
          error: "#ffffff",
        },
      },
    ],
  },
  content: ["./web/template/*.{html,js}"],
  theme: {
    fontFamily: {
      sans: ["Inter", "ui-sans-serif"],
      title: ["'Baloo 2'", "ui-sans-serif"],
    },
    extend: {},
  },
  plugins: [require("daisyui")],
};
