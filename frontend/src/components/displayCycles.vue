<script setup>
import { ref } from 'vue';
import { WatchCycles, DeleteCycle } from '../../wailsjs/go/main/App';
import ConfirmDialog from './ConfirmDialog.vue';

const cycles = ref([]);
const showConfirmDialog = ref(false);
const selectedCycleId = ref(null);

async function watchCycles() {
    const result = await WatchCycles();
    if (!result || result.length === 0 || result[0]?.id === 0) {
        return;
    }
    cycles.value = result;
}

function promptDeleteCycle(id) {
    selectedCycleId.value = id;
    showConfirmDialog.value = true;
}

function confirmDeleteCycle() {
    if (selectedCycleId.value !== null) {
        DeleteCycle(selectedCycleId.value);
        cycles.value = cycles.value.filter(cycle => cycle.id !== selectedCycleId.value);
        showConfirmDialog.value = false;
        selectedCycleId.value = null;
    }
}

function cancelDeleteCycle() {
    showConfirmDialog.value = false;
    selectedCycleId.value = null;
}

watchCycles();
</script>

<template>
    <div class="table-container">
        <table v-if="cycles.length > 0" class="custom-table">
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
                    <td>{{ cycle.name }}</td>
                    <td>{{ cycle.program_name }}</td>
                    <td>
                        <button class="btn-delete" @click="promptDeleteCycle(cycle.id)">Eliminar</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <p v-else>No hay ciclos</p>
    </div>

    <ConfirmDialog
        v-if="showConfirmDialog"
        title="Confirmar eliminación"
        message="¿Estás seguro de que deseas eliminar este ciclo?"
        @confirm="confirmDeleteCycle"
        @cancel="cancelDeleteCycle"
    />
</template>

<style scoped>
.table-container {
    margin: 20px;
    overflow-x: auto;
}

.custom-table {
    width: 100%;
    border-collapse: collapse;
    font-family: Arial, sans-serif;
}

.custom-table th, .custom-table td {
    padding: 12px;
    text-align: center;
    border: 1px solid #ddd;
}

.custom-table thead {
    background-color: #343a40;
    color: #fff;
}

.custom-table tbody tr:hover {
    background-color: #f9f9f9;
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
</style>
