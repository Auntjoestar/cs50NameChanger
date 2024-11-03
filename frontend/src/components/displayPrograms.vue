<script setup>
import { ref } from 'vue'
import { WatchPrograms, DeleteProgram } from '../../wailsjs/go/main/App';

const programs = ref([])

async function watchPrograms() {
    const result = await WatchPrograms(); 
    if (!result || result.length === 0 || result[0]?.id === 0) {
        return;
    }
    programs.value = result; 
}

function deleteProgram(id) {
    DeleteProgram(id);
    programs.value = programs.value.filter(program => program.id !== id);
}

watchPrograms();
</script>

<template>
    <table class="table table-striped table-hover" v-if="programs.length > 0">
        <thead class="table-dark">
            <tr class="text-center">
                <th class="text-center" scope="col" hidden>ID</th>
                <th class="text-center" scope="col">Índice</th>
                <th class="text-center" scope="col">Nombre</th>
                <th class="text-center" scope="col">Acciones</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(program, index) in programs" :key="index" class="text-center">
                <td class="text-center" scope="row" hidden>{{ program.id }}</td>
                <td class="text-center">{{ index + 1 }}</td>
                <td class="text-center">{{ program.name }}</td>
                <td class="text-center">
                    <button class="btn btn-danger" @click="deleteProgram(program.id)">Eliminar</button>
                </td>
            </tr>
        </tbody>
    </table>
    <p v-else>No hay programas</p>
</template>
