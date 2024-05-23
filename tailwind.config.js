import daisyui from "daisyui"

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/*.html", "./views/**/*.html"],
  theme: {
    extend: {},
  },
  plugins: [
    daisyui,
  ],
  darkMode: 'class',
}

