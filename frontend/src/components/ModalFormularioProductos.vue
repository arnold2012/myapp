<template>
  <div v-if="true" class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center p-4 z-50">
    <div class="bg-white rounded-lg shadow-xl max-w-xl w-full p-6">
      <div class="flex justify-between items-center mb-6">
        <h2 class="text-xl font-semibold text-gray-800">
          {{ producto?.id_item ? 'Editar Producto' : 'Agregar Producto' }}
        </h2>
        <button @click="$emit('cerrar')" class="text-gray-500 hover:text-gray-700 text-xl">&times;</button>
      </div>

      <form @submit.prevent="validarYGuardar">
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-5">
          <div>
            <label for="codigo" class="text-sm font-medium text-gray-700 mb-1 block">Código</label>
            <input
              id="codigo"
              v-model="form.cod_item"
              class="w-full p-2 border rounded-md"
              placeholder="Código del producto"
              :class="{'border-red-500': errores.cod_item}"
              :readonly="!producto?.id_item"
              required
            />
            <p v-if="errores.cod_item" class="text-red-500 text-xs mt-1">{{ errores.cod_item }}</p>
          </div>

          <div>
            <label for="descripcion" class="text-sm font-medium text-gray-700 mb-1 block">Descripción</label>
            <input
              id="descripcion"
              v-model="form.descripcion_item"
              class="w-full p-2 border rounded-md"
              placeholder="Descripción"
              :class="{'border-red-500': errores.descripcion_item}"
              required
            />
            <p v-if="errores.descripcion_item" class="text-red-500 text-xs mt-1">{{ errores.descripcion_item }}</p>
          </div>

          <div>
            <label for="precio" class="text-sm font-medium text-gray-700 mb-1 block">Precio Unitario</label>
            <div class="relative">
              <span class="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500">₲</span>
              <input
                id="precio"
                v-model.number="form.precio_unitario"
                type="number"
                min="0"
                class="w-full p-2 pl-8 border rounded-md"
                :class="{'border-red-500': errores.precio_unitario}"
                required
              />
            </div>
            <p v-if="errores.precio_unitario" class="text-red-500 text-xs mt-1">{{ errores.precio_unitario }}</p>
          </div>

          <div>
            <label for="cantidad" class="text-sm font-medium text-gray-700 mb-1 block">Cantidad Inicial</label>
            <input
              id="cantidad"
              v-model.number="form.cantidad_inicial"
              type="number"
              min="0"
              class="w-full p-2 border rounded-md"
              :class="{'border-red-500': errores.cantidad_inicial}"
              required
            />
            <p v-if="errores.cantidad_inicial" class="text-red-500 text-xs mt-1">{{ errores.cantidad_inicial }}</p>
          </div>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-4 mb-5">
          <div>
            <label for="categoria" class="text-sm font-medium text-gray-700 mb-1 block">Categoría</label>
            <select
              id="categoria"
              v-model.number="form.id_categoria"
              class="w-full p-2 border rounded-md"
              :class="{'border-red-500': errores.id_categoria}"
              required
            >
              <option disabled value="">Seleccione categoría</option>
              <option v-for="c in categorias" :key="c.id_categoria" :value="c.id_categoria">
                {{ c.describe_categoria }}
              </option>
            </select>
            <p v-if="errores.id_categoria" class="text-red-500 text-xs mt-1">{{ errores.id_categoria }}</p>
          </div>

          <div>
            <label for="marca" class="text-sm font-medium text-gray-700 mb-1 block">Marca</label>
            <select
              id="marca"
              v-model.number="form.id_marcas"
              class="w-full p-2 border rounded-md"
              :class="{'border-red-500': errores.id_marcas}"
              required
            >
              <option disabled value="">Seleccione marca</option>
              <option v-for="m in marcas" :key="m.id_marcas" :value="m.id_marcas">
                {{ m.descripcion_marcas }}
              </option>
            </select>
            <p v-if="errores.id_marcas" class="text-red-500 text-xs mt-1">{{ errores.id_marcas }}</p>
          </div>

          <div>
            <label for="impuesto" class="text-sm font-medium text-gray-700 mb-1 block">IVA</label>
            <select
              id="impuesto"
              v-model.number="form.id_impuesto"
              class="w-full p-2 border rounded-md"
              :class="{'border-red-500': errores.id_impuesto}"
              required
            >
              <option disabled value="">Seleccione IVA</option>
              <option v-for="i in impuestos" :key="i.id_impuesto" :value="i.id_impuesto">
                {{ i.porcentaje_iva }}%
              </option>
            </select>
            <p v-if="errores.id_impuesto" class="text-red-500 text-xs mt-1">{{ errores.id_impuesto }}</p>
          </div>
        </div>

        <div class="flex items-center gap-3 mb-5">
          <div class="flex items-center">
            <input
              id="activo"
              type="checkbox"
              v-model="form.item_activo"
              class="h-4 w-4 text-blue-600 border-gray-300 rounded"
            />
            <label for="activo" class="ml-2 text-sm font-medium text-gray-700">Producto Activo</label>
          </div>
        </div>

        <div class="flex justify-end gap-3 mt-6">
          <button
            type="button"
            @click="$emit('cerrar')"
            class="bg-gray-200 text-gray-800 px-4 py-2 rounded-md hover:bg-gray-300"
            :disabled="cargando"
          >
            Cancelar
          </button>
          <button
            type="submit"
            class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700"
            :disabled="cargando"
          >
            <span v-if="cargando">Guardando...</span>
            <span v-else>{{ producto?.id_item ? 'Actualizar' : 'Guardar' }}</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref, watch, onMounted } from 'vue'
