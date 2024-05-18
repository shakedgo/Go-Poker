import { createRouter, createWebHistory } from 'vue-router'
import { type RouteRecordRaw } from 'vue-router';
import { useAuthStore } from '@/stores/auth'
import HomeView from '../views/HomeView.vue'
import ProtectedView from '../views/ProtectedView.vue'
import LoginComp from '@/components/Game/LoginComp.vue';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('../views/AboutView.vue'),
    meta:{
      requiresAuth: true
    }
  },
  {
    path: '/game',
    name: 'Game',
    component: () => import('../views/GameView.vue'),
    meta:{
      requiresAuth: true
    }
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginComp
  },
  {
    path: '/protected',
    name: 'Protected',
    component: ProtectedView,
    meta:{
      requiresAuth: true
    }
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});
router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()

  // Check if the user is authenticated
  if (to.meta.requiresAuth && !authStore.isLoggedIn) {
    // Redirect to login page
    next({ name: 'Login' })
  } else {
    next()
  }
})
export default router