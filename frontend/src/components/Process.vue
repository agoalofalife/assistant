<template>
<div>
    <el-col :span="20" class="process">
      <el-menu :theme="theme" :default-active="activeIndex" class="el-menu-demo" mode="horizontal" menu-trigger="">
        <!--<el-menu-item index="1">Processing Center</el-menu-item>-->
        <el-submenu index="1">
          <template slot="title"> <i class="el-icon-setting"></i>Column</template>
          <el-col :span="1">
            <el-menu-item ref="columnHide" v-for="column in listColumn"  :key="column.name" :index="column.name"><el-checkbox @change="chooseHideColumn(column.name)" :label="column.name"></el-checkbox></el-menu-item>
          </el-col>
        </el-submenu>

        <el-col :span="2"  :offset="17" class="el-menu--horizontal el-submenu el-submenu__title">
        <el-switch
                v-model="switchColorHeader"
                on-color="#324157"
                off-color="#324157"
                on-value="dark"
                off-value="light"
                off-text="dark"
                on-text="light"
                @change="changeColorHeader"
        >
        </el-switch>
        </el-col>

      </el-menu>

      <i class="el-icon-loading spinner" v-show="!ready"></i>
      <el-table height="700" v-show="ready" :data="processes" style="width: 100%" @filter-change="selectingUser">
        <el-table-column
              v-for="column in listColumn"
             :key="column.name"
              v-if="column.show"
             :sortable="column.sortable"
             :fixed="column.fixed"
             :prop="column.name"
              header-align="center"
             :label="column.name"
             :render-header="renderHeader"
        >
        </el-table-column>

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
            switchColorHeader:true,
            activeIndex: '1',
            theme:'dark',
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
                {name: 'USER', width:100, fixed:false, sortable:false, show:true, filter:{filters : 'listUsers'}},
                {name: 'CMD', width:120, fixed:false, sortable:true, show:true},
                {name: 'WCHAN', width:120, fixed:false, sortable:true, show:true},
                ],
        }
    },
    methods: {
      // choose for hide column in table
      chooseHideColumn(columnName) {
//          document.querySelector('.el-menu ul').style.display = true

          let checkbox  =     collect( this.listColumn ).where('name', columnName ).first()
          checkbox.show = !checkbox.show
          console.log(   document.querySelector('.el-menu--horizontal .el-submenu .el-menu').style.display)

//          document.querySelector('.el-menu--horizontal .el-submenu .el-menu').style.display = ""
      },
      changeColorHeader (switcher) {
          this.theme = switcher
      },
      renderHeader(createElement, { column }){
          if (column.label === 'USER') {
              column.filters         = this.listUsers
              column.filterPlacement = 'bottom-end'
              column.filterMethod    = this.filterUser
              column.filterOpened    = false
              column.filterable      = []
          }
          return createElement('span', [column.label])
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
        },
    },
    socket: {
        events: {
            listProcess(msg) {
                this.ready                 = true
                this.processes             = JSON.parse(msg)
                this.passThroughFilter()
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
