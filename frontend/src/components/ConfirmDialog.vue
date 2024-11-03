<template>
  <div class="overlay" @click="handleOverlayClick">
      <div class="dialog" @click.stop>
          <h3>{{ title }}</h3>
          <p>{{ message }}</p>
          <div class="buttons">
              <button @click="$emit('confirm')" class="btn btn-primary">Confirmar</button>
              <button @click="$emit('cancel')" class="btn btn-secondary">Cancelar</button>
          </div>
      </div>
  </div>
</template>

<script>
export default {
  props: {
      title: String,
      message: String,
      isVisible: Boolean,
  },
  methods: {
      handleOverlayClick() {
          this.$emit('cancel'); 
      },
  },
}
</script>

<style scoped>
.overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}
.dialog {
  background: white;
  padding: 20px;
  border-radius: 8px;
  max-width: 400px;
  width: 100%;
}
.buttons {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  margin-top: 15px;
}
</style>
