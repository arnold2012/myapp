<template>
  <section class="flex flex-col h-screen max-h-screen overflow-hidden p-2 sm:p-3">
    <!-- Cabecera responsive -->
    <div class="flex flex-col sm:flex-row sm:justify-between sm:items-center mb-2 gap-2">
      <h1 class="text-xl font-bold text-gray-900">Historial de Transacciones</h1>
      <div class="flex items-center space-x-2">
        <div class="relative">
          <input
            v-model="busquedaOrden"
            placeholder="Buscar por número de orden..."
            class="w-full px-3 py-2 pl-8 text-sm border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-emerald-500 focus:border-emerald-500"
            type="text"
            @input="filtrarOrdenes"
          />
          <div class="absolute inset-y-0 left-0 pl-2 flex items-center pointer-events-none">
            <component :is="SearchIcon" class="h-4 w-4 text-gray-400" />
          </div>
          <button 
            @click="busquedaOrden = ''" 
            v-if="busquedaOrden"
            class="absolute inset-y-0 right-0 pr-2 flex items-center"
          >
            <component :is="XIcon" class="h-4 w-4 text-gray-400 hover:text-emerald-500" />
          </button>
        </div>
      </div>
    </div>

    <!-- Estado de carga inicial -->
    <div v-if="loadingOrdenes" class="flex-1 flex items-center justify-center">
      <div class="flex flex-col items-center">
        <component :is="LoaderIcon" class="animate-spin h-8 w-8 text-emerald-500 mb-2" />
        <p class="text-sm text-gray-500">Cargando lista de órdenes...</p>
      </div>
    </div>

    <!-- Error al cargar órdenes -->
    <div v-else-if="errorOrdenes" class="flex-1 flex items-center justify-center">
      <div class="bg-red-50 border border-red-200 rounded-md p-4 max-w-md">
        <div class="flex">
          <div class="flex-shrink-0">
            <component :is="AlertIcon" class="h-5 w-5 text-red-400" />
          </div>
          <div class="ml-3">
            <h3 class="text-sm font-medium text-red-800">Error al cargar órdenes</h3>
            <div class="mt-2 text-sm text-red-700">
              <p>{{ errorOrdenes }}</p>
            </div>
            <div class="mt-4">
              <button
                @click="cargarOrdenes"
                class="inline-flex items-center px-2.5 py-1.5 border border-transparent text-xs font-medium rounded text-red-700 bg-red-100 hover:bg-red-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
              >
                Reintentar
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Lista de órdenes y detalles -->
    <div v-else class="flex-1 flex flex-col min-h-0 overflow-hidden">
      <!-- Vista de lista (cuando no hay detalles seleccionados) -->
      <div v-if="!ordenSeleccionada" class="flex-1 flex flex-col min-h-0">
        <!-- Sin resultados en la búsqueda -->
        <div v-if="ordenesFiltradas.length === 0" class="flex-1 flex items-center justify-center">
          <div class="text-center">
            <component :is="SearchIcon" class="mx-auto h-12 w-12 text-gray-400" />
            <h3 class="mt-2 text-sm font-medium text-gray-900">No se encontraron órdenes</h3>
            <p class="mt-1 text-sm text-gray-500">
              Intenta con otro número de orden o limpia el filtro.
            </p>
          </div>
        </div>

        <!-- Tabla de órdenes -->
        <div v-else class="flex-1 flex flex-col min-h-0 border border-gray-200 rounded-md shadow-sm">
          <!-- Encabezados fijos -->
          <div class="sheets-header bg-gray-50 border-b border-gray-200 sticky top-0 z-10 text-xs sm:text-sm">
            <div class="sheets-cell w-16 sm:w-24 font-medium">Número</div>
            <div class="sheets-cell w-20 sm:w-28 font-medium">Fecha</div>
            <div class="sheets-cell flex-1 font-medium">Cliente</div>
            <div class="sheets-cell w-20 sm:w-28 font-medium">Condición</div>
            <div class="sheets-cell w-20 sm:w-28 font-medium">Total</div>
            <div class="sheets-cell w-14 sm:w-20 font-medium">Acciones</div>
          </div>
          
          <!-- Contenedor con scrollbar -->
          <div class="overflow-y-auto overflow-x-auto flex-1 scrollbar-thin scrollbar-thumb-gray-300 scrollbar-track-gray-100">
            <!-- Filas de órdenes -->
            <div 
              v-for="(orden, index) in ordenesFiltradas" 
              :key="orden.id_order"
              class="sheets-row text-xs sm:text-sm hover:bg-emerald-50 cursor-pointer"
              :class="{'bg-gray-50': index % 2 === 0}"
              @click="verDetallesOrden(orden.id_order)"
            >
              <div class="sheets-cell w-16 sm:w-24">{{ orden.numero_order }}</div>
              <div class="sheets-cell w-20 sm:w-28">{{ formatearFecha(orden.fecha_expedicion) }}</div>
              <div class="sheets-cell flex-1 truncate" :title="orden.razon_social">
                {{ orden.razon_social }}
              </div>
              <div class="sheets-cell w-20 sm:w-28 truncate" :title="orden.describe_condicion">
                {{ orden.describe_condicion }}
              </div>
              <div class="sheets-cell w-20 sm:w-28 font-medium">
                {{ formatearMoneda(orden.total_general) }}
              </div>
              <div class="sheets-cell w-14 sm:w-20">
                <button 
                  @click.stop="verDetallesOrden(orden.id_order)"
                  class="text-emerald-600 hover:text-emerald-800 p-1"
                >
                  <component :is="EyeIcon" class="h-4 w-4" />
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Vista de detalles (cuando hay un orden seleccionada) -->
      <div v-else class="flex-1 flex flex-col min-h-0">
        <!-- Botón volver a la lista -->
        <button 
          @click="ordenSeleccionada = null; historialDetalle = []" 
          class="flex items-center text-sm text-emerald-600 hover:text-emerald-800 mb-3"
        >
          <component :is="ArrowLeftIcon" class="h-4 w-4 mr-1" />
          Volver a la lista
        </button>

        <!-- Estado de carga para detalles -->
        <div v-if="loadingDetalles" class="flex-1 flex items-center justify-center">
          <div class="flex flex-col items-center">
            <component :is="LoaderIcon" class="animate-spin h-8 w-8 text-emerald-500 mb-2" />
            <p class="text-sm text-gray-500">Cargando detalles...</p>
          </div>
        </div>

        <!-- Error al cargar detalles -->
        <div v-else-if="errorDetalles" class="flex-1 flex items-center justify-center">
          <div class="bg-red-50 border border-red-200 rounded-md p-4 max-w-md">
            <div class="flex">
              <div class="flex-shrink-0">
                <component :is="AlertIcon" class="h-5 w-5 text-red-400" />
              </div>
              <div class="ml-3">
                <h3 class="text-sm font-medium text-red-800">Error al cargar detalles</h3>
                <div class="mt-2 text-sm text-red-700">
                  <p>{{ errorDetalles }}</p>
                </div>
                <div class="mt-4">
                  <button
                    @click="verDetallesOrden(ordenSeleccionada)"
                    class="inline-flex items-center px-2.5 py-1.5 border border-transparent text-xs font-medium rounded text-red-700 bg-red-100 hover:bg-red-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500"
                  >
                    Reintentar
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Contenido de detalles -->
        <div v-else-if="historialDetalle.length > 0" class="flex-1 flex flex-col min-h-0">
          <!-- Información del cliente y orden -->
          <div class="bg-gray-50 border border-gray-200 rounded-md p-3 mb-3">
            <div class="grid grid-cols-1 sm:grid-cols-2 gap-2">
              <div>
                <h3 class="text-sm font-semibold text-gray-700">Información de la Orden</h3>
                <p class="text-xs text-gray-600">Número: <span class="font-medium">{{ historialDetalle[0].numero_order }}</span></p>
                <p class="text-xs text-gray-600">Fecha: <span class="font-medium">{{ formatearFecha(historialDetalle[0].fecha_expedicion) }}</span></p>
                <p class="text-xs text-gray-600">Condición: <span class="font-medium">{{ historialDetalle[0].describe_condicion }}</span></p>
              </div>
              <div>
                <h3 class="text-sm font-semibold text-gray-700">Información del Cliente</h3>
                <p class="text-xs text-gray-600">Razón Social: <span class="font-medium">{{ historialDetalle[0].razon_social }}</span></p>
                <p class="text-xs text-gray-600">RUC/CI: <span class="font-medium">{{ historialDetalle[0].ruc_ci }}</span></p>
              </div>
            </div>
          </div>

          <!-- Tabla de items -->
          <div class="flex-1 flex flex-col min-h-0 border border-gray-200 rounded-md shadow-sm mb-2">
            <!-- Encabezados fijos -->
            <div class="sheets-header bg-gray-50 border-b border-gray-200 sticky top-0 z-10 text-xs sm:text-sm">
              <div class="sheets-cell w-10 sm:w-16 font-medium">Cod</div>
              <div class="sheets-cell w-12 sm:w-16 font-medium">Cant</div>
              <div class="sheets-cell flex-1 font-medium">Descripción</div>
              <div class="sheets-cell w-16 sm:w-20 font-medium">Precio</div>
              <div class="sheets-cell w-10 sm:w-14 font-medium">IVA</div>
              <div class="sheets-cell w-20 sm:w-24 font-medium">Total</div>
            </div>
            
            <!-- Contenedor con scrollbar -->
            <div class="overflow-y-auto overflow-x-auto flex-1 scrollbar-thin scrollbar-thumb-gray-300 scrollbar-track-gray-100">
              <!-- Filas de productos -->
              <div 
                v-for="(item, index) in historialDetalle" 
                :key="index"
                class="sheets-row text-xs sm:text-sm"
                :class="{'bg-emerald-50': index % 2 === 0}"
              >
                <div class="sheets-cell w-10 sm:w-16">{{ item.cod_item }}</div>
                <div class="sheets-cell w-12 sm:w-16">{{ item.cantidad }}</div>
                <div class="sheets-cell flex-1 truncate" :title="item.descripcion_item">
                  {{ item.descripcion_item }}
                </div>
                <div class="sheets-cell w-16 sm:w-20">
                  {{ formatearMoneda(item.precio_unitario) }}
                </div>
                <div class="sheets-cell w-10 sm:w-14">{{ item.porcentaje_iva }}%</div>
                <div class="sheets-cell w-20 sm:w-24 font-medium">
                  {{ formatearMoneda(item.total_item) }}
                </div>
              </div>
            </div>
            
            <!-- Resumen de totales -->
            <div class="flex justify-end">
            <div class="bg-gray-50 border-t border-gray-200 p-2">
              <div class="grid grid-cols-2 gap-1 text-xs">
                <div class="text-right text-gray-600">Gravado 10%:</div>
                <div class="font-medium">{{ formatearMoneda(calcularTotal('gravado_10')) }}</div>
                
                <div class="text-right text-gray-600">Gravado 5%:</div>
                <div class="font-medium">{{ formatearMoneda(calcularTotal('gravado_5')) }}</div>
                
                <div class="text-right text-gray-600">Exenta:</div>
                <div class="font-medium">{{ formatearMoneda(calcularTotal('exenta')) }}</div>
                
                <div class="text-right text-gray-600">IVA:</div>
                <div class="font-medium">{{ formatearMoneda(calcularTotal('iva_calculado')) }}</div>
                
                <div class="text-right text-gray-700 font-semibold">TOTAL:</div>
                <div class="font-bold text-emerald-700">{{ formatearMoneda(calcularTotal('total_item')) }}</div>
              </div>
            </div>
          </div>
          </div>
        </div>
      </div>
    </div>

    <!-- Navegación de pie de página -->
    <FooterNavigation />
  </section>
