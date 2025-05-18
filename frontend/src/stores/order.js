import { defineStore } from 'pinia';
import api from '@/api/orders'; // Asegúrate de que este archivo exista y esté configurado correctamente

export const useOrdersStore = defineStore('orders', {
  actions: {
    async fetchOrders() {
      try {
        const response = await api.getOrders();
        return response.data;
      } catch (error) {
        console.error('Error fetching orders:', error);
        return [];
      }
    },
    async createOrder(order) {
      try {
        const response = await api.createOrder(order);
        return response.data;
      } catch (error) {
        console.error('Error creating order:', error);
        throw error;
      }
    },
    async deleteOrder(id) {
      try {
        await api.deleteOrder(id);
      } catch (error) {
        console.error('Error deleting order:', error);
        throw error;
      }
    }
  }
});
