<template>
  <div>
    <!-- 表格数据 -->
    <div class="table-content">
      <el-input placeholder="搜索题目关键字" v-model="keyword" clearable @change="searchProblem">
        <i slot="prefix" class="el-input__icon el-icon-search"></i>
      </el-input>
      <el-table :data="pageList" stripe style="width: 100%;">
        <el-table-column label="题目类别" min-width="19%">
          <template slot-scope="scope">
            <div class="type-big">{{ scope.row.type }}</div>
            <div class="type-small">{{ scope.row.sub_type }}</div>
          </template>
        </el-table-column>
        <el-table-column prop="name" label="题目标题" min-width="56%"></el-table-column>
        <el-table-column label="操作" min-width="25%">
          <template slot-scope="scope">
            <el-button type="info" @click="jumpToInfo(scope.row.ID)" class="el-icon-document">详情</el-button>
            <el-button type="info" @click="doItToday(scope.row.ID)" class="el-icon-circle-plus-outline">加入任务</el-button>
            <el-button type="danger" @click="removeFromPlan(scope.$index)" class="el-icon-delete">移出计划</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="block">
        <el-pagination style="padding-top: 14px"
                       @size-change="handleSizeChange"
                       @current-change="handleCurrentChange"
                       :current-page="currentPage"
                       :page-sizes="[5, 10, 20, 50]"
                       :page-size="10"
                       layout="total, sizes, prev, pager, next, jumper"
                       :total="rawList.length"
        ></el-pagination>
      </div>
    </div>
  </div>
</template>

<script>
import {addToDaily, problemPlan, problemPlanQ, removeFromProblemPlan} from "@api";

export default {
  data() {
    return {
      keyword: "",
      originList: [],
      rawList: [],
      currentPage: 1,
      pageSize: 10,
      pageList: []
    };
  },
  mounted() {
    let planCached = this.$store.state.problemPlan
    if(planCached instanceof Array && planCached.length > 0){
      this.originList = planCached
      this.rawList = planCached
      this.currentChangePage(this.rawList, 1)
      problemPlanQ().then(res => {
        if(res.data.code === 200) {
          this.$store.commit('SET_PROBLEM_PLAN', res.data.data)
          console.log('cache updated')
        }else{
          console.log('fetch cache failed')
          console.log(res.data)
        }
      })
    } else {
      problemPlan().then(data => {
        this.originList = data ? data : []
        this.rawList = data ? data : []
        this.currentChangePage(this.rawList, 1)
        this.$store.commit('SET_PROBLEM_PLAN', data)
      })
    }
  },
  methods: {
    jumpToInfo: function (id) {
      let cur = this.$store.state.curProblem
      cur.id = id
      cur.isInPlan = true
      cur.isInToday = false
      this.$store.commit('SET_CUR_PROBLEM', cur)
      this.$router.push({path: "/problemInfo"})
    },
    doItToday: function (id) {
      addToDaily({problem_id: id}).then(() => {
        this.$messages('success', 'success')
      })
    },
    removeFromPlan: function (idx) {
      this.$confirm('移出即默认完成了此题并不再做此题，确认移出？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        removeFromProblemPlan({problem_id: this.pageList[idx].ID}).then(() => {
          this.$messages('success', 'success')
          this.rawList.splice((this.currentPage - 1) * this.pageSize + idx, 1)
          this.currentChangePage(this.rawList, this.currentPage)
        })
      });
    },
    searchProblem: function () {
      let keyword = this.keyword
      if (keyword) {
        let reg = new RegExp(keyword, 'ig')
        this.rawList = this.originList.filter(function (e) {
          if (e.ID.toString() === keyword || e.type.match(reg) || e.sub_type.match(reg) ||
              e.name.match(reg) || e.link.match(reg) || e.content.match(reg) || e.result.match(reg)) {
            return e
          }
        })
      } else {
        this.rawList = this.originList
      }
      this.currentChangePage(this.rawList, 1)
    },
    handleSizeChange: function (pageSize) {
      this.pageSize = pageSize;
      this.handleCurrentChange(this.currentPage);
    },
    handleCurrentChange: function (currentPage) {
      this.currentPage = currentPage;
      this.currentChangePage(this.rawList, currentPage);
    },
    currentChangePage(list, currentPage) {
      let from = (currentPage - 1) * this.pageSize;
      if (currentPage > 1 && from >= this.rawList.length) {
        this.currentPage = this.rawList.length / this.pageSize;
        from = (this.currentPage - 1) * this.pageSize;
      }
      let to = from + this.pageSize;
      this.pageList = [];
      for (; from < to; from++) {
        if (list[from]) {
          this.pageList.push(list[from]);
        }
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.el-button {
  width: auto;
  height: 25px;
  font-size: 10px;
  padding: 3px 5px;
  margin: 2px 5px;
}

.type-big {
  font-size: 14px;
  font-weight: 550;
}

.type-small {
  font-size: xx-small;
  font-weight: 300;
}
</style>

