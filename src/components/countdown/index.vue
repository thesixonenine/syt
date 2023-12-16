<script setup lang="ts">
import {ref, watch} from "vue";

let props = defineProps(['sendToSon']);
let $emit = defineEmits(['getFromSonEmit']);
let timeCnt = ref<number>(5);
// 监听父组件传递的 countDownFlag 变量
watch(() => props.sendToSon, () => {
  let timer = setInterval(() => {
    timeCnt.value--;
    if (timeCnt.value == 0) {
      $emit('getFromSonEmit', false);
      clearInterval(timer);
    }
  }, 1000)
}, {
  immediate: true
});
</script>

<template>
  <div>
    <span>获取验证码({{ timeCnt }}s)</span>
  </div>
</template>

<style scoped lang="scss">

</style>