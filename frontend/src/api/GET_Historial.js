import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json'
  }
})

export default {
  // Obtener el historial de una orden específica
  async getHistorialPorOrden(idOrder) {
    console.log(`Llamando a API: /historial/${idOrder}`);
    try {
      const response = await api.get(`/historial/${idOrder}`);
      console.log('Respuesta exitosa:', response.data);
      return { success: true, data: response.data };
    } catch (error) {
      console.error('Error en la API:', error);
      
      const errorInfo = {
        message: error.message,
        status: error.response?.status,
        statusText: error.response?.statusText,
        data: error.response?.data,
        url: `/historial/${idOrder}`
      };
      
      console.error('Detalles del error:', errorInfo);
      
      return {
        success: false,
        error: error.response?.data?.error || error.message,
        details: errorInfo
      };
    }
  },
  
  // Obtener la lista de todas las órdenes
  async getOrdenes() {
    try {
      const response = await api.get('/ordenes');
      return { success: true, data: response.data };
    } catch (error) {
      console.error('Error obteniendo órdenes:', error);
      
      return {
        success: false,
        error: error.response?.data?.error || error.message
      };
    }
  }
}