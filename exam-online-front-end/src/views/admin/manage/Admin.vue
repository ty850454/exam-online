<script setup>
import IconQuestionFill from "@/components/icons/IconQuestionFill.vue";
import {ref} from "vue";
import IconSearch from "@/components/icons/IconSearch.vue";
import {get} from "@/utils/http";

const total = ref(0)
const tableData = ref([])

get('/api/admin').then(response => {
  total.value = response.total
  tableData.value = response.data
})

</script>

<template>
    <div class="admin-container">
      <div><IconQuestionFill/>&nbsp;&nbsp;子管理员（{{total}}）</div>
      <el-divider style="margin-top: 10px" />
      <div class="tools">
        <div>

          <el-form :inline="true">
            <el-form-item>
              <el-input placeholder="输入关键字查询" :prefix-icon="IconSearch" clearable />
            </el-form-item>
            <el-form-item label="创建时间：">
              <el-date-picker
                  clearable
                  type="datetimerange"
                  range-separator="至"
                  start-placeholder="起始日期"
                  end-placeholder="结束日期"
              />
            </el-form-item>

            <el-form-item>
              <el-button type="primary">查询</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div style="margin-left: auto;"><el-button type="primary">新增子管理员</el-button></div>
      </div>
      <div>
        <el-table :data="tableData" border style="width: 100%" header-row-class-name="abc" :header-cell-style="{background:'#F8F8F9'}">
          <el-table-column prop="username" label="账号" width="180" />
          <el-table-column prop="name" label="姓名" width="180" />
          <el-table-column prop="address" label="创建时间" />
          <el-table-column prop="address" label="状态" />
          <el-table-column prop="address" label="操作" />
        </el-table>
      </div>
    </div>
</template>

<style scoped>
.admin-container {
  margin: 20px
}

.admin-container > div {
  margin-bottom: 10px;
}

.tools {
  display: flex;
}

.admin-container :deep(.abc) {
  background-color: red;
}


</style>