<script setup>
import { reactive, onMounted, onUnmounted } from 'vue';

import Create from './Create.vue';
import Display from './Display.vue';

const view = reactive({
    create: false,
    display: true,
});


onMounted(() => {
    document.body.style.overflow = 'hidden'; 
});


onUnmounted(() => {
    document.body.style.overflow = ''; 
});
</script>

<template>
    <header class="admin-header">
        <nav class="navbar">
            <a class="navbar-brand" href="#">Panel de administración</a>
            <div class="navbar-buttons">
                <a class="nav-link" href="#ver"
                    :class="{ active: view.display }"
                    @click="() => { view.create = false; view.display = true }">Ver</a>
                <a class="nav-link" href="#crear"
                    :class="{ active: view.create }"
                    @click="() => { view.create = true; view.display = false }">Crear</a>
            </div>
        </nav>
    </header>
    <main class="admin-content">
        <Display v-if="view.display" />
        <Create v-if="view.create" />
    </main>
</template>

<style scoped lang="scss">
.admin-header {
    background-color: #333;
    color: #fff;
    padding: 10px;
    display: flex;
    align-items: center;
    justify-content: space-between;

    .navbar {
        display: flex;
        align-items: center;
        width: 100%;

        .navbar-brand {
            font-size: 1.25rem;
            font-weight: bold;
            color: #fff;
            text-decoration: none;
        }

        .navbar-buttons {
            display: flex;

            .nav-link {
                color: #ccc;
                text-decoration: none;
                padding: 8px 12px;
                margin-right: 10px;
                transition: color 0.3s, background-color 0.3s;
                border-radius: 4px;

                &.active, &:hover {
                    background-color: #007bff;
                    color: #fff;
                }
            }
        }
    }
}

.admin-content {
    background-color: #f8f9fa;
    width: 100%;
}

</style>
