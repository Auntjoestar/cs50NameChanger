<script setup>
import { ref, reactive } from 'vue'
import {
    WatchPrograms,
    WatchCycles,
    WatchWeeks,
    WatchGroups,
} from '../../wailsjs/go/main/App';

import ProgramForm from './createProgramForm.vue';
import CycleForm from './createCycleForm.vue';
import WeekForm from './createWeekForm.vue';
import GroupForm from './createGroupForm.vue';
import DisplayPrograms from './displayPrograms.vue';
import DisplayCycles from './displayCycles.vue';
import DisplayWeeks from './displayWeeks.vue';
import DisplayGroups from './displayGroups.vue';

const creatView = reactive({
    program: true,
    cycle: false,
    week: false,
    group: false,
})

</script>

<template>
    <header>
        <nav class="navbar navbar-expand-lg bg-body-tertiary">
            <div class="container-fluid">
                <a class="navbar-brand" href="#">Panel de administración</a>
                <button class="navbar-toggler" type="button" data-bs-toggle="collapse"
                    data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent"
                    aria-expanded="false" aria-label="Toggle navigation">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                        <li class="nav-item">
                            <a class="nav-link active" aria-current="page" href="#">Ver</a>
                        </li>
                        <li class="nav-item dropdown">
                            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown"
                                aria-expanded="false">
                                Crear
                            </a>
                            <ul class="dropdown-menu">
                                <li>
                                    <a class="dropdown-item"
                                        @click="creatView.program = true; creatView.cycle = false; creatView.week = false; creatView.group = false;">Programa
                                    </a>
                                </li>
                                <li>
                                    <a class="dropdown-item"
                                        @click="creatView.program = false; creatView.cycle = true; creatView.week = false; creatView.group = false;">Ciclo
                                    </a>
                                </li>
                                <li>
                                    <a class="dropdown-item"
                                        @click="creatView.program = false; creatView.cycle = false; creatView.week = true; creatView.group = false;">Semana
                                    </a>
                                </li>
                                <li>
                                    <a class="dropdown-item"
                                        @click="creatView.program = false; creatView.cycle = false; creatView.week = false; creatView.group = true;">Grupo
                                    </a>
                                </li>
                            </ul>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    </header>
    <main>
        <ProgramForm v-if="creatView.program" />

        <CycleForm v-if="creatView.cycle" />

        <WeekForm v-if="creatView.week" />

        <GroupForm v-if="creatView.group" />
        
        <DisplayPrograms />

        <DisplayCycles />
        
        <DisplayWeeks />

        <DisplayGroups />

    </main>
</template>

