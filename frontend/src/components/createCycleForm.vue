<script setup>
import { ref, defineEmits } from 'vue'
import { ListPrograms, CreateCycle } from '../../wailsjs/go/main/App';


const programs = ref([]);
const emit = defineEmits('cycle-created', 'error');

const cycle = ref({
    name: "",
    program: "",
})

async function listPrograms() {
    const response = await ListPrograms();
    programs.value = response;
}

async function createCycle() {
    try {
        const result = await CreateCycle(cycle.value.name, cycle.value.program);
        console.log(result);
        cycle.value.name = "";
        cycle.value.program = "";
        emit('cycle-created');
    } catch (error) {
        error = error.charAt(0).toUpperCase() + error.slice(1);
        emit('error', error || "Error al crear el ciclo");
    }
}

listPrograms();
</script>

<template>
    <form @submit.prevent="createCycle" @keydown.enter="createCycle">
        <h2 class="form-title">Crear Ciclo</h2>
        <input type="text" v-model="cycle.name" placeholder="Ciclo" class="form-control" />
        <select v-model="cycle.program" placeholder="Programa" class="form-control form-select">
            <option value="" disabled>Selecciona un programa</option>
            <option v-if="programs[0] === 'No hay programas'" :value="programs[0]" disabled>{{ programs[0] }}</option>
            <option v-else v-for="(program, index) in programs" :key="index">{{ program }}</option>
        </select>
        <button type="button" @click="createCycle" class="btn btn-dark"
            :disabled="cycle.name === '' || cycle.program === ''">Crear</button>
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

.form-select {
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
    cursor: pointer;
}

.btn-dark {
    background-color: #343a40;
    color: #fff;
}

.btn-dark:disabled {
    background-color: #6c757d;
    color: #fff;
    cursor: not-allowed;
}
</style>