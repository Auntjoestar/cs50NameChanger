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
const messageType = ref("")
const tableVisibility = ref(true);
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

window.addEventListener("keydown", (e) => {
    switch (connected.value) {
        case true:
            if (e.key === "Enter") {
                if (newFilesName.name) {
                    changeFileNames();
                }
            }
            if (e.ctrlKey && e.key === "r") {
                initializeForms();
            }
            break;
        case false:
            if (e.key === "Enter") {
                handleConnect();
            }
            break;
    }
});



window.addEventListener
async function connect() {
    try {
        try {
            await CreateDB();
        } catch (error) {
            message.value = `Error al crear la base de datos: ${error}`;
            messageType.value = "alert-danger";
        }
        try {
            await ListPrograms();
        } catch (error) {
            message.value = `Error al cargar los programas: ${error}`;
            messageType.value = "alert-danger";
        }
        connected.value = true;
        message.value = "Conexión exitosa";
        messageType.value = "alert-success"

    } catch (error) {
        message.value = `Error al conectar con la base de datos: ${error}`;
        messageType.value = "alert-danger";
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
        messageType.value = "alert-danger";
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
        messageType.value = "alert-danger";
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
        messageType.value = "alert-danger";
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
        message.value = "Datos cargados";
        messageType.value = "alert-success";
    } catch (error) {
        message.value = `Error al cargar los datos: ${error.message}`;
        messageType.value = "alert-danger";
    }
}

async function selectFiles() {
    try {
        const result = await OpenFileDialog();
        newFilesName.files = result;
    } catch (error) {
        message.value = error;
        messageType.value = "alert-danger";
    }
}

async function addFiles() {
    try {
        const result = await OpenFileDialog();
        newFilesName.files = [...newFilesName.files, ...result];
    } catch (error) {
        message.value = error;
        messageType.value = "alert-danger";
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
            messageType.value = "alert-success";
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
            messageType.value = "alert-danger";
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

function handleConnect() {
    connect();
    sendLoadedState();
}

initializeForms();
</script>

<template>
    <div class="layout-container">
        <aside class="sidebar">
            <h1>
                CS50 Name Changer
            </h1>
        </aside>

        <div v-if="!connected || !contentLoaded" class="not-loaded">
            <main v-if="!connected" class="container">
                <div class="alert alert-info alert-dismissible fade show" role="alert">
                    Al parecer no estás conectado a la base de datos, por favor, conectate para continuar.
                    <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
                </div>
                <button class="btn-connect" @click="handleConnect">Conectar</button>
            </main>
            <footer>
                <p>&copy; 2024 - CS50 Name Changer</p>
            </footer>
            <button class="btn-load" @click="initializeForms" v-if="connected" v-show="!contentLoaded">Cargar
                datos</button>
        </div>
        <div v-else>
            <main class="main-container">
                <div id="message" v-if="message" :class="['alert', messageType, 'alert-dismissible', 'fade', 'show']">
                    {{ message }}
                    <button type="button" class="btn-close" @click="message = ''" aria-label="Close"></button>
                </div>
                <div class="forms-container">
                    <div class="choices">
                        <div class="input-box">
                            <label for="programs-options">Escoge el programa: </label>
                            <select id="programs-options" autocomplete="off" class="input" type="text"
                                v-model="newFilesName.program" @change="(event) => { listCycles(); makeNewName(); }">
                                <option value="" disabled selected v-if="programs.length == 0">Cargando programas...
                                </option>
                                <option value="" disabled selected v-if="programs[0] == 'No hay programas'">No hay
                                    programas
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
                                v-model="newFilesName.cycle"
                                @change="(event) => { listWeeks(); listGroups(); makeNewName(); }"
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
                            <select id="weeks-options" autocomplete="off" class="input" type="text"
                                v-model="newFilesName.week" @change="makeNewName"
                                v-if="weeks[0] != 'No hay semanas' && cycles[0] != 'No hay ciclos' && newFilesName.cycle != ''">
                                <option value="" disabled selected v-if="weeks.length == 0">Cargando semanas...</option>
                                <option value="" disabled selected v-else>Selecciona una semana</option>
                                <option v-for="(week, index) in weeks" :key="index">
                                    {{ week }}
                                </option>
                            </select>
                            <select id="weeks-options" autocomplete="off" class="input" type="text"
                                v-model="newFilesName.week" disabled v-else-if="newFilesName.cycle == ''">
                                <option value="" disabled selected>Selecciona un ciclo</option>
                            </select>
                            <select id="weeks-options" autocomplete="off" class="input" type="text"
                                v-model="newFilesName.week" disabled v-else>
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

                        <div class="input-box" v-if="newFilesName.name">
                            <label for="new-name">Nombre del archivo: </label>
                            <p class="name-para"><strong>{{ newFilesName.name }}</strong></p>
                        </div>

                        <div class="input-box" v-else>
                            <label for="complete">Nombre del archivo: </label>
                            <p class="name-para"><strong id="complete">Complete todos los campos</strong></p>
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
                    <table class="custom-table" v-if="tableVisibility">
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
                    <div class="actions">
                        <button class="btn-change" @click="changeFileNames" v-if="newFilesName.name">
                            Cambiar nombres
                        </button>
                    </div>
                </div>
            </main>
            <footer>
                <p>&copy; 2024 - CS50 Name Changer</p>
            </footer>
        </div>
    </div>
</template>

<style scoped lang="scss">
.layout-container {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    gap: 1rem;
    height: 100vh;
}

.sidebar {
    width: 200px;
    height: 100vh;
    background-color: #333;
    padding-top: 20px;
    color: white;
    text-align: center;
    overflow-y: auto;
}

.forms-container {
    display: flex;
    gap: 1rem;
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

.container {
    display: flex;
    flex-direction: column;
    width: 100%;
    gap: 1rem;
    padding: 1rem;
    background-color: #f8f9fa;
    border-radius: 8px;
    flex-grow: 1;
}


header {
    text-align: center;
    margin-bottom: 2rem;
}



#message {
    text-align: start;
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

    label {
        text-align: start;
        font-weight: 500;
    }
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

    th,
    td {
        padding: 0.5rem;
        text-align: left;
        border-bottom: 1px solid #ddd;
    }
}

footer {
    text-align: center;
    font-size: 0.8rem;
    color: #666;
}

.not-loaded {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 1rem;
    flex-grow: 1;
}

.main-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    background-color: #f8f9fa;
    flex-grow: 1;
    width: 75vw;
    max-height: 90vh;
    overflow-y: auto;
    padding: 1rem;
    margin: auto;
}

#new-name {
    margin-top: 1rem;
}

.name-para {
    margin-top: 1rem;
}

.custom-table {
    width: 100%;
    border-collapse: collapse;
    font-family: Arial, sans-serif;

    th,
    td {
        padding: 12px;
        text-align: center;
        border: 1px solid #ddd;
        border-top: none;

    }

    th {
        top: 0;
        border-top: none;
    }

    thead th {
        background-color: #343a40;
        color: #fff;
        position: sticky;
        top: 0;
        z-index: 1;
    }

    tbody tr:hover {
        background-color: #f1f1f1;
    }
}
</style>
