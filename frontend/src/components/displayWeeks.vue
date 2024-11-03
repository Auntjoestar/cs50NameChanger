<script setup>
import { ref } from 'vue';
import { WatchWeeks, DeleteWeek } from '../../wailsjs/go/main/App';
import ConfirmDialog from './ConfirmDialog.vue';

const weeks = ref([]);
const showConfirmDialog = ref(false);
const selectedWeekId = ref(null);

async function watchWeeks() {
    const result = await WatchWeeks();
    if (!result || result.length === 0 || result[0]?.id === 0) {
        return;
    }
    weeks.value = result;
}

function promptDeleteWeek(id) {
    selectedWeekId.value = id;
    showConfirmDialog.value = true;
}

function confirmDeleteWeek() {
    if (selectedWeekId.value !== null) {
        DeleteWeek(selectedWeekId.value);
        weeks.value = weeks.value.filter(week => week.id !== selectedWeekId.value);
        showConfirmDialog.value = false;
        selectedWeekId.value = null;
    }
}

function cancelDeleteWeek() {
    showConfirmDialog.value = false;
    selectedWeekId.value = null;
}

watchWeeks();
</script>

<template>
    <div class="table-container">
        <table v-if="weeks.length > 0" class="custom-table">
            <thead>
                <tr>
                    <th hidden>ID</th>
                    <th>Índice</th>
                    <th>Nombre</th>
                    <th>Ciclo</th>
                    <th>Acciones</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="(week, index) in weeks" :key="index">
                    <td hidden>{{ week.id }}</td>
                    <td>{{ index + 1 }}</td>
                    <td>{{ week.name }}</td>
                    <td>{{ week.cycle_name }}</td>
                    <td>
                        <button class="btn-delete" @click="promptDeleteWeek(week.id)">Eliminar</button>
                    </td>
                </tr>
            </tbody>
        </table>
        <p v-else>No hay semanas</p>
    </div>

    <!-- Confirmation Dialog -->
    <ConfirmDialog
        v-if="showConfirmDialog"
        title="Confirmar eliminación"
        message="¿Estás seguro de que deseas eliminar esta semana?"
        @confirm="confirmDeleteWeek"
        @cancel="cancelDeleteWeek"
    />
</template>

<style scoped>
.table-container {
    margin: 20px;
    overflow-x: auto;
}

.custom-table {
    width: 100%;
    border-collapse: collapse;
    font-family: Arial, sans-serif;
}

.custom-table th, .custom-table td {
    padding: 12px;
    text-align: center;
    border: 1px solid #ddd;
}

.custom-table thead {
    background-color: #343a40;
    color: #fff;
}

.custom-table tbody tr:hover {
    background-color: #f9f9f9;
}

.btn-delete {
    padding: 6px 12px;
    background-color: #dc3545;
    border: none;
    color: #fff;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s ease;
}

.btn-delete:hover {
    background-color: #c82333;
}
</style>
