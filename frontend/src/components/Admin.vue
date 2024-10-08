<script setup>
import { ref } from 'vue'
import {
    CreateProgram,
    ListPrograms,
    CreateCycle,
    ListCycles,
    CreateWeek,
    CreateGroup
} from '../../wailsjs/go/main/App';

const program = ref("");
const programs = ref([]);

const cycle = ref({
    name: "",
    program: "",
})
const cycles = ref([]);

const week = ref({
    name: "",
    cycle: "",
})

const group = ref({
    name: "",
    cycle: "",
})

async function createProgram() {
    await CreateProgram(program.value.toLocaleLowerCase());
    program.value = "";
}

async function listPrograms() {
    const response = await ListPrograms();
    programs.value = response;
}

async function listCycles() {
    const response = await ListCycles(cycle.value.program);
    if (response[0] === "No hay ciclos") {
        cycles.value = ["No hay ciclos"];
        return;
    }
    cycles.value = response;
}

async function createCycle() {
    await CreateCycle(cycle.value.name.toLocaleLowerCase(), cycle.value.program);
    cycle.value = {
        name: "",
        program: "",
    }
}

async function createWeek() {
    await CreateWeek(week.value.name.toLocaleLowerCase(), week.value.cycle);
    week.value = {
        name: "",
        cycle: "",
    }
}

async function createGroup() {
    await CreateGroup(group.value.name.toLocaleLowerCase(), group.value.cycle);
    group.value = {
        name: "",
        cycle: "",
    }
}

listPrograms();
</script>

<template>
    <header>
        <h1>Panel de administración</h1>
        <nav>
            <ul>
            </ul>
        </nav>
    </header>
    <main>
        <section>
            <h2>Crear</h2>
            <form>
                <input type="text" v-model="program" placeholder="Programa" />
                <button type="button" @click="createProgram">Crear</button>
            </form>
        </section>
        <section>
            <form>
                <input type="text" v-model="cycle.name" placeholder="Ciclo" />
                <select v-model="cycle.program" type="text" placeholder="Programa">
                    <option value="">Selecciona un programa</option>
                    <option v-for="(program, index) in programs" :key="index">{{ program }}</option>
                </select>
                <button type="button" @click="createCycle">Crear</button>
            </form>
        </section>
        <section>
            <form>
                <input type="text" v-model="week.name" placeholder="Semana" />
                <select v-model="cycle.program" type="text" placeholder="Programa" @change="listCycles">
                    <option value="">Selecciona un programa</option>
                    <option v-for="(program, index) in programs" :key="index">{{ program }}</option>
                </select>
                <select v-model="week.cycle" type="text" placeholder="Ciclo">
                    <option value="">Selecciona un ciclo</option>
                    <option v-for="(cycle, index) in cycles" :key="index">{{ cycle }}</option>
                </select>
                <button type="button" @click="createWeek">Crear</button>
            </form>
        </section>
        <section>
            <form>
                <input type="text" v-model="group.name" placeholder="Grupo" />
                <select v-model="cycle.program" type="text" placeholder="Programa" @change="listCycles">
                    <option value="">Selecciona un programa</option>
                    <option v-for="(program, index) in programs" :key="index">{{ program }}</option>
                </select>
                <select v-model="group.cycle" type="text" placeholder="Ciclo">
                    <option value="">Selecciona un ciclo</option>
                    <option v-for="(cycle, index) in cycles" :key="index">{{ cycle }}</option>
                </select>
                <button type="button" @click="createGroup">Crear</button>
            </form>
        </section>
    </main>
</template>