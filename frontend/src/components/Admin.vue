<script setup>
import { ref, reactive } from 'vue'
import {
    WatchPrograms,
    WatchCycles,
    WatchWeeks,
    WatchGroups,
} from '../../wailsjs/go/main/App';
import ProgramForm from './createProgramForm.vue';
import CycleForm from './createCycleForm.vue';
import WeekForm from './createWeekForm.vue';
import GroupForm from './createGroupForm.vue';

const creatView = reactive({
    program: true,
    cycle: false,
    week: false,
    group: false,
})

const programs = ref([])
const cycles = ref([])
const weeks = ref([])
const groups = ref([])

async function watchPrograms() {
    programs.value.push(await WatchPrograms());
    if (programs.id=== 0) {
        console.log("No hay programas disponibles")
        return;
    }
    console.log(programs.value);
}

async function watchCycles() {
    cycles.value.push(await WatchCycles());
    if (cycles.id === 0) {
        console.log("No hay ciclos disponibles")
        return;
    }
    console.log(cycles.value);
}

async function watchWeeks() {
    weeks.value.push(await WatchWeeks());
    if (weeks.id === 0) {
        console.log("No hay semanas disponibles")
        return;
    }
    console.log(weeks.value);
}

async function watchGroups() {
    groups.value.push(await WatchGroups());
    if (groups.id === 0) {
        console.log("No hay grupos disponibles")
        return;
    }
    console.log(groups.value);
}

watchPrograms();
watchCycles();
watchWeeks();
watchGroups();
</script>

<template>
    <header>
        <nav class="navbar navbar-expand-lg bg-body-tertiary">
            <div class="container-fluid">
                <a class="navbar-brand" href="#">Panel de administración</a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                    data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent"
                    aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                        <li class="nav-item">
                            <a class="nav-link active" aria-current="page" href="#">Ver</a>
                        </li>
                        <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown"
                                aria-expanded="false">
                                Crear
                            </a>
                            <ul class="dropdown-menu">
                                <li>
                                    <a class="dropdown-item"
                                        @click="creatView.program = true; creatView.cycle = false; creatView.week = false; creatView.group = false;">Programa
                                    </a>
                                </li>
                                <li>
                                    <a class="dropdown-item"
                                        @click="creatView.program = false; creatView.cycle = true; creatView.week = false; creatView.group = false;">Ciclo
                                    </a>
                                </li>
                                <li>
                                    <a class="dropdown-item"
                                        @click="creatView.program = false; creatView.cycle = false; creatView.week = true; creatView.group = false;">Semana
                                    </a>
                                </li>
                                <li>
                                    <a class="dropdown-item"
                                        @click="creatView.program = false; creatView.cycle = false; creatView.week = false; creatView.group = true;">Grupo
                                    </a>
                                </li>
                            </ul>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    </header>
    <main>
        <ProgramForm v-if="creatView.program" />

        <CycleForm v-if="creatView.cycle" />

        <WeekForm v-if="creatView.week" />

        <GroupForm v-if="creatView.group" />

        <table v-if="programs[0].id !== 0">
            <thead>
                <tr>
                    <th>ID</th> 
                    <th>Nombre</th>
                    <th>Acciones</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(program, index) in programs" :key="index">
                    <td>{{ program.id }}</td>
                    <td>{{ program.name }}</td>
                    <td>
                        <button class="btn btn-danger">Eliminar</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <table v-if="cycles[0].id !== 0">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Nombre</th>
                    <th>Programa</th>
                    <th>Acciones</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(cycle, index) in cycles" :key="index">
                    <td>{{ cycle.id }}</td>
                    <td>{{ cycle.name }}</td>
                    <td>{{ cycle.program_id }}</td>
                    <td>
                        <button class="btn btn-danger">Eliminar</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <p v-else>No hay ciclos</p>

        <table v-if="weeks[0].id !== 0">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Nombre</th>
                    <th>Ciclo</th>
                    <th>Acciones</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(week, index) in weeks" :key="index">
                    <td>{{ week.id }}</td>
                    <td>{{ week.name }}</td>
                    <td>{{ week.cycle_id }}</td>
                    <td>
                        <button class="btn btn-danger">Eliminar</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <p v-else>No hay semanas</p>

        <table v-if="groups[0].id !== 0">
            <thead>
                <tr>
                    <th>ID</th>
                    <th>Nombre</th>
                    <th>Ciclo</th>
                    <th>Acciones</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(group, index) in groups" :key="index">
                    <td>{{ group.id }}</td>
                    <td>{{ group.name }}</td>
                    <td>{{ group.cicle_id }}</td>
                    <td>
                        <button class="btn btn-danger">Eliminar</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <p v-else>No hay grupos</p>
    </main>
</template>

