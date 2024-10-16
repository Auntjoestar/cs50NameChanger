<script setup>
import { ref } from 'vue'
import { WatchGroups } from '../../wailsjs/go/main/App';

const groups = ref([])

async function watchGroups() {
    const result = await WatchGroups();
    if (!result || result.length === 0 || result[0]?.id === 0) {
        console.log("No hay grupos disponibles");
        return;
    }
    console.log(result);
    groups.value = result; // Assign the result directly to groups
    console.log(groups.value);
}

watchGroups();
</script>

<template>
    <table v-if="groups.length > 0">
        <thead>
            <tr>
                <th>ID</th>
                <th>Nombre</th>
                <th>Semana</th>
                <th>Acciones</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(group, index) in groups" :key="index">
                <td>{{ group.id }}</td>
                <td>{{ group.name }}</td>
                <td>{{ group.week_id }}</td>
                <td>
                    <button class="btn btn-danger">Eliminar</button>
                </td>
            </tr>
        </tbody>
    </table>
    <p v-else>No hay grupos</p>
</template>