<template>
  <div id="app">
    <el-col :span="3">
      <el-menu default-active="2" class="el-menu-vertical-demo" theme="dark" router>
        <el-submenu index="1">
          <template slot="title">General</template>
          <el-menu-item index="process">Process</el-menu-item>
        </el-submenu>
      </el-menu>
    </el-col>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'app'
}
</script>

<style>
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
