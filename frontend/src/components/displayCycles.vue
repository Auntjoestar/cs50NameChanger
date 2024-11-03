<script setup>
import { ref } from 'vue'
import { WatchCycles, DeleteCycle } from '../../wailsjs/go/main/App';

const cycles = ref([])

async function watchCycles() {
    const result = await WatchCycles();
    if (!result || result.length === 0 || result[0]?.id === 0) {
        return;
    }
    cycles.value = result; // Assign the result directly to cycles
    console.log(cycles.value);
}

function deleteCycle(id) {
    DeleteCycle(id);
    cycles.value = cycles.value.filter(cycle => cycle.id !== id);
}

watchCycles();
</script>

<template>
    <table v-if="cycles.length > 0" class="table table-striped table-bordered">
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
                    <button class="btn btn-danger" @click="deleteCycle(cycle.id)">Eliminar</button>
                </td>
            </tr>
        </tbody>
    </table>
    <p v-else>No hay ciclos</p>
</template>

<style>
</style>