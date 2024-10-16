<script setup>
import { ref } from 'vue'
import { ListPrograms, ListCycles, CreateGroup } from '../../wailsjs/go/main/App';

const programs = ref([]);
const cycles = ref([]);
const group = ref({
    program: "",
    name: "",
    cycle: "",
})

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
    await CreateGroup(group.value.name, group.value.cycle);
    group.value = {
        name: "",
        cycle: "",
    }
}

listPrograms();
</script>

<template>
    <form>
        <h2 class="form-title">Crear Grupo</h2>
        <input type="text" v-model="group.name" placeholder="Grupo" class="form-control" />
        <select v-model="group.program" @change="listCycles" class="form-control form-select">
            <option value="" disabled>Selecciona un programa</option>
            <option v-if = "programs[0] === 'No hay programas'" :value="programs[0]" disabled>{{ programs[0] }}</option>
            <option v-else v-for="(program, index) in programs" :key="index">{{ program }}</option>
        </select>
        <select v-model="group.cycle" placeholder="Ciclo" class="form-control form-select">
            <option value="" disabled>Selecciona un ciclo</option>
            <option v-if = "cycles[0] === 'No hay ciclos'" :value="cycles[0]" disabled>{{ cycles[0] }}</option>
            <option v-else v-for="(cycle, index) in cycles" :key="index">{{ cycle }}</option>
        </select>
        <button type="button" class="btn btn-dark" @click="createGroup" >Crear</button>
    </form>
</template>