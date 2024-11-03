<script setup>
import { ref } from 'vue'
import {
    WatchWeeks,
    DeleteWeek,
}
    from '../../wailsjs/go/main/App';

const weeks = ref([])

async function watchWeeks() {
    const result = await WatchWeeks();
    if (!result || result.length === 0 || result[0]?.id === 0) {
        return;
    }
    weeks.value = result;
}

async function deleteWeek(id) {
    DeleteWeek(id);
    weeks.value = weeks.value.filter(week => week.id !== id);
}

watchWeeks();
</script>

<template>
    <table v-if="weeks.length > 0" class="table table-striped table-hover">
        <thead class="table-dark">
            <tr class="text-center">
                <th class="text-center" scope="col" hidden>ID</th>
                <th class="text-center" scope="col">Índice</th>
                <th class="text-center" scope="col">Nombre</th>
                <th class="text-center" scope="col">Ciclo</th>
                <th class="text-center" scope="col">Acciones</th>
            </tr>
        </thead>
        <tbody class="text-center">
            <tr v-for="(week, index) in weeks" :key="index" class="text-center">
                <td class="text-center" scope="row" hidden>{{ week.id }}</td>
                <td class="text-center">{{ index + 1 }}</td>
                <td class="text-center">{{ week.name }}</td>
                <td class="text-center">{{ week.cycle_name }}</td>
                <td class="text-center">
                    <button class="btn btn-danger" @click="deleteWeek(week.id)">Eliminar</button>
                </td>
            </tr>
        </tbody>
    </table>
    <p v-else>No hay semanas</p>
</template>