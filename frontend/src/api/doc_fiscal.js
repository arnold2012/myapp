import axios from 'axios'
const api = axios.create({ baseURL: import.meta.env.VITE_API_BASE_URL + '/api' })

export default {
  getEstablecimientos: () => api.get('/establecimientos'),
  getPuntosExpedicion: () => api.get('/puntos_expedicion'),
}
