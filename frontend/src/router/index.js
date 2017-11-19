import Vue from 'vue'
import Router from 'vue-router'
import Homepage from '@/components/Homepage'
import Score from '@/components/Score'
import Profile from '@/components/Profile'
import Challenge from '@/components/Challenge'
import Login from '@/components/Login'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'homepage',
      component: Homepage
    },
    {
      path: '/score',
      name: 'score',
      component: Score
    },
    {
      path: '/profile',
      name: 'profile',
      component: Profile
    },
    {
      path: '/challenge',
      name: 'challenge',
      component: Challenge
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    }
  ],
  history: true
})
