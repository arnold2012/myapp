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
      const res = await api.get('/marcas')
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al obtener marcas' }
    }
  },

  async getById(id) {
    try {
      const res = await api.get(`/marcas/${id}`)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Marca no encontrada' }
    }
  },

  async create(marca) {
    try {
      const res = await api.post('/marcas', marca)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al crear marca' }
    }
  },

  async update(id, marca) {
    try {
      const res = await api.put(`/marcas/${id}`, marca)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al actualizar marca' }
    }
  },

  async delete(id) {
    try {
      await api.delete(`/marcas/${id}`)
      return { success: true }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al eliminar marca' }
    }
  }
}
