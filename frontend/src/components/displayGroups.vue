<script setup>
import { ref } from 'vue';
import { WatchGroups, DeleteGroup } from '../../wailsjs/go/main/App';
import ConfirmDialog from './ConfirmDialog.vue';

const groups = ref([]);
const showConfirmDialog = ref(false);
const selectedGroupId = ref(null);

async function watchGroups() {
    const result = await WatchGroups();
    if (!result || result.length === 0 || result[0]?.id === 0) {
        return;
    }
    groups.value = result;
}

function promptDeleteGroup(id) {
    selectedGroupId.value = id;
    showConfirmDialog.value = true;
}

function confirmDeleteGroup() {
    if (selectedGroupId.value !== null) {
        DeleteGroup(selectedGroupId.value);
        groups.value = groups.value.filter(group => group.id !== selectedGroupId.value);
        showConfirmDialog.value = false;
        selectedGroupId.value = null;
    }
}

function cancelDeleteGroup() {
    showConfirmDialog.value = false;
    selectedGroupId.value = null;
}

watchGroups();
</script>

<template>
    <div class="table-container">
        <table v-if="groups.length > 0" class="custom-table">
            <thead>
                <tr>
                    <th hidden>ID</th>
                    <th>Índice</th>
                    <th>Nombre</th>
                    <th>Ciclo</th>
                    <th>Acciones</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(group, index) in groups" :key="index">
                    <td hidden>{{ group.id }}</td>
                    <td>{{ index + 1 }}</td>
                    <td>{{ group.name }}</td>
                    <td>{{ group.cycle_name }}</td>
                    <td>
                        <button class="btn-delete" @click="promptDeleteGroup(group.id)">Eliminar</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <p v-else>No hay grupos</p>
    </div>

    <!-- Confirmation Dialog -->
    <ConfirmDialog
        v-if="showConfirmDialog"
        title="Confirmar eliminación"
        message="¿Estás seguro de que deseas eliminar este grupo?"
        @confirm="confirmDeleteGroup"
        @cancel="cancelDeleteGroup"
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
    background-color: #f1f1f1;
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
