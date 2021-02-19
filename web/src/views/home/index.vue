<template>
  <div class="home">
    <!-- 数据概况 start -->
    <el-row :gutter="40">
      <el-col :lg="6" :sm="12">
        <div class="grid-content bg-white">
          <el-row>
            <el-col :span="12">
              <div class="grid-content">
                <div class="icons">
                  <i class="iconfont icon-chenggong"></i>
                </div>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="grid-content">
                <ul class="icons-right">
                  <li class="chain">坚持天数</li>
                  <li>
                    <countTo :startVal="0" :endVal="userinfo.persist_day" :duration="1000"></countTo>
                  </li>
                </ul>
              </div>
            </el-col>
          </el-row>
        </div>
      </el-col>
      <el-col :lg="6" :sm="12">
        <div class="grid-content bg-white">
          <el-row>
            <el-col :span="12">
              <div class="grid-content">
                <div class="icons icons1">
                  <i class="iconfont icon-zanting"></i>
                </div>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="grid-content">
                <ul class="icons-right">
                  <li class="chain">未达标天数</li>
                  <li>
                    <countTo :startVal="0" :endVal="userinfo.interrupt_day" :duration="1000"></countTo>
                  </li>
                </ul>
              </div>
            </el-col>
          </el-row>
        </div>
      </el-col>
      <el-col :lg="6" :sm="12">
        <div class="grid-content bg-white">
          <el-row>
            <el-col :span="12">
              <div class="grid-content">
                <div class="icons icons2">
                  <i class="iconfont icon-baocun"></i>
                </div>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="grid-content">
                <ul class="icons-right">
                  <li class="chain">已做题数</li>
                  <li>
                    <countTo :startVal="0" :endVal="userinfo.persist_num" :duration="1000"></countTo>
                  </li>
                </ul>
              </div>
            </el-col>
          </el-row>
        </div>
      </el-col>
      <el-col :lg="6" :sm="12">
        <div class="grid-content bg-white">
          <el-row>
            <el-col :span="12">
              <div class="grid-content">
                <div class="icons icons3">
                  <i class="iconfont icon-bianji"></i>
                </div>
              </div>
            </el-col>
            <el-col :span="12">
              <div class="grid-content">
                <ul class="icons-right">
                  <li class="chain">做题次数</li>
                  <li>
                    <countTo :startVal="0" :endVal="userinfo.persist_times" :duration="1000"></countTo>
                  </li>
                </ul>
              </div>
            </el-col>
          </el-row>
        </div>
      </el-col>
    </el-row>
    <!-- 数据概况 end -->
    <el-row :gutter="40">
      <!-- 我的账户 start -->
      <el-col :lg="6" :sm="8" :xs="24">
        <div class="main-center clearfix">
          <div class="pull-left center-left">
            <ul>
              <li class="accout">我的账户</li>
              <li class="tou">
                <img src="../../assets/img/tou.jpg">
                <br><br>
                <span>{{ $store.state.user.name }}</span>
              </li>
              <li class="time">当前时间：{{ date }}</li>
            </ul>
          </div>
        </div>
      </el-col>
      <!-- 我的账户 end -->
      <el-col :lg="18" :sm="16" :xs="24">
        <div class="todulist">
          <div class="item"><span>今日待办</span></div>
          <div class="item" v-for="(item,i) in userinfo.todulist" :key="i">
            <el-checkbox v-model="item.done"></el-checkbox>
            <span :class="item.done?'done':''">{{ item.content }}</span>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script>
// 数字滚动插件
import countTo from "vue-count-to"
import {getTodayOverview} from "@api"

export default {
  name: "home",
  components: {countTo},
  data() {
    return {
      date: '',
      userinfo: {
        persist_day: 0,
        interrupt_day: 0,
        persist_num: 0,
        persist_times: 0,
        todulist: []
      },
    };
  },
  mounted() {
    this.date = new Date().toLocaleDateString()
    getTodayOverview().then(data => {
      this.userinfo = data
    })
  }
};
</script>

<style lang="scss" scoped>
.home {
  padding: 40px;
  background: $base-gray1;

  .bg-white {
    background: $base-white;
    -webkit-box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    border-color: rgba(0, 0, 0, 0.05);
    margin-bottom: 20px;
    cursor: pointer;

    .icons {
      text-align: left;
      width: 86px;
      height: 86px;
      margin: 10px;
      padding: 13px;
      transition: 0.3s ease-in-out;
      border-radius: 3px;

      &:hover {
        background: $base-green;

        .icon-chenggong {
          color: $base-white;
        }
      }

      .iconfont {
        font-size: 60px;
      }

      .icon-chenggong {
        font-size: 60px;
        color: $base-green;
      }
    }

    .icons1 {
      &:hover {
        background: $base-pink;

        .icon-zanting {
          color: $base-white;
        }
      }

      .icon-zanting {
        color: $base-pink;
      }
    }

    .icons2 {
      &:hover {
        background: $base-bule1;

        .icon-baocun {
          color: $base-white;
        }
      }

      .icon-baocun {
        color: $base-bule1;
      }
    }

    .icons3 {
      &:hover {
        background: $base-ye;

        .icon-bianji {
          color: $base-white;
        }
      }

      .icon-bianji {
        color: $base-ye;
      }
    }

    .icons-right {
      font-size: 24px;
      margin-top: 16px;
      margin-right: 16px;

      li {
        margin: 10px 0;
      }
    }

    .chain {
      color: rgba(0, 0, 0, 0.45);
      font-size: 18px;
    }
  }

  .main-center {
    width: 100%;
    margin-top: 20px;
  }

  .center-left {
    width: 100%;
    text-align: center;
    background: $base-white;
    font-size: 16px;
    color: $base-666;
    padding-bottom: 30px;
    -webkit-box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    border-color: rgba(0, 0, 0, 0.05);

    .accout {
      text-align: center;
      margin: 20px;
    }

    .tou {
      margin-bottom: 20px;

      img {
        width: 85px;
        height: 85px;
        border-radius: 50%;
      }

      span {
        line-height: 15px;
      }
    }

    .time {
      line-height: 15px;
    }
  }

  .todulist {
    background: $base-white;
    margin-top: 20px;
    min-height: 255px;

    .item {
      line-height: 46px;
      border-bottom: 1px solid #ededed;
      text-align: left;
      padding-left: 40px;
      cursor: pointer;
      font-size: 16px;
    }

    .done {
      text-decoration: line-through;
      color: gray;
    }
  }
}
</style>

