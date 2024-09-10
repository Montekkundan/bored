import { createRouter, createWebHistory } from 'vue-router'
import Changelog from '../views/Changelog.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'changelog',
      component: Changelog
    },
    // Add more routes here in the future
  ]
})

export default router