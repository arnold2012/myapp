<template>
  <section class="flex flex-col h-screen max-h-screen overflow-hidden p-2 sm:p-3">
    <!-- Cabecera responsive -->
    <div class="flex flex-col sm:flex-row sm:justify-between sm:items-center mb-2 gap-2">
      <h1 class="text-xl font-bold text-gray-900">Órdenes</h1>
      <div class="flex items-center space-x-2">
        <button 
          @click="ordenarPor('cod_item')" 
          class="text-xs px-2 py-1 rounded bg-gray-100 hover:bg-gray-200 flex items-center"
        >
          <component :is="ArrowUpDownIcon" class="h-3 w-3 mr-1" /> Ordenar
        </button>
        <button 
          v-if="itemsSeleccionados.length > 0"
          @click="limpiarTodo" 
          class="text-xs px-2 py-1 rounded bg-red-100 hover:bg-red-200 text-red-700 flex items-center"
        >
          <component :is="TrashIcon" class="h-3 w-3 mr-1" /> Limpiar
        </button>
      </div>
    </div>

    <!-- Buscador compacto -->
    <div class="mb-2">
      <div class="relative">
        <input
          id="search-input"
          v-model="busqueda"
          @input="buscarProductos"
          placeholder="Buscar productos..."
          class="w-full px-3 py-2 pl-8 text-sm border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-1 focus:ring-emerald-500 focus:border-emerald-500"
          type="text"
        />
        <div class="absolute inset-y-0 left-0 pl-2 flex items-center pointer-events-none">
          <component :is="SearchIcon" class="h-4 w-4 text-gray-400" />
        </div>
      </div>
      <div class="relative">
        <p v-if="loading" class="mt-1 text-xs text-gray-500 flex items-center">
          <component :is="LoaderIcon" class="animate-spin h-3 w-3 mr-1" /> Buscando...
        </p>
        <ul
          v-if="!loading && productos.length"
          class="mt-1 border border-gray-300 rounded-md bg-white shadow-sm divide-y divide-gray-200 absolute z-20 w-full max-h-40 overflow-y-auto"
        >
          <li
            v-for="producto in productos"
            :key="producto.id_item"
            class="px-3 py-2 hover:bg-emerald-50 cursor-pointer text-sm"
            @click="agregarProducto(producto)"
          >
            <div class="font-medium">
              {{ producto.cod_item }} - {{ producto.descripcion_item }}
            </div>
            <div class="text-xs text-gray-500 flex justify-between mt-1">
              <span>Gs {{ producto.precio_unitario.toLocaleString('es-PY') }}</span>
              <span>Stock: {{ producto.cantidad_inicial }}</span>
            </div>
          </li>
        </ul>
        <p v-else-if="!loading && busqueda.trim().length >= 2" class="mt-1 text-xs text-gray-500">
          Producto o servicio no encontrado.
        </p>
      </div>
    </div>

    <!-- Tabla con altura ajustada -->
    <div class="flex-1 flex flex-col min-h-0 border border-gray-200 rounded-md shadow-sm mb-2" style="max-height: 60vh;">
      <!-- Encabezados fijos -->
      <div class="sheets-header bg-gray-50 border-b border-gray-200 sticky top-0 z-10 text-sm">
        <div class="sheets-cell w-10 sm:w-16 font-medium">Cod</div>
        <div class="sheets-cell w-16 sm:w-20 font-medium">Cant</div>
        <div class="sheets-cell flex-1 font-medium">Desc</div>
        <div class="sheets-cell w-14 sm:w-20 font-medium">Precio</div>
        <div class="sheets-cell w-10 sm:w-16 font-medium">IVA</div>
        <div class="sheets-cell w-20 sm:w-28 font-medium">Total</div>
        <div class="sheets-cell w-10 sm:w-16 font-medium">Acc</div>
      </div>
      
      <!-- Contenedor con scrollbar que ocupa el espacio disponible -->
      <div class="overflow-y-auto overflow-x-auto flex-1 scrollbar-thin scrollbar-thumb-gray-300 scrollbar-track-gray-100">
        <!-- Mensaje cuando no hay productos -->
        <div v-if="!itemsSeleccionados.length" class="sheets-row">
          <div class="col-span-7 py-4 text-center text-gray-500 text-sm">
            No hay productos seleccionados. Busque y agregue productos arriba.
          </div>
        </div>

        <!-- Filas de productos -->
        <div 
          v-for="(item, index) in itemsSeleccionados" 
          :key="index"
          class="sheets-row text-xs sm:text-sm"
          :class="{'bg-emerald-50': index === filaActiva}"
          @mouseenter="filaActiva = index"
          @mouseleave="filaActiva = null"
        >
          <div class="sheets-cell w-10 sm:w-16">{{ item.cod_item }}</div>
          <div class="sheets-cell w-16 sm:w-20">
            <input
              type="number"
              v-model.number="item.cantidad"
              min="1"
              class="w-full px-1 py-0.5 text-xs sm:text-sm border border-transparent rounded focus:border-emerald-500 focus:ring-1 focus:ring-emerald-500"
              @input="actualizarTotal(index)"
              :style="{ width: calcularAnchoCantidad(item.cantidad) }"
            />
          </div>
          <div class="sheets-cell flex-1 truncate" :title="item.descripcion_item">
            {{ item.descripcion_item }}
          </div>
          <div class="sheets-cell w-14 sm:w-20 editable" @click="activarEdicion(index, 'precio_unitario')">
            <span v-if="editando.index !== index || editando.campo !== 'precio_unitario'">
              {{ item.precio_unitario.toLocaleString('es-PY') }}
            </span>
            <input
              v-else
              type="number"
              v-model.number="item.precio_unitario"
              class="w-full px-1 py-0.5 text-xs sm:text-sm border border-emerald-500 rounded focus:ring-1 focus:ring-emerald-500"
              @blur="finalizarEdicion(); actualizarTotal(index)"
              @keyup.enter="finalizarEdicion(); actualizarTotal(index)"
              ref="editInput"
            />
          </div>
          <div class="sheets-cell w-10 sm:w-16">{{ item.porcentaje_iva }}%</div>
          <div class="sheets-cell w-20 sm:w-28 font-medium">{{ item.total.toLocaleString('es-PY') }}</div>
          <div class="sheets-cell w-10 sm:w-16">
            <button @click="eliminarProducto(index)" class="text-red-600 hover:text-red-800 p-1">
              <component :is="TrashIcon" class="h-3 w-3 sm:h-3.5 sm:w-3.5" />
            </button>
          </div>
        </div>
      </div>
      
      <!-- Fila de totales fija en la parte inferior -->
      <div v-if="itemsSeleccionados.length > 0" class="sheets-footer bg-gray-50 border-t border-gray-200 sticky bottom-0 z-10 text-xs sm:text-sm">
        <div class="sheets-cell w-10 sm:w-16"></div>
        <div class="sheets-cell w-16 sm:w-20"></div>
        <div class="sheets-cell flex-1"></div>
        <div class="sheets-cell w-14 sm:w-20"></div>
        <div class="sheets-cell w-10 sm:w-16 font-medium text-gray-700 justify-end">TOTAL:</div>
        <div class="sheets-cell w-20 sm:w-28 font-bold text-emerald-700">
          Gs {{ totalGeneral.toLocaleString('es-PY') }}
        </div>
        <div class="sheets-cell w-10 sm:w-16"></div>
      </div>
    </div>

    <!-- Botón de confirmar -->
    <div class="flex justify-end mt-2">
      <button
        @click="abrirModalProcesar"
        :disabled="!itemsSeleccionados.length"
        class="bg-emerald-600 hover:bg-emerald-700 disabled:bg-gray-400 text-white font-semibold py-2 px-4 text-sm rounded-md flex items-center"
      >
        <component :is="CheckIcon" class="h-4 w-4 mr-1" />
        Confirmar Orden
      </button>
    </div>

    <!-- Modal para procesar orden -->
    <ModelProcesarOrden
      v-if="showModal"
      :items="itemsSeleccionados"
      :total="totalGeneral"
      :numeroOrden="numeroOrden"
      @close="cerrarModalProcesar"
      @ordenProcesada="limpiarTodo"
    />
  </section>
