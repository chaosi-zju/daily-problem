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
    <!-- 我的账户&今日待办 start -->
    <el-row :gutter="40">
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
      <el-col :lg="18" :sm="16" :xs="24">
        <div class="todulist">
          <div class="item"><span>今日待办</span></div>
          <div class="item" v-for="(item,i) in userinfo.todulist" :key="i">
            <el-checkbox v-model="item.done" disabled="true"></el-checkbox>
            <span :class="item.done?'done':''">{{ item.content }}</span>
          </div>
        </div>
      </el-col>
    </el-row>
    <!-- 我的账户&今日待办 end -->
    <!-- 我的近30日完成量 start-->
    <el-row :gutter="40">
      <el-col>
        <div id="charts" ref="charts"></div>
      </el-col>
    </el-row>
    <!-- 我的近30日完成量 end-->
    <!-- 各用户完成量 start -->
    <el-row :gutter="40">
      <div class="finish_info">
        <div class="title">近期各用户完成量</div>
        <el-col>
          <el-table :data="finishInfoPageList">
            <el-table-column prop="user" label="用户" width="250"></el-table-column>
            <el-table-column prop="date" label="日期" width="250"></el-table-column>
            <el-table-column prop="amount" label="完成量"></el-table-column>
          </el-table>
          <div class="block">
            <el-pagination style="padding-top: 14px"
                           @size-change="handleSizeChange"
                           @current-change="handleCurrentChange"
                           :current-page="finishInfoCurPage"
                           :page-sizes="[5, 10, 20, 50]"
                           :page-size="finishInfoPageSize"
                           layout="total, sizes, prev, pager, next, jumper"
                           :total="finishInfoRawList.length"
            ></el-pagination>
          </div>
        </el-col>
      </div>
    </el-row>
    <!-- 各用户完成量 end -->
  </div>
</template>

<script>
// 数字滚动插件
import countTo from "vue-count-to"
import echarts from "echarts"
import {getTodayOverview, getFinishInfo} from "@api"

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
        todulist: [],
      },
      finishInfoCurPage: 1,
      finishInfoPageSize: 10,
      finishInfoRawList: [],
      finishInfoPageList: [],
      finishInfoChartData: {}
    };
  },
  mounted() {
    this.date = new Date().toLocaleDateString()
    getTodayOverview().then(data => {
      this.userinfo = data
    })
    getFinishInfo().then(data => {
      this.finishInfoRawList = data.list_info ? data.list_info : []
      this.currentChangePage(this.finishInfoRawList, 1)
      if(data.chart_info){
        this.finishInfoChartData.xData = data.chart_info.xData
        this.finishInfoChartData.yData = new Map(Object.entries(data.chart_info.yData))
      }
      this.drawChart();
      this.init();
    })
  },
  destroyed(){
    window.onresize=null
  },
  methods: {
    init() {
      //图表自适应
      window.onresize = () => {
        if (this.$refs.charts) {
          echarts.init(this.$refs.charts).resize();
        }
      };
    },
    handleSizeChange: function (pageSize) {
      this.finishInfoPageSize = pageSize;
      this.handleCurrentChange(this.finishInfoCurPage);
    },
    handleCurrentChange: function (currentPage) {
      this.finishInfoCurPage = currentPage;
      this.currentChangePage(this.finishInfoRawList, currentPage);
    },
    currentChangePage(list, currentPage) {
      let from = (currentPage - 1) * this.finishInfoPageSize;
      if (currentPage > 1 && from >= this.finishInfoRawList.length) {
        this.finishInfoCurPage = this.finishInfoRawList.length / this.finishInfoPageSize;
        from = (this.finishInfoCurPage - 1) * this.finishInfoPageSize;
      }
      let to = from + this.finishInfoPageSize;
      this.finishInfoPageList = [];
      for (; from < to; from++) {
        if (list[from]) {
          this.finishInfoPageList.push(list[from]);
        }
      }
    },
    drawChart() {
      let myChart = echarts.init(this.$refs.charts);
      let options = {
        title: {
          text: "最近30天我的完成量",
          x: "center",
          textStyle: {
            fontSize: 16,
            fontWeight: "normal",
            color: "#696969",
          },
        },
        legend: {
          orient: "vertical",
          top: 20,
          left: 10,
          textStyle:{
            color: "#696969",
          },
          data: [],  // legend array
        },
        tooltip: {
          trigger: "axis",
          padding: [4, 10],
          extraCssText: "text-align:left",
          formatter: function(arg) {
            let s1 = "", s2 = "", sum = 0
            for (let i = 0; i < arg.length; i++) {
              if(i === 0){
                s1 = arg[i].name + " 完成 "
              }
              s2 += "<br/>" + arg[i].value + " 道 " + arg[i].seriesName + " 题"
              sum += arg[i].value
            }
            return s1 + sum + " 道题" + s2
          }
        },
        xAxis: {
          type: "category",
          data: this.finishInfoChartData.xData,  // x axis data
        },
        yAxis: {
          type: "value",
        },
        series: []    // y axis data
      }
      this.finishInfoChartData.yData.forEach(function (v, k){
        options.legend.data.push(k)
        let color = "#bd" + (Math.floor(Math.random()*5)+11).toString(16) + "7" + (Math.floor(Math.random()*5)+11).toString(16) + "f"
        options.series.push({
          name: k,
          type: "bar",
          stack:"总量",
          data: v,
          itemStyle: {
            normal: {
              color: color,
              lineStyle: {
                color: color,
              },
            },
          },
        })
      })
      myChart.setOption(options);
    }
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

  #charts {
    height: 350px;
    background: $base-white;
    margin-top: 40px;
    -webkit-box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    box-shadow: 4px 4px 40px rgba(0, 0, 0, 0.05);
    border-color: rgba(0, 0, 0, 0.05);
  }

  .finish_info {
    margin-top: 40px;

    .title {
      font-size: 15px;
      padding: 10px 0;
      color: #696969;
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

