<script setup>
import { ref, defineEmits, onMounted, onUnmounted } from 'vue';
import { ListPrograms, ListCycles, CreateWeek } from '../../wailsjs/go/main/App';

const programs = ref([]);
const cycles = ref([]);
const week = ref({
    program: "",
    name: "",
    cycle: "",
})
const emit = defineEmits('week-created', 'error');

const handleEnter = (e) => {
    if (e.key === 'Enter' && week.value.name !== '' && week.value.program !== '' && week.value.cycle !== '') {
        createWeek();
    }
}

onMounted(() => {
    window.addEventListener('keydown', handleEnter);

});

onUnmounted(() => {
    window.removeEventListener('keydown', handleEnter);
    week.value = {
        program: "",
        name: "",
        cycle: "",
    };
});

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

    const cycleSelect = document.querySelectorAll('.form-select')[1];
    cycleSelect.selectedIndex = 0;
}

async function createWeek() {
    try {
        const result = await CreateWeek(week.value.name, week.value.cycle);
        week.value = {
            program: "",
            name: "",
            cycle: "",
        };
        emit('week-created');
    } catch (error) {
        if (typeof error === 'string') {
            error = error.charAt(0).toUpperCase() + error.slice(1);
        }
        emit('error', error);
    }
}

function handleProgramChange() {
    week.value.cycle = "";
    week.value.name = "";

    listCycles();
}

listPrograms();
</script>

<template>
    <form @submit.prevent="createWeek" @keydown.enter="createWeek">
        <h2 class="form-title">Crear Semana</h2>
        <input type="text" v-model="week.name" placeholder="Semana" class="form-control" />
        <select v-model="week.program" @change="handleProgramChange" class="form-control form-select"
            placeholder="Programa" v-if="programs[0] !== 'No hay programas'">
            <option value="" disabled>Selecciona un programa</option>
            <option v-if="programs[0] === 'No hay programas'" :value="programs[0]" disabled>{{ programs[0] }}</option>
            <option v-else v-for="(program, index) in programs" :key="index">{{ program }}</option>
        </select>
        <select class="form-control form-select" v-else disabled>
            <option value="" disabled selected>No hay programas</option>
        </select>
        <select v-model="week.cycle" placeholder="Ciclo" class="form-control form-select" v-if="week.program && cycles[0] !== 'No hay ciclos'">
            <option value="" disabled>Selecciona un ciclo</option>
            <option v-if="cycles[0] === 'No hay ciclos'" :value="cycles[0]" disabled>{{ cycles[0] }}</option>
            <option v-else v-for="(cycle, index) in cycles" :key="index">{{ cycle }}</option>
        </select>
        <select v-else-if="week.program && cycles[0] === 'No hay ciclos'"
            class="form-control form-select" disabled>
            <option value="" disabled selected>No hay ciclos</option>
        </select>
        <select placeholder="Ciclo" class="form-control form-select" v-else disabled>
            <option value="" disabled selected>Selecciona un programa primero</option>
        </select>
        <button type="button" class="btn btn-dark" @click="createWeek"
            :disabled="week.name === '' || week.program === '' || week.cycle === ''">Crear</button>
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