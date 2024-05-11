import { createRouter, createWebHistory } from 'vue-router'
import { type RouteRecordRaw } from 'vue-router';
import HomeView from '../views/HomeView.vue'
import ProtectedView from '../views/ProtectedView.vue'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/about',
    name: 'about',
    component: () => import('../views/AboutView.vue')
  },
  {
    path: '/game',
    name: 'game',
    component: () => import('../views/GameView.vue')
  },
  {
    path: '/protected',
    name: 'Protected',
    component: ProtectedView,
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router