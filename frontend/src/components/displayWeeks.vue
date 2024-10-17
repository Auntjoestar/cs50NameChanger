<script setup>
import { ref } from 'vue'
import {
    WatchWeeks,
}
    from '../../wailsjs/go/main/App';

const weeks = ref([])

async function watchWeeks() {
    const result = await WatchWeeks();
    console.log("No hay semanas disponibles");
    if (!result || result.length === 0 || result[0]?.id === 0) {
        return;
    }
    weeks.value = result; // Assign the result directly to weeks
    console.log(weeks.value);
}

watchWeeks();
</script>

<template>
    <table v-if="weeks.length > 0" class="table table-striped table-hover">
        <thead class="table-dark">
            <tr class="text-center">
                <th class="text-center" scope="col">ID</th>
                <th class="text-center" scope="col">Nombre</th>
                <th class="text-center" scope="col">ID Ciclo</th>
                <th class="text-center" scope="col">Acciones</th>
            </tr>
        </thead>
        <tbody class="text-center">
            <tr v-for="(week, index) in weeks" :key="index" class="text-center">
                <td class="text-center" scope="row">{{ week.id }}</td>
                <td class="text-center">{{ week.name }}</td>
                <td class="text-center">{{ week.cycle_id }}</td>
                <td class="text-center">
                    <button class="btn btn-danger">Eliminar</button>
                </td>
            </tr>
        </tbody>
    </table>
    <p v-else>No hay semanas</p>
</template>