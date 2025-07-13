<template>
  <div class="max-w-6xl mx-auto p-6">
    <!-- Header con botón para abrir modal -->
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-gray-800">Gestión de Timbrados</h1>
      <button 
        @click="openModal"
        class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-medium transition-colors duration-200 flex items-center gap-2"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
        </svg>
        Nuevo Timbrado
      </button>
    </div>

    <!-- Mensajes de estado -->
    <div v-if="error" class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg mb-6">
      <div class="flex items-center">
        <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd"></path>
        </svg>
        {{ error }}
      </div>
    </div>

    <div v-if="success" class="bg-green-50 border border-green-200 text-green-700 px-4 py-3 rounded-lg mb-6">
      <div class="flex items-center">
        <svg class="w-5 h-5 mr-2" fill="currentColor" viewBox="0 0 20 20">
          <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd"></path>
        </svg>
        {{ successMessage }}
      </div>
    </div>

    <!-- Tabla de timbrados -->
    <div class="bg-white shadow-lg rounded-lg overflow-hidden">
      <div class="px-6 py-4 bg-gray-50 border-b border-gray-200">
        <h2 class="text-lg font-semibold text-gray-800">Lista de Timbrados</h2>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                # Autorización
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Fecha Autorización
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Inicio Vigencia
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Estado
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Establecimiento
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Punto Expedición
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Acciones
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="t in timbrados" :key="t.id_timbrado" class="hover:bg-gray-50 transition-colors duration-150">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ t.numero_autorizacion }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatFecha(t.fecha_autorizacion) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ formatFecha(t.fecha_inicio_vigencia) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="t.estado_timbrado ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" 
                      class="inline-flex px-2 py-1 text-xs font-semibold rounded-full">
                  {{ t.estado_timbrado ? 'Vigente' : 'No vigente' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ getNombreEstablecimiento(t.id_establecimiento) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ getNombrePunto(t.id_punto_expedicion) }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium">
                <div class="flex items-center gap-2">
                  <button 
                    @click="editarTimbrado(t)" 
                    class="text-blue-600 hover:text-blue-900 transition-colors duration-150"
                    title="Editar"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                  </button>
                  <button 
                    @click="eliminar(t.id_timbrado)" 
                    class="text-red-600 hover:text-red-900 transition-colors duration-150"
                    title="Eliminar"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"></path>
                    </svg>
                  </button>
                </div>
              </td>
            </tr>
            <tr v-if="timbrados.length === 0">
              <td colspan="7" class="px-6 py-8 text-center text-gray-500">
                <div class="flex flex-col items-center">
                  <svg class="w-12 h-12 text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
                  </svg>
                  <p class="text-lg font-medium">No hay timbrados registrados</p>
                  <p class="text-sm">Haz clic en "Nuevo Timbrado" para agregar el primero</p>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-lg shadow-xl max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <!-- Header del modal -->
        <div class="flex justify-between items-center p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-800">
            {{ isEditing ? 'Editar Timbrado' : 'Registrar Nuevo Timbrado' }}
          </h2>
          <button @click="closeModal" class="text-gray-400 hover:text-gray-600 transition-colors">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- Formulario -->
        <form @submit.prevent="handleSubmit" class="p-6">
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Número de Autorización *
              </label>
              <input
                v-model="timbrado.numero_autorizacion"
                type="text"
                required
                pattern="[0-9]+"
                minlength="6"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                placeholder="Ej: 123456"
              />
              <p v-if="timbrado.numero_autorizacion && !/^[0-9]+$/.test(timbrado.numero_autorizacion)" 
                 class="text-red-600 text-sm mt-1">
                Solo se permiten números
              </p>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Fecha de Autorización *
              </label>
              <input 
                v-model="timbrado.fecha_autorizacion" 
                type="date" 
                required 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Inicio de Vigencia *
              </label>
              <input 
                v-model="timbrado.fecha_inicio_vigencia" 
                type="date" 
                required 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              />
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Estado
              </label>
              <select 
                v-model="timbrado.estado_timbrado" 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option :value="true">Vigente</option>
                <option :value="false">No vigente</option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Establecimiento *
              </label>
              <select 
                v-model.number="timbrado.id_establecimiento" 
                required 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option disabled value="">Seleccione uno</option>
                <option v-for="e in establecimientos" :key="e.id_establecimiento" :value="e.id_establecimiento">
                  {{ e.numero_establecimiento }}
                </option>
              </select>
            </div>

            <div>
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Punto de Expedición *
              </label>
              <select 
                v-model.number="timbrado.id_punto_expedicion" 
                required 
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              >
                <option disabled value="">Seleccione uno</option>
                <option v-for="p in puntos" :key="p.id_punto_expedicion" :value="p.id_punto_expedicion">
                  {{ p.numero_expedicion }}
                </option>
              </select>
            </div>
          </div>

          <!-- Botones del modal -->
          <div class="flex justify-end gap-3 mt-8 pt-6 border-t border-gray-200">
            <button 
              type="button" 
              @click="closeModal"
              class="px-4 py-2 text-sm font-medium text-gray-700 bg-white border border-gray-300 rounded-md hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 transition-colors"
            >
              Cancelar
            </button>
            <button 
              type="submit" 
              :disabled="isSubmitting"
              class="px-4 py-2 text-sm font-medium text-white bg-blue-600 border border-transparent rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
            >
              <span v-if="isSubmitting" class="flex items-center">
                <svg class="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                  <circle class="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"></circle>
                  <path class="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                </svg>
                Guardando...
              </span>
              <span v-else>{{ isEditing ? 'Actualizar Timbrado' : 'Guardar Timbrado' }}</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getTimbradosVigentes, crearTimbrado, eliminarTimbrado, actualizarTimbrado } from '@/api/timbrado'
