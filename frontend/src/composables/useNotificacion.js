// src/composables/useNotificacion.js
import { ref } from 'vue'

export function useNotificacion() {
  const notificacion = ref(null)
  
  const mostrarNotificacion = (mensaje, tipo = 'success') => {
    // Implementar lógica de notificación (Toast, alerta, etc.)
    console[tipo === 'error' ? 'error' : 'log'](mensaje)
  }

  return { notificacion, mostrarNotificacion }
}