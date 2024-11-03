<script setup>
import { ref } from 'vue'
import { 
    WatchGroups,
    DeleteGroup,
} from '../../wailsjs/go/main/App';

const groups = ref([])

async function watchGroups() {
    const result = await WatchGroups();
    if (!result || result.length === 0 || result[0]?.id === 0) {
        return;
    }
    groups.value = result; // Assign the result directly to groups
}

async function deleteGroup(id) {
    DeleteGroup(id);
    groups.value = groups.value.filter(group => group.id !== id);
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
                    <button class="btn btn-danger" @click="deleteGroup(group.id)">Eliminar</button>
                </td>
            </tr>
        </tbody>
    </table>
    <p v-else>No hay grupos</p>
</template>