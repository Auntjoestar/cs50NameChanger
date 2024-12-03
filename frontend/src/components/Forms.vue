<script setup>
import { ref, reactive, defineEmits, onMounted, onBeforeUnmount } from 'vue';
import {
    CreateDB,
    ListPrograms,
    ListCycles,
    ListWeeks,
    ListGroups,
    MakeNewName,
    OpenFileDialog,
    ChangeFileNames,
    WatchLastCreatedIndex,
} from '../../wailsjs/go/main/App';



const connected = ref(false);
const contentLoaded = ref(false);
const programs = ref([]);
const cycles = ref([]);
const weeks = ref([]);
const groups = ref([]);
const message = ref("");
const messageType = ref("");
const tableVisibility = ref(true);
const emit = defineEmits(['connected']);
const lastCreatedIndex = ref(0);
const index = ref(0);


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

function keydownConnect(e) {
    let shouldPreventDefault = false;

    switch (e.key) {
        case "Enter":
            if (newFilesName.name) {
                if (newFilesName.files.length > 0 && index.value >= lastCreatedIndex.value) {
                    changeFileNames();
                    shouldPreventDefault = true;
                }
            }
            break;
        case "r":
            if (e.ctrlKey) {
                initializeForms();
                shouldPreventDefault = true;
            }
            break;
        case "a":
            if ((e.ctrlKey || e.metaKey) && newFilesName.files.length > 0) {
                addFiles();
                shouldPreventDefault = true;
            }
            break;
        case "s":
            if (e.ctrlKey) {
                selectFiles();
                shouldPreventDefault = true;
            }
            break;
        default:
            break;
    }

    if (shouldPreventDefault) {
        e.preventDefault();
    }
}

function connectEventListeners() {
    const handleKeydown = (e) => {
        switch (connected.value) {
            case true:
                keydownConnect(e);
                break;
            case false:
                if (e.key === "Enter") {
                    handleConnect();
                }
                break;
        }
    };

    window.addEventListener("keydown", handleKeydown);

    onBeforeUnmount(() => {
        window.removeEventListener("keydown", handleKeydown);
    });
}

async function connect() {
    const err = await CreateDB();
    if (err) {
        connected.value = false;
        message.value = `Error al conectar con la base de datos: ${err}`;
        messageType.value = "alert-danger";
        return;
    }
    try {

        try {
            await ListPrograms();
        } catch (error) {
            message.value = `Error al cargar los programas: ${error}`;
            messageType.value = "alert-danger";
        }
        connected.value = true;
        message.value = "Conexión exitosa";
        messageType.value = "alert-success";
    } catch (error) {
        message.value = `Error al conectar con la base de datos: ${error}`;
        messageType.value = "alert-danger";
    }
    connected.value = true;
    localStorage.setItem('connected', connected.value);
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
        const result = await ListWeeks(newFilesName.cycle, newFilesName.program);
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
        const result = await ListGroups(newFilesName.cycle, newFilesName.program);
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
        if (result === null) {
            return;
        }
        newFilesName.files = result;
    } catch (error) {
        message.value = error;
        messageType.value = "alert-danger";
    }
}

async function addFiles() {
    try {
        const result = await OpenFileDialog();
        if (result === null) {
            return;
        }
        newFilesName.files = [...newFilesName.files, ...result.filter(f => !newFilesName.files.includes(f))];
    } catch (error) {
        message.value = error;
        messageType.value = "alert-danger";
    }
}

async function makeNewName() {
    try {
        newFilesName.name = await MakeNewName(newFilesName.program, newFilesName.cycle, newFilesName.week, newFilesName.group);
        if (newFilesName.name !== "Nombre del archivo") {
            const result = await WatchLastCreatedIndex(newFilesName.name);
            lastCreatedIndex.value = parseInt(result) + 1
            index.value = lastCreatedIndex.value;
        }
    } catch (error) {
        message.value = error;
        messageType.value = "alert-danger";
    }
}

function changeFileNames() {
    try {
        ChangeFileNames(newFilesName.files, newFilesName.name, parseInt(index.value));
        newFilesName.files = [];
        message.value = "Nombres cambiados";
        messageType.value = "alert-success";
        newFilesName.name = "";
        newFilesName.program = "";
        newFilesName.cycle = "";
        newFilesName.week = "";
        newFilesName.group = "";
    } catch (error) {
        message.value = error;
        messageType.value = "alert-danger";
    }
}

