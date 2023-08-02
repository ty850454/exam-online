<script setup>
import '@/assets/admin.css'
import {RouterView} from "vue-router";
import IconLogo from "@/components/icons/IconLogo.vue";
import IconHome from "@/components/icons/IconHome.vue";
import IconAddFill from "@/components/icons/IconAddFill.vue";
import IconUserFill from "@/components/icons/IconUserFill.vue";
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()

</script>

<template>

  <div class="main-layout">

    <header class="admin-layout">
      <div>
        <IconLogo class="logo"/>
        <span
            style="font-size: 20px;font-weight: 900;margin-left: 10px;margin-right: 40px;vertical-align: middle;">考试系统</span>
        <div class="menu">
          <RouterLink to="/admin/dashboard">
            <IconHome/>
          </RouterLink>
          <RouterLink to="/admin/exam"><span class="menu-text">考试</span></RouterLink>
<!--          <RouterLink to="/admin/resource"><span class="menu-text">资源</span></RouterLink>-->
          <RouterLink to="/admin/manage"><span class="menu-text">管理</span></RouterLink>
        </div>
      </div>
      <div class="buttons">
        <div class="bo">
          <IconAddFill style="color: #fff"/>
        </div>
        <el-dropdown>
          <div class="bo">
              <IconUserFill style="color: #fff"/>
          </div>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item @click="userStore.logout">退出</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </header>
    <div style="flex-grow: 1">
      <RouterView/>
    </div>
  </div>

</template>

<style scoped>


.main-layout {
  display: flex;
  flex-direction: column;
  height: 100%;
}


.admin-layout {
  display: flex;
  justify-content: space-between;
  background-color: #20287e;
  color: #fff;
  height: 58px;
  align-items: center;
}

.admin-layout :deep(.icon) {
  vertical-align: middle;
}

.admin-layout :deep(.menu) {
  font-size: 20px;
}

.logo {
  font-size: 40px;
  margin-left: 20px;
}

.menu {
  display: inline-block;
  height: 58px;
  vertical-align: middle;
  line-height: 58px;
}


.menu :deep(a) {
  padding: 0 25px;
  display: inline-block;
  color: #fff;
  text-decoration: none;
  height: 58px;
}

.menu-text {
  font-size: 15px;
}

.menu :deep(a.router-link-active) {
  background-color: #353d97;
}

.buttons {
  display: flex;
  height: 58px;
  margin-right: 20px;
}

.buttons .bo {
  height: 58px;
  width: 58px;
  font-size: 25px;
  line-height: 50px;
  text-align: center;
  cursor: pointer;
}

.buttons .bo:hover {
  background-color: #353d97;
}


</style>