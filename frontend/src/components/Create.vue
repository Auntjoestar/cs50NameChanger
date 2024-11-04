<script setup>
import { reactive } from 'vue';
import Sidebar from './Sidebar.vue'; // Import the Sidebar component

import ProgramForm from './CreateProgramForm.vue';
import CycleForm from './CreateCycleForm.vue';
import WeekForm from './CreateWeekForm.vue';
import GroupForm from './CreateGroupForm.vue';

const createView = reactive({
    programs: true,
    cycles: false,
    weeks: false,
    groups: false,
});

// Method to update the view
const updateView = (view) => {
    createView.programs = view === 'programs';
    createView.cycles = view === 'cycles';
    createView.weeks = view === 'weeks';
    createView.groups = view === 'groups';
};
</script>

<template>
    <div class="content">
        <Sidebar :displayView="createView" :updateView="updateView" /> <!-- Use Sidebar -->
        <div class="tables">
            <ProgramForm v-if="createView.programs" />
            <CycleForm v-if="createView.cycles" />
            <WeekForm v-if="createView.weeks" />
            <GroupForm v-if="createView.groups" />
        </div>
    </div>
</template>

<style scoped lang="scss">
.content {
    display: flex;
}

.tables {
    width: 100%;
    padding: 20px; /* Add padding to the content area */
    background-color: #f8f9fa;
}
</style>
