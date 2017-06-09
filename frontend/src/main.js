// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import VueWebsocket from "vue-websocket";

import io from "socket.io-client";
let  socket = io('ws://localhost:7777', {
    // path: '/echo',
    transports: ['websocket'],
    autoConnect:true,
});
console.log( socket );
socket.on('connect', function(){
    console.log(socket.id); // 'G5p5...'
});
// Vue.use(VueWebsocket, "ws://localhost:7777", {
// // Vue.use(VueWebsocket, "ws://localhost:7777/echo", {
//     origins : 'localhost:7777:*")',
//     path: '/echo',
//     transports: ['websocket']
// });

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})
