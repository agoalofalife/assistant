import Vue from 'vue'
import Router from 'vue-router'
import Process from '@/components/Process'
import Task from '@/components/Task'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/process',
      name: 'Process',
      component: Process
    },
    {
      path:'/task',
      name:'Task',
      component:Task
    }
  ]
})
