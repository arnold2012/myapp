<template>
  <nav class="navbar">
    <div class="navbar-container">
      <!-- Botón hamburguesa para móvil -->
      <button 
        class="mobile-toggle"
        @click="toggleMobileMenu"
        :class="{ 'active': mobileMenuOpen }"
      >
        <span></span>
        <span></span>
        <span></span>
      </button>

      <!-- Menú principal -->
      <div class="navbar-menu" :class="{ 'mobile-open': mobileMenuOpen }">
        <!-- Enlaces de navegación normales -->
        <router-link 
          v-for="route in regularRoutes" 
          :key="route.path" 
          :to="route.path"
          class="nav-link"
          :class="{ 'active': $route.path === route.path }"
          @click="closeMobileMenu"
        >
          <span class="nav-icon">{{ route.meta?.icon }}</span>
          <span class="nav-text">{{ route.name }}</span>
        </router-link>
        
        <!-- Menú desplegable de Productos -->
        <div class="dropdown" :class="{ 'active': menuOpen || isProductMenuActive }">
          <button 
            class="dropdown-toggle nav-link" 
            @click="toggleMenu"
            @mouseenter="handleMouseEnter"
            @mouseleave="handleMouseLeave"
            :class="{ 'active': isProductMenuActive }"
          >
            <span class="nav-icon">📋</span>
            <span class="nav-text">Productos y Referenciales</span>
            <i class="chevron" :class="{ 'rotated': menuOpen }">
              <svg width="12" height="12" viewBox="0 0 12 12" fill="currentColor">
                <path d="M2 4l4 4 4-4" stroke="currentColor" stroke-width="1.5" fill="none" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </i>
          </button>
          
          <div 
            class="dropdown-menu" 
            :class="{ 'show': menuOpen }"
            @mouseenter="handleMouseEnter"
            @mouseleave="handleMouseLeave"
          >
            <div class="dropdown-content">
              <router-link 
                v-for="route in productRoutes"
                :key="route.path"
                :to="route.path"
                class="dropdown-item"
                :class="{ 'active': $route.path === route.path }"
                @click="closeMenu"
              >
                <span class="dropdown-icon">{{ route.meta?.icon }}</span>
                <div class="dropdown-text">
                  <span class="dropdown-title">{{ route.name }}</span>
                  <span class="dropdown-description">{{ getRouteDescription(route.name) }}</span>
                </div>
              </router-link>
            </div>
          </div>
        </div>
            
        <!-- Menú desplegable para Ajuste General -->
        <div class="dropdown" :class="{ 'active': ajustesMenuOpen || isAjustesMenuActive }">
          <button 
            class="dropdown-toggle nav-link" 
            @click="toggleAjustesMenu"
            @mouseenter="handleAjustesMouseEnter"
            @mouseleave="handleAjustesMouseLeave"
            :class="{ 'active': isAjustesMenuActive }"
          >
            <span class="nav-icon">⚙️</span>
            <span class="nav-text">Ajuste General</span>
            <i class="chevron" :class="{ 'rotated': ajustesMenuOpen }">
              <svg width="12" height="12" viewBox="0 0 12 12" fill="currentColor">
                <path d="M2 4l4 4 4-4" stroke="currentColor" stroke-width="1.5" fill="none" stroke-linecap="round" stroke-linejoin="round"/>
              </svg>
            </i>
          </button>
          
          <div 
            class="dropdown-menu" 
            :class="{ 'show': ajustesMenuOpen }"
            @mouseenter="handleAjustesMouseEnter"
            @mouseleave="handleAjustesMouseLeave"
          >
            <div class="dropdown-content">
              <router-link 
                v-for="route in ajustesRoutes"
                :key="route.path"
                :to="route.path"
                class="dropdown-item"
                :class="{ 'active': $route.path === route.path }"
                @click="closeAjustesMenu"
              >
                <span class="dropdown-icon">{{ route.meta?.icon }}</span>
                <div class="dropdown-text">
                  <span class="dropdown-title">{{ route.name }}</span>
                  <span class="dropdown-description">{{ getRouteDescription(route.name) }}</span>
                </div>
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'

