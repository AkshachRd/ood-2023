import defaultConfig from 'tailwindcss/defaultConfig'
import formsPlugin from '@tailwindcss/forms';

/** @type {import('tailwindcss').Config} */
export default {
  content: ['index.html', 'src/**/*.tsx'],
  theme: {
    extend: {fontFamily: {
			sans: ['Inter', ...defaultConfig.theme.fontFamily.sans]
		}},
  },
  experimental: {optimizeUniversalDefaults: true},
  plugins: [formsPlugin],
}
