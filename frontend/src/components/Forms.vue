<script setup>
import { ref, reactive, defineEmits } from 'vue';
import {
    CreateDB,
    ListPrograms,
    ListCycles,
    ListWeeks,
    ListGroups,
    MakeNewName,
    OpenFileDialog,
    ChangeFileNames
} from '../../wailsjs/go/main/App';

const connected = ref(false);
const contentLoaded = ref(false);
const programs = ref([]);
const cycles = ref([]);
const weeks = ref([]);
const groups = ref([]);
const message = ref("");
const tableVisibility = ref(false);
const emit = defineEmits('connected');

localStorage.getItem('connected') ? connected.value = true : connected.value = false;
localStorage.setItem('connected', connected.value);

const newFilesName = reactive({
    name: "",
    program: "",
    cycle: "",
    week: "",
    group: "",
    files: [],
});

async function connect() {
    try {
        await CreateDB();
        await initializeForms();
        connected.value = true;
        message.value = "Conexión exitosa";
    } catch (error) {
        message.value = `Error al conectar con la base de datos: ${error}`;
    }
}

async function listCycles() {
    try {
        if (programs.value == "") {
            cycles.value = ["Selecciona un programa"];
            return;
        }
        const result = await ListCycles(newFilesName.program);
        if (result[0] === "No hay ciclos") {
            cycles.value = ["No hay ciclos"];
            return;
        }
        cycles.value = result;
    } catch (error) {
        message.value = `Error al cargar los ciclos: ${error}`;
    }
}

async function listWeeks() {
    try {
        if (cycles.value == "") {
            weeks.value = ["Selecciona un ciclo"];
            return;
        }
        const result = await ListWeeks(newFilesName.cycle);
        if (result[0] === "No hay semanas") {
            weeks.value = ["No hay semanas"];
            return;
        }
        weeks.value = result;
    } catch (error) {
        message.value = `Error al cargar las semanas: ${error}`;
    }
}

async function listGroups() {
    try {
        if (cycles.value == "") {
            groups.value = ["Selecciona un ciclo"];
            return;
        }
        const result = await ListGroups(newFilesName.cycle);
        if (result[0] === "No hay grupos") {
            groups.value = ["No hay grupos"];
            return;
        }
        groups.value = result;
    } catch (error) {
        message.value = `Error al cargar los grupos: ${error}`;
    }
}

async function initializeForms() {
    try {
        const result = await ListPrograms();
        programs.value = result;
        if (programs.value.at(0) === "No hay programas") {
            cycles.value = ["No hay ciclos"];
            weeks.value = ["No hay semanas"];
            groups.value = ["No hay grupos"];
        }
        if (!contentLoaded.value) {
            contentLoaded.value = true;
        }
    } catch (error) {
        message.value = `Error al cargar los datos: ${error.message}`;
    }
}

async function selectFiles() {
    try {
        const result = await OpenFileDialog();
        newFilesName.files = result;
    } catch (error) {
        message.value = error;
    }
}

async function addFiles() {
    try {
        const result = await OpenFileDialog();
        newFilesName.files = [...newFilesName.files, ...result];
    } catch (error) {
        message.value = error;
    }
}

function makeNewName() {
    MakeNewName(newFilesName.program, newFilesName.cycle, newFilesName.week, newFilesName.group)
        .then(result => {
            newFilesName.name = result;
        });
}

function changeFileNames() {
    ChangeFileNames(newFilesName.files, newFilesName.name)
        .then(result => {
            message.value = result;
            for (let key in newFilesName) {
                if (key != 'files') {
                    newFilesName[key] = "";
                    continue;
                }
                newFilesName[key] = [];
            }
        })
        .catch(error => {
            message.value = error;
        });
}

function toggleTableVisibility() {
    tableVisibility.value = !tableVisibility.value;
}

function unselectFile(file) {
    newFilesName.files = newFilesName.files.filter(f => f != file);
}

function sendLoadedState() {
    emit('connected', connected.value);
}

