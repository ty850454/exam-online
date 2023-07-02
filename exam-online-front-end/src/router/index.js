import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: HomeView
    },
    {
      path: '/home',
      component: HomeView
    },
    {
      path: '/admin/login',
      component: () => import('../views/admin/Login.vue')
    },
    {
      path: '/admin',
      component: () => import('../views/admin/Layout.vue'),
      children: [
        {
          path: '/admin/dashboard',
          component: () => import('../views/admin/Dashboard.vue')
        },
        {
          path: '/admin/exam',
          redirect:'/admin/exam/list',
          children: [
            {
              path: '/admin/exam/list',
              component: () => import('../views/admin/exam/List.vue')
            },
            {
              path: '/admin/exam/create',
              component: () => import('../views/admin/exam/Create.vue')
            },
          ]
        },
        {
          path: '/admin/resource',
          redirect:'/admin/resource/question',
          component: () => import('../views/admin/resource/Layout.vue'),
          children: [
            {
              path: '/admin/resource/question',
              component: () => import('../views/admin/resource/question/List.vue')
            },
            {
              path: '/admin/resource/user',
              component: () => import('../views/admin/resource/user/List.vue')
            },
          ]
        },
        {
          path: '/admin/manage',
          redirect:'/admin/manage/admin',
          component: () => import('../views/admin/manage/Layout.vue'),
          children: [
            {
              path: '/admin/manage/admin',
              component: () => import('../views/admin/manage/Admin.vue')
            },
            {
              path: '/admin/manage/api',
              component: () => import('../views/admin/manage/Api.vue')
            },
            {
              path: '/admin/manage/system',
              component: () => import('../views/admin/manage/System.vue')
            },
          ]
        },
      ]
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue')
    }
  ]
})

export default router
