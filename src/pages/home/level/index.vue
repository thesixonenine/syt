<script setup lang="ts">

</script>

<template>
  <div class="level">
    <h1>角色筛选</h1>
    <div class="content">
      <div class="left">星级:</div>
      <ul class="star">
        <li class="active">全部</li>
        <li>四星</li>
        <li>五星</li>
      </ul>
    </div>
  </div>
</template>

<style scoped lang="scss">
.level {
  color: #7f7f7f;

  h1 {
    font-weight: 900;
    margin: 10px 0;
  }

  .content {
    display: flex;

    .left {
      margin-right: 10px;
    }

    .star {
      display: flex;

      li {
        margin-right: 10px;

        .active {
          color: #55a6fe;
        }
      }
      li:hover{
        color: #55a6fe;
        cursor: pointer;
      }
    }
  }
}
</style>