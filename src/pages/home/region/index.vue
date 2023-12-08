<script setup lang="ts">

</script>

<template>
  <div class="region">
    <div class="content">
      <div class="left">地区:</div>
      <ul class="star">
        <li class="active">全部</li>
        <li>银河</li>
        <li>仙舟「罗浮」</li>
        <li>贝洛伯格</li>
        <li>空间站「黑塔」</li>
        <li>星穹列车</li>
        <li>星核猎手</li>
        <li v-for="item in 10" :key="item">测试</li>
      </ul>
    </div>
  </div>
</template>

<style scoped lang="scss">
.region {
  color: #7f7f7f;
  margin-top: 10px;

  h1 {
    font-weight: 900;
    margin: 10px 0;
  }

  .content {
    display: flex;

    .left {
      margin-right: 10px;
      width: 36px;
    }

    .star {
      display: flex;
      flex-wrap: wrap;

      li {
        margin-right: 10px;
        margin-bottom: 10px;

        .active {
          color: #55a6fe;
        }
      }

      li:hover {
        color: #55a6fe;
        cursor: pointer;
      }
    }
  }
}
</style>