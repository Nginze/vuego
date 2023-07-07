<template>
  <div class="container">
    <div class="chat-area">
      <span v-for="(msg, index) in Rcvmessages" :key="index" style="color: white; font-weight: 600"
        ><span :style="{ color: msg.color, fontWeight: 600 }"> {{ msg.username }} </span>:
        {{ msg.message }}</span
      >
    </div>
    <form @submit.prevent="send">
      <input v-model="message" placeholder="Enter a message" />
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      color: '',
      message: '',
      socket: null,
      Rcvmessages: []
    }
  },

  methods: {
    send() {
      const m = {
        message: this.message,
        color: this.color,
        username: this.$route.params.username
      }
      this.socket.send(JSON.stringify(m))
      this.message = ''
    },

    hashUsernameToColor(username) {
      var usernameHash = this.hashCode(username)

      var red = (usernameHash & 0xff0000) >> 16
      var green = (usernameHash & 0x00ff00) >> 8
      var blue = usernameHash & 0x0000ff

      var color = '#' + this.toHex(red) + this.toHex(green) + this.toHex(blue)

      return color
    },

    hashCode(str) {
      var hash = 0
      if (str.length === 0) {
        return hash
      }
      for (var i = 0; i < str.length; i++) {
        var char = str.charCodeAt(i)
        hash = (hash << 5) - hash + char
        hash = hash & hash // Convert to 32-bit integer
      }
      return hash
    },

    toHex(number) {
      var hex = number.toString(16)
      return hex.length === 1 ? '0' + hex : hex
    }
  },

  mounted() {
    this.color = this.hashUsernameToColor(this.$route.params.username)
    this.socket = new WebSocket('ws://localhost:5000/socket')
    this.socket.onmessage = (event) => {
      console.log(event.data)
      this.Rcvmessages.push(JSON.parse(event.data))
    }
  }
}
</script>

<style scoped>
.container {
  height: 100vh;
  width: 300px;
  margin: 1rem 0;
  display: flex;
  flex-direction: column;
  background-color: #1e2124;
  padding: 0.5rem 0.5rem;
  border-radius: 4px;
}
.chat-area {
  flex: 1;
  width: 100%;
  max-height: 100%;
  overflow-y: auto;
  display: flex;
  align-items: start;
  flex-direction: column;
  padding: 1rem;
}

input {
  background-color: #36393e;
  border: none;
  outline: none;
  padding: 0.7rem;
  width: 100%;
  color: white;
  border-radius: 5px;
}

input::placeholder {
  font-weight: 600;
  font-size: 13px;
}

form {
  width: 100%;
}
</style>
