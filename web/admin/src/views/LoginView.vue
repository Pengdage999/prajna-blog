<template>
  <div class="container">
    <div class="loginBox">
      <a-form-model
        ref="loginFormRef"
        :rules="rules"
        :model="formdata"
        class="loginForm"
      >
        <a-form-model-item prop="username">
          <a-input
            v-model="formdata.username"
            placeholder="请输入用户名"
            type="username"
          >
            <a-icon
              slot="prefix"
              type="user"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>
        <a-form-model-item prop="password">
          <a-input
            v-model="formdata.password"
            placeholder="请输入密码"
            type="password"
            v-on:keyup.enter="login"
          >
            <a-icon
              slot="prefix"
              type="lock"
              style="color: rgba(0, 0, 0, 0.25)"
            />
          </a-input>
        </a-form-model-item>
        <a-form-model-item>
          <a-button type="primary" style="margin: 10px" @click="login"
            >登录</a-button
          >
          <a-button type="info" style="margin: 10px" @click="resetForm"
            >取消</a-button
          >
        </a-form-model-item>
      </a-form-model>
    </div>
  </div>
</template>

<script>
// import axios from 'axios'
// 数据绑定
export default {
  data() {
    return {
      formdata: {
        username: '',
        password: '',
      },
      rules: {
        username: [
          { required: true, message: '请输入用户名', trigger: 'blur' },
          {
            min: 4,
            max: 12,
            message: '用户名字符数量必须在4和12之间',
            trigger: 'blur',
          },
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          {
            min: 6,
            max: 20,
            message: '密码字符数量必须在4和12之间',
            trigger: 'blur',
          },
        ],
      },
    }
  },
  methods: {
    resetForm() {
      this.$refs.loginFormRef.resetFields()
    },
    login() {
      this.$refs.loginFormRef.validate(async (valid) => {
        if (!valid) return this.$message.error('输入非法数据，重新输入')
        const { data: res } = await this.$http.post('login', this.formdata)
        console.log(res)
        if (res.status !== 200) return this.$message.error(res.message)
        window.sessionStorage.setItem('token', res.token)
        this.$router.push('index')
      })
    },
  },
}
</script>

<style scoped>
.container {
  height: 100%;
  background-color: #282c34;
  display: flex;
  justify-content: center;
  align-self: center;
}

.loginBox {
  width: 450px;
  height: 300px;
  background-color: #fff;
  position: absolute;
  top: 50%;
  left: 70%;
  transform: translate(-50%, -50%);
  border-radius: 9px;
}

.loginForm {
  width: 100%;
  position: absolute;
  bottom: 10%;
  padding: 0 20px;
  box-sizing: border-box;
}
</style>
