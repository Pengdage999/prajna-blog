<template>
  <a-button type="danger" @click="logout">退出</a-button>
</template>

<script>
export default {
  methods: {
    logout() {
      // 清理tokens
      window.sessionStorage.clear('token')
      this.$router.push('/login')
      // 跳转到login
    },
  },
}
</script>
