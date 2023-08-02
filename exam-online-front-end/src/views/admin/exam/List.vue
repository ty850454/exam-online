<script setup>

import {RouterView} from "vue-router";
import IconAddFill from "@/components/icons/IconAddFill.vue";
import {ref} from "vue";
import IconAwaitFill from "@/components/icons/IconAwaitFill.vue";
import IconSuccessFill from "@/components/icons/IconSuccessFill.vue";
import IconChoiceFill from "@/components/icons/IconChoiceFill.vue";
import IconStarFill from "@/components/icons/IconStarFill.vue";
import IconStoreFill from "@/components/icons/IconStoreFill.vue";
import IconDeleteFill from "@/components/icons/IconDeleteFill.vue";
import IconFundFill from "@/components/icons/IconFundFill.vue";

import { useRouter } from "vue-router"
import {get} from "@/utils/http";
let router = useRouter()


const query = ref({
  date: null,
  type: null,
})

const options = [
  {
    label: '哈哈',
    value: 1,
  }
]

const pageSize = 10
const total = ref(0)
const examPapers = ref([])

function onCurrentChanged(page) {
  get('/api/exam?pageNum='+page+'&pageSize='+pageSize)
      .then(response => {
        examPapers.value = response.data
        total.value = response.total
      })
}
onCurrentChanged(1)

function onCreateClick() {
  router.push('/admin/exam/create')
}

</script>

<template>
  <div class="exam-container">
    <div class="left">

      <div class="above">
        <p>
          <el-button type="primary" size="large" @click="onCreateClick">
            <IconAddFill/>&nbsp;&nbsp;创建考试
          </el-button>
        </p>
        <p>
          <el-button size="large">
            <IconFundFill/>&nbsp;&nbsp;考试分析
          </el-button>
        </p>
      </div>
      <div class="below">
        <div class="button"><IconStoreFill/>&nbsp;&nbsp;所有考试</div>
        <div class="button"><IconStarFill/>&nbsp;&nbsp;星标考试</div>
        <el-divider/>
        <div class="button"><IconAwaitFill/>&nbsp;&nbsp;进行中</div>
        <div class="button"><IconChoiceFill/>&nbsp;&nbsp;未开始</div>
        <div class="button"><IconSuccessFill/>&nbsp;&nbsp;已结束</div>
        <el-divider/>
        <div class="button"><IconDeleteFill/>&nbsp;&nbsp;回收站</div>
      </div>
    </div>
    <div class="right">

      <div class="query">
        <el-date-picker
            v-model="query.date"
            type="date"
            placeholder="创建日期"
            size="large"
        />

        <el-select v-model="query.type" placeholder="考试分类" size="large">
          <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value"
          />
        </el-select>
      </div>
      <div class="body">

        <div class="cards">


          <el-card class="card" shadow="never" header="春季考试1">
            <template #header>
              <div class="card-header">
                <el-tag type="info" effect="dark" disable-transitions>进行中</el-tag>
                <span style="margin-left: 10px;">春季考试1</span>
              </div>
            </template>
            <div class="card-body">
              <div>2023-04-04 20:30 ~ 2023-05-01 20:30</div>
              <div>
                <el-button class="button" text>编辑</el-button>
                <el-button class="button" text>成绩</el-button>
                <el-button class="button" text>复制考试连接</el-button>
              </div>
            </div>
          </el-card>

          <el-card class="card" shadow="never" header="春季考试2">
            <template #header>
              <div class="card-header">
                <el-tag type="success" effect="dark" disable-transitions>已结束</el-tag>
                <span style="margin-left: 10px;">春季考试1</span>
              </div>
            </template>
            <div class="card-body">
              <div>2023-04-04 20:30 ~ 2023-05-01 20:30</div>
              <div>
                <el-button class="button" text>成绩</el-button>
              </div>
            </div>
          </el-card>

          <el-card v-for="examPaper in examPapers" class="card" shadow="never" :header="examPaper.title">
            <template #header>
              <div class="card-header">
                <el-tag type="success" effect="dark" disable-transitions>{{examPaper.status}}</el-tag>
                <span style="margin-left: 10px;">{{examPaper.title}}</span>
              </div>
            </template>
            <div class="card-body">
              <div>2023-04-04 20:30 ~ 2023-05-01 20:30</div>
              <div>
                <el-button class="button" text>成绩</el-button>
              </div>
            </div>
          </el-card>

        </div>

        <div class="pagination">
          <el-pagination background layout="prev, pager, next" :total="total" @current-change="onCurrentChanged"/>
        </div>
      </div>

    </div>
  </div>

  <RouterView/>
</template>

<style scoped>
.exam-container {
  display: grid;
  grid-template-columns: 240px 1fr;
  height: 100%;
}

.left {
  background-color: #fafbfc;
}

.left .above {
  margin: 40px 40px;
}

.left .above :deep(button) {
  width: 100%;
}

.left .above p {
  margin-top: 10px;
}

.left .below > .button {
  padding: 10px 40px;
  cursor: pointer;
}


.left .below > .button:hover {
  background-color: #f3f3f3;
}


.left > .below > :deep(.el-divider) {
  margin: 10px auto;
  padding: 0 20px;
  width: 180px;
}


.right > .query {
  width: 500px;
  margin: 40px auto 20px;

}


.right > .query > div {
  margin-left: 10px;

}

.right > .body {
  margin: 0 20px;
}

.right > .body > .cards {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  align-items: center;
}

.card-body {
  display: flex;
  align-items: center;
  justify-content: space-between;
}


.right > .body > .pagination :deep(.el-pagination) {
  justify-content: center;
}


.right > .body > .cards > .card :deep(.el-card__header) {
  background-color: #f3f5f7;
}


.card+.card {
  margin-top: 20px;
}



</style>