initializeForms();
</script>

<template>
    <header>
        <h1>Bienvenido a CS50 Name Changer</h1>
        <p>El programa de <strong>cs50x.ni</strong> para renombrar las imagenes según el formato oficial.</p>
        <div class="message" v-if="!connected">
            <p>Al parecer no estás conectado a la base de datos, por favor, conectate para continuar.</p>
        </div>
        <div class="message">
            <p>{{ message }}</p>
        </div>
    </header>
    <main v-if="!connected" class="container">
        <button class="btn-connect" @click="() => { connect(); sendLoadedState(); }">Conectar</button>
    </main>
    <button class="btn-load" @click="initializeForms" v-if="connected" v-show="!contentLoaded">Cargar datos</button>
    <main v-if="connected && contentLoaded" class="container">
        <div class="forms-container">
            <div class="choices">
                <div class="input-box">
                    <label for="programs-options">Escoge el programa: </label>
                    <select id="programs-options" autocomplete="off" class="input" type="text"
                        v-model="newFilesName.program" @change="(event) => { listCycles(); makeNewName(); }">
                        <option value="" disabled selected v-if="programs.length == 0">Cargando programas...</option>
                        <option value="" disabled selected v-if="programs[0] == 'No hay programas'">No hay programas
                        </option>
                        <option value="" disabled selected v-else>Selecciona un programa</option>
                        <option v-for="(program, index) in programs" :key="index">
                            {{ program }}
                        </option>
                    </select>
                </div>

                <div class="input-box">
                    <label for="cycles-options">Escoge el ciclo: </label>
                    <select id="cycles-options" autocomplete="off" class="input" type="text"
                        v-model="newFilesName.cycle" @change="(event) => { listWeeks(); listGroups(); makeNewName(); }"
                        v-if="cycles[0] != 'No hay ciclos' && newFilesName.program != ''">
                        <option value="" disabled selected v-if="cycles.length == 0">Cargando ciclos...</option>
                        <option value="" disabled selected v-else>Selecciona un ciclo</option>
                        <option v-for="(cycle, index) in cycles" :key="index">
                            {{ cycle }}
                        </option>
                    </select>
                    <select id="cycles-options" autocomplete="off" class="input" type="text"
                        v-model="newFilesName.cycle" disabled v-else-if="newFilesName.program == ''">
                        <option value="" disabled selected>Selecciona un programa</option>
                    </select>
                    <select id="cycles-options" autocomplete="off" class="input" type="text"
                        v-model="newFilesName.cycle" disabled v-else>
                        <option value="" disabled selected>No hay ciclos</option>
                    </select>
                </div>

                <div class="input-box">
                    <label for="weeks-options">Escoge la semana: </label>
                    <select id="weeks-options" autocomplete="off" class="input" type="text" v-model="newFilesName.week"
                        @change="makeNewName"
                        v-if="weeks[0] != 'No hay semanas' && cycles[0] != 'No hay ciclos' && newFilesName.cycle != ''">
                        <option value="" disabled selected v-if="weeks.length == 0">Cargando semanas...</option>
                        <option value="" disabled selected v-else>Selecciona una semana</option>
                        <option v-for="(week, index) in weeks" :key="index">
                            {{ week }}
                        </option>
                    </select>
                    <select id="weeks-options" autocomplete="off" class="input" type="text" v-model="newFilesName.week"
                        disabled v-else-if="newFilesName.cycle == ''">
                        <option value="" disabled selected>Selecciona un ciclo</option>
                    </select>
                    <select id="weeks-options" autocomplete="off" class="input" type="text" v-model="newFilesName.week"
                        disabled v-else>
                        <option value="" disabled selected>No hay semanas</option>
                    </select>
                </div>

                <div class="input-box">
                    <label for="groups-options">Escoge el grupo: </label>
                    <select id="groups-options" autocomplete="off" class="input" type="text"
                        v-model="newFilesName.group" @change="makeNewName"
                        v-if="groups[0] != 'No hay grupos' && weeks[0] != 'No hay semanas' && cycles[0] != 'No hay ciclos'">
                        <option value="" disabled selected v-if="groups.length == 0">Cargando grupos...</option>
                        <option value="" disabled selected v-else>Selecciona un grupo</option>
                        <option v-for="(group, index) in groups" :key="index">
                            {{ group }}
                        </option>
                    </select>
                    <select id="groups-options" autocomplete="off" class="input" type="text"
                        v-model="newFilesName.group" disabled v-else-if="newFilesName.week == ''">
                        <option value="" disabled selected>Selecciona una semana</option>
                    </select>
                    <select id="groups-options" autocomplete="off" class="input" type="text"
                        v-model="newFilesName.group" disabled v-else>
                        <option value="" disabled selected>No hay grupos</option>
                    </select>
                </div>

                <div class="file-selection">
                    <button class="btn-select" @click="selectFiles">
                        <i class="fas fa-folder-open"></i>
                    </button>
                    <button class="btn-add" @click="addFiles" v-if="newFilesName.files.length > 0">
                        <i class="fas fa-plus"></i>
                    </button>
                </div>
            </div>

            <button class="btn-refresh" @click="initializeForms" v-if="connected && contentLoaded">
                <i class="fas fa-sync-alt"></i>
            </button>
        </div>

        <div v-if="newFilesName.files.length > 0">
            <h2>Archivos Seleccionados:</h2>
            <table>
                <thead>
                    <tr>
                        <th>Nombre</th>
                        <th>Acciones</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(file, index) in newFilesName.files" :key="index">
                        <td>{{ file }}</td>
                        <td>
                            <button class="btn btn-danger" @click="() => unselectFile(file)">
                                <i class="fas fa-trash"></i>
                            </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
        <div class="filename">
            <h2>Nuevo nombre:</h2>
            <p>{{ newFilesName.name }}</p>
        </div>
        <div class="actions">
            <button class="btn-change" @click="changeFileNames">Cambiar nombres</button>
        </div>
    </main>
    <footer>
        <p>&copy; 2024 - CS50 Name Changer</p>
    </footer>
