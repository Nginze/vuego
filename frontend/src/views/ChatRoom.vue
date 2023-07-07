<template>
  <div class="container">
    <ChatArea />
  </div>
</template>

<script>
import ChatArea from '../components/ChatArea.vue'

export default {
  data() {
    return {}
  },

  components: {
    ChatArea
  }
}
</script>

<style scoped>
.container {
  display: flex;
  align-items: center;
  justify-content: center;
}

input::placeholder {
  font-weight: 500;
  font-size: 18px;
}
</style>
