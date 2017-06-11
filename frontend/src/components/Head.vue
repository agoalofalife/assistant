<template>

  <div class="hello">
    <h2 @click="start">Essential Links</h2>
  </div>
</template>

<script>
export default {
  name: 'hello',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },
    methods: {
      start () {
          console.log(this.$socket.emit('chat message'));
      }
    },
    socket: {
        events: {
            listProcess(msg) {
                console.log("listProcess: " + msg);
                console.log(  JSON.parse(msg) );
            },
           connect() {
           console.info("Websocket connected to " + this.$socket.nsp);
           },
           disconnect() {
           console.log("Websocket disconnected from " + this.$socket.nsp);
           },
           error(err) {
           console.error("Websocket error!", err);
           }
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
