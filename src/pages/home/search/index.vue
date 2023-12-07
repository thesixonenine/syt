<script setup lang="ts">
// 引入图标
import {Search} from "@element-plus/icons-vue";

</script>

<template>
  <div class="search">
    <el-autocomplete
        v-model="state1"
        :fetch-suggestions="querySearch"
        clearable
        class="inline-input w-50"
        placeholder="请输入角色名称"
        @select="handleSelect"
    />
    <el-button type="primary" size="default" @click="" :icon="Search">搜索</el-button>
  </div>
</template>

<style scoped lang="scss">
.search {
  width: 100%;
  height: 50px;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 10px 0;
  // 深度选择器(父组件改变子组件的样式)
  :deep(.el-input__wrapper) {
    width: 300px;
    margin-right: 10px;
  }
}
</style>
