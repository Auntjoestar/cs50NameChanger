<script setup>
import { ref, onMounted, defineEmits } from 'vue';
import { WatchGroups, DeleteGroup, EditGroup } from '../../wailsjs/go/main/App';
import ConfirmDialog from './ConfirmDialog.vue';

const groups = ref([]);
const showConfirmDialog = ref(false);
const selectedGroupId = ref(null);
const isLoading = ref(true);
const editIndex = ref(null);
const editedName = ref("");
const emits = defineEmits(['group-edited', 'error']);

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
        EditGroup(id, editedName.value.trim()).then(() => {
            const group = groups.value.find(group => group.id === id);
            group.name = editedName.value.trim();
            editIndex.value = null;
            editedName.value = '';
            emits('group-edited');
        }).catch(error => {
            error = error.charAt(0).toUpperCase() + error.slice(1);
            console.error("Error editing group:", error);
            emits('error', error);
        });
        return;
    }
    emits('error', 'El nombre del grupo no puede estar vacío');
}

function isSavingDisabled(group) {
    return !editedName.value.trim() || editedName.value.trim() === group;
}

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
                        <td>
                            <div v-if="editIndex === index">
                                <input type="text" v-model="editedName" class="form-control" />
                                <button class="btn btn-success" @click="saveEdit(group.id)"
                                    :disabled="isSavingDisabled(group.name)">
                                    <i class="fas fa-save"></i>
                                </button>
                            </div>
                            <div v-else>
                                {{ group.name }}
                            </div>
                        </td>
                        <td>{{ group.cycle_name }}</td>
                        <td>
                            <button class="btn btn-primary" @click="startEdit(index, group.name)"
                                v-if="editIndex !== index">
                                <i class="fas fa-edit"></i>
                            </button>
                            <button class="btn btn-primary " @click="cancelEdit()" v-else>
                                <i class="fas fa-times"></i>
                            </button>
                            <button class="btn btn-danger" @click="promptDeleteGroup(group.id)">
                                <i class="fas fa-trash-alt"></i>
                            </button>
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
h2 {
    margin-bottom: 1.2rem;
}

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