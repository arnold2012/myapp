<template>
    <div v-if="show" class="fixed inset-0 flex items-center justify-center z-50">
      <!-- Modal -->
      <div class="bg-white p-6 rounded-lg shadow-lg max-w-md w-full z-50"> 
        <h2 class="text-xl font-semibold mb-4">¿Confirmar generación de factura?</h2>
  
        <p class="mb-2">Factura: <span class="font-mono text-blue-600">{{ numeroFactura }}</span></p>
        <p class="mb-2">Cliente: <strong>{{ cliente }}</strong></p>
        <p>Total: <strong>Gs {{ total.toLocaleString('es-PY') }}</strong></p>
  
        <div class="flex justify-end gap-2 mt-6">
          <button @click="onCancel" class="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400">
            Cancelar
          </button>
          <button @click="onConfirm" class="px-4 py-2 bg-green-600 text-white rounded hover:bg-green-700">
            Confirmar
          </button>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  const props = defineProps({
    show: Boolean,
    numeroFactura: String,
    cliente: String,
    total: Number
  })
  
  const emit = defineEmits(['confirm', 'cancel'])
  
  const onConfirm = () => emit('confirm')
  const onCancel = () => emit('cancel')
  </script>
  
  <style scoped>
  /* No background oscuro aplicado */
  </style>
  