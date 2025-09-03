/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./src/pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          500: '#154D71',
          600: '#0e3a57',
        },
        secondary: {
          500: '#468A9A',
          600: '#3a6f7c',
        }
      }
    },
  },
  plugins: [],
}