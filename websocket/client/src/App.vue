<script>
  export default {
    name: 'App',
    data() {
      return {
        message: "",
        socket: null,
        receivedMessage: "",
        showMessage: false
      }
    },
    mounted() {
      this.socket = new WebSocket("ws://localhost:9100/socket")
      this.socket.onmessage = (msg) => {
        this.acceptMessage(msg)
      }
    },
    methods: {
      sendMessage() {
        let msg = {
          "greeting": this.message
        }
        this.socket.send(JSON.stringify(msg))
      },
      acceptMessage(msg) {
        this.receivedMessage = msg.data
        this.showMessage = true
      }
    },
}
</script>

<template>
  <form action="sendMessage" @click.prevent="onsubmit">
    <input v-model="message" type="text">  
    <input type="submit" value="Send" @click="sendMessage">
  </form>
  <div v-if="showMessage">
    <h3>Server:</h3>
    <p>
      {{ receivedMessage }}
    </p>
    <button @click="showMessage = !showMessage">Dismiss</button>
  </div>
</template>

<style scoped>
header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
  }
}
</style>
