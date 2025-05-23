<template>
  <section class="p-4 md:p-6 max-w-6xl mx-auto">
    <div class="flex flex-col md:flex-row justify-between items-center mb-6 gap-4">
      <h1 class="text-2xl font-bold text-gray-800">Productos</h1>
      <div class="flex gap-2 w-full md:w-auto">
        <input
          v-model="busqueda"
          @input="buscar"
          type="text"
          placeholder="Buscar producto..."
          class="flex-1 md:w-64 border border-gray-300 px-3 py-2 rounded-md"
        />
        <button
          @click="abrirModal()"
          class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700"
        >
          Nuevo Producto
        </button>
      </div>
    </div>

    <!-- Estado de carga -->
    <div v-if="cargando" class="flex justify-center my-8">
      <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-blue-500"></div>
    </div>

    <!-- Mensaje sin resultados -->
    <div v-else-if="productosFiltrados.length === 0" class="bg-white rounded-lg shadow-md p-8 text-center">
      <h3 class="text-lg font-medium text-gray-700 mb-2">No se encontraron productos</h3>
      <p class="text-gray-500 mb-4">Intenta con otra búsqueda o agrega un nuevo producto</p>
      <button
        @click="abrirModal()"
        class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700"
      >
        Agregar Producto
      </button>
    </div>

    <!-- Tabla de productos -->
    <div v-else class="bg-white rounded-lg shadow-md overflow-hidden">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-100 text-sm text-left">
            <tr>
              <th class="p-3 font-semibold text-gray-600">Código</th>
              <th class="p-3 font-semibold text-gray-600">Descripción</th>
              <th class="p-3 font-semibold text-gray-600">Precio</th>
              <th class="p-3 font-semibold text-gray-600">IVA</th>
              <th class="p-3 font-semibold text-gray-600">Stock</th>
              <th class="p-3 font-semibold text-gray-600">Marca</th>
              <th class="p-3 font-semibold text-gray-600">Categoría</th>
              <th class="p-3 font-semibold text-gray-600">Estado</th>
              <th class="p-3 text-center font-semibold text-gray-600">Acciones</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="p in productosPaginados"
              :key="p.id_item"
              class="border-t text-sm hover:bg-gray-50"
            >
              <td class="p-3 font-medium">{{ p.cod_item }}</td>
              <td class="p-3">{{ p.descripcion_item }}</td>
              <td class="p-3">Gs {{ formatNumber(p.precio_unitario) }}</td>
              <td class="p-3">{{ p.porcentaje_iva }}%</td>
              <td class="p-3">{{ formatNumber(p.cantidad_inicial) }}</td>
              <td class="p-3">{{ p.descripcion_marcas }}</td>
              <td class="p-3">{{ p.describe_categoria }}</td>
              <td class="p-3">
                <span
                  :class="[
                    'px-2 py-1 rounded-full text-xs font-medium',
                    p.item_activo ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
                  ]"
                >
                  {{ p.item_activo ? 'Activo' : 'Inactivo' }}
                </span>
              </td>
              <td class="p-3 text-center whitespace-nowrap">
                <div class="flex justify-center gap-2">
                  <button 
                    @click="abrirModal(p)" 
                    class="bg-blue-600 text-white px-2 py-1 rounded text-xs hover:bg-blue-700"
                  >
                    Editar
                  </button>
                  <button 
                    @click="confirmarEliminar(p)" 
                    class="bg-red-600 text-white px-2 py-1 rounded text-xs hover:bg-red-700"
                  >
                    Eliminar
                  </button>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
      
      <!-- Paginación -->
      <div class="flex items-center justify-between border-t border-gray-200 bg-gray-50 px-4 py-3 sm:px-6">
        <div class="flex flex-1 justify-between sm:hidden">
          <button
            @click="cambiarPagina(paginaActual - 1)"
            :disabled="paginaActual === 1"
            :class="[
              'relative inline-flex items-center rounded-md px-4 py-2 text-sm font-medium',
              paginaActual === 1 
                ? 'bg-gray-100 text-gray-400 cursor-not-allowed' 
                : 'bg-white text-gray-700 hover:bg-gray-50'
            ]"
          >
            Anterior
          </button>
          <button
            @click="cambiarPagina(paginaActual + 1)"
            :disabled="paginaActual === totalPaginas"
            :class="[
              'relative ml-3 inline-flex items-center rounded-md px-4 py-2 text-sm font-medium',
              paginaActual === totalPaginas 
                ? 'bg-gray-100 text-gray-400 cursor-not-allowed' 
                : 'bg-white text-gray-700 hover:bg-gray-50'
            ]"
          >
            Siguiente
          </button>
        </div>
        <div class="hidden sm:flex sm:flex-1 sm:items-center sm:justify-between">
          <div>
            <p class="text-sm text-gray-700">
              Mostrando <span class="font-medium">{{ inicio }}</span> a <span class="font-medium">{{ fin }}</span> de <span class="font-medium">{{ productosFiltrados.length }}</span> resultados
            </p>
          </div>
          <div>
            <nav class="isolate inline-flex -space-x-px rounded-md shadow-sm" aria-label="Pagination">
              <button
                @click="cambiarPagina(paginaActual - 1)"
                :disabled="paginaActual === 1"
                class="relative inline-flex items-center rounded-l-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
                :class="{ 'cursor-not-allowed': paginaActual === 1 }"
              >
                <span class="sr-only">Anterior</span>
                &laquo;
              </button>
              
              <!-- Botones de página -->
              <template v-for="pagina in paginasVisibles" :key="pagina">
                <span
                  v-if="pagina === '...'"
                  class="relative inline-flex items-center px-4 py-2 text-sm font-semibold text-gray-700 ring-1 ring-inset ring-gray-300 focus:outline-offset-0"
                >
                  ...
                </span>
                <button
                  v-else
                  @click="cambiarPagina(pagina)"
                  :class="[
                    'relative inline-flex items-center px-4 py-2 text-sm font-semibold focus:z-20 focus:outline-offset-0',
                    pagina === paginaActual
                      ? 'bg-blue-600 text-white focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600'
                      : 'text-gray-900 ring-1 ring-inset ring-gray-300 hover:bg-gray-50'
                  ]"
                >
                  {{ pagina }}
                </button>
              </template>
              
              <button
                @click="cambiarPagina(paginaActual + 1)"
                :disabled="paginaActual === totalPaginas"
                class="relative inline-flex items-center rounded-r-md px-2 py-2 text-gray-400 ring-1 ring-inset ring-gray-300 hover:bg-gray-50 focus:z-20 focus:outline-offset-0"
                :class="{ 'cursor-not-allowed': paginaActual === totalPaginas }"
              >
                <span class="sr-only">Siguiente</span>
                &raquo;
              </button>
            </nav>
          </div>
        </div>
      </div>
    </div>

    <!-- Modal de formulario -->
    <ModalFormularioProductos
      v-if="mostrarModal"
      :producto="productoSeleccionado"
      @cerrar="cerrarModal"
      @guardado="productoGuardado"
    />

    <!-- Modal de confirmación para eliminar -->
    <div v-if="mostrarConfirmacion" class="fixed inset-0  bg-opacity-40 flex items-center justify-center p-4 z-50">
      <div class="bg-white rounded-lg shadow-xl max-w-md w-full p-6">
        <h2 class="text-xl font-semibold text-center mb-4">¿Eliminar este producto?</h2>
        <p class="text-gray-600 mb-6 text-center">
          Esta acción eliminará permanentemente el producto <strong>{{ productoAEliminar?.descripcion_item }}</strong>
        </p>
        <div class="flex justify-center gap-3">
          <button
            @click="mostrarConfirmacion = false"
            class="bg-gray-200 text-gray-800 px-4 py-2 rounded-md hover:bg-gray-300"
          >
            Cancelar
          </button>
          <button
            @click="eliminarConfirmado"
            class="bg-red-600 text-white px-4 py-2 rounded-md hover:bg-red-700"
          >
            Eliminar
          </button>
        </div>
      </div>
    </div>

    <!-- Toast de notificación -->
    <div 
      v-if="notificacion.mostrar" 
      :class="[
        'fixed bottom-4 right-4 p-4 rounded-lg shadow-lg max-w-xs z-50 transition-all duration-300 transform',
        notificacion.tipo === 'success' ? 'bg-green-600 text-white' : 'bg-red-600 text-white',
        notificacion.mostrar ? 'translate-y-0 opacity-100' : 'translate-y-10 opacity-0'
      ]"
    >
      <p>{{ notificacion.mensaje }}</p>
    </div>
  </section>
