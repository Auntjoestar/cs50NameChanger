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
        <div class="message">
            <p>Al parecer no estás conectado a la base de datos, por favor, conectate para continuar.</p>
        </div>
        <div class="message">
            <p>{{ message }}</p>
        </div>
    </header>
    <main v-if="!connected">
        <button @click="() => { connect(); sendLoadedState(); }">Conectar</button>
    </main>
    <button @click="initializeForms" v-if="connected" v-show="!contentLoaded">Cargar datos</button>
    <button @click="initializeForms" v-if="connected && contentLoaded">Recargar datos</button>
    <main v-if="connected && contentLoaded">
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
                <select id="cycles-options" autocomplete="off" class="input" type="text" v-model="newFilesName.cycle"
                    @change="(event) => { listWeeks(); listGroups(); makeNewName(); }">
                    <option value="" disabled selected v-if="cycles.length == 0">Cargando ciclos...</option>
                    <option value="" disabled selected v-if="cycles[0] == 'No hay ciclos'">No hay ciclos</option>
                    <option value="" disabled selected v-else>Selecciona un ciclo</option>
                    <option v-for="(cycle, index) in cycles" :key="index">
                        {{ cycle }}
                    </option>
                </select>
            </div>

            <div class="input-box">
                <label for="weeks-options">Escoge la semana: </label>
                <select id="weeks-options" autocomplete="off" class="input" type="text" v-model="newFilesName.week"
                    @change="makeNewName">
                    <option value="" disabled selected v-if="weeks.length == 0">Cargando semanas...</option>
                    <option value="" disabled selected v-else>Selecciona una semana</option>
                    <option v-for="(week, index) in weeks" :key="index">
                        {{ week }}
                    </option>
                </select>
            </div>

            <div class="input-box">
                <label for="groups-options">Escoge el grupo: </label>
                <select id="groups-options" autocomplete="off" class="input" type="text" v-model="newFilesName.group"
                    @change="makeNewName">
                    <option value="" disabled selected v-if="groups.length == 0">Cargando grupos...</option>
                    <option value="" disabled selected v-else>Selecciona un grupo</option>
                    <option v-for="(group, index) in groups" :key="index">
                        {{ group }}
                    </option>
                </select>
            </div>

            <div class="input-box">
                <input id="new-files-name" :value="newFilesName.name" autocomplete="off" class="input" type="text"
                    placeholder="Nombre del archivo" disabled />
            </div>

            <div class="input-box">
                <label for="select-files">Selecciona las imagenes: </label>
                <button id="select-files" @click="selectFiles"><i class="fa-regular fa-image"></i></button>

                <button @click="toggleTableVisibility" v-if="newFilesName.files.length > 0">Mostrar archivos
                    seleccionados</button>
                <div class="table" v-if="newFilesName.files.length > 0" v-show="tableVisibility">
                    <h3>Archivos seleccionados:</h3>
                    <div class="table-actions">
                        <button @click="newFilesName.files = []">Limpiar</button>
                        <button @click="addFiles">Agregar</button>
                    </div>
                    <table>
                        <thead>
                            <tr>
                                <th>Índice</th>
                                <th>Nombre</th>
                                <th>Acción</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(file, index) in newFilesName.files" :key="index">
                                <td>{{ index + 1 }}</td>
                                <td>{{ file }}</td>
                                <td><button @click="unselectFile(file)">Eliminar</button></td>
                            </tr>
                        </tbody>
                    </table>
                </div>
            </div>

            <div class="input-box">
                <!-- 
                    "Nombre del archivo" porque la función MakeNewName retorna la previa cadena cuando no se han seleccionado dato 
                     en todos los formularios
                 -->
                <button @click="changeFileNames" v-if="newFilesName.files.length > 0 && newFilesName.name != ''
                    && newFilesName.name != 'Nombre del archivo'">Cambiar nombres</button>
            </div>
        </div>
    </main>
</template>

<style lang="scss">
body {
    background-color: #f2f2f2;
}

main {
    margin: auto;
    width: 90%;
}

table {
    border-collapse: collapse;
    width: 80%;
}

th,
td {
    border: 1px solid black;
    padding: 8px;
    text-align: left;
}

tbody tr:nth-child(even) {
    background-color: #f2f2f2;
}

.table {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: start;
    width: 100% !important;

    h3 {
        margin-left: 1%;
        align-self: start;
    }


}

.choices {
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
}

.input-box {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    margin: 1%;
}

tbody {
    width: 100%;
}
</style>
