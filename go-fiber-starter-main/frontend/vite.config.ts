import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// Vite config: dev server proxies `/api` to your Go Fiber backend on :8080
export default defineConfig({
  plugins: [react()],
  server: {
    port: 3000,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:8080',
        changeOrigin: true,
        secure: false,
      }
    }
  }
})
