<script setup lang="ts">
// 接受传入的props itemInfo
defineProps(['itemInfo']);
</script>

<template>
  <el-card class="box-card" shadow="hover">
    <div class="content">
      <div class="left">
        <div class="name">{{ itemInfo.hosname }}</div>
        <div class="info">
          <div class="star">❤❤❤❤❤</div>
          <div class="type">{{ itemInfo.param?.hostypeString }}</div>
        </div>
      </div>
      <div class="right">
        <!-- <img src="@/assets/img/home/card/kafka.png" alt=""/> -->
        <img :src="`data:image/jpeg;base64,${itemInfo.logoData}`" alt=""/>
      </div>
    </div>
  </el-card>
</template>

<style scoped lang="scss">
.content {
  display: flex;
  justify-content: space-between;

  .left {
    width: 60%;

    .info {
      color: #7f7f7f;
      margin-top: 20px;
      display: flex;
      justify-content: space-between;
    }
  }

  .right {
    img {
      width: 50px;
      height: 50px;
    }
  }
}
</style>
