<script setup>
import { ref, defineEmits, onMounted, onUnmounted } from 'vue';
import { ListPrograms, ListCycles, CreateGroup } from '../../wailsjs/go/main/App';

const programs = ref([]);
const cycles = ref([]);
const group = ref({
    program: "",
    name: "",
    cycle: "",
})
const emit = defineEmits('group-created', 'error');

const handleEnter = (e) => {
    if (e.key === 'Enter' && group.value.name !== '' && group.value.program !== '' && group.value.cycle !== '') {
        createGroup();
    }
}

onMounted(() => {
    window.addEventListener('keydown', handleEnter);
});

onUnmounted(() => {
    window.removeEventListener('keydown', handleEnter);
    group.value = {
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
    const response = await ListCycles(group.value.program);
    if (response[0] === "No hay ciclos") {
        cycles.value = ["No hay ciclos"];
        return;
    }
    cycles.value = response;
}

async function createGroup() {
    try {
        const result = await CreateGroup(group.value.name, group.value.cycle);
        group.value = {
            program: "",
            name: "",
            cycle: "",
        };
        emit('group-created');
    } catch (error) {
        if (typeof error === 'string') {
            error = error.charAt(0).toUpperCase() + error.slice(1);
        }
        emit('error', error);
    }
}

function handleProgramChange() {
    group.value.cycle = "";
    group.value.name = "";

    listCycles();
}

listPrograms();
</script>

<template>
    <form @submit.prevent="createGroup" @keydown.enter="createGroup">
        <h2 class="form-title">Crear Grupo</h2>
        <input type="text" v-model="group.name" placeholder="Grupo" class="form-control" />
        <select v-model="group.program" @change="handleProgramChange" class="form-control form-select">
            <option value="" disabled>Selecciona un programa</option>
            <option v-if="programs[0] === 'No hay programas'" :value="programs[0]" disabled>{{ programs[0] }}</option>
            <option v-else v-for="(program, index) in programs" :key="index">{{ program }}</option>
        </select>
        <select v-model="group.cycle" placeholder="Ciclo" class="form-control form-select" v-if="group.program && cycles[0] !== 'No hay ciclos'">
            <option value="" disabled>Selecciona un ciclo</option>
            <option v-if="cycles[0] === 'No hay ciclos'" :value="cycles[0]" disabled>{{ cycles[0] }}</option>
            <option v-else v-for="(cycle, index) in cycles" :key="index">{{ cycle }}</option>
        </select>
        <select placeholder="Ciclo" class="form-control form-select" v-else-if="group.program && cycles[0] === 'No hay ciclos'" disabled>
            <option value="" disabled selected>No hay ciclos</option>
        </select>
        <select v-else placeholder="Ciclo" class="form-control form-select" disabled>
            <option value="" disabled selected>Selecciona un programa primero</option>
        </select>
        <button type="button" class="btn btn-dark" @click="createGroup" @keydown.enter="createGroup"
            :disabled="group.name === '' || group.program === '' || group.cycle === ''">Crear</button>
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