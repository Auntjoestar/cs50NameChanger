<script setup>
import { ref, onMounted } from 'vue';
import { WatchGroups, DeleteGroup } from '../../wailsjs/go/main/App';
import ConfirmDialog from './ConfirmDialog.vue';

const groups = ref([]);
const showConfirmDialog = ref(false);
const selectedGroupId = ref(null);
const isLoading = ref(true); // Loading state

async function watchGroups() {
    try {
        const result = await WatchGroups();
        if (!result || result.length === 0 || result[0]?.id === 0) {
            return;
        }
        groups.value = result;
    } catch (error) {
        console.error("Error fetching groups:", error);
    } finally {
        isLoading.value = false; // Set loading to false once the fetching is done
    }
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

// Fetch groups when the component is mounted
onMounted(() => {
    watchGroups();
});
</script>

<template>
    <h2>Grupos (Total: {{ groups.length }})</h2>
    <div class="table-container" v-if="groups.length > 0 || isLoading">
        <div v-if="isLoading" class="loading-indicator">Cargando...</div>
        <div class="table-wrapper">
            <table v-if="!isLoading && groups.length > 0" class="custom-table">
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
                            <button class="btn-delete" @click="promptDeleteGroup(group.id)"
                                aria-label="Eliminar grupo">Eliminar</button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
    <div v-else class="no-groups">
        <p>No hay grupos</p>
    </div>

    <ConfirmDialog v-if="showConfirmDialog" title="Confirmar eliminación"
        message="¿Estás seguro de que deseas eliminar este grupo?" @confirm="confirmDeleteGroup"
        @cancel="cancelDeleteGroup" />
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

.no-groups {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 90%;
    color: #666;
    font-size: 1.1rem;
}
</style>