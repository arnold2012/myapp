<template>
  <Navbar />
  <main class="container">
    <router-view />
  </main>
</template>

<script setup>
import Navbar from '@/components/Navbar.vue'
</script>

<style>
/* Estilos globales */
body {
  margin: 0;
  font-family: Arial, sans-serif;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 1rem;
}
</style>