</template>

<script setup>
import { ref, computed, watch, onMounted, nextTick, h, onUnmounted } from 'vue';
import api from '@/api/productos';
import ModelProcesarOrden from '@/components/ModelProcesarOrden.vue';

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

const TrashIcon = () => h('svg', { 
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
  h('path', { d: 'M3 6h18' }),
  h('path', { d: 'M19 6v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6' }),
  h('path', { d: 'M8 6V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2' })
]);

const CheckIcon = () => h('svg', { 
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
  h('path', { d: 'M20 6L9 17l-5-5' })
]);

const ArrowUpDownIcon = () => h('svg', { 
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
  h('path', { d: 'M7 3v18' }),
  h('path', { d: 'M11 7h10' }),
  h('path', { d: 'M11 17h10' }),
  h('path', { d: 'M3 15l4 4 4-4' }),
  h('path', { d: 'M3 9l4-4 4 4' })
]);

// Estados reactivos
const busqueda = ref('');
const productos = ref([]);
const loading = ref(false);
const itemsSeleccionados = ref([]);
const showModal = ref(false);
const numeroOrden = ref(Math.floor(Math.random() * 10000));
const filaActiva = ref(null);
const ordenActual = ref({ campo: null, direccion: 'asc' });
const editando = ref({ index: null, campo: null });
const editInput = ref(null);
const tablaContainer = ref(null);

