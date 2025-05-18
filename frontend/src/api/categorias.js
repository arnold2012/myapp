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
      const res = await api.get('/categorias')
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al obtener categorías' }
    }
  },

  async getById(id) {
    try {
      const res = await api.get(`/categorias/${id}`)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Categoría no encontrada' }
    }
  },

  async create(categoria) {
    try {
      const res = await api.post('/categorias', categoria)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al crear categoría' }
    }
  },

  async update(id, categoria) {
    try {
      const res = await api.put(`/categorias/${id}`, categoria)
      return { success: true, data: res.data }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al actualizar categoría' }
    }
  },

  async delete(id) {
    try {
      await api.delete(`/categorias/${id}`)
      return { success: true }
    } catch (error) {
      return { success: false, error: error.response?.data?.error || 'Error al eliminar categoría' }
    }
  }
}
