<script setup lang="ts">

</script>

<template>
  <div class="warp">
    <div class="top">
      <div class="hosname">xxx</div>
      <div class="line"></div>
      <div>xxx</div>
      <div class="dot">.</div>
      <div class="depname">xxx</div>
    </div>
    <div class="center">
      <h1 class="time">xxxx-xx-xx</h1>
      <div class="container">
        <div class="item" v-for="item in 6" :key="item">
          <div class="up">123</div>
          <div class="down">456</div>
        </div>
      </div>
      <el-pagination></el-pagination>
    </div>
  </div>
</template>

<style scoped lang="scss">
.warp {
  .top {
    display: flex;
    color: #7f7f7f;

    .line {
      width: 5px;
      height: 20px;
      margin: 0 5px;
      background-color: red;
    }

    .dot {
      margin: 0 5px;
      color: red;
    }
  }

  .center {
    margin: 20px 0;
    display: flex;
    flex-direction: column;
    align-items: center;

    .time {
      font-weight: 900;
    }

    .container {
      display: flex;
      width: 100%;
      margin: 30px 0;
      .item{
        flex: 1;
        width: 100%;
        border: 1px solid #7f7f7f;
        margin: 0 5px;
        display: flex;
        flex-direction: column;
        align-items: center;
        .up{
          background-color: #e8f2ff;
          height: 30px;
          width: 100%;
          text-align: center;
          padding-top: 6px;
        }
        .down{
          height: 60px;
          width: 100%;
          text-align: center;
          padding-top: 16px;
        }
      }
    }
  }
}
</style>