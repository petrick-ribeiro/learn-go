import { createApp } from "vue";
import App from "./App.vue";

import "./assets/main.css";

createApp(App).mount("#app");

  export default {
    name: 'App',
    data() {
      return {
        message: "",
        socket: null
      }
    },
    mounted() {
      this.websocket = new WebSocket("ws://localhost:9100/socket")
    },
    methods: {
      sendMessage() {
        let msg = {
          "greeting": this.message
        }
        this.websocket.send(JSON.stringify(msg))
      }
    },
  }