</template>

<style scoped>
.forms-container {
    display: flex;
    flex-direction: row;
    gap: 1rem;
    width: 100%;
    justify-content: space-between;
}

.btn-refresh {
    padding: 0.5rem;
    border: none;
    border-radius: 4px;
    background-color: #007bff;
    color: white;
    cursor: pointer;
    transition: background-color 0.3s;
    height: 1%;
}

body {
    background-color: #f8f9fa;
}

header {
    text-align: center;
    margin-bottom: 2rem;
}

.message {
    color: red;
}

.main {
    padding: 1rem;
    background-color: #f8f9fa;
    border-radius: 8px;
}

.choices {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    flex-grow: 1;
}

.input-box {
    display: flex;
    flex-direction: column;
}

.input {
    padding: 0.5rem;
    border-radius: 4px;
    border: 1px solid #ccc;
}

.file-selection {
    display: flex;
    gap: 1rem;
}

.btn-select,
.btn-add,
.btn-remove,
.btn-change,
.btn-connect,
.btn-load,
.btn-refresh {
    padding: 0.5rem 1rem;
    border: none;
    border-radius: 4px;
    background-color: #007bff;
    color: white;
    cursor: pointer;
    transition: background-color 0.3s;
}

.btn-select:hover,
.btn-add:hover,
.btn-remove:hover,
.btn-change:hover,
.btn-connect:hover,
.btn-load:hover,
.btn-refresh:hover {
    background-color: #0056b3;
}

table {
    width: 100%;
    border-collapse: collapse;
}

table th,
table td {
    padding: 0.5rem;
    text-align: left;
    border-bottom: 1px solid #ddd;
}

footer {
    text-align: center;
    margin-top: 2rem;
    font-size: 0.8rem;
    color: #666;
}
</style>
