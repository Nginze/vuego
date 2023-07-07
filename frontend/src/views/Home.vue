<template>
  <div class="container">
    <OnBoard />
  </div>
</template>

<script>
import OnBoard from '../components/OnBoard.vue'

export default {
  data() {
    return {}
  },

  components: {
    OnBoard
  }
}
</script>

<style scoped>

.container {
  width: 100%;
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}
</style>
