import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json'
  }
})

// Obtener todos los timbrados vigentes
export const getTimbradosVigentes = () => api.get('/timbrados/vigentes')

// Obtener id_timbrado por establecimiento y expedición
export const getTimbradoPorEstYExp = (idEstablecimiento, idExpedicion) =>
  api.get(`/timbrado/${idEstablecimiento}/${idExpedicion}`)

// Crear nuevo timbrado
export const crearTimbrado = (timbrado) =>
  api.post('/timbrado', timbrado)

// Actualizar un timbrado existente
export const actualizarTimbrado = (id, timbrado) =>
  api.put(`/timbrado/${id}`, timbrado)

// Eliminar (soft-delete) un timbrado
export const eliminarTimbrado = (id) =>
  api.delete(`/timbrado/${id}`)
