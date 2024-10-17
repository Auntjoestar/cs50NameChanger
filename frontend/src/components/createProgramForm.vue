<script setup>
import { ref } from 'vue'
import {
    CreateProgram,
    ListPrograms
} from '../../wailsjs/go/main/App';

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
        <button type="button" class="btn btn-dark" @click="createProgram" :disabled="program === ''">Crear</button>
    </form>
</template>