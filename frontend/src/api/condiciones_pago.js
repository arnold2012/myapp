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
      const res = await api.get('/condiciones_pago')
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al obtener condiciones de pago' }
    }
  },

  async getById(id) {
    try {
      const res = await api.get(`/condiciones_pago/${id}`)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al obtener la condición de pago' }
    }
  },

  async create(condicion) {
    try {
      const res = await api.post('/condiciones_pago', condicion)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al crear condición de pago' }
    }
  },

  async update(id, condicion) {
    try {
      const res = await api.put(`/condiciones_pago/${id}`, condicion)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al actualizar condición de pago' }
    }
  },

  async delete(id) {
    try {
      const res = await api.delete(`/condiciones_pago/${id}`)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al eliminar condición de pago' }
    }
  }
}
