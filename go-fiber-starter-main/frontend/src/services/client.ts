import axios from 'axios'

const api = axios.create({
  baseURL: '/api', // proxied to http://127.0.0.1:8080 by Vite dev server
  headers: {
    'Content-Type': 'application/json'
  }
})

// attach token from localStorage if present
const token = typeof window !== 'undefined' ? localStorage.getItem('token') : null
if(token){
  api.defaults.headers.common['Authorization'] = `Bearer ${token}`
}

export default api