// Control de vista para móviles
const anchoPantalla = ref(window.innerWidth);
const esPantallaMovil = computed(() => anchoPantalla.value < 640); // 640px es el breakpoint 'sm' en Tailwind

// Función para calcular el ancho del input de cantidad
const calcularAnchoCantidad = (cantidad) => {
  const cantidadStr = cantidad.toString();
  // Calcular un ancho basado en la cantidad de dígitos (mínimo 2 dígitos)
  const digitos = Math.max(cantidadStr.length, 2);
  return `${digitos * 0.6 + 1.2}rem`; // Ajustar según necesidad
};

// Actualizar ancho de pantalla al redimensionar
const actualizarAnchoPantalla = () => {
  anchoPantalla.value = window.innerWidth;
};

// Persistencia en sessionStorage
const cargarDesdeSessionStorage = () => {
  const itemsGuardados = sessionStorage.getItem('ordenItems');
  const numeroGuardado = sessionStorage.getItem('ordenNumero');
  
  if (itemsGuardados) {
    try {
      itemsSeleccionados.value = JSON.parse(itemsGuardados);
    } catch (e) {
      sessionStorage.removeItem('ordenItems');
    }
  }
  
  if (numeroGuardado) {
    numeroOrden.value = Number(numeroGuardado);
  }
};

const guardarEnSessionStorage = () => {
  sessionStorage.setItem('ordenItems', JSON.stringify(itemsSeleccionados.value));
  sessionStorage.setItem('ordenNumero', numeroOrden.value.toString());
};

// Watchers para persistencia automática
watch(itemsSeleccionados, guardarEnSessionStorage, { deep: true });
watch(numeroOrden, guardarEnSessionStorage);

// Cargar datos y configurar eventos al montar el componente
onMounted(() => {
  cargarDesdeSessionStorage();
  window.addEventListener('resize', actualizarAnchoPantalla);
});

// Limpiar eventos al desmontar
onUnmounted(() => {
  window.removeEventListener('resize', actualizarAnchoPantalla);
});

// Métodos
const abrirModalProcesar = () => showModal.value = true;
const cerrarModalProcesar = () => showModal.value = false;

const limpiarTodo = () => {
  sessionStorage.removeItem('ordenItems');
  sessionStorage.removeItem('ordenNumero');
  itemsSeleccionados.value = [];
  productos.value = [];
  busqueda.value = '';
  numeroOrden.value = Math.floor(Math.random() * 10000);
  cerrarModalProcesar();
};

