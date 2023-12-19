/** @type {import('tailwindcss').Config} */
module.exports = {
  important: true,
  content: ['./index.templ', './src/templates/**/*.templ'],
  theme: {
    extend: {
      fontFamily: {
        'cabin': ['"Cabin"', 'sans-serif']
      }
    }
  },
  plugins: [],
}