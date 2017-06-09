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
          console.log(this.$socket.emit('hello'));
      }
    },
    socket: {
        // Prefix for event names
        // prefix: "/counter/",

        // If you set `namespace`, it will create a new socket connection to the namespace instead of `/`
        // namespace: "/counter",

        events: {

            // Similar as this.$socket.on("changed", (msg) => { ... });
            // If you set `prefix` to `/counter/`, the event name will be `/counter/changed`
            //
            changed(msg) {
                console.log("Something changed: " + msg);
            },


           connect() {
           console.log("Websocket connected to " + this.$socket.nsp);
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
