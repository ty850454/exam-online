<script setup>
import {ref} from "vue";
import IconUserFill from "@/components/icons/IconUserFill.vue";
import IconLock from "@/components/icons/IconLock.vue";
import { useRouter } from "vue-router"
import { useUserStore } from '@/stores/user'
import { ElMessage } from 'element-plus'

const userStore = useUserStore()

let router = useRouter()


const loginForm = ref({
  username: "",
  password: "",
})

let loginRules = {

}

const loading = ref(false)

function onLoginClick() {
  userStore.login(loginForm.value.username, loginForm.value.password).then(() => {
    router.push('/admin')
  }).catch(err => {
    err.then(ej => {
      ElMessage({
        showClose: true,
        message: ej.error,
        type: 'error',
      })
    })
  })
}


</script>

<template>

  <div class="login-container">

    <div class="login">
      <el-form :model="loginForm" :rules="loginRules" class="login-form" autocomplete="on" label-position="left">

        <h3 class="title">系统登录</h3>

        <el-form-item prop="username">
          <el-input
              v-model="loginForm.username"
              class="w-50 m-2"
              placeholder="用户名"
              :prefix-icon="IconUserFill"
              auto-complete="off"
              tabindex="1"
              size="large"
          />
        </el-form-item>

        <el-form-item prop="password">
          <el-input
              v-model="loginForm.password"
              class="w-50 m-2"
              placeholder="密码"
              :prefix-icon="IconLock"
              auto-complete="off"
              autocomplete-="off"
              tabindex="2"
              show-password
              size="large"
          />
        </el-form-item>

        <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click="onLoginClick">登录</el-button>

      </el-form>


    </div>

  </div>


</template>

<style scoped>

.login-container {
  min-height: 100%;
  background-color: #2d3a4b;
  overflow: hidden;
}

.login {
  width: 520px;
  margin: 0 auto;
  padding: 160px 35px 0;
}

.title {
  color: #eeeeee;
  margin: 0 auto 40px auto;
  text-align: center;
  font-size: 26px;
  font-weight: 700;
}

</style>