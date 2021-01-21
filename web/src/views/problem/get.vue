<template>
  <div>
    <!-- 表格数据 -->
    <div class="table-content">
      <el-table :data="pageList" stripe style="width: 100%;">
        <el-table-column prop="name" label="1"></el-table-column>
        <el-table-column prop="content" label="2"></el-table-column>
        <el-table-column prop="link" label="3"></el-table-column>
        <el-table-column prop="result" label="4"></el-table-column>
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
import {dailyProblem} from "@api";

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
    dailyProblem().then(data => {
      this.rawList = data;
      this.currentChangePage(this.rawList, 1);
    })
  },

  methods: {
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