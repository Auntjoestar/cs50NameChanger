import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  define: {
    __VUE_OPTIONS_API__: true, // Enable Options API if you're using it
    __VUE_PROD_DEVTOOLS__: false, // Disable devtools in production
    __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: true, // Enable detailed hydration mismatch debugging
  },
});
