<!-- frontend/src/components/Navbar.vue -->
<template>
  <nav class="navbar">
    <template v-for="route in routes" :key="route.path">
      <!-- Dropdown para Productos -->
      <div v-if="route.name === 'Productos'" class="dropdown">
        <button class="nav-link dropdown-toggle">
          <span class="icon">{{ route.meta?.icon }}</span>
          {{ route.name }}
        </button>
        <div class="dropdown-menu">
          <router-link
            :to="{ path: '/productos', query: { view: 'list' } }"
            class="dropdown-item"
            :class="{ 'active': $route.path === '/productos' && $route.query.view === 'list' }"
          >
            <span class="icon">📋</span>
            Ver Productos
          </router-link>
          <router-link
            :to="{ path: '/productos', query: { view: 'filter' } }"
            class="dropdown-item"
            :class="{ 'active': $route.path === '/productos' && $route.query.view === 'filter' }"
          >
            <span class="icon">🔍</span>
            Filtrar Productos
          </router-link>
        </div>
      </div>
      <!-- Enlaces normales para otras rutas -->
      <router-link
        v-else
        :to="route.path"
        class="nav-link"
        :class="{ 'active': $route.path === route.path }"
      >
        <span class="icon">{{ route.meta?.icon }}</span>
        {{ route.name }}
      </router-link>
    </template>
  </nav>
</template>

<script setup>
import { computed } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

const routes = computed(() => {
  return router.options.routes.filter(route => route.name); // Solo rutas con nombre
});
</script>

<style scoped>
.navbar {
  display: flex;
  background: #2c3e50;
  padding: 1rem;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
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

/* Estilos para el dropdown */
.dropdown {
  position: relative;
  display: inline-block;
}

.dropdown-toggle {
  background: none;
  border: none;
  cursor: pointer;
  color: white;
  padding: 0.5rem 1rem;
  margin: 0 0.5rem;
  border-radius: 4px;
  display: flex;
  align-items: center;
}

.dropdown-toggle:hover {
  background: #34495e;
}

.dropdown-menu {
  display: none;
  position: absolute;
  background: #34495e;
  min-width: 160px;
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.2);
  z-index: 1;
  border-radius: 4px;
  margin-top: 0.2rem;
}

.dropdown:hover .dropdown-menu {
  display: block;
}

.dropdown-item {
  color: white;
  padding: 0.5rem 1rem;
  text-decoration: none;
  display: flex;
  align-items: center;
  transition: background 0.3s ease;
}

.dropdown-item:hover {
  background: #42b983;
}

.dropdown-item.active {
  background: #42b983;
  font-weight: bold;
}
</style>