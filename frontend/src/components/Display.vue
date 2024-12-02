<script setup>
import { reactive, ref } from 'vue';
import Sidebar from './Sidebar.vue'; // Import the Sidebar component

import DisplayPrograms from './displayPrograms.vue';
import DisplayCycles from './displayCycles.vue';
import DisplayWeeks from './displayWeeks.vue';
import DisplayGroups from './displayGroups.vue';

const message = ref('');
const messageType = ref('alert-success');

const displayView = reactive({
    programs: true,
    cycles: false,
    weeks: false,
    groups: false,
});

// Method to update the view
const updateView = (view) => {
    message.value = '';
    for (const key in displayView) {
        displayView[key] = false;
    }
    displayView[view] = true;
};

function handleProgramEdited() {
    showMessage('alert-success', 'Programa editado correctamente');
}

function handleCycleEdited() {
    showMessage('alert-success', 'Ciclo editado correctamente');
}

function handleWeekEdited() {
    showMessage('alert-success', 'Semana editada correctamente');
}

function handleGroupEdited() {
    showMessage('alert-success', 'Grupo editado correctamente');
}

function handleError(error) {
    showMessage('alert-danger', error);
}

function showMessage(type, msg) {
    message.value = msg;
    messageType.value = type;
}

</script>

<template>
    <div class="content">
        <Sidebar :displayView="displayView" :updateView="updateView" />
        <div class="tables">
            <div id="message" v-if="message" :class="['alert', messageType, 'alert-dismissible', 'fade', 'show']"
                role="alert">
                {{ message }}
                <button type="button" class="btn-close" @click="message = ''"></button>
            </div>
            <DisplayPrograms v-if="displayView.programs" @program-edited="handleProgramEdited" @error="handleError" />
            <DisplayCycles v-if="displayView.cycles" @cycle-edited="handleCycleEdited" @error="handleError" />
            <DisplayWeeks v-if="displayView.weeks" @week-edited="handleWeekEdited" @error="handleError" />
            <DisplayGroups v-if="displayView.groups" @group-edited="handleGroupEdited" @error="handleError" />
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
