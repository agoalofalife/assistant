<template>

  <el-col :span="24" >
    <el-col :span="18" class="process">
      <!--<el-switch-->
      <!--v-model="switchUser"-->
      <!--on-color="#13ce66"-->
      <!--off-color="#ff4949">-->
      <!--</el-switch>-->

      <el-table  v-show="ready" :data="processes" style="width: 100%">
        <el-table-column prop="PID" header-align="center" label="PID" width="90"></el-table-column>
        <el-table-column prop="TTY" header-align="center" label="TTY" width="90"></el-table-column>
        <el-table-column prop="TIME" header-align="center" label="TIME" width="90"></el-table-column>
        <el-table-column prop="CMD" header-align="center" label="CMD" ></el-table-column>
      </el-table>
    </el-col>
  </el-col>



</template>

<script>
export default {
  name: 'hello',
    data() {
        return {
            ready:false,
            processes: [],

            switchUser : true
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
                this.processes = JSON.parse(msg)
                this.ready = true

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
  .process{
    margin: 1% 0;
  }
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