const buscarProductos = async () => {
  const termino = busqueda.value.trim();
  if (termino.length < 2) {
    productos.value = [];
    loading.value = false;
    return;
  }
  loading.value = true;
  try {
    const res = await api.buscarProductos(termino);
    productos.value = res.success ? res.data : [];
  } catch (err) {
    productos.value = [];
  } finally {
    loading.value = false;
  }
};

const agregarProducto = (producto) => {
  // Siempre agregar como una nueva fila, incluso si el producto ya existe
  const nuevoItem = {
    id_item: producto.id_item,
    cod_item: producto.cod_item,
    descripcion_item: producto.descripcion_item,
    precio_unitario: producto.precio_unitario,
    cantidad: 1,
    cantidad_inicial: producto.cantidad_inicial,
    porcentaje_iva: producto.porcentaje_iva,
    total: producto.precio_unitario * 1,
  };
  
  itemsSeleccionados.value.push(nuevoItem);
  filaActiva.value = itemsSeleccionados.value.length - 1;
  
  // Desplazar al final para ver el nuevo elemento
  nextTick(() => {
    if (tablaContainer.value) {
      tablaContainer.value.scrollTop = tablaContainer.value.scrollHeight;
    }
  });
  
  busqueda.value = '';
  productos.value = [];
};

const actualizarTotal = (index) => {
  const item = itemsSeleccionados.value[index];
  // Eliminar la restricción de cantidad máxima
  if (item.cantidad < 1) {
    item.cantidad = 1;
  }
  item.total = item.precio_unitario * item.cantidad;
};

const eliminarProducto = (index) => {
  itemsSeleccionados.value.splice(index, 1);
};

const ordenarPor = (campo) => {
  // Si es el mismo campo, cambiar dirección
  if (ordenActual.value.campo === campo) {
    ordenActual.value.direccion = ordenActual.value.direccion === 'asc' ? 'desc' : 'asc';
  } else {
    ordenActual.value = { campo, direccion: 'asc' };
  }
  
  // Ordenar los items
  itemsSeleccionados.value.sort((a, b) => {
    let valorA = a[campo];
    let valorB = b[campo];
    
    // Si son strings, ignorar mayúsculas/minúsculas
    if (typeof valorA === 'string') {
      valorA = valorA.toLowerCase();
      valorB = valorB.toLowerCase();
    }
    
    if (ordenActual.value.direccion === 'asc') {
      return valorA > valorB ? 1 : -1;
    } else {
      return valorA < valorB ? 1 : -1;
    }
  });
};

const activarEdicion = (index, campo) => {
  editando.value = { index, campo };
  nextTick(() => {
    if (editInput.value) {
      editInput.value.focus();
      editInput.value.select();
    }
  });
};

const finalizarEdicion = () => {
  editando.value = { index: null, campo: null };
};

// Computed
const totalGeneral = computed(() => 
  itemsSeleccionados.value.reduce((sum, item) => sum + item.total, 0)
);
</script>

<style scoped>
.sheets-header, .sheets-row, .sheets-footer {
  display: flex;
  width: 100%;
  border-bottom: 1px solid #e5e7eb;
}

.sheets-cell {
  padding: 0.5rem 0.25rem;
  display: flex;
  align-items: center;
  min-height: 2.25rem;
  overflow: hidden;
  border-right: 1px solid #f0f0f0;
}

.sheets-cell:last-child {
  border-right: none;
}

.editable {
  cursor: pointer;
}

.editable:hover {
  background-color: #f0f9f6;
}

.sheets-row:hover {
  background-color: #f0f9f6;
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

/* Estilos para inputs de cantidad */
input[type="number"] {
  -moz-appearance: textfield; /* Firefox */
}

input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}
</style>

<!-- <script setup>
import { ref, computed } from 'vue';
import api from '@/api/productos';
import ModelProcesarOrden from '@/components/ModelProcesarOrden.vue'

