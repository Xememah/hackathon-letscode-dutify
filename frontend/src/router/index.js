import Vue from 'vue'
import Router from 'vue-router'
import Homepage from '@/components/Homepage'
import Score from '@/components/Score'
import Profile from '@/components/Profile'
import Challenge from '@/components/Challenge'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Homepage',
      component: Homepage
    },
    {
      path: '/score',
      name: 'Score',
      component: Score
    },
    {
      path: '/profile',
      name: 'Profile',
      component: Profile
    },
    {
      path: '/challenge',
      name: 'Challenge',
      component: Challenge
    }
  ]
})
