<template>
  <div class="max-w-6xl mx-auto p-4">
    <div class="flex justify-between items-center mb-6">
      <h2 class="text-2xl font-semibold">Contactos Registrados</h2>
      <button @click="openCreateModal" class="bg-emerald-500 hover:bg-emerald-600 text-white px-4 py-2 rounded">
        + Agregar Contacto
      </button>
    </div>

    <!-- Input de búsqueda -->
    <div class="mb-4">
      <input
        v-model="busqueda"
        @input="buscarContactos"
        type="text"
        placeholder="Buscar por razón social o RUC/CI..."
        class="w-full max-w-md px-4 py-2 border border-gray-300 rounded shadow-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
    </div>

    <!-- Tabla de contactos -->
    <table class="w-full text-left border-collapse">
      <thead class="bg-gray-100">
        <tr>
          <th class="p-3">Razón Social</th>
          <th class="p-3">RUC/CI</th>
          <th class="p-3">Teléfono</th>
          <th class="p-3">Dirección</th>
          <th class="p-3">Acciones</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="contacto in contactos" :key="contacto.id_contacto" class="hover:bg-gray-50">
          <td class="p-3">{{ contacto.razon_social }}</td>
          <td class="p-3">{{ contacto.ruc_ci }}</td>
          <td class="p-3">{{ contacto.telefono_cto || 'N/A' }}</td>
          <td class="p-3">{{ contacto.direccion_cto || 'N/A' }}</td>
          <td class="p-3 space-x-2">
            <button @click="openEditModal(contacto)" class="bg-blue-500 hover:bg-blue-600 text-white px-3 py-1 rounded">Editar</button>
            <button @click="confirmDelete(contacto.id_contacto)" class="bg-red-500 hover:bg-red-600 text-white px-3 py-1 rounded">Eliminar</button>
          </td>
        </tr>
      </tbody>
    </table>

    <div v-if="contactos.length === 0" class="mt-6 p-4 bg-yellow-50 text-yellow-700 text-center rounded">
      No hay contactos registrados
    </div>

    <!-- Modal para creación -->
    <ModelFormularioContacto
      :isOpen="createModalOpen"
      @close="closeCreateModal"
      @submit="createContacto"
    />

    <!-- Modal para edición -->
    <ModelEdicion
      :isOpen="editModalOpen"
      :contacto="contactoSeleccionado"
      @close="closeEditModal"
      @submit="updateContacto"
    />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import ModelFormularioContacto from '@/components/ModelFormularioContacto.vue'
import ModelEdicion from '@/components/ModelEdicion.vue'

const contactos = ref([])
const busqueda = ref('')
const createModalOpen = ref(false)
const editModalOpen = ref(false)
const contactoSeleccionado = ref(null)

const fetchContactos = async () => {
  try {
    const response = await axios.get('http://localhost:8080/api/contactos')
    contactos.value = response.data
  } catch (error) {
    console.error('Error obteniendo contactos:', error)
  }
}

const buscarContactos = async () => {
  const query = busqueda.value.trim()
  if (query.length < 2) {
    fetchContactos()
    return
  }
  try {
    const response = await axios.get(`http://localhost:8080/api/contactos/search?q=${query}`)
    contactos.value = response.data
  } catch (error) {
    console.error('Error buscando contactos:', error)
  }
}

const openCreateModal = () => {
  createModalOpen.value = true
}
const closeCreateModal = () => {
  createModalOpen.value = false
}
const createContacto = async (formData) => {
  try {
    await axios.post('http://localhost:8080/api/contactos', formData)
    await fetchContactos()
    closeCreateModal()
  } catch (error) {
    console.error('Error creando contacto:', error)
  }
}

const openEditModal = (contacto) => {
  contactoSeleccionado.value = { ...contacto }
  editModalOpen.value = true
}
const closeEditModal = () => {
  editModalOpen.value = false
}
const updateContacto = async (formData) => {
  try {
    await axios.put(`http://localhost:8080/api/contactos/${formData.id_contacto}`, formData)
    await fetchContactos()
    closeEditModal()
  } catch (error) {
    console.error('Error actualizando contacto:', error)
  }
}

const confirmDelete = async (id) => {
  if (confirm('¿Estás seguro de eliminar este contacto?')) {
    try {
      await axios.delete(`http://localhost:8080/api/contactos/${id}`)
      await fetchContactos()
    } catch (error) {
      console.error('Error eliminando contacto:', error)
    }
  }
}

onMounted(fetchContactos)
</script>
