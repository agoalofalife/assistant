<template>
  <div id="app">
    <el-col class="icon-menu" :span="1" v-show="!stateMenu" >
    <i class=" el-icon-more" @click="changeStateMenu"> </i>
      </el-col>
    <el-col :span="3" v-show="stateMenu">
      <el-menu default-active="2" class="el-menu-vertical-demo" theme="dark" router>
        <el-submenu index="1">
          <template slot="title"><i class="el-icon-close" @click="changeStateMenu"></i>General</template>
          <el-menu-item index="process">Process</el-menu-item>
        </el-submenu>
      </el-menu>
    </el-col>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app',
  data() {
      return {
          stateMenu :false
      }
  },
    methods : {
        changeStateMenu() {
            this.stateMenu = !this.stateMenu
        }
    }
}
</script>

<style>
  .el-icon-close{
    position: absolute;
    top: 35%;
    right: 80%;
  }

  .icon-menu{
    font-size: 25px;
  }

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
  body{
    background-color: #F9FAFC;
  }
</style>
