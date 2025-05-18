import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json'
  },
  timeout: 5000 // 5 segundos timeout
})

// Interceptor para manejar errores globalmente
api.interceptors.response.use(
  response => response,
  error => {
    if (error.response) {
      console.error('Error en la respuesta:', {
        status: error.response.status,
        data: error.response.data
      })
    } else if (error.request) {
      console.error('No se recibió respuesta:', error.request)
    } else {
      console.error('Error al configurar la petición:', error.message)
    }
    return Promise.reject(error)
  }
)

export default {
  async getContactos() {
    try {
      const response = await api.get('/contactos')
      return {
        success: true,
        data: response.data,
        status: response.status
      }
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Error al obtener contactos',
        status: error.response?.status
      }
    }
  },

  async buscarContactos(query) {
    try {
      const response = await api.get('/contactos/search', {
        params: { q: query }
      })
      return {
        success: true,
        data: response.data,
        status: response.status
      }
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Error al buscar contactos',
        status: error.response?.status
      }
    }
  },

  async createContacto(contacto) {
    try {
      const response = await api.post('/contactos', contacto)
      return {
        success: true,
        data: response.data,
        status: response.status
      }
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Error al crear contacto',
        status: error.response?.status
      }
    }
  },

  async deleteContacto(id) {
    try {
      const response = await api.delete(`/contactos/${id}`)
      return {
        success: true,
        status: response.status
      }
    } catch (error) {
      return {
        success: false,
        error: error.response?.data?.error || 'Error al eliminar contacto',
        status: error.response?.status
      }
    }
  }
}
