export default function RedirectAuthenticated(router) {
  router.beforeEach((to, from, next) => {
    let token = window.localStorage.getItem('token')
    if ((token) && (to.name === 'login' || to.name === 'register')) {
      router.push('/')
    }
    next()
  })
}
