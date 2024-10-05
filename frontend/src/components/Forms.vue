<script setup>
import { ref, reactive } from 'vue';
import { ListPrograms, ListCycles, ListWeeks, ListGroups, MakeNewName, OpenFileDialog, ChangeFileNames } from '../../wailsjs/go/main/App';

const programs = ref([]);
const cycles = ref([]);
const weeks = ref([]);
const groups = ref([]);
const message = ref("");

const newFilesName = reactive({
    name: "",
    program: "",
    cycle: "",
    week: "",
    group: "",
    files: [],
});

async function listPrograms() {
    const result = await ListPrograms();
    programs.value = result;
}

async function listCycles() {
    const result = await ListCycles();
    cycles.value = result;
}

async function listWeeks() {
    const result = await ListWeeks();
    weeks.value = result;
}

async function listGroups() {
    const result = await ListGroups();
    groups.value = result;
}

async function selectFiles() {
    try {
        const result = await OpenFileDialog();
        newFilesName.files= result;
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
        })
        .catch(error => {
            message.value = error;
        });
}

listPrograms();
listCycles();
listWeeks();
listGroups();
</script>

<template>
    <main>
        <div class="message">
            <p>{{ message }}</p>
        </div>
        <div class="choices">
            <div class="input-box">
                <label for="programs-options">Escoge el programa: </label>
                <select id="programs-options" autocomplete="off" class="input" type="text"
                    v-model="newFilesName.program" @change="makeNewName">
                    <option value="" disabled selected>Selecciona un programa</option>
                    <option v-for="(program, index) in programs" :key="index">
                        {{ program }}
                    </option>
                </select>
            </div>

            <div class="input-box">
                <label for="cycles-options">Escoge el ciclo: </label>
                <select id="cycles-options" autocomplete="off" class="input" type="text" v-model="newFilesName.cycle"
                    @change="makeNewName">
                    <option value="" disabled selected>Selecciona un ciclo</option>
                    <option v-for="(cycle, index) in cycles" :key="index">
                        {{ cycle }}
                    </option>
                </select>
            </div>

            <div class="input-box">
                <label for="weeks-options">Escoge la semana: </label>
                <select id="weeks-options" autocomplete="off" class="input" type="text" v-model="newFilesName.week"
                    @change="makeNewName">
                    <option value="" disabled selected>Selecciona una semana</option>
                    <option v-for="(week, index) in weeks" :key="index">
                        {{ week }}
                    </option>
                </select>
            </div>

            <div class="input-box">
                <label for="groups-options">Escoge el grupo: </label>
                <select id="groups-options" autocomplete="off" class="input" type="text" v-model="newFilesName.group"
                    @change="makeNewName">
                    <option value="" disabled selected>Selecciona un grupo</option>
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

                <div class="table" v-if="newFilesName.files.length > 0">
                    <h3>Archivos seleccionados:</h3>
                    <table>
                        <thead>
                            <tr>
                                <th>Índice</th>
                                <th>Nombre</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="(file, index) in newFilesName.files" :key="index">
                                <td>{{ index + 1 }}</td>
                                <td>{{ file }}</td>
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
                <button @click="changeFileNames" v-if="newFilesName.files.length > 0 && newFilesName.name != '' && newFilesName.name != 'Nombre del archivo'">Cambiar
                    nombres</button>
            </div>
        </div>
    </main>
</template>

<style lang="scss">
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
