// File: tailwind.config.js
/** @type {import('tailwindcss').Config} */
module.exports = {
        // 1. Add dark mode strategy ('selector' is recommended)
        darkMode: 'selector', // Enables dark mode based on a `.dark` class on a parent (usually <html>)

        // 2. Specify where Tailwind should scan for classes
        content: [
                "./web/template/layouts/base.html",
        ],

        // 3. Theme customizations (optional, keep empty for now or add later)
        theme: {
                extend: {
                        // You can add custom colors, fonts, etc. here later
                        // colors: {
                        //   'brand-blue': '#1DA1F2',
                        // },
                },
        },

        // 4. Plugins (optional, keep empty for now)
        plugins: [],
}

