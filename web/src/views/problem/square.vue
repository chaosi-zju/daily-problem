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
            <el-button type="success" @click="addToPlan(scope.$index)" class="el-icon-circle-check">加入学习计划</el-button>
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
import {addToProblemPlan, problemSquare} from "@api";

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
    problemSquare().then(data => {
      this.originList = data ? data : []
      this.rawList = data ? data : []
      this.currentChangePage(this.rawList, 1)
    })
  },

  methods: {
    jumpToInfo: function (id) {
      let cur = this.$store.state.curProblem
      cur.id = id
      cur.isInPlan = false
      cur.isInToday = false
      this.$store.commit('SET_CUR_PROBLEM', cur)
      this.$router.push({path: "/problemInfo"})
    },
    addToPlan: function (idx) {
      addToProblemPlan({problem_id: this.pageList[idx].ID}).then(() => {
        this.$messages('success', 'success')
        this.rawList.splice((this.currentPage - 1) * this.pageSize + idx, 1)
        this.currentChangePage(this.rawList, this.currentPage)
      })
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

