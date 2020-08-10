import VueRouter from 'vue-router'
// import { getToken } from '@/utils/auth' // getToken from cookie
// import { getUserInfo } from '@/api/user'
import store from '../store'

const routes = [
  { path: '/login', name: 'login', component: () => import('@/views/login') },
  { path: '/signup', component: () => import('@/views/login/signup') },
  { path: '/', component: () => import('@/views/index') },
  { path: '/project', component: () => import('@/views/project') },
  { path: '/screenshot/', component: () => import('@/views/Dashboard/screenshot') },

  { path: '/dashboard', component: () => import('@/views/Dashboard') },
  { path: '/chartpanel/:id', component: () => import('@/views/ChartPanel'), meta: { activeMenu: '/chartpanel/create' }},
  { path: '/fullscreendb/:id', component: () => import('@/views/Dashboard/fullScreenDb') },
  { path: '/source', component: () => import('@/views/source/index') },
  { path: '/source/:id/table', component: () => import('@/views/source/table') },
  { path: '/lab', component: () => import('@/views/lab/index') },
  { path: '/task', component: () => import('@/views/task/index') },

  { path: '/:projectid/dashboard', component: () => import('@/views/Dashboard') },
  { path: '/:projectid/chartpanel/:id', component: () => import('@/views/ChartPanel'), meta: { activeMenu: '/chartpanel/create' }},
  { path: '/:projectid/fullscreendb/:id', component: () => import('@/views/Dashboard/fullScreenDb') },
  { path: '/:projectid/source', component: () => import('@/views/source/index') },
  { path: '/:projectid/source/:id/table', component: () => import('@/views/source/table') },
  { path: '/:projectid/lab', component: () => import('@/views/lab/index') },
  { path: '/:projectid/task', component: () => import('@/views/task/index') },

  { path: '*', component: () => import('@/views/NotFound') }
]

export const menuRoutes = routes

const router = new VueRouter({
  routes
})

const whiteList = ['/login', '/auth-redirect', '/signup', '/screenshot']

router.beforeEach(async(to, from, next) => {
  if (store.state.user.username) {
    if (to.name === 'login') {
      next({ path: '/' })
    } else {
      next()
    }
  } else if (whiteList.includes(to.path)) { // 在免登录白名单，直接进入
    next()
  } else {
    try {
      await store.dispatch('user/GetUserInfo')
      next()
    } catch (err) {
      next(`/login?redirect=${to.path}`) // 否则全部重定向到登录页
    }
  }
})

export default router
