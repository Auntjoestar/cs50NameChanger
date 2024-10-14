<script setup>
import { ref } from 'vue'
import { ListPrograms, ListCycles, CreateWeek } from '../../wailsjs/go/main/App';

const programs = ref([]);
const cycles = ref([]);
const week = ref({
    program: "",
    name: "",
    cycle: "",
})

async function listPrograms() {
    const response = await ListPrograms();
    programs.value = response;
}

async function listCycles() {
    const response = await ListCycles(week.value.program);
    if (response[0] === "No hay ciclos") {
        cycles.value = ["No hay ciclos"];
        return;
    }
    cycles.value = response;
}

async function createWeek() {
    await CreateWeek(week.value.name.toLowerCase(), week.value.cycle);
    week.value = {
        name: "",
        cycle: "",
    }
}

listPrograms();
</script>

<template>
    <form>
        <h2 class="form-title">Crear Semana</h2>
        <input type="text" v-model="week.name" placeholder="Semana" class="form-control" />
        <select v-model="week.program" @change="listCycles" class="form-control form-select">
            <option value="" disabled>Selecciona un programa</option>
            <option v-if = "programs[0] === 'No hay programas'" :value="programs[0]" disabled>{{ programs[0] }}</option>
            <option v-else v-for="(program, index) in programs" :key="index">{{ program }}</option>
        </select>
        <select v-model="week.cycle" placeholder="Ciclo" class="form-control form-select">
            <option value="" disabled>Selecciona un ciclo</option>
            <option v-if = "cycles[0] === 'No hay ciclos'" :value="cycles[0]" disabled>{{ cycles[0] }}</option>
            <option v-else v-for="(cycle, index) in cycles" :key="index">{{ cycle }}</option>
        </select>
        <button type="button" class="btn btn-dark" @click="createWeek">Crear</button>
    </form>
</template>