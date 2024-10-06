<script setup>
import Forms from './components/Forms.vue'
import Admin from './components/Admin.vue'
import { ref } from 'vue'

const showAdmin = ref(false)

</script>

<template>
  <Admin v-if="showAdmin" />
  <Forms v-else />
  <button @click="showAdmin = !showAdmin">Toggle Admin</button>
</template>

<style></style>
