import { createRouter, createWebHistory } from 'vue-router'
import FileView from '../views/FileView.vue'
import HomeView from '@/views/HomeView.vue'
import AllView from '@/views/AllView.vue'
import RecentsView from '@/views/RecentsView.vue'
import TrashView from '@/views/TrashView.vue'
import BoxView from '@/views/all/BoxView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'main',
      redirect: 'home'
    },
    {
      path: '/home',
      name: 'home',
      component: HomeView,
      redirect: '/home/all',
      children: [
        {
          path: 'all',
          name: 'all',
          component: AllView,
          redirect: '/home/all/box',
          children: [
            {
              path: 'box',
              name: 'box',
              component: BoxView,
            }
          ]
        },
        {
          path: 'recents',
          name: 'recents',
          component: RecentsView,
        },
        {
          path: 'trash',
          name: 'trash',
          component: TrashView,
        }
      ]
    }
  ]
})

export default router
