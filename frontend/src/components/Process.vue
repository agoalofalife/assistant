<template>
<div>
    <el-col :span="20" class="process">
      <el-menu theme="dark" :default-active="activeIndex" class="el-menu-demo" mode="horizontal" @select="handleSelect">
        <el-menu-item index="1">Processing Center</el-menu-item>
        <el-submenu index="2">
          <template slot="title"> <i class="el-icon-setting"></i> Column</template>
          <el-menu-item index="2-1"><el-checkbox label="Option A"></el-checkbox></el-menu-item>
          <el-menu-item index="2-2"><el-checkbox label="Option A"></el-checkbox></el-menu-item>
          <el-menu-item index="2-3"><el-checkbox label="Option A"></el-checkbox></el-menu-item>
         <!---->
          <!--<el-menu-item index="2-2">item two</el-menu-item>-->
          <!--<el-menu-item index="2-3">item three</el-menu-item>-->
        </el-submenu>

      </el-menu>
      <!--<el-switch-->
      <!--v-model="switchUser"-->
      <!--on-color="#13ce66"-->
      <!--off-color="#ff4949">-->
      <!--</el-switch>-->
      <i class="el-icon-loading spinner" v-show="!ready"></i>
      <el-table   height="700" v-show="ready" :data="processes" style="width: 100%" @filter-change="selectingUser">
        <el-table-column v-for="column in listColumn" :key="column.name"  v-if="column.show" :sortable="column.sortable" :fixed="column.fixed" :prop="column.name" header-align="center" :width="column.width" :label="column.name"></el-table-column>
        <!--<el-table-column sortable fixed prop="PID" header-align="center" label="PID" width="70"></el-table-column>-->
        <!--<el-table-column sortable fixed prop="CPU" header-align="center" label="CPU" width="70"></el-table-column>-->
        <!--<el-table-column sortable  prop="F" header-align="center" label="F" width="90"></el-table-column>-->
        <!--<el-table-column sortable prop="UID" header-align="center" label="UID" width="70"></el-table-column>-->
        <!--<el-table-column sortable prop="PPID" header-align="center" label="PPID" width="70"></el-table-column>-->


        <!--<el-table-column sortable prop="NI" header-align="center" label="NI" width="90"></el-table-column>-->
        <!--<el-table-column sortable prop="RSS" header-align="center" label="RSS" width="90"></el-table-column>-->
        <!--<el-table-column sortable prop="S" header-align="center" label="S" width="90"></el-table-column>-->
        <!--<el-table-column prop="TTY" header-align="center" label="TTY" width="90"></el-table-column>-->
        <!--<el-table-column sortable prop="TIME" header-align="center" label="TIME" width="100"></el-table-column>-->
        <!--<el-table-column sortable prop="STIME" header-align="center" label="STIME" width="100"></el-table-column>-->
        <!--<el-table-column prop="USER" header-align="center" label="USER"-->
                         <!--:filters="listUsers"-->
                         <!--:filter-method="filterUser"-->
                          <!--filter-placement="bottom-end"-->
                          <!--width="100">-->
        <!--</el-table-column>-->
        <!--<el-table-column prop="CMD" header-align="center" label="CMD" ></el-table-column>-->
        <!--<el-table-column prop="WCHAN" header-align="center" label="WCHAN" width="120"></el-table-column>-->
        <!--OPERATIONS-->
        <el-table-column fixed="right"
                label="OPERATIONS">
          <template scope="scope">
            <el-button
                    size="small"
                    type="danger"
                    @click="killPs(scope.row)">KILL</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-col>
</div>
</template>

<script>
  import collect from 'collect.js'
export default {
  name: 'hello',
    data() {
        return {
            activeIndex: '1',
            ready:false,
            processes: [],
            switchUser : true,
            selectUser:[],
            listColumn : [
                {name: 'PID', width:80, fixed:true, sortable:true, show:true},
                {name: 'CPU', width:70, fixed:true, sortable:true, show:true},
                {name: 'F', width:90, fixed:false, sortable:true, show:true},
                {name: 'UID', width:70, fixed:false, sortable:true, show:true},
                {name: 'PPID', width:70, fixed:false, sortable:true, show:true},
                {name: 'NI', width:90, fixed:false, sortable:true, show:true},
                {name: 'RSS', width:90, fixed:false, sortable:true, show:true},
                {name: 'S', width:90, fixed:false, sortable:true, show:true},
                {name: 'TTY', width:90, fixed:false, sortable:true, show:true},
                {name: 'TIME', width:100, fixed:false, sortable:true, show:true},
                {name: 'STIME', width:100, fixed:false, sortable:true, show:true},
                {name: 'USER', width:100, fixed:false, sortable:true, show:true},
                {name: 'CMD', width:120, fixed:false, sortable:true, show:true},
                {name: 'WCHAN', width:120, fixed:false, sortable:true, show:true},
                ],
//            listColumn:['PID', 'CPU', 'F', 'UID', 'PPID', 'NI', 'RSS', 'S', 'TTY', 'TIME', 'STIME', 'USER', 'CMD', 'WCHAN'],
        }
    },
    methods: {
        handleSelect(key, keyPath) {
            console.log(key, keyPath);
        },
      testNewFunctional (e,r) {},
        handleSizeChange(val) {
            console.log(`${val} handleSizeChange`);
        },
      handleCurrentChange(page) {
          this.currentPage = page
      },
      selectingUser(arr){
          this.selectUser  = Object.values(arr)[0]
       },
      // filter User ps
       passThroughFilter(){
        if (this.selectUser.length > 0 ){
            this.processes = this.processes.filter(function (ps) {
               return  this.selectUser.indexOf(ps.USER) !== -1
            }.bind(this))
        }
          return this.processes
      },
      // push emit kill process from list
      killPs(row) {
          this.$socket.emit('kill:ps', row.PID, function(data){
              console.log('ACK from server wtih data: ', data)
          })
      },
      filterUser(value, row) {
          return row.USER === value;
      },
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
                this.processes             = JSON.parse(msg)
                this.passThroughFilter()

                this.ready                 = true
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
  .el-pagination{
    margin-top: 2%;
  }
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
