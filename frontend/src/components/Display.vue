<script setup>
import { reactive } from 'vue';

import DisplayPrograms from './DisplayPrograms.vue';
import DisplayCycles from './DisplayCycles.vue';
import DisplayWeeks from './DisplayWeeks.vue';
import DisplayGroups from './DisplayGroups.vue';

const displayView = reactive({
    programs: true,
    cycles: false,
    weeks: false,
    groups: false,
})
</script>

<template>
    <div class="content">
        <div class="navbar">
            <nav>
                <ul>
                    <li @click="displayView.programs = true; displayView.cycles = false; 
                    displayView.weeks = false; displayView.groups = false">Programas</li>
                    <li @click="displayView.programs = false; displayView.cycles = true; 
                    displayView.weeks = false; displayView.groups = false">Ciclos</li>
                    <li @click="displayView.programs = false; displayView.cycles = false; 
                    displayView.weeks = true; displayView.groups = false">Semanas</li>
                    <li @click="displayView.programs = false; displayView.cycles = false; 
                    displayView.weeks = false; displayView.groups = true">Grupos</li>
                </ul>
            </nav>
        </div>
        <div class="tables">
            <DisplayPrograms v-if="displayView.programs" />

            <DisplayCycles v-if="displayView.cycles" />

            <DisplayWeeks v-if="displayView.weeks" />

            <DisplayGroups v-if="displayView.groups" />
        </div>
    </div>
</template>

<style type="scss">
.content {
    display: flex;
}

.tables {
    width: 100%;
}
</style>