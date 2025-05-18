// frontend/src/api/facturar.js
import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json'
  }
})

export default {
  crearFactura: async (data) => {
    try {
      const res = await api.post('/facturar', data)
      return { success: true, data: res.data }
    } catch (err) {
      return { success: false, error: err.response?.data?.error || err.message }
    }
  },

  buscarOrdenPorNumero: async (numeroOrden) => {
    try {
      const res = await api.get(`/order/search/${numeroOrden}`)
      return { success: true, data: res.data }
    } catch (err) {
      return { success: false, error: err.response?.data?.error || err.message }
    }
  },

  obtenerProximoNumeroFactura: async (codigoEstablecimiento, codigoExpedicion) => {
    try {
      const res = await api.get(`/facturar/next-numero/${codigoEstablecimiento}/${codigoExpedicion}`)
      return { success: true, data: res.data }
    } catch (err) {
      return { success: false, error: err.response?.data?.error || err.message }
    }
  },

  getEstablecimientos: async () => {
    try {
      const res = await api.get('/establecimientos')
      return { success: true, data: res.data }
    } catch (err) {
      return { success: false, error: err.response?.data?.error || err.message }
    }
  },

  getPuntosExpedicion: async () => {
    try {
      const res = await api.get('/puntos_expedicion')
      return { success: true, data: res.data }
    } catch (err) {
      return { success: false, error: err.response?.data?.error || err.message }
    }
  },

  // ✅ Obtener todos los timbrados vigentes
  getTimbradosVigentes: async () => {
    try {
      const res = await api.get('/timbrados/vigentes')
      return { success: true, data: res.data }
    } catch (err) {
      return { success: false, error: err.response?.data?.error || err.message }
    }
  },

  // ✅ Obtener el id_timbrado según establecimiento y punto de expedición
  getTimbradoPorEstablecimientoYExpedicion: async (idEst, idExp) => {
    try {
      const res = await api.get(`/timbrado/${idEst}/${idExp}`)
      return { success: true, data: res.data.id_timbrado }
    } catch (err) {
      return { success: false, error: err.response?.data?.error || err.message }
    }
  }
}


// // frontend/src/api/facturar.js
// import axios from 'axios'

// const api = axios.create({
//   baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
//   headers: { 'Content-Type': 'application/json', Accept: 'application/json' }
// })

// export default {
//   crearFactura: async (data) => {
//     try {
//       const res = await api.post('/facturar', data)
//       return { success: true, data: res.data }
//     } catch (err) {
//       console.error('crearFactura error:', err.response || err.message)
//       return { success: false, error: err.response?.data?.error || err.message }
//     }
//   },
//   buscarOrdenPorNumero: async (numeroOrden) => {
//     try {
//       const res = await api.get(`/order/search/${numeroOrden}`)
//       return { success: true, data: res.data }
//     } catch (err) {
//       console.error('buscarOrdenPorNumero error:', err.response || err.message)
//       return { success: false, error: err.response?.data?.error || err.message }
//     }
//   },
//   obtenerProximoNumeroFactura: async (est, exp) => {
//     try {
//       const res = await api.get(`/facturar/next-numero/${est}/${exp}`)
//       return { success: true, data: res.data.next_numero_factura }
//     } catch (err) {
//       console.error('obtenerProximoNumeroFactura error:', err.response || err.message)
//       return { success: false, error: err.response?.data?.error || err.message }
//     }
//   }
// }
