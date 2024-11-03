<script setup>
import { ref } from 'vue'
import { WatchPrograms, DeleteProgram } from '../../wailsjs/go/main/App'
import ConfirmDialog from './ConfirmDialog.vue'

const programs = ref([])
const isConfirmVisible = ref(false)
const programIdToDelete = ref(null)

async function watchPrograms() {
    const result = await WatchPrograms()
    if (!result || result.length === 0 || result[0]?.id === 0) {
        return
    }
    programs.value = result
}

function openConfirmDialog(id) {
    programIdToDelete.value = id
    isConfirmVisible.value = true
}

function handleConfirm() {
    DeleteProgram(programIdToDelete.value)
    watchPrograms()
    closeConfirmDialog()
}

function closeConfirmDialog() {
    isConfirmVisible.value = false
    programIdToDelete.value = null
}

watchPrograms()
</script>

<template>
    <table class="table table-striped table-hover table-bordered" v-if="programs.length > 0">
        <thead class="table-dark">
            <tr class="text-center">
                <th class="text-center" scope="col" hidden>ID</th>
                <th class="text-center" scope="col">Índice</th>
                <th class="text-center" scope="col">Nombre</th>
                <th class="text-center" scope="col">Acciones</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(program, index) in programs" :key="program.id" class="text-center">
                <td class="text-center" scope="row" hidden>{{ program.id }}</td>
                <td class="text-center">{{ index + 1 }}</td>
                <td class="text-center">{{ program.name }}</td>
                <td class="text-center">
                    <button class="btn btn-danger" @click="openConfirmDialog(program.id)">Eliminar</button>
                </td>
            </tr>
        </tbody>
    </table>
    <p v-else>No hay programas</p>

    <!-- Confirmation Dialog -->
    <ConfirmDialog
        v-if="isConfirmVisible"
        :title="'Confirmación de Eliminación'"
        :message="'¿Estás seguro de que deseas eliminar este programa?'"
        :isVisible="isConfirmVisible"
        @confirm="handleConfirm"
        @cancel="closeConfirmDialog"
    />
</template>