</template>

<script>
import { ref, computed, onMounted, watch } from 'vue'
import productosApi from '@/api/productos'
import ModalFormularioProductos from '@/components/ModalFormularioProductos.vue'

export default {
  name: "ProductoView",
  components: {
    ModalFormularioProductos
  },
  setup() {
    const productos = ref([])
    const busqueda = ref('')
    const mostrarModal = ref(false)
    const productoSeleccionado = ref(null)
    const cargando = ref(true)
    const mostrarConfirmacion = ref(false)
    const productoAEliminar = ref(null)
    const notificacion = ref({
      mostrar: false,
      mensaje: '',
      tipo: 'success'
    })

    // Variables para paginación
    const paginaActual = ref(1)
    const itemsPorPagina = ref(10)

    // Productos filtrados según la búsqueda
    const productosFiltrados = computed(() => {
      if (!busqueda.value.trim()) return productos.value
      
      const busquedaLower = busqueda.value.toLowerCase()
      return productos.value.filter(p => 
        p.cod_item.toLowerCase().includes(busquedaLower) || 
        p.descripcion_item.toLowerCase().includes(busquedaLower)
      )
    })

    // Total de páginas
    const totalPaginas = computed(() => {
      return Math.ceil(productosFiltrados.value.length / itemsPorPagina.value)
    })

    // Productos de la página actual
    const productosPaginados = computed(() => {
      const inicio = (paginaActual.value - 1) * itemsPorPagina.value
      const fin = inicio + itemsPorPagina.value
      return productosFiltrados.value.slice(inicio, fin)
    })

    // Índices para mostrar en la paginación
    const inicio = computed(() => {
      if (productosFiltrados.value.length === 0) return 0
      return (paginaActual.value - 1) * itemsPorPagina.value + 1
    })

    const fin = computed(() => {
      if (productosFiltrados.value.length === 0) return 0
      return Math.min(paginaActual.value * itemsPorPagina.value, productosFiltrados.value.length)
    })

    // Páginas visibles en la paginación
    const paginasVisibles = computed(() => {
      const totalPags = totalPaginas.value
      const actual = paginaActual.value
      const delta = 1 // Número de páginas a mostrar a cada lado de la página actual
      
      if (totalPags <= 5) {
        // Si hay 5 o menos páginas, mostrar todas
        return Array.from({ length: totalPags }, (_, i) => i + 1)
      }
      
      let paginas = []
      
      // Siempre incluir la primera página
      paginas.push(1)
      
      // Agregar puntos suspensivos si la página actual está lejos del inicio
      if (actual > 2 + delta) {
        paginas.push('...')
      }
      
      // Agregar páginas alrededor de la página actual
      const rangoInicio = Math.max(2, actual - delta)
      const rangoFin = Math.min(totalPags - 1, actual + delta)
      
      for (let i = rangoInicio; i <= rangoFin; i++) {
        paginas.push(i)
      }
      
      // Agregar puntos suspensivos si la página actual está lejos del final
      if (actual < totalPags - 1 - delta) {
        paginas.push('...')
      }
      
      // Siempre incluir la última página
      if (totalPags > 1) {
        paginas.push(totalPags)
      }
      
      return paginas
    })

    const cargarProductos = async () => {
      cargando.value = true
      try {
        const res = await productosApi.buscarProductos(busqueda.value || '')
        if (res.success) {
          productos.value = [...res.data]
          // Resetear a la primera página cuando cambia la búsqueda
          paginaActual.value = 1
        } else {
          mostrarNotificacion('Error al cargar productos', 'error')
        }
      } catch (error) {
        mostrarNotificacion('Error de conexión', 'error')
        console.error('Error al cargar productos:', error)
      } finally {
        cargando.value = false
      }
    }

    const buscar = () => {
      cargarProductos()
    }

    const cambiarPagina = (pagina) => {
      if (pagina < 1 || pagina > totalPaginas.value) return
      paginaActual.value = pagina
    }

    const abrirModal = (producto = null) => {
      productoSeleccionado.value = producto ? { ...producto } : null
      mostrarModal.value = true
    }

    const cerrarModal = () => {
      mostrarModal.value = false
      productoSeleccionado.value = null
    }

    const productoGuardado = () => {
      mostrarNotificacion('Producto guardado con éxito', 'success')
      cerrarModal()
      cargarProductos()
    }

    const confirmarEliminar = (producto) => {
      productoAEliminar.value = producto
      mostrarConfirmacion.value = true
    }

    const eliminarConfirmado = async () => {
      if (!productoAEliminar.value) return
      
      try {
        const res = await productosApi.deleteProducto(productoAEliminar.value.id_item)
        if (res.success) {
          mostrarNotificacion('Producto eliminado con éxito', 'success')
          cargarProductos()
        } else {
          mostrarNotificacion(`Error al eliminar: ${res.error}`, 'error')
        }
      } catch (error) {
        mostrarNotificacion('Error de conexión', 'error')
      } finally {
        mostrarConfirmacion.value = false
        productoAEliminar.value = null
      }
    }

    const mostrarNotificacion = (mensaje, tipo = 'success') => {
      notificacion.value = {
        mostrar: true,
        mensaje,
        tipo
      }
      
      setTimeout(() => {
        notificacion.value.mostrar = false
      }, 3000)
    }

    const formatNumber = (num) => {
      return num.toLocaleString('es-PY')
    }

    // Resetear la página cuando cambia la búsqueda
    watch(busqueda, () => {
      paginaActual.value = 1
    })

    onMounted(cargarProductos)

    return {
      productos,
      busqueda,
      mostrarModal,
      productoSeleccionado,
      cargando,
      mostrarConfirmacion,
      productoAEliminar,
      notificacion,
      paginaActual,
      itemsPorPagina,
      productosFiltrados,
      totalPaginas,
      productosPaginados,
      inicio,
      fin,
      paginasVisibles,
      cargarProductos,
      buscar,
      cambiarPagina,
      abrirModal,
      cerrarModal,
      productoGuardado,
      confirmarEliminar,
      eliminarConfirmado,
      mostrarNotificacion,
      formatNumber
    }
  }
}
</script>