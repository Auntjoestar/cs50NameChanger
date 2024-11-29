<template>
  <div class="input-box">
    <label :for="id">{{ label }}</label>
    <select :id="id" autocomplete="off" class="input" 
        v-model="modelValue" @change="onChange"
        v-if="items[0] != emptyMessage && cycles[0] != 'No hay ciclos' && newFilesName.cycle != ''">
      <option value="" disabled selected v-if="items.length == 0">{{ loadingMessage }}</option>
      <option value="" disabled selected v-else>{{ selectMessage }}</option>
      <option v-for="(item, index) in items" :key="index">
        {{ item }}
      </option>
    </select>
    <select :id="id" autocomplete="off" class="input" v-model="modelValue" disabled v-else-if="dependencyMissing">
      <option value="" disabled selected>{{ dependencyMessage }}</option>
    </select>
    <select :id="id" autocomplete="off" class="input" v-model="modelValue" disabled v-else>
      <option value="" disabled selected>{{ emptyMessage }}</option>
    </select>
  </div>
</template>

<script>
export default {
  props: {
    id: String,
    label: String,
    modelValue: String,
    items: Array,
    dependencyMissing: Boolean,
    cycles: Array,
    emptyMessage: String,
    loadingMessage: String,
    selectMessage: String,
    dependencyMessage: String,
    newFilesName: Object,
    onChange: Function,
  },
};
</script>
