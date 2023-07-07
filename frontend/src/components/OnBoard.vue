<template>
  <div class="container">
    <div class="header">
      Enter your Username✨
      <div v-if="username" style="font-size: 13px">👋 Hello, {{ username }}</div>
      <div style="color: red" v-if="error">{{ this.error }}</div>
    </div>
    <input v-model="username" placeholder="Username" />
    <div
      @click="
        () => {
          if (this.username) {
            this.$router.push({ path: `/chat/${this.username}` })
          } else {
            this.error = 'Please enter a username'
          }
        }
      "
      class="button"
    >
      Okay
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: '',
      error: ''
    }
  },

  methods: {}
}
</script>

<style scoped>
@import '../assets/base.css';
.container {
  color: white;
}
.header {
  font-size: 25px;
  font-weight: 700;
  margin-bottom: 1rem;
  text-align: center;
}

input {
  background-color: var(--color-background-soft);
  border: none;
  outline: none;
  padding: 1rem;
  border-radius: 0.4rem;
  width: 300px;
  color: white;
  margin-bottom: 1rem;
}

input::placeholder {
  font-weight: 500;
  font-size: 18px;
}

.button {
  padding: 0.5rem 1rem;
  background-color: #7289da;
  font-weight: 700;
  border-radius: 0.4rem;
  width: 200px;
  text-align: center;
  cursor: pointer;
}
</style>
