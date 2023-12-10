<script setup lang="ts">
import useDetailStore from "@/store/modules/hospital.ts";

let detailStore = useDetailStore();

</script>

<template>
  <div class="reg">
    <div class="top">
      <div class="title">{{ detailStore.hospitalDetail.hospital?.hosname }}</div>
      <div class="level">
        <span>{{ detailStore.hospitalDetail.hospital?.param.hostypeString }}</span>
      </div>
    </div>
    <div class="content">
      <div class="left">
        <img :src="`data:jpeg;base64,`+detailStore.hospitalDetail.hospital?.logoData" alt="">
      </div>
      <div class="right">
        <div>挂号规则</div>
        <div class="time">
          周期: 放号时间:{{
            detailStore.hospitalDetail.bookingRule?.releaseTime
          }},停挂时间:{{ detailStore.hospitalDetail.bookingRule?.stopTime }}
        </div>
        <div class="addr">地址:{{ detailStore.hospitalDetail.hospital?.param.fullAddress }}</div>
        <div class="route">路线:{{ detailStore.hospitalDetail.hospital?.route }}</div>
        <div class="release">退号:{{ detailStore.hospitalDetail.bookingRule?.quitTime }}</div>
        <div class="rule">挂号规则</div>
        <div class="ruleInfo" v-for="(item,index) in detailStore.hospitalDetail.bookingRule?.rule" :key="index">
          {{ item }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.reg {
  .top {
    display: flex;

    .title {
      font-size: 30px;
      font-weight: 900;
    }

    .level {
      color: #7f7f7f;
      margin-left: 10px;
      height: 40px;
      text-align: center;
      line-height: 40px;

      span {
        margin-left: 5px;
      }
    }
  }

  .content {
    display: flex;
    margin-top: 20px;
    margin-left: 20px;

    .left {
      width: 80px;

      img {
        width: 80px;
        height: 80px;
        border-radius: 50%;
      }
    }

    .right {
      flex: 1;
      font-size: 14px;
      margin-left: 20px;

      .time, .addr, .route, .release, .ruleInfo {
        margin-top: 10px;
        color: #7f7f7f;
      }

      .rule {
        margin: 10px 0;
      }
    }
  }
}
</style>
