import { createRouter, createWebHistory } from 'vue-router'
import HelloWorld from '../components/HelloWorld.vue'
import UsersPage from '../pages/user/UsersPage.vue'
import UserCreatePage from '../pages/user/UserCreatePage.vue'
import UserInfoPage from '../pages/user/UserInfoPage.vue'
import ProfilePage from '../pages/user/ProfilePage.vue'
import UserEditPage from '../pages/user/UserEditPage.vue'
import ProjectPage from '../pages/project/ProjectPage.vue'
import ProjectCreatePage from '../pages/project/ProjectCreatePage.vue'
import ProjectInfoPage from '../pages/project/ProjectInfoPage.vue'
import ProjectEditPage from '../pages/project/ProjectEditPage.vue'


const routes = [
  { path: '/', component: HelloWorld },
  { path: '/profile', component: ProfilePage },
  { path: '/users', component: UsersPage },
  { path: '/users/new', component: UserCreatePage },
  { path: '/users/:id', component: UserInfoPage },
  { path: '/users/:id/edit', component: UserEditPage },
  { path: '/projects', component: ProjectPage },
  { path: '/project/new', component: ProjectCreatePage },
  { path: '/project/:id', component: ProjectInfoPage },
  { path: '/project/:id/edit', component: ProjectEditPage },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
