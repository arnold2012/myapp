<template>
    <div v-if="isOpen" class="modal-overlay" @click.self="closeModal">
      <div class="modal-content">
        <h2>{{ contactoEdit.id_contacto ? 'Editar' : 'Nuevo' }} Contacto</h2>
        <button class="close-btn" @click="closeModal">×</button>
        
        <form @submit.prevent="handleSubmit">
          <div class="form-group">
            <label>Razón Social:</label>
            <input v-model="formData.razon_social" required>
          </div>
          
          <div class="form-group">
            <label>RUC/CI:</label>
            <input v-model="formData.ruc_ci" required>
          </div>
          
          <div class="form-group">
            <label>Teléfono:</label>
            <input v-model="formData.telefono_cto">
          </div>
          
          <div class="form-group">
            <label>Dirección:</label>
            <input v-model="formData.direccion_cto">
          </div>
  
          <button type="submit" class="submit-btn">
            {{ contactoEdit.id_contacto ? 'Actualizar' : 'Guardar' }}
          </button>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, watch } from 'vue'
  
  const props = defineProps({
    isOpen: Boolean,
    contactoEdit: {
      type: Object,
      default: () => ({})
    }
  })
  
  const emit = defineEmits(['close', 'submit'])
  
  const formData = ref({
    razon_social: '',
    ruc_ci: '',
    telefono_cto: '',
    direccion_cto: ''
  })
  
  // Actualizar formData cuando cambia contactoEdit
  watch(() => props.contactoEdit, (newVal) => {
    formData.value = { ...newVal }
  }, { deep: true, immediate: true })
  
  const closeModal = () => {
    emit('close')
  }
  
  const handleSubmit = () => {
    emit('submit', formData.value)
  }
  </script>
  
  <style scoped>
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
  }
  
  .modal-content {
    background: white;
    padding: 2rem;
    border-radius: 8px;
    width: 90%;
    max-width: 500px;
    position: relative;
  }
  
  .close-btn {
    position: absolute;
    top: 10px;
    right: 10px;
    background: none;
    border: none;
    font-size: 1.5rem;
    cursor: pointer;
  }
  
  .form-group {
    margin-bottom: 1rem;
  }
  
  .form-group label {
    display: block;
    margin-bottom: 0.5rem;
    font-weight: bold;
  }
  
  .form-group input {
    width: 100%;
    padding: 0.5rem;
    border: 1px solid #ddd;
    border-radius: 4px;
  }
  
  .submit-btn {
    background-color: #42b983;
    color: white;
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    margin-top: 1rem;
  }
  </style>