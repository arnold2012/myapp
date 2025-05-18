import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json'
  }
});

export default {
  async buscarProductos(query) {
    try {
      const response = await api.get(`/productos/search`, {
        params: { q: query }
      });
      return { success: true, data: response.data };
    } catch (error) {
      return { success: false, error: error.message };
    }
  },

  async getAll() {
    try {
      const response = await api.get('/productos');
      return { success: true, data: response.data };
    } catch (error) {
      return { success: false, error: error.message };
    }
  },

  async getById(id) {
    try {
      const response = await api.get(`/productos/${id}`);
      return { success: true, data: response.data };
    } catch (error) {
      return { success: false, error: error.message };
    }
  },

  async createProducto(producto) {
    try {
      const response = await api.post('/productos', producto);
      return { success: true, data: response.data };
    } catch (error) {
      return { success: false, error: error.response?.data?.error || error.message };
    }
  },

  async updateProducto(id, producto) {
    try {
      const response = await api.put(`/productos/${id}`, producto);
      return { success: true, data: response.data };
    } catch (error) {
      return { success: false, error: error.response?.data?.error || error.message };
    }
  },

  async deleteProducto(id) {
    try {
      const response = await api.delete(`/productos/${id}`);
      return { success: true, data: response.data };
    } catch (error) {
      return { success: false, error: error.response?.data?.error || error.message };
    }
  },

  async getNextCodItem() {
    try {
      const response = await api.get('/productos/next-code');
      return { success: true, data: response.data.next_cod_item };
    } catch (error) {
      return { success: false, error: error.message };
    }
  }
};