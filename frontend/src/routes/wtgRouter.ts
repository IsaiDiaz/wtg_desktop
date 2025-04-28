import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '../components/HelloWorld.vue'
import UsersPage from '../pages/UsersPage.vue'

const routes = [
  { path: '/', component: HelloWorld },
  { path: '/users', component: UsersPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
