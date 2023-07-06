<script setup>

import {ref} from "vue";
import IconRight from "@/components/icons/IconRight.vue";

const form = ref({})
const activeStep = ref(1)


const questionSet = ref([
])

function onAddQuestionGroup() {
  questionSet.value.push({
    title: '分组' + (questionSet.value.length + 1),
    questions: []
  })
}

function onSaveInfoClick(nextStep) {
  if (nextStep) {
    activeStep.value = 1
  }
}

const dialogVisible = ref(false)


const dialogQuestion = ref({})


function onAddQuestionSave() {
  dialogVisible.value = false
}

function onAddQuestion() {
  dialogVisible.value = true
  dialogQuestion.value = {
    type: 1,
    score: 0,
    title: '',
    answer: '',
    options: [
      {
        code: 'A',
        content: ''
      },
      {
        code: 'B',
        content: ''
      },
      {
        code: 'C',
        content: ''
      },
      {
        code: 'D',
        content: ''
      },
    ]
  }
}

</script>

<template>
  <div class="create-exam">
    <div class="step card">
      <el-steps :active="activeStep" finish-status="success" align-center>
        <el-step title="考试信息"/>
        <el-step title="设计试卷"/>
        <el-step title="发布考试"/>
      </el-steps>
    </div>


    <el-tabs
        v-model="activeStep"
        type="card">
      <el-tab-pane label="step1" :name="0">

        <el-form :model="form" label-width="120px">

          <div class="card input-200">

            <div class="title">
              基本信息
            </div>

            <el-form-item label="考试名称">
              <el-input v-model="form.name"/>
            </el-form-item>

            <el-form-item label="考试分类">
              <el-select v-model="form.region" placeholder="">
                <el-option label="Zone one" value="shanghai"/>
                <el-option label="Zone two" value="beijing"/>
              </el-select>
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
              <el-input v-model="form.name"/>
            </el-form-item>
          </div>


          <div class="card">

            <div class="title">
              参加方式
            </div>


            <el-form-item label="考生参加方式">
              <el-radio-group v-model="form.resource">
                <el-radio label="免登录"/>
                <el-radio label="免登录+口令"/>
              </el-radio-group>
            </el-form-item>


          </div>


          <el-form-item>
            <el-button type="primary" @click="onSaveInfoClick(true)">保存&设计试卷</el-button>
            <el-button @click="onSaveInfoClick(false)">仅保存</el-button>
            <el-button>取消</el-button>
          </el-form-item>
        </el-form>

      </el-tab-pane>

      <el-tab-pane label="step2" :name="1" class="exam-designer">

        <div class="left">

          <el-affix target=".left" :offset="80">
            <div class="navigator">
              <div class="add-group" @click="onAddQuestionGroup">添加题目分组
                <IconRight/>
              </div>
              <div class="info">共 0 题 0 分</div>

              <div class="foot">
                <el-button type="primary" @click="activeStep=2">保存与发布</el-button>
                <el-button>仅保存</el-button>
              </div>
            </div>
          </el-affix>

        </div>

        <div class="right">
          <div class="card" v-if="questionSet.length === 0">
            <div class="empty">
              请添加题目分组
            </div>
          </div>
          <div class="card" v-if="questionSet.length > 0" v-for="question in questionSet">
            <div>{{ question.title }} <span>共0题，0分</span></div>
            <div class="empty" v-if="question.questions.length === 0">
              请向分组添加题目
            </div>
            <div>

              <el-dropdown @command="onAddQuestion">
                <el-button type="primary">
                  <el-button type="primary">新增题目</el-button>
                </el-button>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item>单选题</el-dropdown-item>
                    <el-dropdown-item>多选题</el-dropdown-item>
                    <el-dropdown-item>判断题</el-dropdown-item>
                    <el-dropdown-item>填空题</el-dropdown-item>
                    <el-dropdown-item>问答题</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>


            </div>
          </div>
        </div>


      </el-tab-pane>
      <el-tab-pane label="step3" :name="2">Role</el-tab-pane>
    </el-tabs>

    <el-dialog v-model="dialogVisible" title="新增题目" :width="800" class="question-dialog">
      <el-form :model="dialogQuestion">

        <el-row>
          <el-col :span="12">
            <el-form-item label="题目类型" :label-width="100">
              <el-select v-model="dialogQuestion.type" placeholder="选择题目类型">
                <el-option label="单选题" :value="1"/>
                <el-option label="多选题" :value="2"/>
                <el-option label="判断题" :value="3"/>
                <el-option label="填空题" :value="4"/>
                <el-option label="问答题" :value="5"/>
              </el-select>
            </el-form-item>
          </el-col>
          <el-col :span="12">
            <el-form-item label="分数" :label-width="100">
              <el-input-number v-model="dialogQuestion.score"/>
            </el-form-item>
          </el-col>
        </el-row>


        <el-form-item label="题目内容" :label-width="100">
          <el-input v-model="dialogQuestion.title" autocomplete="off" type="textarea"/>
        </el-form-item>

        <div style="margin-left: 20px" v-if="dialogQuestion.type === 1">
          <el-row>
            <el-col :span="4">
              设置答案
            </el-col>
            <el-col :span="20">
              选项内容
            </el-col>
          </el-row>

          <el-radio-group v-model="dialogQuestion.answer" class="radios">

            <el-radio class="radio" v-for="option in dialogQuestion.options" :label="option.code" size="large">
              <el-row style="align-items: center">
                <el-col :span="2">
                  {{ option.code }}
                </el-col>
                <el-col :span="22">
                  <el-input v-model="option.content" placeholder="输入选项内容"/>
                </el-col>
              </el-row>
            </el-radio>

          </el-radio-group>

        </div>






      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button>取消</el-button>
          <el-button type="primary" @click="onAddQuestionSave">保存</el-button>
        </span>
      </template>
    </el-dialog>


  </div>


</template>

<style scoped>


.create-exam {
  background-color: #f3f5f7;
  height: 100%;
  padding: 30px;
}

.create-exam :deep(.el-tabs__header) {
  display: none;
}

.exam-designer {
  display: grid;
  grid-template-columns: 4fr 9fr;
  min-height: 400px;
  grid-column-gap: 20px;
  align-content: start;
}

.exam-designer > .left {
}

.exam-designer .navigator {
  background-color: #fff;
  border-radius: 8px;
  margin-bottom: 20px;
  min-height: 200px;
  position: relative;
}


.exam-designer .navigator > .foot {
  position: absolute;
  bottom: 0;
  text-align: center;
  border-top: #e7e7e7 1px solid;
  right: 20px;
  left: 20px;
  padding: 10px;
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

.exam-designer > .right {
  position: relative;
}

.card {
  background-color: #fff;
  border-radius: 5px;
  padding: 20px 20px 10px;
  margin-bottom: 20px;
}

.card > .title {
  margin-bottom: 20px;
}

.card > .title:before {
  content: " ";
  width: 4px;
  height: 18px;
  background-color: #128beb;
  display: inline-block;
  vertical-align: text-top;
  margin-right: 10px;
}

.input-200 :deep(.el-input) {
  width: 200px !important;
}

.exam-designer .card {
  min-height: 200px;
}


.exam-designer .card .empty {
  text-align: center;
  height: 100%;
}

.question-dialog .radios {
  flex-direction: column;
  width: 100%;
}

.question-dialog .radios .radio {
  width: 100%;
}


.question-dialog .radios :deep(.el-radio__label) {
  flex-grow: 1;
}


.question-dialog .radios :deep(.el-radio:last-child) {
  margin-right: 32px;
}

</style>