<template>

    <el-col :span="20" class="process">
      <!--<el-switch-->
      <!--v-model="switchUser"-->
      <!--on-color="#13ce66"-->
      <!--off-color="#ff4949">-->
      <!--</el-switch>-->
      <i class="el-icon-loading spinner" v-show="!ready"></i>
      <el-table  v-show="ready" :data="processes" style="width: 100%">

        <el-table-column sortable fixed prop="PID" header-align="center" label="PID" width="70"></el-table-column>
        <el-table-column sortable fixed prop="CPU" header-align="center" label="CPU" width="70"></el-table-column>
        <el-table-column sortable  prop="F" header-align="center" label="F" width="90"></el-table-column>
        <el-table-column sortable prop="UID" header-align="center" label="UID" width="70"></el-table-column>
        <el-table-column sortable prop="PPID" header-align="center" label="PPID" width="70"></el-table-column>


        <el-table-column sortable prop="NI" header-align="center" label="NI" width="90"></el-table-column>
        <el-table-column sortable prop="RSS" header-align="center" label="RSS" width="90"></el-table-column>
        <el-table-column sortable prop="S" header-align="center" label="S" width="90"></el-table-column>
        <el-table-column prop="TTY" header-align="center" label="TTY" width="90"></el-table-column>
        <el-table-column sortable prop="TIME" header-align="center" label="TIME" width="100"></el-table-column>
        <el-table-column sortable prop="STIME" header-align="center" label="STIME" width="100"></el-table-column>
        <el-table-column prop="USER" header-align="center" label="USER"
                         :filters="listUsers"
                         :filter-method="filterUser"
                         filter-placement="bottom-end"width="100"></el-table-column>
        <el-table-column prop="CMD" header-align="center" label="CMD" ></el-table-column>
        <el-table-column prop="WCHAN" header-align="center" label="WCHAN" width="120"></el-table-column>
        <el-table-column fixed="right"
                label="Operations">
          <template scope="scope">
            <el-button
                    size="small"
                    type="danger"
                    @click="killPs(scope.row)">Kill</el-button>
          </template>
        </el-table-column>
      </el-table>
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
      testNewFunctional () {},

      // push emit kill process from list
      killPs(row) {
          this.$socket.emit('kill:ps', row.PID, function(data){
              console.log('ACK from server wtih data: ', data)
          })
      },

      filterUser(value, row) {
          return row.USER === value;
      }
    },
    computed:{
        listUsers(){
            let resultCompare = []
            let result        = []
            this.processes.forEach((el) => {

                if (!result.includes(el.USER)) {
                    result.push( el.USER)
                    resultCompare.push({ text: el.USER, value: el.USER })
                }
            });
            return resultCompare
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
  .spinner{
    font-size: 50px;
    position: absolute;
    top: 45%;
  }
  .process{
    margin: 0 1%;
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
