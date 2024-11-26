<script setup>
import { reactive, ref } from 'vue';
import Sidebar from './Sidebar.vue'; // Import the Sidebar component

import ProgramForm from './CreateProgramForm.vue';
import CycleForm from './CreateCycleForm.vue';
import WeekForm from './CreateWeekForm.vue';
import GroupForm from './CreateGroupForm.vue';

const message = ref(''); 
const messageType = ref('alert-success');

const createView = reactive({
    programs: true,
    cycles: false,
    weeks: false,
    groups: false,
});


function updateView(view) {
    message.value = '';
    for (const key in createView) {
        createView[key] = false;
    }
    createView[view] = true;
}

function showMessage(type, msg) {
    messageType.value = type;
    message.value = msg;
}

function handleProgramCreated() {
    showMessage('alert-success', 'Programa creado correctamente');
}


function handleCycleCreated() {
    showMessage('alert-success', 'Ciclo creado correctamente');
}

function handleWeekCreated() {
    showMessage('alert-success', 'Semana creada correctamente');
}

function handleGroupCreated() {
    showMessage('alert-success', 'Grupo creado correctamente');
}

function handleError(error) {
    showMessage('alert-danger', error);
}
</script>

<template>
    <div class="content">
        <Sidebar :displayView="createView" :updateView="updateView" /> <!-- Use Sidebar -->
        <div class="tables">
            <div id="message" v-if="message" :class="['alert', messageType, 'alert-dismissible', 'fade', 'show']"
                role="alert">
                {{ message }}
                <button type="button" class="btn-close" @click="message = ''" aria-label="Close"></button>
            </div>
            <ProgramForm v-if="createView.programs" @program-created="handleProgramCreated" @error="handleError" />
            <CycleForm v-if="createView.cycles" @cycle-created="handleCycleCreated" @error="handleError" />
            <WeekForm v-if="createView.weeks" @week-created="handleWeekCreated" @error="handleError" />
            <GroupForm v-if="createView.groups" @group-created="handleGroupCreated" @error="handleError" />
        </div>
    </div>
</template>

<style scoped lang="scss">
.content {
    display: flex;
}

.tables {
    width: 100%;
    padding: 20px;
    /* Add padding to the content area */
    background-color: #f8f9fa;
}
</style>