function toggleTableVisibility() {
    tableVisibility.value = !tableVisibility.value;
    document.querySelector('.custom-table').classList.toggle('collapsed');
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

function handleProgramChange() {
    document.getElementById("cycles-options").selectedIndex = 0;
    document.getElementById("weeks-options").selectedIndex = 0;
    document.getElementById("groups-options").selectedIndex = 0;
    newFilesName.name = "";
    newFilesName.cycle = "";
    newFilesName.week = "";
    newFilesName.group = "";
    listCycles();
    makeNewName();
};

function handleCycleChange() {
    document.getElementById("weeks-options").selectedIndex = 0;
    document.getElementById("groups-options").selectedIndex = 0;
    newFilesName.week = "";
    newFilesName.group = "";
    listWeeks();
    listGroups();
    makeNewName();
};

onMounted(() => {
    connectEventListeners();
    initializeForms();
});
</script>


<template>
    <div class="layout-container">
        <aside class="sidebar">
            <aside class="sidebar">
                <div class="logo">
                    <img src="./../assets/images/logo.svg" alt="CS50 Name Changer Logo" class="logo-image">
                    <h1>CS50 Name Changer</h1>
                </div>
            </aside>

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
        <div class="layout-container-2" v-else>
            <main class="main-container">
                <div class="forms-container">
                    <div class="choices">
                        <div id="message" v-if="message"
                            :class="['alert', messageType, 'alert-dismissible', 'fade', 'show']">
                            {{ message }}
                            <button type="button" class="btn-close" @click="message = ''" aria-label="Close"></button>
                        </div>
                        <div class="input-box">
                            <label for="programs-options">Escoge el programa: </label>
                            <select id="programs-options" autocomplete="off" class="input"
                                v-model="newFilesName.program" @change="handleProgramChange">
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
                            <select id="cycles-options" autocomplete="off" class="input" v-model="newFilesName.cycle"
                                @change="handleCycleChange"
                                v-if="cycles[0] != 'No hay ciclos' && newFilesName.program != ''">
                                <option value="" disabled selected v-if="cycles.length == 0">Cargando ciclos...</option>
                                <option value="" disabled selected v-else>Selecciona un ciclo</option>
                                <option v-for="(cycle, index) in cycles" :key="index">
                                    {{ cycle }}
                                </option>
                            </select>
                            <select id="cycles-options" autocomplete="off" class="input" v-model="newFilesName.cycle"
                                disabled v-else-if="newFilesName.program == ''">
                                <option value="" disabled selected>Selecciona un programa</option>
                            </select>
                            <select id="cycles-options" autocomplete="off" class="input" v-model="newFilesName.cycle"
                                disabled v-else>
                                <option value="" disabled selected>No hay ciclos</option>
                            </select>
                        </div>

                        <div class="input-box">
                            <label for="weeks-options">Escoge la semana: </label>
                            <select id="weeks-options" autocomplete="off" class="input" v-model="newFilesName.week"
                                @change="makeNewName"
                                v-if="weeks[0] != 'No hay semanas' && cycles[0] != 'No hay ciclos' && newFilesName.cycle != ''">
                                <option value="" disabled selected v-if="weeks.length == 0">Cargando semanas...</option>
                                <option value="" disabled selected v-else>Selecciona una semana</option>
                                <option v-for="(week, index) in weeks" :key="index">
                                    {{ week }}
                                </option>
                            </select>
                            <select id="weeks-options" autocomplete="off" class="input" v-model="newFilesName.week"
                                disabled v-else-if="newFilesName.cycle == ''">
                                <option value="" disabled selected>Selecciona un ciclo</option>
                            </select>
                            <select id="weeks-options" autocomplete="off" class="input" v-model="newFilesName.week"
                                disabled v-else>
                                <option value="" disabled selected>No hay semanas</option>
                            </select>
                        </div>

                        <div class="input-box">
                            <label for="groups-options">Escoge el grupo: </label>
                            <select id="groups-options" autocomplete="off" class="input" v-model="newFilesName.group"
                                @change="makeNewName"
                                v-if="groups[0] != 'No hay grupos' && cycles[0] != 'No hay ciclos' && newFilesName.cycle != ''">
                                <option value="" disabled selected v-if="groups.length == 0">Cargando grupos...</option>
                                <option value="" disabled selected v-else>Selecciona un grupo</option>
                                <option v-for="(group, index) in groups" :key="index">
                                    {{ group }}
                                </option>
                            </select>
                            <select id="groups-options" autocomplete="off" class="input" v-model="newFilesName.group"
                                disabled v-else-if="newFilesName.cycle == ''">
                                <option value="" disabled selected>Selecciona un ciclo</option>
                            </select>
                            <select id="groups-options" autocomplete="off" class="input" v-model="newFilesName.group"
                                disabled v-else>
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

                <div v-if="newFilesName.files.length > 0" class="table-container">
                    <h2>Archivos Seleccionados:</h2>
                    <button class="btn-change-visbility btn btn-primary" @click="toggleTableVisibility">
                        <i class="fas fa-eye"></i>
                    </button>
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
                        <button class="btn-change" @click="changeFileNames"
                            :disabled="!newFilesName.name || newFilesName.name === 'Nombre del archivo' || index < lastCreatedIndex">
                            Cambiar nombres
                        </button>
                        <div class="select-index">
                            <label for="index">Índice de inicio:</label>
                            <div class="index">
                            <input type="number" v-model="index" :min=lastCreatedIndex
                                v-if="newFilesName.name && newFilesName.name !== 'Nombre del archivo'"
                                title="Índice de inicio" id="index" class="input" />
                            <i class="fas fa-info-circle"
                                title="Él índice comenzara con respecto al último archivo creado"
                                data-bs-toggle="tooltip" data-bs-placement="top" data-bs-trigger="hover focus"></i>
                            </div>
                        </div>
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
    background-color: #282c34;
    color: #fff;
    display: flex;
    flex-direction: column;
    align-items: center;
    box-shadow: 2px 0 10px rgba(0, 0, 0, 0.2);

}

