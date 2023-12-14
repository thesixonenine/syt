<script setup lang="ts">
defineOptions({name: "Login"});

import {ref} from "vue";
import userStore from "@/store/modules/user.ts";
import {User, Lock} from "@element-plus/icons-vue";

let user = userStore();
let visibleNum = ref<number>(0);
const changeVisible = () => {
  visibleNum.value = 1;
};
</script>

<template>
  <div class="login">
    <el-dialog v-model="user.visible" title="用户登录">
      <el-row>
        <el-col :span="12">
          <div class="loginForm">
            <div v-show="visibleNum==0">
              <el-form>
                <el-form-item>
                  <el-input placeholder="手机号码" :prefix-icon="User"/>
                </el-form-item>
                <el-form-item>
                  <el-input placeholder="验证码" :prefix-icon="Lock"/>
                </el-form-item>
                <el-form-item>
                  <el-button>获取验证码</el-button>
                </el-form-item>
              </el-form>
              <el-button style="width: 100%;" type="primary" size="default">用户登录</el-button>
              <div class="bottom" @click="changeVisible">
                <p>微信扫码登录</p>
                <svg style="width: 32px;" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024"
                     data-v-ea893728="">
                  <path fill="currentColor"
                        d="m174.72 855.68 135.296-45.12 23.68 11.84C388.096 849.536 448.576 864 512 864c211.84 0 384-166.784 384-352S723.84 160 512 160 128 326.784 128 512c0 69.12 24.96 139.264 70.848 199.232l22.08 28.8-46.272 115.584zm-45.248 82.56A32 32 0 0 1 89.6 896l58.368-145.92C94.72 680.32 64 596.864 64 512 64 299.904 256 96 512 96s448 203.904 448 416-192 416-448 416a461.056 461.056 0 0 1-206.912-48.384l-175.616 58.56z"></path>
                  <path fill="currentColor"
                        d="M512 563.2a51.2 51.2 0 1 1 0-102.4 51.2 51.2 0 0 1 0 102.4m192 0a51.2 51.2 0 1 1 0-102.4 51.2 51.2 0 0 1 0 102.4m-384 0a51.2 51.2 0 1 1 0-102.4 51.2 51.2 0 0 1 0 102.4"></path>
                </svg>
              </div>
            </div>
            <div class="qrcode" v-show="visibleNum==1">
            </div>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="content">
            <div class="top">
              <div class="item">
                <img src="@/assets/img/login/qrcode.jpg" alt="">
                <svg style="width: 16px;" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024" data-v-ea893728="">
                  <path fill="currentColor"
                        d="m174.72 855.68 135.296-45.12 23.68 11.84C388.096 849.536 448.576 864 512 864c211.84 0 384-166.784 384-352S723.84 160 512 160 128 326.784 128 512c0 69.12 24.96 139.264 70.848 199.232l22.08 28.8-46.272 115.584zm-45.248 82.56A32 32 0 0 1 89.6 896l58.368-145.92C94.72 680.32 64 596.864 64 512 64 299.904 256 96 512 96s448 203.904 448 416-192 416-448 416a461.056 461.056 0 0 1-206.912-48.384l-175.616 58.56z"></path>
                  <path fill="currentColor"
                        d="M512 563.2a51.2 51.2 0 1 1 0-102.4 51.2 51.2 0 0 1 0 102.4m192 0a51.2 51.2 0 1 1 0-102.4 51.2 51.2 0 0 1 0 102.4m-384 0a51.2 51.2 0 1 1 0-102.4 51.2 51.2 0 0 1 0 102.4"></path>
                </svg>
                <p>微信扫一扫关注</p>
                <p>快速预约挂号</p>
              </div>
              <div class="item">
                <img src="@/assets/img/login/qrcode.jpg" alt="">
                <svg style="width: 16px;" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024" data-v-ea893728="">
                  <path fill="currentColor"
                        d="M224 768v96.064a64 64 0 0 0 64 64h448a64 64 0 0 0 64-64V768zm0-64h576V160a64 64 0 0 0-64-64H288a64 64 0 0 0-64 64zm32 288a96 96 0 0 1-96-96V128a96 96 0 0 1 96-96h512a96 96 0 0 1 96 96v768a96 96 0 0 1-96 96zm304-144a48 48 0 1 1-96 0 48 48 0 0 1 96 0"></path>
                </svg>
                <p>扫一扫下载</p>
                <p>预约挂号APP</p>
              </div>
            </div>
            <p class="tip">官方指定平台</p>
            <p class="tip">快速挂号,安全放心</p>
          </div>
        </el-col>
      </el-row>
      <template #footer>
        <el-button type="primary" size="default">关闭窗口</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<style scoped lang="scss">
.login {
  // 深度选择器(父组件改变子组件的样式)
  :deep(.el-dialog__body) {
    border-top: 1px solid #ccc;
    border-bottom: 1px solid #ccc;
  }

  .loginForm {
    padding: 20px;
    border: 1px solid #ccc;

    .bottom {
      display: flex;
      flex-direction: column;
      align-items: center;

      p {
        margin: 10px 0;
        cursor: pointer;
      }
      svg{
        cursor: pointer;
      }
    }
  }

  .content {
    .tip {
      text-align: center;
      margin: 20px 0;
      font-size: 20px;
      font-weight: 900;
    }

    .top {
      display: flex;
      justify-content: space-around;

      .item {
        display: flex;
        flex-direction: column;
        align-items: center;

        p {
          margin: 5px 0;
        }

        img {
          width: 130px;
        }
      }
    }
  }
}
</style>