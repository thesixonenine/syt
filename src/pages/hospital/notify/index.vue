<script setup lang="ts">
defineOptions({name: "Notify"})
import useDetailStore from "@/store/modules/hospital.ts";

let detailStore = useDetailStore();
</script>

<template>
  <div class="notify">
    <h1 class="hosname">{{ detailStore.hospitalDetail.hospital?.hosname }}预约挂号通知</h1>
    <div class="top">须知:</div>
    <div class="info">
      <div class="tip">xxx</div>
      <div class="text">xxxxxxxxxxxx</div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.notify {
  .hosname {
    text-align: center;
    font-size: 30px;
  }
  .top {
    color: #7f7f7f;
    margin: 10px 0;
  }
  .info {
    .tip {
      color: red;
      font-weight: 900;
      .text {
        color: #7f7f7f;
        margin: 10px 0;
        line-height: 30px;
      }
    }
  }
}
</style>
