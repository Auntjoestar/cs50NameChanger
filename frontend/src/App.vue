<script setup>
import { ref, onMounted } from 'vue';
import Forms from './components/Forms.vue';
import Admin from './components/Admin.vue';
import ToggleButton from './components/ToggleButton.vue';

const showAdmin = ref(false);
const buttonVisible = ref(false);

window.addEventListener('beforeunload', () => {
  localStorage.clear();
});

window.addEventListener('keydown', (e) => {
  if (e.ctrlKey && e.key === 'a') {
    toggleAdminView();
  }
});

function showButton() {
  buttonVisible.value = true;
}


onMounted(() => {
  const savedState = localStorage.getItem('showAdmin');
  showAdmin.value = savedState === 'true';
});

const toggleAdminView = () => {
  showAdmin.value = !showAdmin.value;
  localStorage.setItem('showAdmin', showAdmin.value);
};

window.addEventListener('beforeunload', () => {
  localStorage.clear();
});
</script>

<template>
  <div class="main-container">
    <Admin v-if="showAdmin" />
    <Forms  @connected="showButton" v-else />

    <!-- Circular Toggle Button -->
    <ToggleButton @toggle="toggleAdminView" :element="showAdmin ? 'fa-image' : 'fa-database'" v-if="buttonVisible" />
  </div>
</template>

<style>
html {
  background-color: #f8f9fa;
}

body {
  background-color: #f8f9fa !important;
}

.main-container {
  position: relative;
  background-color: #f8f9fa;
}

</style>
