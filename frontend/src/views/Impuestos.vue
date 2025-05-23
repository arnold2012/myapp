<template>
  <section class="max-w-4xl mx-auto p-6">
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold">Impuestos</h1>
      <button @click="abrirModal()" class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
        Nuevo Impuesto
      </button>
    </div>

    <table class="w-full bg-white shadow rounded-md">
      <thead class="bg-gray-100 text-left">
        <tr>
          <th class="p-3">Cod</th>
          <th class="p-3">Porcentaje IVA</th>
          <th class="p-3 text-center">Acciones</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="imp in impuestos" :key="imp.id_impuesto" class="border-t hover:bg-gray-50">
          <td class="p-3">{{ imp.id_impuesto }}</td>
          <td class="p-3">{{ imp.porcentaje_iva }}%</td>
          <td class="p-3 text-center">
            <button @click="abrirModal(imp)" class="text-blue-600 hover:underline mr-2">Editar</button>
            <button @click="eliminarImpuesto(imp.id_impuesto)" class="text-red-600 hover:underline">Eliminar</button>
          </td>
        </tr>
      </tbody>
    </table>

    <div v-if="mostrarModal" class="fixed inset-0  bg-opacity-40 flex items-center justify-center">
      <div class="bg-white rounded-lg p-6 max-w-md w-full">
        <h2 class="text-lg font-semibold mb-4">
          {{ impuestoSeleccionado ? 'Editar Impuesto' : 'Nuevo Impuesto' }}
        </h2>
        <form @submit.prevent="guardar">
          <div class="mb-4">
            <label class="block text-sm font-medium text-gray-700">Porcentaje IVA</label>
            <input
              v-model.number="form.porcentaje_iva"
              type="number"
              step="0.01"
              min="0"
              required
              class="w-full p-2 border rounded"
            />
          </div>

          <div class="flex justify-end gap-3">
            <button type="button" @click="cerrarModal" class="px-4 py-2 bg-gray-200 rounded hover:bg-gray-300">
              Cancelar
            </button>
            <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
              Guardar
            </button>
          </div>
        </form>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import impuestosApi from '@/api/impuestos'

const impuestos = ref([])
const mostrarModal = ref(false)
const impuestoSeleccionado = ref(null)
const form = ref({ porcentaje_iva: 0 })

const cargarImpuestos = async () => {
  const res = await impuestosApi.getAll()
  if (res.success) impuestos.value = res.data
}

const abrirModal = (imp = null) => {
  impuestoSeleccionado.value = imp
  form.value = imp ? { ...imp } : { porcentaje_iva: 0 }
  mostrarModal.value = true
}

const cerrarModal = () => {
  mostrarModal.value = false
  impuestoSeleccionado.value = null
  form.value = { porcentaje_iva: 0 }
}

const guardar = async () => {
  const res = impuestoSeleccionado.value
    ? await impuestosApi.update(impuestoSeleccionado.value.id_impuesto, form.value)
    : await impuestosApi.create(form.value)

  if (res.success) {
    cerrarModal()
    cargarImpuestos()
  } else {
    alert('❌ ' + res.error)
  }
}

const eliminarImpuesto = async (id) => {
  if (confirm('¿Desea eliminar este impuesto?')) {
    const res = await impuestosApi.delete(id)
    if (res.success) cargarImpuestos()
    else alert('❌ ' + res.error)
  }
}

onMounted(cargarImpuestos)
</script>
<style scoped>
th, td {
  text-align: left;
}
</style>
