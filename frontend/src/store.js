import Vue from 'vue'
const jwtDecode = require('jwt-decode')

let store = {
  bus: new Vue(),
  user: {
    token: () => {
      window.localStorage.getItem('token')
    },
    getUser: () => {
      if (window.localStorage.getItem('token')) {
        return jwtDecode(window.localStorage.getItem('token')).User
      }
      return null;
    },
  },
  project: {
      users: [],
      ranking: [],
      max: 10,
  }
}

export default {
  store,
  install(Vue, options) {
    //window.localStorage.setItem('token', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTExNDQ3NTksIlVzZXIiOnsiaWQiOjMsImNyZWF0ZWRfYXQiOiIyMDE3LTExLTE5VDAyOjI1OjU5LjI1ODkyNDc5NVoiLCJuYW1lIjoiTWFjaWVqIE1pb25za293c2tpIiwiZW1haWwiOiJ0ZXN0MEBtYWNpZWttbS5uZXQifX0.DFccP5wc3QKmgQC-7PoNf6D-Ui29IdIJUGXALF30-kc')
    Vue.prototype.$store = store;
  }
}
