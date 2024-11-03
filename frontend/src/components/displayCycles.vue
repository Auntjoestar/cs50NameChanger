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
    console.log(cycles.value);
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
    <table v-if="cycles.length > 0" class="table table-striped table-bordered table-hover">
        <thead class="table-dark">
            <tr>
                <th scope="col" hidden>ID</th>
                <th scope="col">Índice</th>
                <th scope="col">Nombre</th>
                <th scope="col">Programa</th>
                <th scope="col">Acciones</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(cycle, index) in cycles" :key="index">
                <td scope="row" hidden>{{ cycle.id }}</td>
                <td>{{ index + 1 }}</td>
                <td>{{ cycle.name }}</td>
                <td>{{ cycle.program_name }}</td>
                <td>
                    <button class="btn btn-danger" @click="promptDeleteCycle(cycle.id)">Eliminar</button>
                </td>
            </tr>
        </tbody>
    </table>
    <p v-else>No hay ciclos</p>

    <ConfirmDialog
        v-if="showConfirmDialog"
        title="Confirmar eliminación"
        message="¿Estás seguro de que deseas eliminar este ciclo?"
        @confirm="confirmDeleteCycle"
        @cancel="cancelDeleteCycle"
    />
</template>

<style>
</style>
