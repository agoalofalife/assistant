import Vue from 'vue'
import Router from 'vue-router'
import Process from '@/components/Process'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/process',
      name: 'Process',
      component: Process
    }
  ]
})
