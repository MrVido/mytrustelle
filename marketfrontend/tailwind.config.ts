import { Config } from 'tailwindcss';
import defaultTheme from 'tailwindcss/defaultTheme';

const config: Config = {
  darkMode: 'class', // This enables dark mode based on a class, which you can toggle
  content: [
    './pages/**/*.{ts,tsx}',
    './components/**/*.{ts,tsx}',
    './app/**/*.{ts,tsx}',
    './src/**/*.{ts,tsx}',
  ],
  theme: {
    container: {
      center: true,
      padding: '2rem',
      screens: {
        '2xl': '1400px',
      },
    },
    extend: {
      fontFamily: {
        sans: ['Segoe UI', 'Tahoma', 'Geneva', 'Verdana', 'sans-serif', ...defaultTheme.fontFamily.sans],
      },
      colors: {
        // Light mode colors
        h1Color: '#007bff',
        navColor: '#0056b3',
        headingColor: '#004085',
        highlightColor: '#cce5ff',
        successColor: '#28a745',
        successDark: '#155724',
        successLight: '#d4edda',
        warningColor: '#ffc107',
        warningLight: '#fff3cd',
        errorColor: '#dc3545',
        errorLight: '#f8d7da',
        backgroundColor: '#ffffff',
        borderColor: '#f8f9fa',
        textColor: '#6c757d',

        // Dark mode colors
        'dark:h1Color': '#91d5ff', // Lighter blue for dark mode
        'dark:navColor': '#409eff', // More vibrant blue for visibility
        'dark:headingColor': '#aab8c2', // Softer dark blue
        'dark:highlightColor': '#4d7298', // Darker highlight for contrast
        'dark:successColor': '#73d13d', // Brighter green
        'dark:successDark': '#237804', // Dark green
        'dark:successLight': '#95de64', // Lighter success color for dark backgrounds
        'dark:warningColor': '#ffd666', // Softened yellow
        'dark:warningLight': '#fff566', // Brighter yellow for dark mode readability
        'dark:errorColor': '#ff7875', // Softer red
        'dark:errorLight': '#ffa39e', // Light red for background hints
        'dark:backgroundColor': '#434342', // Deep dark for main background
        'dark:borderColor': '#499999', // Dark border for separation
        'dark:textColor': '#e0e0e0', // Light text for dark background readability
      },
      keyframes: {
        'accordion-down': {
          from: { height: '0' },
          to: { height: 'var(--radix-accordion-content-height)' },
        },
        'accordion-up': {
          from: { height: 'var(--radix-accordion-content-height)' },
          to: { height: '0' },
        },
      },
      animation: {
        'accordion-down': 'accordion-down 0.2s ease-out',
        'accordion-up': 'accordion-up 0.2s ease-out',
      },
    },
  },
  plugins: [require('tailwindcss-animate')],
};

export default config;
