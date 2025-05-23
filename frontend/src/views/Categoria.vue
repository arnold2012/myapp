<template>
  <div class="p-6">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-semibold text-gray-800">Categorías</h1>
      <button
        @click="abrirModal()"
        class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
      >
        + Nueva Categoría
      </button>
    </div>

    <div v-if="categorias.length" class="overflow-x-auto">
      <table class="min-w-full bg-white border border-gray-200">
        <thead class="bg-gray-100 text-left">
          <tr>
            <th class="px-4 py-2 border">Cod</th>
            <th class="px-4 py-2 border">Descripción</th>
            <th class="px-4 py-2 border text-center">Acciones</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="cat in categorias" :key="cat.id_categoria">
            <td class="px-4 py-2 border">{{ cat.id_categoria }}</td>
            <td class="px-4 py-2 border">{{ cat.describe_categoria }}</td>
            <td class="px-4 py-2 border text-center">
              <button @click="abrirModal(cat)" class="text-blue-600 hover:underline mr-2">Editar</button>
              <button @click="eliminar(cat.id_categoria)" class="text-red-600 hover:underline">Eliminar</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div v-else class="text-gray-600">No hay categorías registradas.</div>

    <!-- Modal -->
    <div v-if="modalVisible" class="fixed inset-0 bg-opacity-50 flex justify-center items-center z-50">
      <div class="bg-white rounded-lg p-6 w-full max-w-md">
        <h2 class="text-lg font-semibold mb-4">
          {{ form.id_categoria ? 'Editar Categoría' : 'Nueva Categoría' }}
        </h2>
        <form @submit.prevent="guardar">
          <label class="block mb-2 text-sm text-gray-700">Descripción</label>
          <input
            v-model="form.describe_categoria"
            type="text"
            class="w-full border px-3 py-2 rounded mb-4"
            required
          />

          <div class="flex justify-end gap-3">
            <button @click="cerrarModal" type="button" class="bg-gray-300 text-gray-800 px-4 py-2 rounded">Cancelar</button>
            <button type="submit" class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
              {{ form.id_categoria ? 'Actualizar' : 'Guardar' }}
            </button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '@/api/categorias'

const categorias = ref([])
const modalVisible = ref(false)
const form = ref({ id_categoria: null, describe_categoria: '' })

const cargar = async () => {
  const res = await api.getAll()
  if (res.success) categorias.value = res.data
}

const abrirModal = (cat = null) => {
  modalVisible.value = true
  form.value = cat ? { ...cat } : { id_categoria: null, describe_categoria: '' }
}

const cerrarModal = () => {
  modalVisible.value = false
  form.value = { id_categoria: null, describe_categoria: '' }
}

const guardar = async () => {
  const { id_categoria, describe_categoria } = form.value
  const payload = { describe_categoria }
  const res = id_categoria
    ? await api.update(id_categoria, payload)
    : await api.create(payload)

  if (res.success) {
    await cargar()
    cerrarModal()
  } else {
    alert(res.error)
  }
}

const eliminar = async (id) => {
  if (confirm('Deseas eliminar esta categoría?')) {
    const res = await api.delete(id)
    if (res.success) await cargar()
    else alert(res.error)
  }
}

onMounted(cargar)
</script>

<style scoped>
th, td {
  white-space: nowrap;
}
</style>
