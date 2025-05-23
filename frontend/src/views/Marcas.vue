<template>
  <div class="p-6">
    <div class="flex justify-between items-center mb-4">
      <h1 class="text-2xl font-bold">Gestión de Marcas</h1>
      <button
        @click="abrirFormulario()"
        class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        + Nueva Marca
      </button>
    </div>

    <div v-if="error" class="text-red-600 mb-4">{{ error }}</div>

    <div class="overflow-x-auto">
      <table class="min-w-full border">
        <thead>
          <tr class="bg-gray-100">
            <th class="px-4 py-2 border">Cod</th>
            <th class="px-4 py-2 border">Descripción</th>
            <th class="px-4 py-2 border">Acciones</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="marca in marcas" :key="marca.id_marcas">
            <td class="px-4 py-2 border">{{ marca.id_marcas }}</td>
            <td class="px-4 py-2 border">{{ marca.descripcion_marcas }}</td>
            <td class="px-4 py-2 border">
              <button @click="abrirFormulario(marca)" class="text-blue-600 hover:underline mr-2">Editar</button>
              <button @click="eliminarMarca(marca.id_marcas)" class="text-red-600 hover:underline">Eliminar</button>
            </td>
          </tr>
          <tr v-if="marcas.length === 0">
            <td colspan="3" class="text-center py-4 text-gray-500">No hay marcas registradas</td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- Modal Formulario -->
    <div v-if="mostrarModal" class="fixed inset-0  bg-opacity-50 flex items-center justify-center z-50">
      <div class="bg-white p-6 rounded shadow-xl w-full max-w-md">
        <h2 class="text-lg font-semibold mb-4">{{ form.id_marcas ? 'Editar Marca' : 'Nueva Marca' }}</h2>
        <form @submit.prevent="guardarMarca">
          <label class="block mb-2 text-sm">Descripción:</label>
          <input
            v-model="form.descripcion_marcas"
            class="w-full p-2 border rounded mb-4"
            placeholder="Descripción de la marca"
            required
          />

          <div class="flex justify-end gap-3">
            <button type="button" @click="cerrarFormulario" class="bg-gray-300 px-4 py-2 rounded">Cancelar</button>
            <button type="submit" class="bg-blue-600 text-white px-4 py-2 rounded">Guardar</button>
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
const error = ref(null)
const mostrarModal = ref(false)
const form = ref({ id_marcas: null, descripcion_marcas: '' })

const cargarMarcas = async () => {
  const res = await marcasApi.getAll()
  if (res.success) {
    marcas.value = res.data
  } else {
    error.value = res.error
  }
}

const abrirFormulario = (marca = null) => {
  form.value = marca ? { ...marca } : { id_marcas: null, descripcion_marcas: '' }
  mostrarModal.value = true
}

const cerrarFormulario = () => {
  mostrarModal.value = false
  error.value = null
}

const guardarMarca = async () => {
  const res = form.value.id_marcas
    ? await marcasApi.update(form.value.id_marcas, form.value)
    : await marcasApi.create(form.value)

  if (res.success) {
    cerrarFormulario()
    cargarMarcas()
  } else {
    error.value = res.error
  }
}

const eliminarMarca = async (id) => {
  if (!confirm('¿Desea eliminar esta marca?')) return
  const res = await marcasApi.delete(id)
  if (res.success) cargarMarcas()
  else error.value = res.error
}

onMounted(cargarMarcas)
</script>

<style scoped>
th, td {
  text-align: left;
}
</style>
