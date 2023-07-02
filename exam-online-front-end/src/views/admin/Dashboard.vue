<script setup>


import {ref} from "vue";

const activeTab = ref('1')
</script>

<template>
  <div class="dashboard-container">
    <div class="above">
      <div class="card buttons">
        <div>创建考试</div>
        <div>试题库</div>
        <div>考生管理</div>
        <div>批量导入</div>
      </div>
      <div class="card">
        二维码
      </div>
      <div class="card" style="padding: 20px">
        <table style="width: 100%; text-align: center">
          <tr>
            <td>在线人数</td>
            <td>考生数量</td>
            <td>考试次数</td>
          </tr>
          <tr>
            <td>1</td>
            <td>2</td>
            <td>8</td>
          </tr>
        </table>
      </div>
      <div class="card">
        <el-tabs v-model="activeTab">
          <el-tab-pane label="系统公告" name="1">示例1</el-tab-pane>
          <el-tab-pane label="快速入门" name="2">示例2</el-tab-pane>
        </el-tabs>
      </div>
    </div>
  </div>

</template>

<style scoped>

.dashboard-container {
  //padding: 20px;
  background-color: #f3f5f7;
  height: 100%;
}

.dashboard-container > .above {
  display: grid;
  grid-template-columns: 4fr 2fr;
  padding: 20px;
  background-color: #f3f5f7;
  grid-gap: 20px;
}

.card {
  background-color: #ffffff;
}

.card.buttons {
  padding: 20px;
  display: grid;
  grid-template-columns: 1fr 1fr;
}

.card.buttons > div {
  height: 60px;
}

</style>