const busqueda = ref('');
const productos = ref([]);
const loading = ref(false);
const itemsSeleccionados = ref([]);
const showModal = ref(false);
const numeroOrden = ref(Math.floor(Math.random() * 10000))

const abrirModalProcesar = () => showModal.value = true
const cerrarModalProcesar = () => showModal.value = false

const limpiarTodo = () => {
  itemsSeleccionados.value = []
  productos.value = []
  busqueda.value = ''
  numeroOrden.value = Math.floor(Math.random() * 10000)
  cerrarModalProcesar()
}

const buscarProductos = async () => {
  const termino = busqueda.value.trim();
  if (termino.length < 2) {
    productos.value = [];
    loading.value = false;
    return;
  }
  loading.value = true;
  try {
    const res = await api.buscarProductos(termino);
    if (res.success) productos.value = res.data;
    else productos.value = [];
  } catch (err) {
    productos.value = [];
  } finally {
    loading.value = false;
  }
};

const agregarProducto = (producto) => {
  const total = producto.precio_unitario * 1;
  itemsSeleccionados.value.push({
    id_item: producto.id_item,
    cod_item: producto.cod_item,
    descripcion_item: producto.descripcion_item,
    precio_unitario: producto.precio_unitario,
    cantidad: 1,
    cantidad_inicial: producto.cantidad_inicial,
    porcentaje_iva: producto.porcentaje_iva,
    total: total,
  });
  busqueda.value = '';
  productos.value = [];
};

const actualizarTotal = (index) => {
  const item = itemsSeleccionados.value[index];
  if (item.cantidad > item.cantidad_inicial) item.cantidad = item.cantidad_inicial;
  item.total = item.precio_unitario * item.cantidad;
};

const eliminarProducto = (index) => {
  itemsSeleccionados.value.splice(index, 1);
};

const totalGeneral = computed(() => {
  return itemsSeleccionados.value.reduce((sum, item) => sum + item.total, 0);
});
</script> -->


<!-- <script setup>
import { ref, computed } from 'vue';
import api from '@/api/productos';
import ModelProcesarOrden from '@/components/ModelProcesarOrden.vue'

const busqueda = ref('');
const productos = ref([]);
const loading = ref(false);
const itemsSeleccionados = ref([]);
const showModal = ref(false);
const numeroOrden = ref(Math.floor(Math.random() * 10000))

const abrirModalProcesar = () => showModal.value = true
const cerrarModalProcesar = () => showModal.value = false

const limpiarTodo = () => {
  itemsSeleccionados.value = []
  productos.value = []
  busqueda.value = ''
  numeroOrden.value = Math.floor(Math.random() * 10000)
  cerrarModalProcesar()
}

const buscarProductos = async () => {
  const termino = busqueda.value.trim();
  if (termino.length < 2) {
    productos.value = [];
    loading.value = false;
    return;
  }
  loading.value = true;
  try {
    const res = await api.buscarProductos(termino);
    if (res.success) productos.value = res.data;
    else productos.value = [];
  } catch (err) {
    productos.value = [];
  } finally {
    loading.value = false;
  }
};

const agregarProducto = (producto) => {
  const total = producto.precio_unitario * 1;
  itemsSeleccionados.value.push({
    id_item: producto.id_item,
    cod_item: producto.cod_item,
    descripcion_item: producto.descripcion_item,
    precio_unitario: producto.precio_unitario,
    cantidad: 1,
    cantidad_inicial: producto.cantidad_inicial,
    porcentaje_iva: producto.porcentaje_iva,
    total: total,
  });
  busqueda.value = '';
  productos.value = [];
};

const actualizarTotal = (index) => {
  const item = itemsSeleccionados.value[index];
  if (item.cantidad > item.cantidad_inicial) item.cantidad = item.cantidad_inicial;
  item.total = item.precio_unitario * item.cantidad;
};

const eliminarProducto = (index) => {
  itemsSeleccionados.value.splice(index, 1);
};

const totalGeneral = computed(() => {
  return itemsSeleccionados.value.reduce((sum, item) => sum + item.total, 0);
});
</script> -->
