<!-- Sidebar.vue -->
<script setup>
import { toRefs, defineProps } from 'vue';

const props = defineProps({
    displayView: Object,
    updateView: Function
});
</script>

<template>
    <div class="sidebar">
        <nav>
            <ul>
                <li :class="{ active: displayView.programs }" 
                    @click="() => updateView('programs')">Programas</li>
                <li :class="{ active: displayView.cycles }" 
                    @click="() => updateView('cycles')">Ciclos</li>
                <li :class="{ active: displayView.weeks }" 
                    @click="() => updateView('weeks')">Semanas</li>
                <li :class="{ active: displayView.groups }" 
                    @click="() => updateView('groups')">Grupos</li>
            </ul>
        </nav>
    </div>
</template>

<style scoped lang="scss">
.sidebar {
    width: 200px;
    height: 100vh;
    background-color: #333;
    padding-top: 20px;

    nav {
        ul {
            list-style: none;
            padding: 0;
            margin: 0;

            li {
                padding: 15px;
                color: white;
                cursor: pointer;
                transition: background-color 0.3s, color 0.3s;

                &:hover {
                    background-color: #555;
                }

                &.active {
                    background-color: #007bff;
                    color: #fff;
                    font-weight: bold;
                }
            }
        }
    }
}
</style>