export default {
  name: 'Navbar',
  setup() {
    const router = useRouter()
    const route = useRoute()
    const menuOpen = ref(false)
    const ajustesMenuOpen = ref(false)
    const mobileMenuOpen = ref(false)
    const hoverTimeout = ref(null)
    
    // Obtener todas las rutas
    const allRoutes = computed(() => {
      return router.getRoutes().filter(route => 
        route.name && 
        route.path !== '/:pathMatch(.*)*' &&
        route.meta
      )
    })
    
    // Rutas que se mostrarán como enlaces normales (sin dropdown)
    const regularRoutes = computed(() => {
      return allRoutes.value.filter(route => 
        !route.meta?.dropdown
      )
    })
    
    // Rutas del menú de productos
    const productRoutes = computed(() => {
      return allRoutes.value.filter(route => 
        route.meta?.dropdown === 'productos'
      )
    })
    
    // Rutas del menú de ajustes
    const ajustesRoutes = computed(() => {
      return allRoutes.value.filter(route => 
        route.meta?.dropdown === 'ajustes'
      )
    })
    
    // Verificar si alguna ruta del menú desplegable de Productos está activa
    const isProductMenuActive = computed(() => {
      return productRoutes.value.some(route => route.path === route.path)
    })
    
    // Verificar si alguna ruta del menú desplegable de Ajustes está activa
    const isAjustesMenuActive = computed(() => {
      return ajustesRoutes.value.some(route => route.path === route.path)
    })
    
    // Función para obtener descripciones de rutas
    const getRouteDescription = (routeName) => {
      const descriptions = {
        'Productos': 'Gestiona tu inventario',
        'Impuestos': 'Configuración fiscal',
        'Marcas': 'Administra las marcas',
        'Categoría': 'Organiza por categorías',
        'Timbrado': 'Ajustes del sistema'
      }
      return descriptions[routeName] || 'Configuración'
    }
    
    // Funciones del menú desplegable de Productos
    const toggleMenu = () => {
      menuOpen.value = !menuOpen.value
      if (menuOpen.value) ajustesMenuOpen.value = false
    }
    
    const closeMenu = () => {
      menuOpen.value = false
      mobileMenuOpen.value = false
    }
    
    const handleMouseEnter = () => {
      if (window.innerWidth > 768) {
        clearTimeout(hoverTimeout.value)
        menuOpen.value = true
        ajustesMenuOpen.value = false
      }
    }
    
    const handleMouseLeave = () => {
      if (window.innerWidth > 768) {
        hoverTimeout.value = setTimeout(() => {
          menuOpen.value = false
        }, 150)
      }
    }
    
    // Funciones del menú desplegable de Ajustes
    const toggleAjustesMenu = () => {
      ajustesMenuOpen.value = !ajustesMenuOpen.value
      if (ajustesMenuOpen.value) menuOpen.value = false
    }
    
    const closeAjustesMenu = () => {
      ajustesMenuOpen.value = false
      mobileMenuOpen.value = false
    }
    
    const handleAjustesMouseEnter = () => {
      if (window.innerWidth > 768) {
        clearTimeout(hoverTimeout.value)
        ajustesMenuOpen.value = true
        menuOpen.value = false
      }
    }
    
    const handleAjustesMouseLeave = () => {
      if (window.innerWidth > 768) {
        hoverTimeout.value = setTimeout(() => {
          ajustesMenuOpen.value = false
        }, 150)
      }
    }
    
    // Funciones del menú móvil
    const toggleMobileMenu = () => {
      mobileMenuOpen.value = !mobileMenuOpen.value
    }
    
    const closeMobileMenu = () => {
      mobileMenuOpen.value = false
    }
    
    // Cerrar menús al hacer clic fuera
    const handleOutsideClick = (event) => {
      const navbar = document.querySelector('.navbar')
      if (navbar && !navbar.contains(event.target)) {
        menuOpen.value = false
        ajustesMenuOpen.value = false
        mobileMenuOpen.value = false
      }
    }
    
    // Cerrar menú móvil al cambiar el tamaño de ventana
    const handleResize = () => {
      if (window.innerWidth > 768) {
        mobileMenuOpen.value = false
        menuOpen.value = false
        ajustesMenuOpen.value = false
      }
    }
    
    onMounted(() => {
      document.addEventListener('click', handleOutsideClick)
      window.addEventListener('resize', handleResize)
    })
    
    onUnmounted(() => {
      document.removeEventListener('click', handleOutsideClick)
      window.removeEventListener('resize', handleResize)
      clearTimeout(hoverTimeout.value)
    })
    
    return {
      regularRoutes,
      productRoutes,
      ajustesRoutes,
      isProductMenuActive,
      isAjustesMenuActive,
      menuOpen,
      ajustesMenuOpen,
      mobileMenuOpen,
      toggleMenu,
      closeMenu,
      toggleAjustesMenu,
      closeAjustesMenu,
      toggleMobileMenu,
      closeMobileMenu,
      handleMouseEnter,
      handleMouseLeave,
      handleAjustesMouseEnter,
      handleAjustesMouseLeave,
      getRouteDescription
    }
  }
}
</script>

