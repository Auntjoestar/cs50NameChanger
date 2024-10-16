<script setup>
import { ref } from 'vue'
import { CreateProgram } from '../../wailsjs/go/main/App';

const program = ref("");

async function createProgram() {
    await CreateProgram(program.value);
    program.value = "";
}

</script>

<template>
    <form>
        <h2 class="form-title">Crear Programa</h2>
        <input type="text" v-model="program" placeholder="Programa" class="form-control" />
        <button type="button" class="btn btn-dark" @click="createProgram">Crear</button>
    </form>
</template>