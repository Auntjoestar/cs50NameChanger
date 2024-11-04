<script setup>
import { ref, onMounted } from 'vue';
import Forms from './components/Forms.vue';
import Admin from './components/Admin.vue';
import ToggleButton from './components/ToggleButton.vue';

const showAdmin = ref(false);

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
    <Forms v-else />

    <!-- Circular Toggle Button -->
    <ToggleButton @toggle="toggleAdminView" :element="showAdmin ? 'fa-image' : 'fa-database'" />
  </div>
</template>

<style>
body {
  background-color: transparent;
}

.main-container {
  position: relative;
  background-color: rgba(255, 255, 255, 0.1); /* Match app background */
}


#app { 
  background-color: rgba(255, 255, 255, 0.1); /* Slightly translucent white background */
  backdrop-filter: blur(5px);
}

.main-container {
  position: relative;
}
</style>