<style scoped>
.navbar {
  background: linear-gradient(135deg, #2c3e50 0%, #2c3e50 100%);
  box-shadow: 0 2px 20px rgba(0, 0, 0, 0.1);
  position: sticky;
  top: 0;
  z-index: 1000;
  backdrop-filter: blur(10px);
}

.navbar-container {
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 1rem;
  height: 64px;
}

.mobile-toggle {
  display: none;
  flex-direction: column;
  background: none;
  border: none;
  cursor: pointer;
  padding: 0.5rem;
  gap: 4px;
}

.mobile-toggle span {
  width: 25px;
  height: 3px;
  background: white;
  transition: all 0.3s ease;
  border-radius: 2px;
}

.mobile-toggle.active span:nth-child(1) {
  transform: rotate(45deg) translate(6px, 6px);
}

.mobile-toggle.active span:nth-child(2) {
  opacity: 0;
}

.mobile-toggle.active span:nth-child(3) {
  transform: rotate(-45deg) translate(6px, -6px);
}

.navbar-menu {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.nav-link {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.75rem 1rem;
  text-decoration: none;
  color: rgba(255, 255, 255, 0.9);
  border-radius: 8px;
  transition: all 0.3s ease;
  font-weight: 500;
  position: relative;
  background: none;
  border: none;
  cursor: pointer;
  font-size: 0.95rem;
}

.nav-link:hover {
  background: rgba(255, 255, 255, 0.15);
  color: white;
  transform: translateY(-1px);
}

.nav-link.active {
  background: rgba(255, 255, 255, 0.2);
  color: white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

.nav-icon {
  font-size: 1.1rem;
}

.dropdown {
  position: relative;
}

.dropdown-toggle {
  background: none !important;
}

.chevron {
  margin-left: 0.5rem;
  transition: transform 0.3s ease;
  display: flex;
  align-items: center;
}

.chevron.rotated {
  transform: rotate(180deg);
}

.dropdown-menu {
  position: absolute;
  top: 100%;
  left: 0;
  background: white;
  border-radius: 12px;
  box-shadow: 0 10px 40px rgba(0, 0, 0, 0.15);
  opacity: 0;
  visibility: hidden;
  transform: translateY(-10px);
  transition: all 0.3s ease;
  min-width: 280px;
  z-index: 1000;
  border: 1px solid rgba(0, 0, 0, 0.05);
}

.dropdown-menu.show {
  opacity: 1;
  visibility: visible;
  transform: translateY(0);
}

.dropdown-content {
  padding: 0.5rem;
}

.dropdown-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.75rem;
  text-decoration: none;
  color: #374151;
  border-radius: 8px;
  transition: all 0.2s ease;
  margin-bottom: 2px;
}

.dropdown-item:hover {
  background: #f3f4f6;
  color: #1f2937;
  transform: translateX(4px);
}

.dropdown-item.active {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.dropdown-icon {
  font-size: 1.25rem;
  flex-shrink: 0;
}

.dropdown-text {
  display: flex;
  flex-direction: column;
  gap: 0.125rem;
}

.dropdown-title {
  font-weight: 600;
  font-size: 0.95rem;
}

.dropdown-description {
  font-size: 0.8rem;
  opacity: 0.7;
}

/* Responsive Design */
@media (max-width: 768px) {
  .mobile-toggle {
    display: flex;
  }
  
  .navbar-menu {
    position: fixed;
    top: 64px;
    left: 0;
    right: 0;
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    flex-direction: column;
    align-items: stretch;
    padding: 1rem;
    gap: 0.5rem;
    transform: translateX(-100%);
    transition: transform 0.3s ease;
    max-height: calc(100vh - 64px);
    overflow-y: auto;
  }
  
  .navbar-menu.mobile-open {
    transform: translateX(0);
  }
  
  .nav-link {
    justify-content: flex-start;
    width: 100%;
    padding: 1rem;
    border-radius: 8px;
  }
  
  .dropdown-menu {
    position: static;
    box-shadow: none;
    background: rgba(255, 255, 255, 0.1);
    margin-top: 0.5rem;
    border: 1px solid rgba(255, 255, 255, 0.2);
  }
  
  .dropdown-item {
    color: rgba(255, 255, 255, 0.9);
  }
  
  .dropdown-item:hover {
    background: rgba(255, 255, 255, 0.15);
    color: white;
  }
  
  .dropdown-item.active {
    background: rgba(255, 255, 255, 0.25);
    color: white;
  }
}

@media (max-width: 480px) {
  .navbar-container {
    padding: 0 0.75rem;
  }
  
  .nav-text {
    font-size: 0.9rem;
  }
}
</style>