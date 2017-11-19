// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuetify from 'vuetify'
import {
  API_URL
} from './constants.js'
import Resource from 'vue-resource'
import 'vuetify/dist/vuetify.css'
import App from './App'
import router from './router'
import middleware from '@/middleware/middleware'
import store from './store.js'

Vue.use(store)
Vue.use(Resource)


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

Vue.http.interceptors.push((request, next) => {
  let token = window.localStorage.getItem('token')

  if (token) {
    request.headers.set('Authorization', `Bearer ${token}`)
  }
  next((response) => {
    if (response.status === 401) {
      return Vue.http.post(API_URL + 'accounts/token/').then((result) => {
          return Vue.http(request).then((response) => {
            return response
          })
        })
        .catch(() => {
          return router.push('login')
        })
    }
  })
})

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {
    App
  }
})

