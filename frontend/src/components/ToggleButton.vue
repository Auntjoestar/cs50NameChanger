<!-- ToggleButton.vue -->
<script setup>
import { defineEmits, defineProps } from 'vue';

const props = defineProps({
  element: {
    type: String,
    default: 'fa-sync-alt', // Default Font Awesome icon
  },
});

const emit = defineEmits(['toggle']);
</script>

<template>
  <button class="toggle-button" @click="$emit('toggle')">
    <!-- Dynamically bind the Font Awesome icon class -->
    <i :class="['fas', element]"></i>
  </button>
</template>

<style scoped lang="scss">
.toggle-button {
  position: fixed;
  bottom: 20px;
  right: 20px;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  background-color: #007bff;
  color: #fff;
  border: none;
  font-size: 1.5rem;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
  transition: background-color 0.3s;
  z-index: 1000;

  &:hover {
    background-color: #0056b3;
  }
}
</style>
