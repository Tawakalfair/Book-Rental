// tailwind.config.js
/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.html"], // Scans all .html files in the views folder
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'), 
  ],
}