</template>

<script setup>
import { ref, computed, h, onMounted } from 'vue';
import historialApi from '@/api/GET_Historial';
import FooterNavigation from '@/components/FooterNavigation.vue';

// Componentes de iconos locales
const SearchIcon = () => h('svg', { 
  xmlns: 'http://www.w3.org/2000/svg', 
  width: '24', 
  height: '24', 
  viewBox: '0 0 24 24', 
  fill: 'none', 
  stroke: 'currentColor', 
  strokeWidth: '2', 
  strokeLinecap: 'round', 
  strokeLinejoin: 'round' 
}, [
  h('circle', { cx: '11', cy: '11', r: '8' }),
  h('path', { d: 'M21 21l-4.3-4.3' })
]);

const LoaderIcon = () => h('svg', { 
  xmlns: 'http://www.w3.org/2000/svg', 
  width: '24', 
  height: '24', 
  viewBox: '0 0 24 24', 
  fill: 'none', 
  stroke: 'currentColor', 
  strokeWidth: '2', 
  strokeLinecap: 'round', 
  strokeLinejoin: 'round' 
}, [
  h('path', { d: 'M21 12a9 9 0 1 1-6.219-8.56' })
]);

const AlertIcon = () => h('svg', { 
  xmlns: 'http://www.w3.org/2000/svg', 
  width: '24', 
  height: '24', 
  viewBox: '0 0 24 24', 
  fill: 'none', 
  stroke: 'currentColor', 
  strokeWidth: '2', 
  strokeLinecap: 'round', 
  strokeLinejoin: 'round' 
}, [
  h('circle', { cx: '12', cy: '12', r: '10' }),
  h('line', { x1: '12', y1: '8', x2: '12', y2: '12' }),
  h('line', { x1: '12', y1: '16', x2: '12.01', y2: '16' })
]);

