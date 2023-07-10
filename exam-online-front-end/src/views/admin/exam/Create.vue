<script setup>

import {ref} from "vue";
import IconRight from "@/components/icons/IconRight.vue";

const form = ref({})
const activeStep = ref(1)

function getTypeTitle(type) {
  switch (type) {
    case 1:
      return '单选题'
    case 2:
      return '多选题'
    case 3:
      return '判断题'
    case 4:
      return '填空题'
    case 5:
      return '问答题'
    default:
      return '未知'
  }
}

const questionSet = ref([
  {
    title: '分组1',
    score: 0,
    questions: [
      {
        type: 1,
        score: 0,
        title: '123',
        answer: 'A,B',
        options: [
          {
            code: 'A',
            content: 'asdasd'
          },
          {
            code: 'B',
            content: 'asdasd'
          },
          {
            code: 'C',
            content: 'asdasd'
          },
          {
            code: 'D',
            content: 'asdas'
          },
        ]
      }
    ]
  }
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

function onInsertCha() {
  dialogQuestion.value.title += ' ___ '
  onType4TitleChange(dialogQuestion.value.title)
}

function onType4TitleChange(text) {
  let sp = text.split('___')
  let l = text.split('___').length - 1
  dialogQuestion.value.options = sp
  dialogQuestion.value.answer.length = l
  console.log(dialogQuestion.value)
}

function onAddQuestionSave() {
  dialogVisible.value = false
}

function onAddQuestion(command) {
  dialogVisible.value = true

  switch (command) {
    case 1:
      dialogQuestion.value = {
        type: command,
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
      break;
    case 2:
      dialogQuestion.value = {
        type: command,
        score: 0,
        title: '',
        answer: [],
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
      break;
    case 3:
      dialogQuestion.value = {
        type: command,
        score: 0,
        title: '',
        answer: '',
      }
      break;
    case 4:
      dialogQuestion.value = {
        type: command,
        score: 0,
        title: '',
        answer: [],
      }
      break;
    case 5:
      dialogQuestion.value = {
        type: command,
        score: 0,
        title: '',
      }
      break;

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

              <div class="navigator-body">
                <div class="info">共 0 题 0 分</div>
                <div style="padding: 20px 20px 20px 28px" v-for="set in questionSet">
                  <div>{{ set.title }}<span>（共{{ set.questions.length }}题{{ set.score }}分）</span></div>
                </div>
              </div>

              <div style="height: 30px;"></div>
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
          <div class="card" v-if="questionSet.length > 0" v-for="(set, si) in questionSet">
            <div>{{ set.title }} <span>共0题，0分</span></div>
            <div class="empty" v-if="set.questions.length === 0">
              请向分组添加题目
            </div>
            <table class="options">
              <tr class="title">
                <td>序号</td>
                <td>题型</td>
                <td>题目</td>
                <td>答案</td>
                <td>分数</td>
                <td>操作</td>
              </tr>
              <tr v-for="(question, i) in set.questions">
                <td>{{ i + 1 }}</td>
                <td>{{ getTypeTitle(question.type) }}</td>
                <td>{{ question.title }}</td>
                <td>{{ question.answer }}</td>
                <td>{{ question.score }}</td>
              </tr>
            </table>
            <div>

              <el-dropdown @command="onAddQuestion">
                <el-button type="primary">
                  新增题目
                </el-button>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item :command="1">单选题</el-dropdown-item>
                    <el-dropdown-item :command="2">多选题</el-dropdown-item>
                    <el-dropdown-item :command="3">判断题</el-dropdown-item>
                    <el-dropdown-item :command="4">填空题</el-dropdown-item>
                    <el-dropdown-item :command="5">问答题</el-dropdown-item>
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
              <el-select v-model="dialogQuestion.type" placeholder="选择题目类型" @change="onAddQuestion">
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
              <el-input-number v-model="dialogQuestion.score" :min="0"/>
            </el-form-item>
          </el-col>
        </el-row>


        <el-form-item label="题目内容" :label-width="100">
          <el-input v-if="dialogQuestion.type !== 4" v-model="dialogQuestion.title" autocomplete="off" type="textarea"/>
          <el-input v-if="dialogQuestion.type === 4" v-model="dialogQuestion.title" autocomplete="off" type="textarea"
                    @input="onType4TitleChange"/>
          <el-button v-if="dialogQuestion.type === 4" type="primary" link @click="onInsertCha">插入填空符</el-button>
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


        <div style="margin-left: 20px" v-if="dialogQuestion.type === 2">
          <el-row>
            <el-col :span="4">
              设置答案
            </el-col>
            <el-col :span="20">
              选项内容
            </el-col>
          </el-row>
          <el-checkbox-group v-model="dialogQuestion.answer" class="check-box-list">
            <el-checkbox class="check-box" v-for="option in dialogQuestion.options" :label="option.code" size="large">
              <el-row style="align-items: center">
                <el-col :span="2">
                  {{ option.code }}
                </el-col>
                <el-col :span="22">
                  <el-input v-model="option.content" placeholder="输入选项内容"/>
                </el-col>
              </el-row>
            </el-checkbox>
          </el-checkbox-group>
        </div>


        <div v-if="dialogQuestion.type === 3">
          <el-form-item label="选项" :label-width="100">
            <el-radio-group v-model="dialogQuestion.answer">
              <el-radio label="T" size="large">正确</el-radio>
              <el-radio label="F" size="large">错误</el-radio>
            </el-radio-group>
          </el-form-item>
        </div>


        <div v-if="dialogQuestion.type === 4">
          <el-form-item label="答案" :label-width="100">
            <div style="width: 100%;">
              <el-row v-for="(a, i) in dialogQuestion.answer" style="margin-bottom: 10px">
                <el-col :span="4">选项{{ i + 1 }}答案：</el-col>
                <el-col :span="20">
                  <el-input v-model="dialogQuestion.answer[i]" autocomplete="off" type="textarea"/>
                </el-col>
              </el-row>
            </div>
          </el-form-item>
        </div>


      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
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
  min-height: 200px;
  position: relative;
}


.exam-designer .navigator .navigator-body {
  margin-bottom: 40px;
}

.exam-designer .navigator .navigator-body > * {
  padding: 10px;
}

.exam-designer .navigator .navigator-body .info {
  position: relative;
  border-bottom: #e7e7e7 1px solid;
  right: 20px;
  left: 20px;
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
  height: 100px;
  line-height: 100px;
}

.question-dialog .check-box-list {
  display: flex;
  flex-direction: column;
  width: 100%;
}

.question-dialog .check-box-list {
  display: flex;
  flex-direction: column;
  width: 100%;
}


.question-dialog .check-box-list .check-box {
  width: 100%;
}

.question-dialog .check-box-list :deep(.el-checkbox__label) {
  flex-grow: 1;
}

.question-dialog .check-box-list :deep(.el-checkbox__label) {
  flex-grow: 1;
}

.question-dialog .check-box-list :deep(.el-checkbox:last-child) {
  margin-right: 32px;
}

.question-dialog .radios .radio {
  width: 100%;
}


.question-dialog .radios :deep(.el-radio) {
  margin-right: 0;
}


.question-dialog .radios :deep(.el-radio__label) {
  flex-grow: 1;
}


.card :deep(.el-button:focus) {
//background-color: #20287e;
}

.card .options {
  width: 100%;
  border-collapse: collapse;
}


.card .options td {
  padding: 10px;
}

.card .options > .title {
  background-color: #e5e5e5;
}

.card .options > .op {
  margin-right: 10px;
}

</style>