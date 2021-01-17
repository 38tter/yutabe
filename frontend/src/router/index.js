import Vue from 'vue'
import Router from 'vue-router'
import index from '@/components/index'
import recipe from '@/components/recipe'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'index',
      component: index
    },
    {
      path: '/recipe',
      name: 'recipe',
      component: recipe
    }
  ]
})
