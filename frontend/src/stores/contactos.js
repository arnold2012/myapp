import { defineStore } from 'pinia'
import api from '@/api/contactos'

export const useContactosStore = defineStore('contactos', {
  actions: {
    async fetchContactos() {
      try {
        const response = await api.getContactos()
        return response.data
      } catch (error) {
        console.error('Error fetching contactos:', error)
        return []
      }
    },
    async createContacto(contacto) {
      await api.createContacto(contacto)
    },
    async deleteContacto(id) {
      await api.deleteContacto(id)
    }
  }
})