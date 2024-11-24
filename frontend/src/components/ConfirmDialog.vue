<template>
  <div class="dialog-overlay" @click.self="onCancel">
    <div class="dialog-content">
      <h3>{{ title }}</h3>
      <p>{{ message }}</p>
      <div class="dialog-actions">
        <button class="btn btn-confirm" @click="onConfirm" @keydown.enter="onConfirm">Confirmar</button>
        <button class="btn btn-cancel" @click="onCancel" @keydown.escape="onCancel">Cancelar</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineProps, defineEmits } from 'vue';

const props = defineProps({
  title: String,
  message: String,
});
const emit = defineEmits(['confirm', 'cancel']);

function onConfirm() {
  emit('confirm');
}

function onCancel() {
  emit('cancel');
}
</script>

<style scoped>
.dialog-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  animation: fadeIn 0.3s forwards;
  z-index: 1000;
}

.dialog-content {
  background: white;
  padding: 20px 30px;
  border-radius: 8px;
  width: 90%;
  max-width: 400px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.2);
  text-align: center;
  transform: scale(0.95);
  animation: slideUp 0.3s forwards;
}

.dialog-content h3 {
  margin-top: 0;
  color: #333;
}

.dialog-actions {
  margin-top: 20px;
  display: flex;
  gap: 10px;
  justify-content: center;
}

.btn {
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.btn-confirm {
  background-color: #007bff;
  color: white;
}

.btn-confirm:hover {
  background-color: #0056b3;
}

.btn-cancel {
  background-color: #dc3545;
  color: white;
}

.btn-cancel:hover {
  background-color: #c82333;
}

@keyframes fadeIn {
  to { opacity: 1; }
}

@keyframes slideUp {
  to { transform: scale(1); }
}
</style>
