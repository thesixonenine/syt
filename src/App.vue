<script lang="ts">
// import HelloWorld from './components/HelloWorld.vue'
import {defineAsyncComponent, defineComponent} from "vue";
export default defineComponent({
  name: "App",
  components: {
    "WishList": defineAsyncComponent(() => import('@/components/WishList.vue')),
  }
})
</script>

<template>
  <!-- <HelloWorld msg="Dashboard" /> -->
  <WishList />
</template>

<style scoped>
</style>
