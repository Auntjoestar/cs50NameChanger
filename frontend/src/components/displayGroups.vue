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
    <table v-if="groups.length > 0" class="table table-striped table-hover table-bordered">
        <thead class="table-dark">
            <tr>
                <th scope="col" hidden>ID</th>
                <th scope="col">Índice</th>
                <th scope="col">Nombre</th>
                <th scope="col">Ciclo</th>
                <th scope="col">Acciones</th>
            </tr>
        </thead>
        <tbody class="text-center">
            <tr v-for="(group, index) in groups" :key="index" class="text-center">
                <td scope="row" hidden>{{ group.id }}</td>
                <td>{{ index + 1 }}</td>
                <td>{{ group.name }}</td>
                <td>{{ group.cycle_name }}</td>
                <td>
                    <button class="btn btn-danger" @click="promptDeleteGroup(group.id)">Eliminar</button>
                </td>
            </tr>
        </tbody>
    </table>
    <p v-else>No hay grupos</p>

    <!-- Confirmation Dialog -->
    <ConfirmDialog
        v-if="showConfirmDialog"
        title="Confirmar eliminación"
        message="¿Estás seguro de que deseas eliminar este grupo?"
        @confirm="confirmDeleteGroup"
        @cancel="cancelDeleteGroup"
    />
</template>

<style>
</style>
