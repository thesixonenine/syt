<script setup lang="ts">
import {onMounted} from "vue";
import wishJson from '@/assets/data/star-rail-wish.json'

class WishItem {
  count: string
  gacha_id: string
  gacha_type: string
  id: string
  item_id: string
  item_type: string
  lang: string
  name: string
  rank_type: string
  time: string
  uid: string

  constructor(count: string, gacha_id: string, gacha_type: string, id: string, item_id: string, item_type: string, lang: string, name: string, rank_type: string, time: string, uid: string) {
    this.count = count;
    this.gacha_id = gacha_id;
    this.gacha_type = gacha_type;
    this.id = id;
    this.item_id = item_id;
    this.item_type = item_type;
    this.lang = lang;
    this.name = name;
    this.rank_type = rank_type;
    this.time = time;
    this.uid = uid;
  }
}

let wishList = new Array<WishItem>();
for (const idx in wishJson[1]) {
  const dx = <WishItem>(wishJson[1][idx]);
  wishList.push(dx);
}
for (const idx in wishJson[2]) {
  const dx = <WishItem>(wishJson[2][idx]);
  wishList.push(dx);
}
for (const idx in wishJson[11]) {
  const dx = <WishItem>(wishJson[11][idx]);
  wishList.push(dx);
}
for (const idx in wishJson[12]) {
  const dx = <WishItem>(wishJson[12][idx]);
  wishList.push(dx);
}
console.log("length = " + wishList.length);
onMounted(() => {
  console.log("WishList.vue is now mounted.")
})
</script>

<template>
  <div class="wishes">
    <div class="wishTable" v-for="wishItem in wishList">
      <div id="wishTime">{{ wishItem.time }}</div>
      <div id="wishName">{{ wishItem.name }}</div>
      <div id="wishStar">{{ wishItem.item_type }}</div>
      <div id="wishType">{{ wishItem.gacha_type }}</div>
      <div id="wishRoll">{{ wishItem.rank_type }}</div>
      <div id="uid">{{ wishItem.uid }}</div>
    </div>
  </div>
</template>

<style scoped>
.wishes {
  display: flex;
  flex-direction: column;
}
.wishTable {
  height: 100%;
  display: flex;
  flex-wrap: nowrap;
  justify-content: space-around;
  align-items: center;
}
</style>