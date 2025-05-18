<template>
  <section class="p-6 max-w-4xl mx-auto">
    <h1 class="text-2xl font-bold mb-6">Generar Factura</h1>

    <!-- Próximo número -->
    <div class="mb-4">
      <label class="block text-sm font-medium">Próximo Nº Factura</label>
      <div class="mt-1 text-lg font-mono text-blue-600">{{ numeroFactura }}</div>
    </div>

    <!-- Búsqueda de orden -->
    <div class="flex items-center gap-2 mb-6">
      <input
        v-model="numeroOrden"
        type="number"
        placeholder="Número de Orden"
        class="border px-3 py-2 rounded flex-1"
        @keyup.enter="buscarOrden"
      />
      <button
        @click="buscarOrden"
        class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
        :disabled="isLoading"
      >
        {{ isLoading ? 'Buscando...' : 'Buscar' }}
      </button>
    </div>

    <!-- Detalle de la orden -->
    <div v-if="detalles.length" class="bg-white shadow rounded p-4 mb-6">
      <h2 class="text-lg font-semibold mb-2">Datos de la Orden</h2>
      <p class="text-gray-700">Cliente: {{ detalles[0].razon_social }} ({{ detalles[0].ruc_ci }})</p>
      <p class="text-gray-700 mb-4">Fecha: {{ detalles[0].fecha }}</p>

      <table class="w-full text-sm border-collapse">
        <thead class="bg-gray-100">
          <tr>
            <th class="p-2">Código</th>
            <th class="p-2">Descripción</th>
            <th class="p-2">Cantidad</th>
            <th class="p-2">Precio Unitario</th>
            <th class="p-2">% IVA</th>
            <th class="p-2">Subtotal</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, i) in detalles" :key="i" class="border-t">
            <td class="p-2">{{ item.codigo_item }}</td>
            <td class="p-2">{{ item.descripcion_item }}</td>
            <td class="p-2">{{ item.cantidad }}</td>
            <td class="p-2">Gs {{ item.precio_unitario.toLocaleString('es-PY') }}</td>
            <td class="p-2">{{ item.porcentaje_iva }}%</td>
            <td class="p-2">Gs {{ item.subtotal.toLocaleString('es-PY') }}</td>
          </tr>
        </tbody>
      </table>

      <div class="mt-4 text-right font-semibold">
        Total: Gs {{ totalGeneral.toLocaleString('es-PY') }}
      </div>
    </div>

    <!-- Botón generar factura -->
    <div v-if="detalles.length" class="flex justify-between items-center">
      <div class="text-sm text-gray-600 font-mono">
         Factura Nro: {{ numeroFactura }}
      </div>
      <button
        @click="confirmarGenerarFactura"
        class="bg-green-600 text-white px-6 py-2 rounded hover:bg-green-700"
        :disabled="isLoading"
      >
        {{ isLoading ? 'Procesando...' : 'Generar Factura' }}
      </button>
    </div>

    <!-- Modal de confirmación -->
    <ModelConfirmarFactura
    :show="mostrarModalConfirmacion"
    :numeroFactura="numeroFactura"
    :cliente="clienteNombre"
    :total="totalGeneral"
    @confirm="confirmarGeneracion"
    @cancel="cancelarYLimpiar"
  />

  </section>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import facturarApi from '@/api/facturar'
import ModelConfirmarFactura from '@/components/ModelConfirmarFactura.vue'
const cancelarYLimpiar = () => {
  mostrarModalConfirmacion.value = false
  numeroOrden.value = ''
  detalles.value = []
  idOrder.value = null
}


const numeroOrden = ref('')
const detalles = ref([])
const idOrder = ref(null)
const numeroFactura = ref('...')
const isLoading = ref(false)

const establecimientos = ref([])
const puntosExpedicion = ref([])
const selectedEstablecimiento = ref('')
const selectedExpedicion = ref('')

const mostrarModalConfirmacion = ref(false)

onMounted(async () => {
  const [estRes, peRes] = await Promise.all([
    facturarApi.getEstablecimientos(),
    facturarApi.getPuntosExpedicion()
  ])
  if (estRes.success) establecimientos.value = estRes.data
  if (peRes.success) puntosExpedicion.value = peRes.data

  selectedEstablecimiento.value = establecimientos.value[0]?.id_establecimiento || ''
  selectedExpedicion.value = puntosExpedicion.value[0]?.id_punto_expedicion || ''

  await cargarProximoNumero()
})

const cargarProximoNumero = async () => {
  const estObj = establecimientos.value.find(e => e.id_establecimiento === selectedEstablecimiento.value)
  const expObj = puntosExpedicion.value.find(p => p.id_punto_expedicion === selectedExpedicion.value)
  if (!estObj || !expObj) return

  try {
    const res = await facturarApi.obtenerProximoNumeroFactura(
      estObj.numero_establecimiento,
      expObj.numero_expedicion
    )
    numeroFactura.value = res.success ? res.data.next_numero_factura : 'Error'
  } catch (error) {
    numeroFactura.value = 'Error'
    console.error(error)
  }
}

const buscarOrden = async () => {
  if (!numeroOrden.value) return
  isLoading.value = true
  try {
    const res = await facturarApi.buscarOrdenPorNumero(numeroOrden.value)
    if (res.success && res.data.length) {
      detalles.value = res.data
      idOrder.value = res.data[0].id_order
    } else {
      alert('❌ ' + (res.error || 'Orden no encontrada'))
      detalles.value = []
      idOrder.value = null
    }
  } catch (error) {
    alert('❌ Error de conexión')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}

const totalGeneral = computed(() =>
  detalles.value.reduce((sum, d) => sum + d.subtotal, 0)
)

const confirmarGenerarFactura = () => {
  mostrarModalConfirmacion.value = true
}

const cerrarModal = () => {
  mostrarModalConfirmacion.value = false
}

const confirmarGeneracion = async () => {
  mostrarModalConfirmacion.value = false
  await generarFactura()
}

const generarFactura = async () => {
  if (!idOrder.value) {
    alert('Primero buscá una orden válida.')
    return
  }
  if (!selectedEstablecimiento.value || !selectedExpedicion.value) {
    alert('Selecciona establecimiento y punto de expedición.')
    return
  }

  isLoading.value = true
  try {
    const timbradoRes = await facturarApi.getTimbradoPorEstablecimientoYExpedicion(
      selectedEstablecimiento.value,
      selectedExpedicion.value
    )

    if (!timbradoRes.success) {
      alert('❌ No se pudo obtener el timbrado: ' + timbradoRes.error)
      isLoading.value = false
      return
    }

    const cab = detalles.value[0]
    const payload = {
      id_contacto: cab.id_contacto,
      idcondicion_pago: cab.idcondicion_pago,
      id_punto_expedicion: selectedExpedicion.value,
      id_timbrado: timbradoRes.data,
      id_establecimiento: selectedEstablecimiento.value,
      id_order: idOrder.value,
      items: detalles.value.map(d => ({
        id_item: d.id_item,
        cantidad: d.cantidad,
        unit_price: d.precio_unitario
      }))
    }

    const res = await facturarApi.crearFactura(payload)
    if (res.success) {
      alert('✅ Factura generada correctamente')
      numeroOrden.value = ''
      detalles.value = []
      idOrder.value = null
      await cargarProximoNumero()
    } else {
      alert('❌ ' + res.error)
    }
  } catch (error) {
    alert('❌ Error al generar factura')
    console.error(error)
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
/* estilos opcionales */
</style>
