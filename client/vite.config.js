import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/users': {
        target: 'http://localhost:3000', // Reemplaza esto con la URL de tu servidor de backend
        changeOrigin: true,
       
      },
    },
  },
})
