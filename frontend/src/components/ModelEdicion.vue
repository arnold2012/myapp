<template>
    <div v-if="isOpen" class="modal-overlay" @click.self="closeModal">
      <div class="modal-content">
        <h2>Editar Contacto</h2>
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
  
          <div class="form-actions">
            <button type="submit" class="btn-submit">Guardar Cambios</button>
            <button type="button" class="btn-cancel" @click="closeModal">Cancelar</button>
          </div>
        </form>
      </div>
    </div>
  </template>
  
  <script setup>
  import { ref, watch } from 'vue'
  
  const props = defineProps({
    isOpen: Boolean,
    contacto: {
      type: Object,
      required: true
    }
  })
  
  const emit = defineEmits(['close', 'submit'])
  
  const formData = ref({ ...props.contacto })
  
  watch(() => props.contacto, (newVal) => {
    formData.value = { ...newVal }
  }, { deep: true })
  
  const closeModal = () => {
    emit('close')
  }
  
  const handleSubmit = () => {
    emit('submit', formData.value)
  }
  </script>
  
  <style scoped>
  /* Estilos similares a ModelFormularioContacto.vue */
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
  
  .form-actions {
    display: flex;
    justify-content: flex-end;
    gap: 10px;
    margin-top: 1rem;
  }
  
  .btn-submit {
    background-color: #42b983;
    color: white;
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  
  .btn-cancel {
    background-color: #e74c3c;
    color: white;
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }
  </style>