// File: tailwind.config.js
/** @type {import('tailwindcss').Config} */
const defaultTheme = require('tailwindcss/defaultTheme')
module.exports = {
  content: [
    "./web/template/**/*.html", // Scan templates
  ],
  darkMode: 'selector', // Enable dark mode
  theme: {
    extend: {
      colors: {

        accent: {
          DEFAULT: '#297838',
          50: '#C1EAC8',
          100: '#B1E4BB',
          200: '#93DAA0',
          300: '#75D086',
          400: '#56C56B',
          500: '#3EB554',
          600: '#339646',
          700: '#297838',
          800: '#1B4E24',
          900: '#0C2411',
          950: '#050F07'
        },

      },
      fontFamily: {
        sans: ['Nunito', ...defaultTheme.fontFamily.sans], // Use Nunito as the primary sans-serif
        serif: ['Merriweather', ...defaultTheme.fontFamily.serif], // Replace 'Merriweather' with your chosen serif
      },
    },
  },
  plugins: [],
}
