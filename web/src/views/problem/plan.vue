<template>
  <div>
    <!-- 表格数据 -->
    <div class="table-content">
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
            <el-button type="danger" @click="removeFromPlan(scope.$index)" class="el-icon-delete">移出学习计划</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="block">
        <el-pagination
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
import {problemPlan, removeFromProblemPlan} from "@api";

export default {
  data() {
    return {
      rawList: [],
      currentPage: 1,
      pageSize: 10,
      pageList: []
    };
  },

  mounted() {
    problemPlan().then(data => {
      this.rawList = data;
      this.currentChangePage(this.rawList, 1);
    })
  },

  methods: {
    jumpToInfo: function () {

    },
    removeFromPlan: function (idx) {
      removeFromProblemPlan({problem_id: this.rawList[idx].ID}).then(() => {
        this.$messages('success', 'success')
        this.rawList.splice(idx, 1)
        this.currentChangePage(this.rawList, 1)
      })
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
      let to = currentPage * this.pageSize;
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