.actions {
    display: flex;
    gap: 2rem;
    justify-content: center;
    align-items: flex-start;

    button {
        height: 100%;
    }

    
    .btn-change {
                align-self: flex-end !important;
            }

    .select-index {
        width: 20%;
        display: flex;
        flex-direction: column;
        gap: 0.5rem;

        .index {
            display: flex;
            gap: 0.5rem;
            align-items: center;
            justify-content: center;
            input {
                width: 40%;
            }
        }
    }
}

.logo {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin-bottom: 2rem;
    padding: 1rem;
    height: 100vh;
}

.logo-image {
    width: 70%;
    max-width: 150px;
    border-radius: 50%;
    margin-bottom: 1rem;
}

.logo-image:hover {
    animation: rotation 2s ease-out;
}

@keyframes rotation {
    0% {
        transform: rotate(0deg);
    }

    50% {
        transform: rotate(150deg);
    }

    100% {
        transform: rotate(360deg);
    }
}


.logo h1 {
    font-size: 1.5rem;
    font-weight: bold;
    text-transform: uppercase;
}

.sidebar-nav {
    width: 100%;
}

.sidebar-nav ul {
    list-style: none;
    padding: 0;
    margin: 0;
}

.sidebar-nav ul li {
    margin: 1rem 0;
}

.sidebar-nav ul li a {
    text-decoration: none;
    font-size: 1rem;
    color: #ddd;
    padding: 0.5rem 1rem;
    border-radius: 8px;
    display: block;
    transition: background-color 0.3s, color 0.3s;
}

.sidebar-nav ul li a:hover {
    background-color: #61dafb;
    color: #282c34;
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

.btn-change:disabled {
    background-color: #ccc;
    cursor: not-allowed;
}

.btn-change:disabled:hover {
    background-color: #ccc;
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
    padding: 1rem;
    ;
    width: 100%;
    position: relative;
    /* Default positioning */
    bottom: 0;
    /* Ensures footer is at the bottom */
    flex-shrink: 0;
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
    transition: max-height 0.5s ease-in-out;
    max-height: 1000px;

    th,
    td {
        padding: 12px;
        text-align: center;
        border: 1px solid #ddd;
        border-top: none;
        text-align: left;
        text-wrap: normal;
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

.custom-table.collapsed {
    max-height: 0;
}

.btn-change-visbility {
    width: fit-content;
    background-color: transparent;
    border: none;
    color: #007bff;
    margin-right: 100%;
}




.layout-container {
    display: flex;
    flex-direction: row;
    justify-content: flex-start;
    gap: 1rem;
    height: 100vh;
}


.main-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    background-color: #f8f9fa;
    flex-grow: 1;
    width: 100%;
    max-height: 90vh;
    overflow-y: auto;
    padding: 1rem;
    margin: auto;
}

.layout-container-2 {
    width: 100%;
}

.table-container {
    display: flex;
    flex-direction: column;
    gap: 1rem;
}

@media (max-width: 768px) {

    .layout-container {
        flex-direction: column;
        /* Stack elements vertically on smaller screens */
    }

    .sidebar {
        width: 100%;
        /* Full-width for mobile */
        height: auto;
        padding: 10px;
    }

    .main-container {
        /* 75% width for better readability */
        max-height: none;
        /* Remove max-height for better responsiveness */
    }

    .name-para {
        margin-top: 1rem;
    }

    .custom-table {
        width: 100%;
        border-collapse: collapse;
        font-family: Arial, sans-serif;
        transition: max-height 0.5s ease-in-out;
        max-height: 1000px;
        margin-bottom: 1rem;
    }

    .custom-table th,
    .custom-table td {
        padding: 12px;
        text-align: center;
        border: 1px solid #ddd;
        border-top: none;
        text-align: left;
        word-wrap: break-word;
        word-break: break-word;
        /* Allows text to wrap inside cells */
    }

    .custom-table th {
        top: 0;
        border-top: none;
    }

    .custom-table thead th {
        background-color: #343a40;
        color: #fff;
        position: sticky;
        top: 0;
        z-index: 1;
    }

    .custom-table tbody tr:hover {
        background-color: #f1f1f1;
    }

    /* Make the table scrollable horizontally on small screens */
    @media screen and (max-width: 768px) {
        .custom-table {
            display: block;
            width: 100%;
            overflow-x: auto;
            -webkit-overflow-scrolling: touch;
            /* Smooth scrolling on iOS */
        }

        .custom-table th,
        .custom-table td {
            padding: 8px;
            /* Reduced padding on smaller screens */
        }
    }

    @media screen and (max-width: 480px) {

        .custom-table th,
        .custom-table td {
            font-size: 12px;
            /* Smaller font size for very small screens */
        }
    }

}

@media (min-width: 990px) {

    .sidebar {
        width: 280px;

        h1 {
            font-size: 2rem;
        }
    }

    .main-container {
        width: calc(100% - 200px);
    }
}
</style>
