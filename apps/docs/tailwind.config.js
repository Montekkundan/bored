import { createPreset } from 'fumadocs-ui/tailwind-plugin';
import path from 'path';
/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './components/**/*.{ts,tsx}',
    './app/**/*.{ts,tsx}',
    './content/**/*.{md,mdx}',
    './mdx-components.{ts,tsx}',
    path.join(process.cwd(), '../../node_modules/fumadocs-ui/dist/**/*.js'),
  ],
  presets: [createPreset()],
};
