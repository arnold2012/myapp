<template>
  <nav class="navbar">
    <!-- Enlaces de navegación normales (excluyendo los del menú desplegable) -->
    <router-link 
      v-for="route in regularRoutes" 
      :key="route.path" 
      :to="route.path"
      class="nav-link"
      :class="{ 'active': $route.path === route.path }"
    >
      <span class="icon">{{ route.meta?.icon }}</span>
      {{ route.name }}
    </router-link>
    
    <!-- Menú desplegable para Productos y Referenciales -->
    <div class="menu-container">
      <button 
        class="menu-button nav-link" 
        @click="toggleMenu"
        :class="{ 'active': isProductMenuActive }"
      >
        <span class="icon">📋</span>
        Productos y Referenciales
        <i class="chevron" :class="{ 'rotated': menuOpen }">▼</i>
      </button>
      <div class="dropdown-content" :class="{ 'show': menuOpen }">
        <router-link 
          to="/productos"
          class="dropdown-link"
          :class="{ 'active': $route.path === '/productos' }"
          @click="menuOpen = false"
        >
          <span class="icon">📋</span>
          Productos
        </router-link>
        <router-link 
          to="/impuestos"
          class="dropdown-link"
          :class="{ 'active': $route.path === '/impuestos' }"
          @click="menuOpen = false"
        >
          <span class="icon">💰</span>
          Impuestos
        </router-link>
        <router-link 
          to="/marcas"
          class="dropdown-link"
          :class="{ 'active': $route.path === '/marcas' }"
          @click="menuOpen = false"
        >
          <span class="icon">🏷️</span>
          Marcas
        </router-link>
        <router-link 
          to="/categoria"
          class="dropdown-link"
          :class="{ 'active': $route.path === '/categoria' }"
          @click="menuOpen = false"
        >
          <span class="icon">📂</span>
          Categoría
        </router-link>
      </div>
    </div>
  </nav>
</template>

<script>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

export default {
  setup() {
    const router = useRouter()
    const route = useRoute()
    const menuOpen = ref(false)
    
    // Rutas que se mostrarán como enlaces normales (excluyendo las del menú desplegable)
    const regularRoutes = computed(() => {
      return router.options.routes.filter(route => 
        route.name && 
        !['Productos', 'Impuestos', 'Marcas', 'Categoría'].includes(route.name)
      )
    })
    
    // Verificar si alguna ruta del menú desplegable está activa
    const isProductMenuActive = computed(() => {
      const productPaths = ['/productos', '/impuestos', '/marcas', '/categoria']
      return productPaths.includes(route.path)
    })
    
    // Función para alternar el menú
    const toggleMenu = () => {
      menuOpen.value = !menuOpen.value
    }
    
    // Cerrar el menú cuando se hace clic fuera
    const closeMenuOnOutsideClick = (event) => {
      const menuContainer = document.querySelector('.menu-container')
      if (menuOpen.value && menuContainer && !menuContainer.contains(event.target)) {
        menuOpen.value = false
      }
    }
    
    // Agregar/eliminar event listeners
    onMounted(() => {
      document.addEventListener('click', closeMenuOnOutsideClick)
    })
    
    onUnmounted(() => {
      document.removeEventListener('click', closeMenuOnOutsideClick)
    })
    
    return {
      regularRoutes,
      isProductMenuActive,
      menuOpen,
      toggleMenu
    }
  }
}
</script>

<style scoped>
.navbar {
  display: flex;
  background: #2c3e50;
  padding: 1rem;
  box-shadow: 0 2px 4px rgba(0,0,0,0.1);
}

.nav-link {
  color: white;
  text-decoration: none;
  padding: 0.5rem 1rem;
  margin: 0 0.5rem;
  border-radius: 4px;
  display: flex;
  align-items: center;
  transition: all 0.3s ease;
}

.nav-link:hover {
  background: #34495e;
}

.nav-link.active {
  background: #42b983;
  font-weight: bold;
}

.icon {
  margin-right: 8px;
  font-size: 1.1em;
}

/* Estilos para el menú desplegable */
.menu-container {
  position: relative;
  display: inline-block;
}

.menu-button {
  background: none;
  border: none;
  cursor: pointer;
  font-size: 1em;
  font-family: inherit;
  display: flex;
  align-items: center;
}

.chevron {
  margin-left: 8px;
  font-size: 0.7em;
  transition: transform 0.3s ease;
}

.chevron.rotated {
  transform: rotate(180deg);
}

.dropdown-content {
  display: none;
  position: absolute;
  top: 100%;
  left: 0;
  background-color: #2c3e50;
  min-width: 200px;
  box-shadow: 0 8px 16px rgba(0,0,0,0.2);
  z-index: 1;
  border-radius: 4px;
  margin-top: 5px;
}

.dropdown-content.show {
  display: block;
}

.dropdown-link {
  color: white;
  padding: 12px 16px;
  text-decoration: none;
  display: flex;
  align-items: center;
  transition: background-color 0.3s;
}

.dropdown-link:hover {
  background-color: #34495e;
}

.dropdown-link.active {
  background-color: #42b983;
  font-weight: bold;
}
</style>