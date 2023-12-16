<script setup lang="ts">
defineOptions({name: 'Top'})

import userStore from "@/store/modules/user.ts";
import {useRouter} from "vue-router";
import { ArrowDown } from '@element-plus/icons-vue'
let $router = useRouter();
const home = () => {
  $router.push({path: '/home'});
};
let user = userStore();
const login = () => {
  user.visible = true;
};
const logout = ()=>{
  user.logout();
  home();
};
</script>

<template>
  <div class="top">
    <div class="container">
      <div class="left" @click="home">
        <img class="logo" src="@/assets/img/ts.jpg" alt="">
        <p>Dashboard</p>
      </div>
      <div class="right">
        <p class="github">GitHub</p>
        <p class="login" v-if="!user.info.name" @click="login">登录/注册</p>
        <el-dropdown v-else>
          <span class="el-dropdown-link">
            {{user.info.name}}
            <el-icon class="el-icon--right">
              <arrow-down/>
            </el-icon>
          </span>
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item>实名认证</el-dropdown-item>
              <el-dropdown-item>挂号订单</el-dropdown-item>
              <el-dropdown-item>就诊人管理</el-dropdown-item>
              <el-dropdown-item @click="logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
    </div>

  </div>
</template>

<style scoped lang="scss">
.top {
  position: fixed;
  z-index: 999;
  width: 100%;
  height: 70px;
  background-color: white;
  display: flex;
  justify-content: center;

  .container {
    width: 1200px;
    height: 70px;
    background-color: white;
    display: flex;
    justify-content: space-between;

    .left {
      display: flex;
      justify-content: center;
      align-items: center;
      cursor: pointer;

      .logo {
        width: 50px;
        height: 50px;
        margin-right: 10px;
      }

      p {
        font-size: 20px;
        color: #55a6fe;
      }
    }

    .right {
      display: flex;
      justify-content: center;
      align-items: center;
      font-size: 14px;
      color: #bbb;

      .github {
        margin-right: 10px;
      }
    }
  }
}
</style>