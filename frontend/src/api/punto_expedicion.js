// frontend/src/api/punto_expedicion.js
import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json'
  }
})

export default {
  async getAll() {
    try {
      const res = await api.get('/puntos')
      return { success: true, data: res.data }
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Error al obtener puntos de expedición'
      }
    }
  }
}