const EyeIcon = () => h('svg', { 
  xmlns: 'http://www.w3.org/2000/svg', 
  width: '24', 
  height: '24', 
  viewBox: '0 0 24 24', 
  fill: 'none', 
  stroke: 'currentColor', 
  strokeWidth: '2', 
  strokeLinecap: 'round', 
  strokeLinejoin: 'round' 
}, [
  h('path', { d: 'M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z' }),
  h('circle', { cx: '12', cy: '12', r: '3' })
]);

const ArrowLeftIcon = () => h('svg', { 
  xmlns: 'http://www.w3.org/2000/svg', 
  width: '24', 
  height: '24', 
  viewBox: '0 0 24 24', 
  fill: 'none', 
  stroke: 'currentColor', 
  strokeWidth: '2', 
  strokeLinecap: 'round', 
  strokeLinejoin: 'round' 
}, [
  h('line', { x1: '19', y1: '12', x2: '5', y2: '12' }),
  h('polyline', { points: '12 19 5 12 12 5' })
]);

const XIcon = () => h('svg', { 
  xmlns: 'http://www.w3.org/2000/svg', 
  width: '24', 
  height: '24', 
  viewBox: '0 0 24 24', 
  fill: 'none', 
  stroke: 'currentColor', 
  strokeWidth: '2', 
  strokeLinecap: 'round', 
  strokeLinejoin: 'round' 
}, [
  h('line', { x1: '18', y1: '6', x2: '6', y2: '18' }),
  h('line', { x1: '6', y1: '6', x2: '18', y2: '18' })
]);

