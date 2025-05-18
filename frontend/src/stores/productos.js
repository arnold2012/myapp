// frontend/src/stores/productos.js
import { defineStore } from 'pinia';
import api from '@/api/productos';

export const useProductosStore = defineStore('productos', {
  state: () => ({
    resultados: [],
  }),
  actions: {
    async buscarProductos(query) {
      try {
        const res = await api.searchProductos(query);
        this.resultados = res;
      } catch (error) {
        console.error('Error al buscar productos:', error);
      }
    }
  }
});
