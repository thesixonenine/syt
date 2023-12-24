<script setup lang="ts">
import {computed, onMounted, ref} from "vue";
import {useRoute} from "vue-router";
import {cardDepartBookingReq, doctorDetailReq} from "@/api/hospital.ts";
import {Response} from "@/api/common_type.ts";
import {Booking, BookingSchedule, Doctor} from "@/api/hospital_type.ts";

let $route = useRoute();
const hoscode = $route.query.hoscode as string;
const depcode = $route.query.depcode as string;
let pageNo = ref<number>(1);
let pageSize = ref<number>(6);
let booking = ref<Booking>();
onMounted(() => {
  bookingPage();
});
let workTime = ref<BookingSchedule>();
const bookingPage = async () => {
  let result: Response<Booking> = await cardDepartBookingReq(pageNo.value, pageSize.value, hoscode, depcode);
  if (result.code == 200) {
    console.log('BookingReq', result.data);
    booking.value = result.data;
    workTime.value = booking.value?.bookingScheduleList[0];
    await doctorWorkTime();
  } else {
    console.error('BookingReq', result.message);
  }
};
let doctorList = ref<Doctor[]>([]);
const doctorWorkTime = async () => {
  let result = await doctorDetailReq(hoscode, depcode, workTime.value?.workDate as string);
  if (result.code == 200) {
    console.log('doctorList', result.data);
    doctorList.value = result.data;
  }
};
// 切换时间
const changeTime = (item: BookingSchedule) => {
  workTime.value = item;
  doctorWorkTime();
};
// 计算
let amDoctorList = computed(() => {
  return doctorList.value.filter((doctor: Doctor) => {
    return doctor.workTime == 0;
  })
});
let pmDoctorList = computed(() => {
  return doctorList.value.filter((doctor: Doctor) => {
    return doctor.workTime == 1;
  })
});
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
        <div class="item" :class="{active:item.status==-1||item.availableNumber==-1,cur:item.workDate==workTime?.workDate}"
             v-for="item in booking?.bookingScheduleList" @click="changeTime(item)">
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
      <div class="will" v-if="workTime?.status == 1">
        <span class="time">xxxxxxxx</span>
        <span class="willText">放号</span>
      </div>
      <div class="doctor" v-else>
        <div class="am"  v-if="amDoctorList.length > 0">
          <div class="tip">
            <span>上午号源</span>
          </div>
          <div class="doc-info" v-for="doctor in amDoctorList" :key="doctor.id">
            <div class="left">
              <div class="info">
                <span>{{doctor.title}}</span>
                <span>|</span>
                <span>{{doctor.docname}}</span>
              </div>
              <div class="skill">{{doctor.skill}}</div>
            </div>
            <div class="right">
              <div class="money">${{doctor.amount}}</div>
              <el-button type="primary" size="default" @click="">{{ doctor.availableNumber }}</el-button>
            </div>
          </div>
        </div>
        <br/>
        <div class="am" v-if="pmDoctorList.length > 0">
          <div class="tip">
            <span>下午号源</span>
          </div>
          <div class="doc-info" v-for="doctor in pmDoctorList" :key="doctor.id">
            <div class="left">
              <div class="info">
                <span>{{doctor.title}}</span>
                <span>|</span>
                <span>{{doctor.docname}}</span>
              </div>
              <div class="skill">{{doctor.skill}}</div>
            </div>
            <div class="right">
              <div class="money">${{doctor.amount}}</div>
              <el-button type="primary" size="default" @click="">{{ doctor.availableNumber }}</el-button>
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
        transition: all .5s;

        &.active {
          border: 1px solid #ccc;
          color: #7f7f7f;

          .up {
            background: #ccc;
          }
        }
        &.cur {
          transform: scale(1.1);
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

          .left {
            .info {
              color: skyblue;
              margin: 10px 0;

              span {
                margin: 0 5px;
                font-size: 18px;
                font-weight: 900;
              }
            }

            .skill {
              margin: 10px 0;
              color: #7f7f7f;
            }
          }

          .right {
            width: 150px;
            display: flex;
            justify-content: space-between;
            align-items: center;

            .money {
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