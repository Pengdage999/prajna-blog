<template>
  <a-layout-sider breakpoint="lg" v-model="collapsed">
    <div class="log">
      <span>{{ collapsed ? 'Blog' : 'Prajna Blog' }}</span>
    </div>
    <a-menu class="navigator" theme="dark" mode="inline" @click="goToPage">
      <a-menu-item key="index">
        <a-icon type="dashboard" /><span>仪表盘</span>
      </a-menu-item>
      <a-sub-menu>
        <span slot="title">
          <a-icon type="file-word" /><span>文章管理</span>
        </span>
        <a-menu-item key="addart">
          <a-icon type="edit" /><span>撰写文章</span>
        </a-menu-item>
        <a-menu-item key="artlist">
          <a-icon type="ordered-list" /><span>文章列表</span>
        </a-menu-item>
      </a-sub-menu>
      <a-menu-item key="catelist">
        <a-icon type="book" /><span>分类列表</span>
      </a-menu-item>
      <a-menu-item key="userlist">
        <a-icon type="user" /><span>用户列表</span>
      </a-menu-item>
      <a-menu-item key="profile">
        <a-icon type="setting" />
        <span>个人设置</span>
      </a-menu-item>
    </a-menu>
  </a-layout-sider>
</template>

<script>
export default {
  data() {
    return {
      collapsed: false,
    }
  },
  methods: {
    goToPage(item) {
      this.$router.push('/' + item.key).catch((err) => err)
    },
  },
}
</script>

<style scoped>
.log {
  height: 32px;
  margin: 16px;
  background-color: #fff;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 20px;
}

.navigator {
  text-align: left;
}
</style>