import establecimientoApi from '@/api/establecimiento'
import puntoExpApi from '@/api/punto_expedicion'

const isEditing = ref(false)
const editingId = ref(null)

const establecimientos = ref([])
const puntos = ref([])
const timbrados = ref([])
const error = ref('')
const success = ref(false)
const successMessage = ref('')
const showModal = ref(false)
const isSubmitting = ref(false)

const timbrado = ref({
  numero_autorizacion: '',
  fecha_autorizacion: '',
  fecha_inicio_vigencia: '',
  estado_timbrado: true,
  id_establecimiento: '',
  id_punto_expedicion: ''
})

const openModal = () => {
  showModal.value = true
  // Limpiar errores previos
  error.value = ''
  success.value = false
}

const editarTimbrado = (timbradoData) => {
  isEditing.value = true
  editingId.value = timbradoData.id_timbrado
  
  // Cargar los datos en el formulario
  timbrado.value = {
    numero_autorizacion: timbradoData.numero_autorizacion,
    fecha_autorizacion: timbradoData.fecha_autorizacion.split('T')[0], // Convertir a formato date input
    fecha_inicio_vigencia: timbradoData.fecha_inicio_vigencia.split('T')[0],
    estado_timbrado: timbradoData.estado_timbrado,
    id_establecimiento: timbradoData.id_establecimiento,
    id_punto_expedicion: timbradoData.id_punto_expedicion
  }
  
  showModal.value = true
  error.value = ''
  success.value = false
}

const closeModal = () => {
  showModal.value = false
  isEditing.value = false
  editingId.value = null
  // Resetear formulario
  timbrado.value = {
    numero_autorizacion: '',
    fecha_autorizacion: '',
    fecha_inicio_vigencia: '',
    estado_timbrado: true,
    id_establecimiento: '',
    id_punto_expedicion: ''
  }
  error.value = ''
}

const handleSubmit = async () => {
  error.value = ''
  success.value = false
  successMessage.value = ''
  isSubmitting.value = true

  // Validación adicional para numero_autorizacion
  if (!/^[0-9]+$/.test(timbrado.value.numero_autorizacion)) {
    error.value = 'El número de autorización debe contener solo números'
    isSubmitting.value = false
    return
  }
  if (timbrado.value.numero_autorizacion.length < 6) {
    error.value = 'El número de autorización debe tener al menos 6 dígitos'
    isSubmitting.value = false
    return
  }

  try {
    const payload = {
      ...timbrado.value,
      fecha_autorizacion: new Date(timbrado.value.fecha_autorizacion).toISOString(),
      fecha_inicio_vigencia: new Date(timbrado.value.fecha_inicio_vigencia).toISOString()
    }

    let res
    if (isEditing.value) {
      res = await actualizarTimbrado(editingId.value, payload)
    } else {
      res = await crearTimbrado(payload)
    }

    // El PUT devuelve 204 (No Content), el POST devuelve 201 (Created)
    if (res.status === 204 || res.status === 201) {
      success.value = true
      successMessage.value = isEditing.value ? '✅ Timbrado actualizado correctamente' : '✅ Timbrado guardado correctamente'
      closeModal()
      setTimeout(() => {
        success.value = false
        successMessage.value = ''
      }, 3000)
      await cargarDatos()
    }
  } catch (err) {
    error.value = err.response?.data?.error || (isEditing.value ? 'Error al actualizar timbrado' : 'Error al guardar timbrado')
  } finally {
    isSubmitting.value = false
  }
}

const eliminar = async (id) => {
  if (confirm('¿Seguro que quieres eliminar este timbrado?')) {
    try {
      const res = await eliminarTimbrado(id)
      if (res.status === 204) {
        success.value = true
        successMessage.value = '✅ Timbrado eliminado correctamente'
        setTimeout(() => {
          success.value = false
          successMessage.value = ''
        }, 3000)
        await cargarDatos()
      } else {
        error.value = 'Error al eliminar timbrado'
      }
    } catch (err) {
      error.value = err.response?.data?.error || 'Error al eliminar timbrado'
    }
  }
}

const formatFecha = (iso) => {
  if (!iso) return '—'
  const date = new Date(iso)
  return isNaN(date.getTime()) ? '—' : date.toLocaleDateString()
}

const getNombreEstablecimiento = (id) => {
  const est = establecimientos.value.find(e => e.id_establecimiento === id)
  return est ? est.numero_establecimiento : '—'
}

const getNombrePunto = (id) => {
  const pto = puntos.value.find(p => p.id_punto_expedicion === id)
  return pto ? pto.numero_expedicion : '—'
}

const cargarDatos = async () => {
  try {
    const [est, pto, tim] = await Promise.all([
      establecimientoApi.getAll(),
      puntoExpApi.getAll(),
      getTimbradosVigentes()
    ])

    if (est.success) establecimientos.value = est.data
    else error.value = est.error

    if (pto.success) puntos.value = pto.data
    else error.value = pto.error

    if (tim.status === 200) timbrados.value = tim.data
    else error.value = 'Error al obtener timbrados'
  } catch (err) {
    error.value = 'Error al cargar datos: ' + err.message
  }
}

onMounted(cargarDatos)
</script>

<style scoped>
/* Animaciones para el modal */
.modal-enter-active, .modal-leave-active {
  transition: opacity 0.3s ease;
}
.modal-enter-from, .modal-leave-to {
  opacity: 0;
}
</style>