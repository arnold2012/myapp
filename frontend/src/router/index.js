// frontend/src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/',
    name: 'Inicio',
    component: () => import('@/views/HomeView.vue'),
    meta: { icon: '🏠' }
  },
  {
    path: '/contactos',
    name: 'Contactos',
    component: () => import('@/views/ContactosView.vue'),
    meta: { icon: '📋' }
  },
  {
    path: '/productos',
    name: 'Productos',
    component: () => import('@/views/ProductoView.vue'),
    meta: { icon: '📋' }
  },
  {
    path: '/impuestos',
    name: 'Impuestos',
    component: () => import('@/views/Impuestos.vue'),
    meta: { icon: '💰' }
  },
  {
    path: '/marcas',
    name: 'Marcas',
    component: () => import('@/views/Marcas.vue'),
    meta: { icon: '🏷️' }
  },
  {
    path: '/categoria',
    name: 'Categoría',
    component: () => import('@/views/Categoria.vue'),
    meta: { icon: '📂' }
  },
  {
    path: '/orders',
    name: 'Crear Orden De Ventas',
    component: () => import('@/views/OrderView.vue'),
    meta: { icon: '📦' }
  },
  {
    path: '/order-transactions',
    name: 'Historial de Transacciones',
    component: () => import('@/views/Order_Transaction_View.vue'),
    meta: { icon: '📊' }
  },
  {
    path: '/facturar',
    name: 'Crear factura',
    component: () => import('@/views/FacturarOrderView.vue'),
    meta: { icon: '🧾' }
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;