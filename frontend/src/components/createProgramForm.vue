<script setup>
import { ref } from 'vue'
import {
    CreateProgram,
    ListPrograms
} from '../../wailsjs/go/main/App';

const program = ref("");

async function createProgram() {
    await CreateProgram(program.value);
    program.value = "";
}

</script>

<template>
    <form>
        <h2 class="form-title">Crear Programa</h2>
        <input type="text" v-model="program" placeholder="Programa" class="form-control" />
        <button type="button" class="btn btn-dark" @click="createProgram" :disabled="program === ''">Crear</button>
    </form>
</template>

<style scoped lang="scss">
h2 {
    margin-bottom: 1rem;
}

.form-control {
    width: 100%;
    padding: 0.5rem;
    margin-bottom: 1rem;
    border: 1px solid #ced4da;
    border-radius: 0.25rem;
}

.btn {
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 0.25rem;
    background-color: #007bff;
    color: #fff;
    cursor: pointer;
}

.btn-dark {
    background-color: #343a40;
}

.btn-dark:disabled {
    background-color: #6c757d;
    color: #fff;
    cursor: not-allowed;
}
</style>