<script setup lang="ts">
import {onMounted, ref} from "vue";
import {useRoute} from "vue-router";
import {cardDepartBookingReq} from "@/api/hospital.ts";
import {Response} from "@/api/common_type.ts";
import {Booking} from "@/api/hospital_type.ts";

let $route = useRoute();
let pageNo = ref<number>(1);
let pageSize = ref<number>(6);
let booking = ref<Booking>();
onMounted(() => {
  bookingPage();
});
const bookingPage = async () => {
  let result: Response<Booking> = await cardDepartBookingReq(pageNo.value, pageSize.value,
      $route.query.hoscode as string, $route.query.depcode as string);
  if (result.code == 200) {
    console.log('BookingReq', result.data);
    booking.value = result.data;
  } else {
    console.error('BookingReq', result.message);
  }
};
</script>

<template>
  <div class="warp">
    <div class="top">
      <div class="hosname">{{ booking?.baseMap.hosname }}</div>
      <div class="line"></div>
      <div>{{ booking?.baseMap.bigname }}</div>
      <div class="dot">.</div>
      <div class="depname">{{ booking?.baseMap.depname }}</div>
    </div>
    <div class="center">
      <h1 class="time">{{ booking?.baseMap.workDateString }}</h1>
      <div class="container">
        <div class="item" :class="{active:item.status==-1||item.availableNumber==-1}"
             v-for="item in booking?.bookingScheduleList">
          <div class="up">{{ item.workDate }}-{{ item.dayOfWeek }}</div>
          <div class="down">
            <div v-if="item.status == -1">停止挂号</div>
            <div v-if="item.status == 0">
              {{ item.availableNumber == -1 ? '约满了' : `有号(${item.availableNumber})` }}
            </div>
            <div v-if="item.status==1">即将放号</div>
          </div>
        </div>
      </div>
      <el-pagination layout="prev,pager,next"
                     v-model:current-page="pageNo"
                     :total="booking?.total"
                     @click="bookingPage"/>
    </div>
    <div class="bottom">
      <div class="will">
        <span class="time">xxxxxxxx</span>
        <span class="willText">放号</span>
      </div>
      <div class="doctor">
        <div class="am">
          <div class="tip">
            <span>上午号源</span>
          </div>
          <div class="doc-info">
            <div class="left">
              <div class="info">
                <span>xxxx</span>
                <span>|</span>
                <span>xxxx</span>
              </div>
              <div class="skill">xxxxxxxxxxxx</div>
            </div>
            <div class="right">
              <div class="money">$100</div>
              <el-button type="primary" size="default" @click="">100</el-button>
            </div>
          </div>
        </div>
      </div>
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

      .item {
        flex: 1;
        width: 150px;
        border: 1px solid #7f7f7f;
        margin: 0 5px;
        display: flex;
        flex-direction: column;
        align-items: center;

        &.active {
          border: 1px solid #ccc;
          color: #7f7f7f;

          .up {
            background: #ccc;
          }
        }

        .up {
          font-size: 10px;
          background-color: #e8f2ff;
          height: 30px;
          width: 100%;
          text-align: center;
          padding-top: 10px;
        }

        .down {
          height: 60px;
          width: 100%;
          text-align: center;
          padding-top: 16px;
        }
      }
    }
  }

  .bottom {
    .will {
      text-align: center;
      font-size: 30px;
      font-weight: 900;

      .time {
        color: red;
      }

      .willText {
        color: skyblue;
      }
    }

    .doctor {
      .am {
        .tip {
          color: #7f7f7f;
          font-weight: 900;
        }

        .doc-info {
          display: flex;
          justify-content: space-between;
          margin: 10px 0;
          border-bottom: 1px solid skyblue;
          .left{
            .info{
              color: skyblue;
              margin: 10px 0;
              span{
                margin: 0 5px;
                font-size: 18px;
                font-weight: 900;
              }
            }
            .skill{
              margin: 10px 0;
              color: #7f7f7f;
            }
          }
          .right{
            width: 150px;
            display: flex;
            justify-content: space-between;
            align-items: center;
            .money{
              color: #7f7f7f;
              font-weight: 900;
            }
          }
        }
      }

      .pm {
        color: #7f7f7f;
        font-weight: 900;
      }
    }
  }
}
</style>