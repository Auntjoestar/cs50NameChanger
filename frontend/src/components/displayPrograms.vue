<script setup>
import { ref } from 'vue'
import { WatchPrograms } from '../../wailsjs/go/main/App';

const programs = ref([])

async function watchPrograms() {
    const result = await WatchPrograms(); // Await result from WatchPrograms
    if (!result || result.length === 0 || result[0]?.id === 0) {
        console.log("No hay programas disponibles");
        return;
    }
    console.log(result);
    programs.value = result; // Assign the result directly to programs
    console.log(programs.value); // Safely log the first program's ID
}

watchPrograms();
</script>

<template>
    <table class="table table-striped table-hover" v-if="programs.length > 0">
        <thead class="table-dark">
            <tr class="text-center">
                <th class="text-center">ID</th>
                <th class="text-center">Nombre</th>
                <th class="text-center">Acciones</th>
            </tr>
        </thead>
        <tbody class="text-center">
            <tr v-for="(program, index) in programs" :key="index" class="text-center">
                <td class="text-center">{{ program.id }}</td>
                <td class="text-center">{{ program.name }}</td>
                <td class="text-center">
                    <button class="btn btn-danger">Eliminar</button>
                </td>
            </tr>
        </tbody>
    </table>
    <p v-else>No hay programas</p>
</template>
