<template>
  <div class="footer">
    <v-footer padless color="indigo darken-2">
      <v-row justify="center" no-gutters>
        <v-col class="py-4 text-center white--text" cols="12">{{ new Date().getFullYear() }} - Prajna Blog</v-col>
  
        <div class="text-center white--text">
          <a class="text-center white--text" href="https://beian.miit.gov.cn/">{{icp_record}}</a>
        </div>
      </v-row>
    </v-footer>
  </div>
</template>

<script>
export default {
  data() {
    return {
      icp_record: ''
    }
  },
  mounted() {
    this.icp_record = "渝ICP备2024028100号-1"
  }
}
</script>

<style scoped>
.footer {
  height: 100%;
  background-color: #eee;
  text-align: center;
}
.footer span {
  height: 100%;
  font-size: 14px;
}
</style>