// Estados reactivos
const busquedaOrden = ref('');
const ordenSeleccionada = ref(null);
const historialDetalle = ref([]);
const ordenes = ref([]);
const loadingOrdenes = ref(true);
const loadingDetalles = ref(false);
const errorOrdenes = ref(null);
const errorDetalles = ref(null);

// Ordenes filtradas basadas en la búsqueda
const ordenesFiltradas = computed(() => {
  if (!busquedaOrden.value) {
    return ordenes.value;
  }
  
  const termino = busquedaOrden.value.toLowerCase();
  return ordenes.value.filter(orden => 
    orden.numero_order.toLowerCase().includes(termino) ||
    orden.razon_social.toLowerCase().includes(termino)
  );
});

// Métodos
const cargarOrdenes = async () => {
  loadingOrdenes.value = true;
  errorOrdenes.value = null;
  
  try {
    const res = await historialApi.getOrdenes();
    if (res.success) {
      ordenes.value = res.data;
    } else {
      errorOrdenes.value = res.error || 'Error al obtener las órdenes';
    }
  } catch (err) {
    console.error('Error cargando órdenes:', err);
    errorOrdenes.value = err.message || 'Error inesperado';
  } finally {
    loadingOrdenes.value = false;
  }
};

const verDetallesOrden = async (idOrder) => {
  ordenSeleccionada.value = idOrder;
  loadingDetalles.value = true;
  errorDetalles.value = null;
  historialDetalle.value = [];
  
  try {
    const res = await historialApi.getHistorialPorOrden(idOrder);
    if (res.success) {
      historialDetalle.value = res.data;
    } else {
      errorDetalles.value = res.error || 'Error al obtener el historial';
    }
  } catch (err) {
    console.error('Error cargando detalles:', err);
    errorDetalles.value = err.message || 'Error inesperado';
  } finally {
    loadingDetalles.value = false;
  }
};

const filtrarOrdenes = () => {
  // La función se ejecuta automáticamente gracias al computed
};

const formatearFecha = (fechaStr) => {
  if (!fechaStr) return '';
  
  try {
    const fecha = new Date(fechaStr);
    return fecha.toLocaleDateString('es-PY', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric'
    });
  } catch (e) {
    return fechaStr;
  }
};

const formatearMoneda = (valor) => {
  return `Gs ${Number(valor).toLocaleString('es-PY')}`;
};

const calcularTotal = (campo) => {
  return historialDetalle.value.reduce((sum, item) => sum + Number(item[campo] || 0), 0);
};

// Cargar la lista de órdenes al montar el componente
onMounted(() => {
  cargarOrdenes();
});
</script>

<style scoped>
.sheets-header, .sheets-row {
  display: flex;
  width: 100%;
  border-bottom: 1px solid #e5e7eb;
}

.sheets-cell {
  padding: 0.5rem 0.25rem;
  display: flex;
  align-items: center;
  min-height: 2rem;
  overflow: hidden;
  border-right: 1px solid #f0f0f0;
}

.sheets-cell:last-child {
  border-right: none;
}

/* Estilos personalizados para scrollbar */
.scrollbar-thin::-webkit-scrollbar {
  width: 6px;
}

.scrollbar-thin::-webkit-scrollbar-track {
  background: #f1f1f1;
  border-radius: 3px;
}

.scrollbar-thin::-webkit-scrollbar-thumb {
  background: #d1d5db;
  border-radius: 3px;
}

.scrollbar-thin::-webkit-scrollbar-thumb:hover {
  background: #9ca3af;
}

/* Ajustes para pantallas pequeñas */
@media (max-width: 640px) {
  .sheets-cell {
    padding: 0.25rem 0.125rem;
    min-height: 1.75rem;
  }
}
</style>
