<script setup>
import { ref, onMounted, defineEmits } from 'vue';
import { WatchCycles, DeleteCycle, EditCycle } from '../../wailsjs/go/main/App';
import ConfirmDialog from './ConfirmDialog.vue';

const cycles = ref([]);
const showConfirmDialog = ref(false);
const selectedCycleId = ref(null);
const isLoading = ref(true); // Loading state
const errorMessage = ref(null); // Error handling
const editIndex = ref(null)
const editedName = ref("")
const emits = defineEmits(['cycle-edited', 'error']);


async function watchCycles() {
    try {
        const result = await WatchCycles();
        if (!result || result.length === 0 || result[0]?.id === 0) {
            cycles.value = [];
            errorMessage.value = null; // Reset error message
            return;
        }
        cycles.value = result;
    } catch (error) {
        console.error("Error fetching cycles:", error);
        errorMessage.value = "Error al cargar los ciclos. Por favor, inténtelo de nuevo más tarde."; // User-friendly error message
    } finally {
        isLoading.value = false; // Set loading to false once fetching is done
    }
}

function promptDeleteCycle(id) {
    selectedCycleId.value = id;
    showConfirmDialog.value = true;
}

async function confirmDeleteCycle() {
    if (selectedCycleId.value !== null) {
        try {
            await DeleteCycle(selectedCycleId.value); // Awaiting deletion
            cycles.value = cycles.value.filter(cycle => cycle.id !== selectedCycleId.value);
            showConfirmDialog.value = false;
            selectedCycleId.value = null;
            emits('cycle-edited');
        } catch (error) {
            console.error("Error deleting cycle:", error);
            emits('error', error);
        }
    }
}

function cancelDeleteCycle() {
    showConfirmDialog.value = false;
    selectedCycleId.value = null;
}

function startEdit(index, name) {
    editIndex.value = index;
    editedName.value = name
}

function cancelEdit() {
    editIndex.value = null;
    editedName.value = '';
}

async function saveEdit(id) {
    if (editedName.value.trim()) {
        try {
            await EditCycle(id, editedName.value.trim());
            const cycle = cycles.value.find(cycle => cycle.id === id);
            if (cycle) {
                cycle.name = editedName.value.trim();
            }
            cancelEdit();
            emits('cycle-edited');
        } catch (error) {
            error = error.charAt(0).toUpperCase() + error.slice(1);
            console.error("Error updating cycle:", error);
            emits('error', error);
        }
        return;
    }
    emits('error', 'El nombre del ciclo no puede estar vacío');
}

function isSavingDisabled(program) {
    return !editedName.value.trim() || editedName.value.trim() === program;
}

onMounted(() => {
    watchCycles();
});
</script>

<template>
    <h2>Ciclos (Total: {{ cycles.length }})</h2>
    <div class="table-container" v-if="cycles.length > 0 || isLoading || errorMessage">
        <div v-if="isLoading" class="loading-indicator">Cargando...</div>
        <table v-else-if="cycles.length > 0" class="custom-table">
            <thead>
                <tr>
                    <th scope="col" hidden>ID</th>
                    <th scope="col">Índice</th>
                    <th scope="col">Nombre</th>
                    <th scope="col">Programa</th>
                    <th scope="col">Acciones</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(cycle, index) in cycles" :key="cycle.id">
                    <td hidden>{{ cycle.id }}</td>
                    <td>{{ index + 1 }}</td>
                    <td>
                        <div v-if="editIndex === index">
                            <input type="text" v-model="editedName" class="form-control" />
                            <button class="btn btn-success" @click="saveEdit(cycle.id)"
                                :disabled="isSavingDisabled(cycle.name)">
                                <i class="fas fa-check"></i>
                            </button>
                        </div>
                        <div v-else>
                            {{ cycle.name }}
                        </div>
                    </td>
                    <td>{{ cycle.program_name }}</td>
                    <td>
                        <button class="btn btn-primary" @click="startEdit(index, cycle.name)"
                            v-if="editIndex !== index"><i class="fas fa-edit"></i>
                        </button>
                        <button class="btn btn-primary" @click="cancelEdit()" v-else>
                            <i class="fas fa-times"></i>
                        </button>
                        <button class="btn btn-danger" @click="promptDeleteCycle(cycle.id)">
                            <i class="fas fa-trash-alt"></i>
                        </button>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
    <div class="no-cycles" v-else>
        <p>No hay ciclos</p>
    </div>

    <ConfirmDialog v-if="showConfirmDialog" title="Confirmar eliminación"
        message="¿Estás seguro de que deseas eliminar este ciclo?" @confirm="confirmDeleteCycle"
        @cancel="cancelDeleteCycle" />
</template>

<style scoped lang="scss">
h2 {
    margin-bottom: 1.2rem;
}

.table-container {
    max-height: 400px;
    overflow-y: auto;
    margin: 1%;
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
        background-color: #f9f9f9;
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

    &:hover {
        background-color: #c82333;
    }
}

#error-message {
    color: #dc3545;
    /* Red color for error messages */
    text-align: center;
    margin: 10px 0;
}

.loading-indicator {
    text-align: center;
    font-size: 1.2rem;
    color: #007bff;
    padding: 20px;
}

.no-cycles {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 90%;
    color: #666;
    font-size: 1.1rem;
}
</style>
