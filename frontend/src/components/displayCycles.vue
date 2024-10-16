<script setup>
import { ref } from 'vue'
import { WatchCycles } from '../../wailsjs/go/main/App';

const cycles = ref([])

async function watchCycles() {
    const result = await WatchCycles();
    console.log("No hay ciclos disponibles");
    if (!result || result.length === 0 || result[0]?.id === 0) {
        return;
    }
    cycles.value = result; // Assign the result directly to cycles
    console.log(cycles.value);
}

watchCycles();
</script>

<template>
    <table v-if="cycles.length > 0">
        <thead>
            <tr>
                <th>ID</th>
                <th>Nombre</th>
                <th>Programa</th>
                <th>Acciones</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(cycle, index) in cycles" :key="index">
                <td>{{ cycle.id }}</td>
                <td>{{ cycle.name }}</td>
                <td>{{ cycle.program_id }}</td>
                <td>
                    <button class="btn btn-danger">Eliminar</button>
                </td>
            </tr>
        </tbody>
    </table>
    <p v-else>No hay ciclos</p>
</template>
