<script setup>

import {ref} from "vue";
import IconRight from "@/components/icons/IconRight.vue";

const form = ref({})
const activeStep = ref(0)


function onSaveInfoClick(nextStep) {
  if (nextStep) {
    activeStep.value = 1
  }
}


</script>

<template>
<div class="create-exam">
  <div>
    <el-steps :active="activeStep" finish-status="success" align-center>
      <el-step title="考试信息" />
      <el-step title="设计试卷" />
      <el-step title="发布考试" />
    </el-steps>
  </div>


  <el-tabs
      v-model="activeStep"
      type="card">
    <el-tab-pane label="step1" :name="0">

      <el-form :model="form" label-width="120px">

        <el-form-item label="考试名称">
          <el-input v-model="form.name" />
        </el-form-item>

        <el-form-item label="考试分类">
          <el-select v-model="form.region" placeholder="">
            <el-option label="Zone one" value="shanghai" />
            <el-option label="Zone two" value="beijing" />
          </el-select>
        </el-form-item>

        <el-form-item label="考生参加方式">
          <el-radio-group v-model="form.resource">
            <el-radio label="免登录" />
            <el-radio label="免登录+口令" />
          </el-radio-group>
        </el-form-item>

        <el-form-item label="开考时间">
          <el-date-picker
              v-model="form.date1"
              type="datetime"
              placeholder="选择时间"
              style="width: 100%"
          />
        </el-form-item>

        <el-form-item label="考试时长">
          <el-input v-model="form.name" />
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="onSaveInfoClick(true)">保存&设计试卷</el-button>
          <el-button  @click="onSaveInfoClick(false)">仅保存</el-button>
          <el-button>取消</el-button>
        </el-form-item>
      </el-form>

    </el-tab-pane>

    <el-tab-pane label="step2" :name="1" class="exam-designer">

      <div class="left">
        <div class="add-group">添加题目分组<IconRight/></div>
        <div class="info">共 0 题 0 分</div>

      </div>

      <div class="right">
        <div class="empty">请添加题目分组</div>
      </div>


    </el-tab-pane>
    <el-tab-pane label="step3" :name="2">Role</el-tab-pane>
  </el-tabs>





</div>








</template>

<style scoped>


.create-exam {
  background-color: #f3f5f7;
  min-height: 500px;
}

.create-exam :deep(.el-tabs__header) {
  display: none;
}

.exam-designer {
  display: grid;
  grid-template-columns: 4fr 9fr;
  min-height: 400px;
  grid-column-gap: 20px;
  padding: 0 20px;
}

.exam-designer > div {
  background-color: #fff;
  border-radius: 8px;
}

.exam-designer > .left {
  background-color: #fff;
}

.exam-designer .add-group {
  position: relative;
  background-color: #128beb;
  color: #fff;
  height: 45px;
  line-height: 45px;
  font-size: 16px;
  text-align: center;
  border-top-left-radius: 8px;
  border-top-right-radius: 8px;
  cursor: pointer;
}

.exam-designer .add-group:hover {
  background-color: #1f96f4;
}

.exam-designer .add-group :deep(.icon) {
  position: absolute;
  right: 5px;
  top: 15px;
}

.exam-designer > .left > .info {
  padding: 5px 10px;
  border-bottom: #ececec 1px solid;
}

.exam-designer > .right {
  position: relative;
}

.exam-designer > .right > .empty {
  position: absolute;
  left: 0;
  right: 0;
  top: 50%;
  text-align: center;
  font-size: 16px;
  font-weight: lighter;
  color: gray;
}

</style>