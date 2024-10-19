<script setup>
import { reactive } from 'vue';

import ProgramForm from './CreateProgramForm.vue';
import CycleForm from './CreateCycleForm.vue';
import WeekForm from './CreateWeekForm.vue';
import GroupForm from './CreateGroupForm.vue';

const createView = reactive({
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
                    <li @click="createView.programs = true; createView.cycles = false;
                    createView.weeks = false; createView.groups = false">Programas</li>
                    <li @click="createView.programs = false; createView.cycles = true;
                    createView.weeks = false; createView.groups = false">Ciclos</li>
                    <li @click="createView.programs = false; createView.cycles = false;
                    createView.weeks = true; createView.groups = false">Semanas</li>
                    <li @click="createView.programs = false; createView.cycles = false;
                    createView.weeks = false; createView.groups = true">Grupos</li>
                </ul>
            </nav>
        </div>
        <div class="tables">
            <ProgramForm v-if="createView.programs" />

            <CycleForm v-if="createView.cycles" />

            <WeekForm v-if="createView.weeks" />

            <GroupForm v-if="createView.groups" />
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