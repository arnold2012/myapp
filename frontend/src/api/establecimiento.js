// frontend/src/api/establecimiento.js
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
      const res = await api.get('/establecimientos')
      return { success: true, data: res.data }
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Error al obtener establecimientos'
      }
    }
  },

  async getById(id) {
    try {
      const res = await api.get(`/establecimientos/${id}`)
      return { success: true, data: res.data }
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Establecimiento no encontrado'
      }
    }
  },

  async create(data) {
    try {
      const res = await api.post('/establecimientos', data)
      return { success: true, data: res.data }
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Error al crear establecimiento'
      }
    }
  },

  async update(id, data) {
    try {
      const res = await api.put(`/establecimientos/${id}`, data)
      return { success: true, data: res.data }
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Error al actualizar establecimiento'
      }
    }
  },

  async delete(id) {
    try {
      await api.delete(`/establecimientos/${id}`)
      return { success: true }
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Error al eliminar establecimiento'
      }
    }
  }
}
