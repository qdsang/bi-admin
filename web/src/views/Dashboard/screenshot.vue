<template>
  <div>
    <dashboardItem :dashboard="currentDashboard" mode="view" />
  </div>
</template>
<script>
import dashboardItem from './dashboardItem'
import { getdDashboardById } from '@/api/dashboard'

export default {
  components: { dashboardItem },
  data() {
    return {
      currentDashboard: undefined
    }
  },
  created() {
    // console.log(this.$route)
    const code = this.$route.query.code
    if (!code) {
      return
    }

    this.$store.dispatch('user/LoginByCode', code).then(() => {
    })

    getdDashboardById(this.$route.query.id).then(resp => {
      resp.data.content = JSON.parse(resp.data.content)
      this.currentDashboard = resp.data
    })
  }
}
</script>
