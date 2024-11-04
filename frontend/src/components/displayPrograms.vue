<script setup>
import { ref, onMounted, onUnmounted } from 'vue';
import { WatchPrograms, DeleteProgram } from '../../wailsjs/go/main/App'; // Assuming these are your actual imports
import ConfirmDialog from './ConfirmDialog.vue';

const programs = ref([]);
const showConfirmDialog = ref(false);
const selectedProgramId = ref(null);
const isLoading = ref(true); // Loading state

async function watchPrograms() {
    try {
        const result = await WatchPrograms();
        if (result && result.length > 0 && result[0]?.id !== 0) {
            programs.value = result;
        }
    } catch (error) {
        console.error("Error fetching programs:", error);
    } finally {
        isLoading.value = false; // Set loading to false once the fetching is done
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

watchPrograms();
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
                        <td>{{ program.name }}</td>
                        <td>
                            <button class="btn-delete" @click="promptDeleteProgram(program.id)"
                                aria-label="Eliminar programa">Eliminar</button>
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
.table-container {
    margin: 1%;
}

.table-wrapper {
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
</style>
