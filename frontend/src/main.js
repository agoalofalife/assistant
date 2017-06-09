// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueWebsocket from "vue-websocket";

Vue.use(VueWebsocket, "http://localhost:7777/echo", {
// Vue.use(VueWebsocket, "ws://localhost:7777/echo", {
//     path: '/echo',
    transports: ['websocket']
});

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
