<script setup>
import { ref, onMounted } from 'vue';
import { WatchWeeks, DeleteWeek } from '../../wailsjs/go/main/App';
import ConfirmDialog from './ConfirmDialog.vue';

const weeks = ref([]);
const showConfirmDialog = ref(false);
const selectedWeekId = ref(null);
const isLoading = ref(true); // Add loading state

async function watchWeeks() {
    try {
        const result = await WatchWeeks();
        if (!result || result.length === 0 || result[0]?.id === 0) {
            return;
        }
        weeks.value = result;
    } catch (error) {
        console.error("Error fetching weeks:", error);
    } finally {
        isLoading.value = false; // Set loading to false once the fetching is done
    }
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

// Fetch weeks when the component is mounted
onMounted(() => {
    watchWeeks();
});
</script>

<template>
    <h2>Semanas (Total: {{ weeks.length }})</h2>
    <div class="table-container" v-if="weeks.length > 0 || isLoading">
        <div v-if="isLoading" class="loading-indicator">Cargando...</div>
        <div class="table-wrapper">
            <table v-if="!isLoading && weeks.length > 0" class="custom-table">
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
                            <button class="btn-delete" @click="promptDeleteWeek(week.id)"
                                aria-label="Eliminar semana">Eliminar</button>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
    <div class="no-weeks" v-else>
        <p class="empty-state">No hay semanas para mostrar.</p>
    </div>

    <ConfirmDialog v-if="showConfirmDialog" title="Confirmar eliminación"
        message="¿Estás seguro de que deseas eliminar esta semana?" @confirm="confirmDeleteWeek"
        @cancel="cancelDeleteWeek" />
</template>

<style scoped lang="scss">
.table-container {
    margin: 1%;
}

.table-wrapper {
    max-height: 400px;
    overflow-y: auto;
    border: 1px solid #ddd;
}

.custom-table {
    width: 100%;
    border-collapse: collapse;
    font-family: Arial, sans-serif;

    th,
    td {
        padding: 12px;
        text-align: center;
        border: 1px solid #ddd;
        border-top: none;

    }

    th {
        top: 0;
        border-top: none;
    }

    thead th {
        background-color: #343a40;
        color: #fff;
        position: sticky;
        top: 0;
        z-index: 1;
    }

    tbody tr:hover {
        background-color: #f1f1f1;
    }
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

.loading-indicator {
    text-align: center;
    font-size: 1.2rem;
    color: #007bff;
    padding: 20px;
}

.no-weeks {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 90%;
    color: #666;
    font-size: 1.1rem;
}
</style>
