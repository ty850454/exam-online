<script setup>

import {RouterView} from "vue-router";
import IconHome from "@/components/icons/IconHome.vue";
import IconQuestionFill from "@/components/icons/IconQuestionFill.vue";
import IconUserGroupFill from "@/components/icons/IconUserGroupFill.vue";
</script>

<template>
  <div class="resource-layout">
    <div class="left">
      <div class="title">
        资源
      </div>
      <div class="menus">
        <RouterLink class="menu" to="/admin/resource/question"><IconQuestionFill/>&nbsp;&nbsp;试题库</RouterLink>
        <RouterLink class="menu" to="/admin/resource/user"><IconUserGroupFill/>&nbsp;&nbsp;考生</RouterLink>
      </div>
    </div>
    <div style="flex-grow: 1"><RouterView /></div>
  </div>
</template>

<style scoped>

.resource-layout {
  display: flex;
  height: 100%;
}


.resource-layout > .left {
  width: 200px;
  background-color: #fafbfc;
  border-right: #e9ebed 1px solid;
}


.resource-layout > .left > .title {
  font-size: 16px;
  padding: 20px 30px;
  border-bottom: 1px solid #e9ebed;
  color: #333333;
}

.resource-layout > .left > .title::before {
  content: " ";
  width: 3px;
  height: 14px;
  background-color: #20287e;
  margin-right: 10px;
  display: inline-block;
  vertical-align: middle;
}

.resource-layout > .left > .menus {

}

.resource-layout > .left > .menus > .menu {
  display: block;
  text-decoration: none;
  height: 45px;
  line-height: 45px;
  padding-left: 30px;
}

.resource-layout > .left > .menus > .menu.router-link-active {
  background-color: #e9ebed;
  border-right: 3px solid #20287e;
}



</style>