import productosApi from '@/api/productos'
import categoriasApi from '@/api/categorias'
import marcasApi from '@/api/marcas'
import impuestosApi from '@/api/impuestos'

const props = defineProps({
  producto: {
    type: Object,
    default: () => null
  }
})
const emit = defineEmits(['cerrar', 'guardado'])

const defaultForm = {
  cod_item: '',
  descripcion_item: '',
  precio_unitario: 0,
  item_activo: true,
  id_categoria: '',
  id_marcas: '',
  id_impuesto: '',
  cantidad_inicial: 0
}

const form = ref({ ...defaultForm })
const categorias = ref([])
const marcas = ref([])
const impuestos = ref([])
const cargando = ref(false)
const errores = ref({})

const cargarReferenciales = async () => {
  try {
    console.log('Cargando datos referenciales...')
    const [cat, marc, imp] = await Promise.all([
      categoriasApi.getAll(),
      marcasApi.getAll(),
      impuestosApi.getAll()
    ])
    categorias.value = cat.data || []
    marcas.value = marc.data || []
    impuestos.value = imp.data || []
    console.log('Datos referenciales cargados:', {
      categorias: categorias.value.length,
      marcas: marcas.value.length,
      impuestos: impuestos.value.length
    })
  } catch (error) {
    console.error('Error al cargar datos referenciales:', error)
  }
}

const obtenerSiguienteCodigo = async () => {
  try {
    cargando.value = true
    const res = await productosApi.getNextCodItem()
    if (res.success) {
      form.value.cod_item = res.data
      console.log('Código generado automáticamente:', form.value.cod_item)
    } else {
      console.error('Error al obtener el siguiente código:', res.error)
      form.value.cod_item = ''
    }
  } catch (error) {
    console.error('Error al obtener el siguiente código:', error)
    form.value.cod_item = ''
  } finally {
    cargando.value = false
  }
}

const validarFormulario = () => {
  const nuevoErrores = {}
  
  if (!form.value.cod_item.trim()) {
    nuevoErrores.cod_item = 'El código es obligatorio'
  }
  
  if (!form.value.descripcion_item.trim()) {
    nuevoErrores.descripcion_item = 'La descripción es obligatoria'
  }
  
  if (form.value.precio_unitario <= 0) {
    nuevoErrores.precio_unitario = 'El precio debe ser mayor a 0'
  }
  
  if (form.value.cantidad_inicial < 0) {
    nuevoErrores.cantidad_inicial = 'La cantidad no puede ser negativa'
  }
  
  if (!form.value.id_categoria) {
    nuevoErrores.id_categoria = 'Seleccione una categoría'
  }
  
  if (!form.value.id_marcas) {
    nuevoErrores.id_marcas = 'Seleccione una marca'
  }
  
  if (!form.value.id_impuesto) {
    nuevoErrores.id_impuesto = 'Seleccione un impuesto'
  }
  
  errores.value = nuevoErrores
  return Object.keys(nuevoErrores).length === 0
}

const validarYGuardar = () => {
  if (validarFormulario()) {
    guardarProducto()
  }
}

const guardarProducto = async () => {
  cargando.value = true
  const productoActual = { ...form.value }
  
  // Asegurarse de que los IDs sean números
  productoActual.id_categoria = Number(productoActual.id_categoria)
  productoActual.id_marcas = Number(productoActual.id_marcas)
  productoActual.id_impuesto = Number(productoActual.id_impuesto)
  
  try {
    console.log('Guardando producto:', productoActual)
    const res = props.producto?.id_item
      ? await productosApi.updateProducto(props.producto.id_item, productoActual)
      : await productosApi.createProducto(productoActual)

    if (res.success) {
      console.log('Producto guardado con éxito')
      emit('guardado')
    } else {
      console.error('Error al guardar producto:', res.error)
      alert('❌ Error: ' + res.error)
    }
  } catch (error) {
    console.error('Error de conexión:', error)
    alert('❌ Error de conexión')
  } finally {
    cargando.value = false
  }
}

const inicializarFormulario = async () => {
  if (props.producto) {
    // Edición: cargar datos del producto existente
    form.value = { 
      ...props.producto,
      id_categoria: Number(props.producto.id_categoria),
      id_marcas: Number(props.producto.id_marcas),
      id_impuesto: Number(props.producto.id_impuesto),
      item_activo: Boolean(props.producto.item_activo)
    }
  } else {
    // Creación: reiniciar formulario y obtener nuevo código
    form.value = { ...defaultForm }
    await obtenerSiguienteCodigo()
  }
  errores.value = {}
}

watch(() => props.producto, async () => {
  console.log('Producto seleccionado cambiado:', props.producto)
  await inicializarFormulario()
}, { immediate: true })

onMounted(async () => {
  console.log('Modal montado')
  await cargarReferenciales()
  if (!props.producto) {
    await obtenerSiguienteCodigo()
  }
})
</script>
