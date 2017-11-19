import Vue from 'vue'
import Router from 'vue-router'
import Homepage from '@/components/Homepage'
import Project from '@/components/project/Project'
import Overview from '@/components/project/Overview'
import Duty from '@/components/project/Duty'
import Score from '@/components/project/Score'
import Profile from '@/components/project/Profile'
import Login from '@/components/Login'

Vue.use(Router)

export default new Router({
  routes: [{
      path: '/',
      component: Homepage
    },
    {
      path: '/project/:projectId',
      component: Project,
      children: [
        {
          path: '',
          redirect: 'home'
        },
        {
          path: 'home',
          name: 'home',
          component: Overview
        },
        {
          path: 'score',
          name: 'score',
          component: Score
        },
        {
          path: 'profile',
          name: 'profile',
          component: Profile
        },
        {
          path: 'duty/:dutyId',
          name: 'duty',
          component: Duty
        }
      ]
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    }
  ],
  history: true
})
