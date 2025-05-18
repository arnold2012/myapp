// frontend/src/api/impuestos.js
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
      const res = await api.get('/impuestos')
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al obtener impuestos' }
    }
  },

  async getById(id) {
    try {
      const res = await api.get(`/impuestos/${id}`)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Impuesto no encontrado' }
    }
  },

  async create(impuesto) {
    try {
      const res = await api.post('/impuestos', impuesto)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al crear impuesto' }
    }
  },

  async update(id, impuesto) {
    try {
      const res = await api.put(`/impuestos/${id}`, impuesto)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al actualizar impuesto' }
    }
  },

  async delete(id) {
    try {
      await api.delete(`/impuestos/${id}`)
      return { success: true }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al eliminar impuesto' }
    }
  }
}
