import axios from 'axios'
import { Message } from 'element-ui'
// import Cookies from 'js-cookie'
import store from '../store'

const fetchInstance = axios.create({
  baseURL: process.env.VUE_APP_BASE_API
  // timeout: 2000,
})

// request拦截器
fetchInstance.interceptors.request.use(config => {
  // console.log(store, store.getters)
  if (store.getters.token) {
    config.headers['x-Token'] = store.getters.token
  }
  if (store.getters.code) {
    config.headers['x-code'] = store.getters.code
  }
  // if (Cookies.get('x-Token')) {
  //   config.headers['x-token'] = Cookies.get('x-Token')
  // }
  config.withCredentials = true
  return config
}, error => {
  console.log(error) // for debug
  Promise.reject(error)
})

// response拦截器
fetchInstance.interceptors.response.use(
  response => {
    const res = response.data
    if (res.code !== 20000) {
      Message({
        message: res.message,
        type: 'error',
        duration: 5 * 1000
      })
      return Promise.reject('request error')
    } else {
      return response.data
    }
  },
  error => {
    if (error.message !== 'cancel') {
      console.log('err:' + error)// for debug
      Message({
        message: error.message,
        type: 'error',
        duration: 5 * 1000
      })
    }

    // Raven.captureException(error)
    return Promise.reject(error)
  }
)

export default fetchInstance
