<script setup lang="ts">
defineOptions({name: 'Hospital'});
import {Response} from "@/api/common_type.ts";
import {useRoute, useRouter} from "vue-router";
import {onMounted} from "vue";
// 引入图标
import {Document, Menu as IconMenu, Search, InfoFilled, HomeFilled, Setting} from "@element-plus/icons-vue";
import {cardDetailReq} from "@/api/hospital.ts";
import {HospitalDetail} from "@/api/hospital_type.ts";
import useDetailStore from "@/store/modules/hospital.ts";

let detailStore = useDetailStore();
let $router = useRouter();
const changeActive = (p: string) => {
  $router.push({path: p});
};
// 获取当前 URL 上的路由信息
let $route = useRoute();
onMounted(() => {
  detailStore.getHospitalDetail($route.query.hoscode as string);
});
</script>

<template>
  <div class="hospital">
    <div class="menu">
      <div class="top">
        <el-icon>
          <HomeFilled/>
        </el-icon>
        <span>/ 医院信息</span>
      </div>
      <el-menu :default-active="$route.path" class="el-menu--vertical-demo">
        <el-menu-item index="/hospital/reg" @click="changeActive('/hospital/reg')">
          <el-icon>
            <icon-menu/>
          </el-icon>
          <span>预约挂号</span>
        </el-menu-item>
        <el-menu-item index="/hospital/detail" @click="changeActive('/hospital/detail')">
          <el-icon>
            <document/>
          </el-icon>
          <span>医院详情</span>
        </el-menu-item>
        <el-menu-item index="/hospital/notify" @click="changeActive('/hospital/notify')">
          <el-icon>
            <setting/>
          </el-icon>
          <span>预约通知</span>
        </el-menu-item>
        <el-menu-item index="/hospital/stop" @click="changeActive('/hospital/stop')">
          <el-icon>
            <info-filled/>
          </el-icon>
          <span>停诊信息</span>
        </el-menu-item>
        <el-menu-item index="/hospital/cancel" @click="changeActive('/hospital/cancel')">
          <el-icon>
            <search/>
          </el-icon>
          <span>查询和取消</span>
        </el-menu-item>
      </el-menu>
    </div>
    <div class="content">
      <!-- 二级路由 -->
      <router-view/>
    </div>
  </div>
</template>

<style scoped lang="scss">
.hospital {
  display: flex;

  .menu {
    flex: 2;
    display: flex;
    flex-direction: column;
    align-items: center;

    .top {
      color: #7f7f7f;
    }
  }

  .content {
    flex: 8;
  }
}
</style>
