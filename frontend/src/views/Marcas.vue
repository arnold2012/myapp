<template>
  <div class="max-w-6xl mx-auto p-6">
    <!-- Header con botón para abrir modal -->
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-3xl font-bold text-gray-800">Gestión de Marcas</h1>
      <button 
        @click="abrirFormulario()"
        class="bg-blue-600 hover:bg-blue-700 text-white px-6 py-3 rounded-lg font-medium transition-colors duration-200 flex items-center gap-2"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4"></path>
        </svg>
        Nueva Marca
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

    <!-- Tabla de marcas -->
    <div class="bg-white shadow-lg rounded-lg overflow-hidden">
      <div class="px-6 py-4 bg-gray-50 border-b border-gray-200">
        <h2 class="text-lg font-semibold text-gray-800">Lista de Marcas</h2>
      </div>
      
      <div class="overflow-x-auto">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Código
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Descripción
              </th>
              <th class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider">
                Acciones
              </th>
            </tr>
          </thead>
          <tbody class="bg-white divide-y divide-gray-200">
            <tr v-for="marca in marcas" :key="marca.id_marcas" class="hover:bg-gray-50 transition-colors duration-150">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">
                {{ marca.id_marcas }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                {{ marca.descripcion_marcas }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm font-medium text-right">
                <div class="flex items-center justify-end gap-2">
                  <button 
                    @click="abrirFormulario(marca)" 
                    class="text-blue-600 hover:text-blue-900 transition-colors duration-150"
                    title="Editar"
                  >
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"></path>
                    </svg>
                  </button>
                  <button 
                    @click="eliminarMarca(marca.id_marcas)" 
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
            <tr v-if="marcas.length === 0">
              <td colspan="3" class="px-6 py-8 text-center text-gray-500">
                <div class="flex flex-col items-center">
                  <svg class="w-12 h-12 text-gray-300 mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 11V7a4 4 0 00-8 0v4M5 9h14l1 12H4L5 9z"></path>
                  </svg>
                  <p class="text-lg font-medium">No hay marcas registradas</p>
                  <p class="text-sm">Haz clic en "Nueva Marca" para agregar la primera</p>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Modal -->
    <div v-if="mostrarModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
      <div class="bg-white rounded-lg shadow-xl max-w-md w-full max-h-[90vh] overflow-y-auto">
        <!-- Header del modal -->
        <div class="flex justify-between items-center p-6 border-b border-gray-200">
          <h2 class="text-xl font-semibold text-gray-800">
            {{ form.id_marcas ? 'Editar Marca' : 'Nueva Marca' }}
          </h2>
          <button @click="cerrarFormulario" class="text-gray-400 hover:text-gray-600 transition-colors">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
            </svg>
          </button>
        </div>

        <!-- Formulario -->
        <form @submit.prevent="guardarMarca" class="p-6">
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700 mb-2">
              Descripción *
            </label>
            <input
              v-model="form.descripcion_marcas"
              type="text"
              required
              class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
              placeholder="Ej: Samsung"
            />
            <p class="mt-1 text-sm text-gray-500">
              Ingrese el nombre de la marca
            </p>
          </div>

          <!-- Botones del modal -->
          <div class="flex justify-end gap-3 mt-8 pt-6 border-t border-gray-200">
            <button 
              type="button" 
              @click="cerrarFormulario"
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
              <span v-else>{{ form.id_marcas ? 'Actualizar' : 'Guardar' }}</span>
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import marcasApi from '@/api/marcas'

const marcas = ref([])
const error = ref('')
const success = ref(false)
const successMessage = ref('')
const mostrarModal = ref(false)
const form = ref({ id_marcas: null, descripcion_marcas: '' })
const isSubmitting = ref(false)

const cargarMarcas = async () => {
  try {
    const res = await marcasApi.getAll()
    if (res.success) {
      marcas.value = res.data
    } else {
      error.value = res.error || 'Error al cargar marcas'
      setTimeout(() => { error.value = '' }, 5000)
    }
  } catch (err) {
    error.value = 'Error de conexión al servidor'
    setTimeout(() => { error.value = '' }, 5000)
  }
}

const abrirFormulario = (marca = null) => {
  form.value = marca ? { ...marca } : { id_marcas: null, descripcion_marcas: '' }
  mostrarModal.value = true
  error.value = ''
  success.value = false
}

const cerrarFormulario = () => {
  mostrarModal.value = false
  form.value = { id_marcas: null, descripcion_marcas: '' }
}

const guardarMarca = async () => {
  isSubmitting.value = true
  error.value = ''
  success.value = false
  
  try {
    const res = form.value.id_marcas
      ? await marcasApi.update(form.value.id_marcas, form.value)
      : await marcasApi.create(form.value)

    if (res.success) {
      cerrarFormulario()
      await cargarMarcas()
      success.value = true
      successMessage.value = form.value.id_marcas 
        ? '✅ Marca actualizada correctamente' 
        : '✅ Marca creada correctamente'
      setTimeout(() => { success.value = false }, 5000)
    } else {
      error.value = res.error || 'Error al guardar'
    }
  } catch (err) {
    error.value = 'Error de conexión al servidor'
  } finally {
    isSubmitting.value = false
  }
}

const eliminarMarca = async (id) => {
  if (!confirm('¿Está seguro que desea eliminar esta marca?')) return
  
  try {
    const res = await marcasApi.delete(id)
    if (res.success) {
      await cargarMarcas()
      success.value = true
      successMessage.value = '✅ Marca eliminada correctamente'
      setTimeout(() => { success.value = false }, 5000)
    } else {
      error.value = res.error || 'Error al eliminar'
      setTimeout(() => { error.value = '' }, 5000)
    }
  } catch (err) {
    error.value = 'Error de conexión al servidor'
    setTimeout(() => { error.value = '' }, 5000)
  }
}

onMounted(cargarMarcas)
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