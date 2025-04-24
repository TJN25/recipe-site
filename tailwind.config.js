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
          DEFAULT: '#C02F00',
          50: '#FF9A79',
          100: '#FF8A64',
          200: '#FF6B3B',
          300: '#FF4C13',
          400: '#E93900',
          500: '#C02F00',
          600: '#882100',
          700: '#501400',
          800: '#180600',
          900: '#000000',
          950: '#000000'
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
