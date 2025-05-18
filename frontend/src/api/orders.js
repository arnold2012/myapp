import axios from "axios"

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || "http://localhost:8080/api",
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
  timeout: 10000, // Aumentado a 10 segundos
})

export default {
  async createOrder(orderData) {
    try {
      console.log("Enviando solicitud para crear orden:", orderData)
      const response = await api.post("/orders", orderData)
      console.log("Respuesta del servidor (createOrder):", response.data)

      // Si el backend no devuelve id_order, podemos simular uno para pruebas
      // En producción, asegúrate de que el backend devuelva el ID correcto
      const responseData = response.data
      if (!responseData.id_order && process.env.NODE_ENV === "development") {
        console.warn("Backend no devolvió id_order, usando ID simulado para desarrollo")
        responseData.id_order = 1 // Solo para pruebas en desarrollo
      }

      return {
        success: true,
        data: responseData,
        status: response.status,
      }
    } catch (error) {
      console.error("Error al crear orden:", error)
      return {
        success: false,
        error: error.response?.data?.error || "Error al crear orden",
        status: error.response?.status,
      }
    }
  },

  async getOrderForPrint(orderId) {
    try {
      console.log("Solicitando datos para imprimir orden ID:", orderId)
      const response = await api.get(`/orders/${orderId}/imprimir`)
      console.log("Respuesta del servidor (getOrderForPrint):", response.data)
      return {
        success: true,
        data: response.data,
      }
    } catch (error) {
      console.error("Error al obtener orden para imprimir:", error)
      return {
        success: false,
        error: error.response?.data?.error || "Error al obtener orden",
      }
    }
  },
}



// import axios from 'axios'

// const api = axios.create({
//   baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
//   headers: {
//     'Content-Type': 'application/json',
//     Accept: 'application/json'
//   },
//   timeout: 5000
// })

// export default {
//   async createOrder(orderData) {
//     try {
//       const response = await api.post('/orders', orderData)
//       return {
//         success: true,
//         data: response.data,
//         status: response.status
//       }
//     } catch (error) {
//       return {
//         success: false,
//         error: error.response?.data?.error || 'Error al crear orden',
//         status: error.response?.status
//       }
//     }
//   }
// }



// import axios from 'axios';

// const api = axios.create({
//   baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api',
//   headers: {
//     'Content-Type': 'application/json',
//     'Accept': 'application/json'
//   },
//   timeout: 5000 // 5 segundos timeout
// });

// // Interceptor para manejar errores globalmente
// api.interceptors.response.use(
//   response => response,
//   error => {
//     if (error.response) {
//       console.error('Error en la respuesta:', {
//         status: error.response.status,
//         data: error.response.data
//       });
//     } else if (error.request) {
//       console.error('No se recibió respuesta:', error.request);
//     } else {
//       console.error('Error al configurar la petición:', error.message);
//     }
//     return Promise.reject(error);
//   }
// );

// export default {
//   async getOrders() {
//     try {
//       const response = await api.get('/api/orders');
//       return {
//         success: true,
//         data: response.data,
//         status: response.status
//       };
//     } catch (error) {
//       return {
//         success: false,
//         error: error.response?.data?.error || 'Error al obtener órdenes',
//         status: error.response?.status
//       };
//     }
//   },

//   async createOrder(order) {
//     try {
//       const response = await api.post('/api/orders', order);
//       return {
//         success: true,
//         data: response.data,
//         status: response.status
//       };
//     } catch (error) {
//       return {
//         success: false,
//         error: error.response?.data?.error || 'Error al crear orden',
//         status: error.response?.status
//       };
//     }
//   },

//   async deleteOrder(id) {
//     try {
//       const response = await api.delete(`/api/orders/${id}`);
//       return {
//         success: true,
//         status: response.status
//       };
//     } catch (error) {
//       return {
//         success: false,
//         error: error.response?.data?.error || 'Error al eliminar orden',
//         status: error.response?.status
//       };
//     }
//   }
// };
