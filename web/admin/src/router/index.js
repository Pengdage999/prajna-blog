import Vue from 'vue'
import VueRouter from 'vue-router'
import LoginView from '../views/LoginView.vue'
import AdminView from '../views/AdminView.vue'

// 页面路由组件
import Index from '../components/admin/Index.vue'
import AddArt from '../components/articles/AddArt.vue'
import ArtList from '../components/articles/ArtList.vue'
import CateList from '../components/categories/CateList.vue'
import UserList from '../components/users/UserList.vue'
import Profile from '../components/users/Profile.vue'
Vue.use(VueRouter)

const routes = [
  {
    path: '/login',
    name: 'login',
    component: LoginView,
    meta:{title: '管理页面'}
  },
  {
    path: '/',
    name: 'admin',
    component: AdminView,
    children: [
      { path: 'index', component: Index, meta:{title: '管理页面'} },
      { path: 'addart', component: AddArt },
      { path: 'addart/:id', component: AddArt, props: true },
      { path: 'artlist', component: ArtList },
      { path: 'catelist', component: CateList },
      { path: 'userlist', component: UserList },
      { path: 'profile', component: Profile },
    ],
  },
]

const router = new VueRouter({
  routes,
})

// 路由导航守卫
router.beforeEach((to, from, next) => {
  const token = window.sessionStorage.getItem('token')
  if (to.path === '/login') return next()
  if (!token && to.path === '/') {
    next('/login')
  } else {
    next()
  }
})
export default router
