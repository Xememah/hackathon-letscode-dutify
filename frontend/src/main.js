// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuetify from 'vuetify'
import {API_URL} from './constants.js'
import Resource from 'vue-resource'
Vue.use(Resource)

import 'vuetify/dist/vuetify.css'
import App from './App'
import router from './router'
import middleware from '@/middleware/middleware'


// Vue.use(Vuetify)
Vue.use(Vuetify, {
  theme: {
    primary: '#000',
    secondary: '#b0bec5',
    accent: '#8c9eff',
    error: '#b71c1c'
  }
})
Vue.config.productionTip = false

middleware(router)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: { App }
})

Vue.http.interceptors.push((request, next) => {
  let token = window.localStorage.getItem('token')

  if (token) {
      request.headers = request.headers || {}
      request.headers.Authorization = `Bearer ${token}`
  }
  next((response) => {
      if (response.status === 401) {
          return Vue.http.get(API_URL+'token').then((result) => {
              window.localStorage.setItem('token', result.data.token)
              return Vue.http(request).then((response) => {
                  return response
              })
          })
          .catch(() => {
              return router.go({name: 'login'})
          })
      }
  })
})