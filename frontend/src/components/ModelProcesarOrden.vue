<template>
  <div class="fixed inset-0 bg-opacity-50 flex items-center justify-center z-50">
    <div v-if="!ordenProcesadaId" class="bg-white rounded-lg shadow-lg w-full max-w-xl p-6 relative">
      <h2 class="text-xl font-semibold mb-4 text-gray-800">Procesar Orden</h2>

      <div class="mb-4">
        <p class="text-gray-700">Número de Orden: <strong>{{ numeroOrden }}</strong></p>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium text-gray-700 mb-1">Condición de Pago</label>
        <select
          v-model="condicionSeleccionada"
          class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary"
          :class="{ 'border-red-500': errores.condicion }"
        >
          <option disabled value="">Seleccione una condición</option>
          <option v-for="cond in condiciones" :key="cond.idcondicion_pago" :value="cond.idcondicion_pago">
            {{ cond.describe_condicion }}
          </option>
        </select>
        <p v-if="errores.condicion" class="mt-1 text-sm text-red-600">{{ errores.condicion }}</p>
      </div>

      <div class="mb-4">
        <label class="block text-sm font-medium text-gray-700 mb-1">Buscar Contacto</label>
        <div class="relative">
          <input
            v-model="busquedaContacto"
            @input="buscarContacto"
            type="text"
            placeholder="Ingrese nombre o RUC..."
            class="w-full border border-gray-300 rounded-md px-3 py-2 focus:outline-none focus:ring-2 focus:ring-primary"
            :class="{ 'border-red-500': errores.contacto }"
          />
          <div v-if="buscando" class="absolute right-3 top-2">
            <div class="animate-spin h-5 w-5 border-2 border-primary border-t-transparent rounded-full"></div>
          </div>
        </div>
        <p v-if="errores.contacto" class="mt-1 text-sm text-red-600">{{ errores.contacto }}</p>
        
        <ul v-if="contactos.length" class="border border-gray-200 rounded mt-1 bg-white shadow max-h-48 overflow-y-auto">
          <li
            v-for="c in contactos"
            :key="c.id_contacto"
            @click="seleccionarContacto(c)"
            class="px-3 py-2 cursor-pointer hover:bg-gray-100"
          >
            {{ c.razon_social }} ({{ c.ruc_ci }})
          </li>
        </ul>
        <p v-if="contactoSeleccionado" class="text-sm text-green-600 mt-1">
          Contacto seleccionado: {{ contactoSeleccionado.razon_social }}
        </p>
      </div>

      <div class="flex justify-end gap-2 mt-6">
        <button @click="cancelarOrden" class="px-4 py-2 bg-gray-300 text-gray-800 rounded hover:bg-gray-400">
          Cancelar
        </button>
        <button 
          @click="confirmarOrden" 
          class="px-4 py-2 bg-primary text-white rounded hover:bg-primary-dark transition-colors"
          :disabled="procesando"
        >
          <span v-if="procesando" class="flex items-center">
            <span class="animate-spin h-4 w-4 border-2 border-white border-t-transparent rounded-full mr-2"></span>
            Procesando...
          </span>
          <span v-else>Confirmar</span>
        </button>
      </div>
    </div>

    <!-- Componente para generar PDF al confirmar -->
    <OrdersPdfCanvas
      v-if="ordenProcesadaId"
      :order-id="ordenProcesadaId"
      @cerrar="finalizarProceso"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import condicionesApi from '@/api/condiciones_pago'
import contactosApi from '@/api/contactos'
import ordersApi from '@/api/orders'
import OrdersPdfCanvas from '@/components/html2pdf/html2pdf_orders_pdf.vue'

const props = defineProps({
  items: Array,
  total: Number,
  numeroOrden: Number
})

const emit = defineEmits(['close', 'ordenProcesada'])

const condiciones = ref([])
const condicionSeleccionada = ref('')
const busquedaContacto = ref('')
const contactos = ref([])
const contactoSeleccionado = ref(null)
const ordenProcesadaId = ref(null)
const procesando = ref(false)
const buscando = ref(false)
const errores = ref({
  condicion: '',
  contacto: ''
})

// Debounce function to limit API calls during typing
const debounce = (fn, delay) => {
  let timeout
  return function(...args) {
    clearTimeout(timeout)
    timeout = setTimeout(() => fn.apply(this, args), delay)
  }
}

const buscarContactoDebounced = debounce(async () => {
  if (busquedaContacto.value.length < 2) {
    contactos.value = []
    buscando.value = false
    return
  }
  
  try {
    const res = await contactosApi.buscarContactos(busquedaContacto.value)
    if (res.success) contactos.value = res.data
  } catch (error) {
    console.error('Error al buscar contactos:', error)
    contactos.value = []
  } finally {
    buscando.value = false
  }
}, 300)

const buscarContacto = () => {
  buscando.value = true
  errores.value.contacto = ''
  buscarContactoDebounced()
}

const seleccionarContacto = (c) => {
  contactoSeleccionado.value = c
  contactos.value = []
  busquedaContacto.value = c.razon_social
  errores.value.contacto = ''
}

const validarFormulario = () => {
  let valido = true
  errores.value = {
    condicion: '',
    contacto: ''
  }
  
  if (!condicionSeleccionada.value) {
    errores.value.condicion = 'Seleccione una condición de pago'
    valido = false
  }
  
  if (!contactoSeleccionado.value) {
    errores.value.contacto = 'Seleccione un contacto'
    valido = false
  }
  
  return valido
}

const confirmarOrden = async () => {
  if (procesando.value) return
  
  if (!validarFormulario()) return
  
  procesando.value = true
  
  const payload = {
    id_contacto: contactoSeleccionado.value.id_contacto,
    numero_order: props.numeroOrden,
    fecha_expedicion: new Date().toISOString().split('T')[0],
    idcondicion_pago: condicionSeleccionada.value,
    items: props.items.map(i => ({ id_item: i.id_item, cantidad: i.cantidad }))
  }

  console.log('Enviando payload para crear orden:', payload)
  
  try {
    const res = await ordersApi.createOrder(payload)
    console.log('Respuesta de createOrder:', res)
    
    if (res.success) {
      const orderId = res.data.id_order || 1
      console.log('Orden creada con ID:', orderId)
      ordenProcesadaId.value = orderId
    } else {
      throw new Error(res.error || 'Error desconocido al procesar orden')
    }
  } catch (error) {
    console.error('Error al procesar orden:', error)
    alert('❌ Error al procesar orden: ' + (error.message || 'Error desconocido'))
    procesando.value = false
  }
}

const finalizarProceso = () => {
  console.log('PDF generado, cerrando modal')
  ordenProcesadaId.value = null
  procesando.value = false
  emit('ordenProcesada')
}

const cancelarOrden = () => {
  emit('close')
}

onMounted(async () => {
  try {
    const res = await condicionesApi.getAll()
    if (res.success) condiciones.value = res.data
  } catch (error) {
    console.error('Error al cargar condiciones de pago:', error)
  }
})
</script>

<style>
:root {
  --color-primary: #4f46e5;
  --color-primary-dark: #4338ca;
}

.bg-primary {
  background-color: var(--color-primary);
}

.bg-primary-dark {
  background-color: var(--color-primary-dark);
}

.text-primary {
  color: var(--color-primary);
}

.border-primary {
  border-color: var(--color-primary);
}

.focus\:ring-primary:focus {
  --tw-ring-color: var(--color-primary);
}

.hover\:bg-primary-dark:hover {
  background-color: var(--color-primary-dark);
}
</style>