<script setup>
import { ref } from 'vue';
import { WatchPrograms, DeleteProgram } from '../../wailsjs/go/main/App';
import ConfirmDialog from './ConfirmDialog.vue';

const programs = ref([]);
const isConfirmVisible = ref(false);
const programIdToDelete = ref(null);

async function watchPrograms() {
    const result = await WatchPrograms();
    if (!result || result.length === 0 || result[0]?.id === 0) {
        return;
    }
    programs.value = result;
}

function openConfirmDialog(id) {
    programIdToDelete.value = id;
    isConfirmVisible.value = true;
}

function handleConfirm() {
    DeleteProgram(programIdToDelete.value);
    watchPrograms();
    closeConfirmDialog();
}

function closeConfirmDialog() {
    isConfirmVisible.value = false;
    programIdToDelete.value = null;
}

watchPrograms();
</script>

<template>
    <div class="table-container">
        <table v-if="programs.length > 0" class="custom-table">
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
                        <button class="btn-delete" @click="openConfirmDialog(program.id)">Eliminar</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <p v-else>No hay programas</p>
    </div>

    <ConfirmDialog
        v-if="isConfirmVisible"
        :title="'Confirmación de Eliminación'"
        :message="'¿Estás seguro de que deseas eliminar este programa?'"
        :isVisible="isConfirmVisible"
        @confirm="handleConfirm"
        @cancel="closeConfirmDialog"
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
