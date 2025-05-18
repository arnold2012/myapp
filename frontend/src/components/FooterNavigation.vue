<template>
  <div class="footer-nav">
    <div class="footer-nav-container">
      <router-link 
        to="/orders" 
        class="footer-nav-item"
        :class="{ 'active': $route.path === '/orders' }"
      >
        <div class="footer-icon">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M22 12H2"/>
            <path d="M5.45 5.11L2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11z"/>
            <line x1="6" y1="16" x2="6.01" y2="16"/>
            <line x1="10" y1="16" x2="10.01" y2="16"/>
          </svg>
        </div>
        <span>Órdenes</span>
      </router-link>
      <router-link 
        to="/order-transactions" 
        class="footer-nav-item"
        :class="{ 'active': $route.path === '/order-transactions' }"
      >
        <div class="footer-icon">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <rect x="2" y="3" width="20" height="14" rx="2" ry="2"/>
            <line x1="8" y1="21" x2="16" y2="21"/>
            <line x1="12" y1="17" x2="12" y2="21"/>
          </svg>
        </div>
        <span>Historial</span>
      </router-link>
      <router-link 
        to="/facturar" 
        class="footer-nav-item"
        :class="{ 'active': $route.path === '/facturar' }"
      >
        <div class="footer-icon">
          <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/>
            <polyline points="14 2 14 8 20 8"/>
            <line x1="16" y1="13" x2="8" y2="13"/>
            <line x1="16" y1="17" x2="8" y2="17"/>
            <polyline points="10 9 9 9 8 9"/>
          </svg>
        </div>
        <span>Facturar</span>
      </router-link>
    </div>
  </div>
</template>

<script setup>
// No se requiere lógica adicional para este componente
</script>

<style scoped>
.footer-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background-color: #2c3e50;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.1);
  z-index: 100;
}

.footer-nav-container {
  display: flex;
  justify-content: space-around;
  padding: 0.5rem;
}

.footer-nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  color: white;
  text-decoration: none;
  padding: 0.5rem;
  border-radius: 4px;
  transition: all 0.3s ease;
  width: 100%;
  max-width: 120px;
}

.footer-nav-item:hover {
  background: #34495e;
}

.footer-nav-item.active {
  background: #42b983;
  font-weight: bold;
}

.footer-icon {
  margin-bottom: 4px;
}

.footer-icon svg {
  width: 20px;
  height: 20px;
}

@media (min-width: 768px) {
  .footer-nav {
    display: none; /* Ocultar en pantallas grandes */
  }
}
</style>