<script setup lang="ts">
// 引入轮播图
import Carousel from './carousel/index.vue'
// 引入搜索
import Search from './search/index.vue'
// 引入等级
import Level from './level/index.vue'
// 引入地区
import Region from './region/index.vue'
// 引入角色卡片
import Card from './card/index.vue'

import {ref} from 'vue';

let pageNo = ref<number>(1)
let pageSize = ref<number>(10)

</script>

<template>
  <div>
    <!-- 首页轮播图 -->
    <Carousel/>
    <!-- 首页搜索 -->
    <Search/>
    <!-- 内容展示 -->
    <el-row :gutter="20">
      <el-col :span="20">
        <!-- 星级 -->
        <Level/>
        <!-- 地区 -->
        <Region/>
        <!-- 角色 -->
        <div class="cardList">
          <Card class="card" v-for="item in 10" :key="item"/>
        </div>
        <!-- 分页器 -->
        <el-pagination
            v-model:current-page="pageNo"
            v-model:page-size="pageSize"
            :page-sizes="[10, 20, 30, 40]"
            :small="true"
            :background="true"
            layout="prev, pager, next, jumper, ->, total, sizes"
            :total="100"
        />
      </el-col>
      <el-col :span="4">123</el-col>
    </el-row>
  </div>
</template>

<style scoped lang="scss">
.cardList {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;

  .card {
    width: 48%;
    margin: 10px 0;
  }
}
</style>