const getters = {
  username: state => state.user.username,
  avatar: state => state.user.avatar,
  token: state => state.user.token,
  code: state => state.user.code
}
export default getters
