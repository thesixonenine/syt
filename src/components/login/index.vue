<script setup lang="ts">
defineOptions({name:"Login"});
import userStore from "@/store/modules/user.ts";

let user = userStore();
</script>

<template>
<div class="login">
  <el-dialog v-model="user.visible" title="用户登录">

  </el-dialog>
</div>
</template>

<style scoped lang="scss">
.login{
  color: red;
}
</style>