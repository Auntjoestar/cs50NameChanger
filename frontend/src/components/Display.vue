<script setup>
import { reactive } from 'vue';
import Sidebar from './Sidebar.vue'; // Import the Sidebar component

import DisplayPrograms from './displayPrograms.vue';
import DisplayCycles from './displayCycles.vue';
import DisplayWeeks from './displayWeeks.vue';
import DisplayGroups from './displayGroups.vue';

const displayView = reactive({
    programs: true,
    cycles: false,
    weeks: false,
    groups: false,
});

// Method to update the view
const updateView = (view) => {
    displayView.programs = view === 'programs';
    displayView.cycles = view === 'cycles';
    displayView.weeks = view === 'weeks';
    displayView.groups = view === 'groups';
};
</script>

<template>
    <div class="content">
        <Sidebar :displayView="displayView" :updateView="updateView" /> <!-- Use Sidebar -->
        <div class="tables">
            <DisplayPrograms v-if="displayView.programs" />
            <DisplayCycles v-if="displayView.cycles" />
            <DisplayWeeks v-if="displayView.weeks" />
            <DisplayGroups v-if="displayView.groups" />
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
