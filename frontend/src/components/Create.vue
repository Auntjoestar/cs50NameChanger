<script setup>
import { reactive } from 'vue';

import ProgramForm from './CreateProgramForm.vue';
import CycleForm from './CreateCycleForm.vue';
import WeekForm from './CreateWeekForm.vue';
import GroupForm from './CreateGroupForm.vue';

const creatView = reactive({
    program: true,
    cycle: false,
    week: false,
    group: false,
})
</script>

<template>
    <ProgramForm v-if="creatView.program" />

    <CycleForm v-if="creatView.cycle" />

    <WeekForm v-if="creatView.week" />

    <GroupForm v-if="creatView.group" />
</template>