<script setup>
import { reactive } from 'vue'

import Create from './Create.vue';
import Display from './Display.vue';

const view = reactive({
    create: false,
    display: true,
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
                            <a class="nav-link active" aria-current="page" href="#ver"
                                @click="() => { view.create = false; view.display = true }">Ver</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link active" aria-current="page" href="#crear"
                                @click="() => { view.create = true; view.display = false }">Crear
                            </a>
                        </li>
                    </ul>
                </div>
            </div>
        </nav>
    </header>
    <main>
        <Display v-if="view.display" />

        <Create v-if="view.create" />
    </main>
</template>
