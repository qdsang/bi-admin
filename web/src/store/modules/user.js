import { getToken, setToken } from '@/utils/auth'
import { getUserInfo, logUserOut } from '@/api/user'

const user = {
  namespaced: true,
  state: {
    username: undefined,
    token: getToken(),
    code: '',
    avatar: 'https://wpimg.wallstcn.com/3fce7247-d863-4e3d-a0de-d30aaef7358a'
  },
  mutations: {
    SET_USERINFO: (state, userInfo) => {
      state.username = userInfo.username
    },
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_CODE: (state, code) => {
      state.code = code
    }
  },
  actions: {
    async GetUserInfo({ commit }) {
      const { data } = await getUserInfo()
      commit('SET_USERINFO', data)
    },
    // 用户名登录
    Login({ commit }, token) {
      return new Promise((resolve, reject) => {
        commit('SET_TOKEN', token)
        setToken(token)
        resolve()
      })
    },
    async FedLogOut({ commit }) {
      commit('SET_TOKEN', '')
      return await logUserOut()
    },
    LoginByCode({ commit }, code) {
      return new Promise((resolve, reject) => {
        commit('SET_CODE', code)
        resolve()
      })
    }
  }
}

export default user
