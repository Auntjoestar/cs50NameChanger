<script setup>
import { ref, onMounted, defineEmits } from 'vue';
import { WatchPrograms, DeleteProgram, EditProgram } from '../../wailsjs/go/main/App';
import ConfirmDialog from './ConfirmDialog.vue';

const programs = ref([]);
const isLoading = ref(true);
const showConfirmDialog = ref(false);
const selectedProgramId = ref(null);
const editIndex = ref(null); // Track currently edited index
const editedName = ref(""); // Store the value of the edited name
const emits = defineEmits(['program-edited', 'error']);

async function watchPrograms() {
    try {
        const result = await WatchPrograms();
        if (result && result.length > 0 && result[0]?.id !== 0) {
            programs.value = result;
        }
    } catch (error) {
        console.error("Error fetching programs:", error);
    } finally {
        isLoading.value = false;
    }
}

function promptDeleteProgram(id) {
    selectedProgramId.value = id;
    showConfirmDialog.value = true;
}

function confirmDeleteProgram() {
    if (selectedProgramId.value !== null) {
        DeleteProgram(selectedProgramId.value);
        programs.value = programs.value.filter(program => program.id !== selectedProgramId.value);
        showConfirmDialog.value = false;
        selectedProgramId.value = null;
    }
}

function cancelDeleteProgram() {
    showConfirmDialog.value = false;
    selectedProgramId.value = null;
}

function startEdit(index, name) {
    editIndex.value = index;
    editedName.value = name;
}

function cancelEdit() {
    editIndex.value = null;
    editedName.value = '';
}

function saveEdit(id) {
    if (editedName.value.trim()) {
        EditProgram(id, editedName.value.trim()).then(() => {
            const program = programs.value.find(program => program.id === id);
            if (program) {
                program.name = editedName.value.trim();
            }
            cancelEdit();
            emits('program-edited');
        }).catch(error => {
            error = error.charAt(0).toUpperCase() + error.slice(1);
            console.error("Error actualizando el programa:", error);
            emits('error', error);
        });
        return;
    }
    emits('error', 'El nombre del programa no puede estar vacío');
}

function isSavingDisabled(program) {
    return !editedName.value.trim() || editedName.value.trim() === program;
}

onMounted(() => {
    watchPrograms();
});
</script>


<template>
    <h2>Programas (Total: {{ programs.length }})</h2>
    <div class="table-container" v-if="programs.length > 0 || isLoading">
        <div v-if="isLoading" class="loading-indicator">Cargando...</div>
        <div class="table-wrapper">
            <table v-if="!isLoading && programs.length > 0" class="custom-table">
                <thead>
                    <tr>
                        <th hidden>ID</th>
                        <th>Índice</th>
                        <th>Nombre</th>
                        <th>Acciones</th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="(program, index) in programs" :key="program.id">
                        <td hidden>{{ program.id }}</td>
                        <td>{{ index + 1 }}</td>
                        <td class="name">
                            <div v-if="editIndex === index">
                                <input type="text" v-model="editedName" class="form-control" />
                                <button @click="saveEdit(program.id)" class="btn btn-success"
                                    :disabled="isSavingDisabled(program.name)">
                                    <i class="fas fa-check"></i>
                                </button>
                            </div>
                            <div v-else>
                                {{ program.name }}
                            </div>
                        </td>
                        <td class="actions">
                            <button class="btn btn-primary" @click="startEdit(index, program.name)"
                                v-if="editIndex !== index">
                                <i class="fas fa-edit"></i>
                            </button>
                            <button class="btn btn-primary " @click="cancelEdit()" v-else>
                                <i class="fas fa-times"></i>
                            </button>
                            <button class="btn btn-danger" @click="promptDeleteProgram(program.id)">
                                <i class="fas fa-trash-alt"></i>
                            </button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>

    </div>
    <div class="no-programs" v-else>
        <p>No hay programas disponibles.</p>
    </div>

    <ConfirmDialog v-if="showConfirmDialog" title="Confirmar eliminación"
        message="¿Estás seguro de que deseas eliminar este programa?" @confirm="confirmDeleteProgram"
        @cancel="cancelDeleteProgram" />
</template>

<style scoped lang="scss">
h2 {
    margin-bottom: 1.2rem;
}

.table-wrapper {
    width: 100%;
    max-height: 400px;
    overflow-y: auto;
    border: 1px solid #ddd;
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

.btn-delete {
    padding: 6px 12px;
    background-color: #dc3545;
    border: none;
    color: #fff;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s ease;
}

.btn-delete:hover {
    background-color: #c82333;
}

.name {
    input {
        display: inline-block;
        width:80%;
        padding: 0.5rem;
        border: 1px solid #ccc;
        border-radius: 4px;
        margin-right: 0.5rem;
    }
}

.actions {
    button {
        margin: 0 0.2rem;
    }
}

.loading-indicator {
    text-align: center;
    font-size: 1.2rem;
    color: #007bff;
    padding: 20px;
}

.no-programs {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 90%;
    color: #666;
    font-size: 1.1rem;
}

@media (max-width: 670px) {
.name {
    input {
        width: 60%;
    }
}


.actions {
    button {
        margin-bottom: 0.5rem;
    }
}
}
</style>
