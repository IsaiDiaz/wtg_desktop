import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '../components/HelloWorld.vue'
import UsersPage from '../pages/UsersPage.vue'
import UserInfoPage from '../pages/UserInfoPage.vue'
import UserEditAdminPage from '../pages/UserEditAdminPage.vue'

const routes = [
  { path: '/', component: HelloWorld },
  { path: '/users', component: UsersPage },
  { path: '/users/:id', component: UserInfoPage },
  { path: '/users/:id/edit', component: UserEditAdminPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
