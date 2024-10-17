<script setup>
import { ref } from 'vue'
import { ListPrograms, CreateCycle } from '../../wailsjs/go/main/App';


const programs = ref([]);

const cycle = ref({
    name: "",
    program: "",
})

async function listPrograms() {
    const response = await ListPrograms();
    programs.value = response;
}

async function createCycle() {
    await CreateCycle(cycle.value.name, cycle.value.program);
    cycle.value = {
        name: "",
        program: "",
    }
}

listPrograms();
</script>

<template>
    <form>
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