<template>
  <div>
    <!-- 表格数据 -->
    <div class="table-content">
      <el-table :data="pageList" stripe style="width: 100%;">
        <el-table-column prop="1" label="1"></el-table-column>
        <el-table-column prop="2" label="2"></el-table-column>
        <el-table-column prop="3" label="3"></el-table-column>
        <el-table-column prop="4" label="4"></el-table-column>
      </el-table>
      <div class="block">
        <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="currentPage"
            :page-sizes="[5, 10, 12, 20]"
            :page-size="10"
            layout="total, sizes, prev, pager, next, jumper"
            :total="rawList.length"
        ></el-pagination>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      rawList:  [
        {
          "1": "1协信01",
          "2": 691861.0999317318,
          "3": "信用债",
          "4": 0.0027959958264152343
        }, {
          "1": "2朗诗01",
          "2": 690131.4471819025,
          "3": "信用债",
          "4": 0.002789005836849196
        }, {
          "1": "3三盛01",
          "2": 688816.9110920322,
          "3": "信用债",
          "4": 0.0027836934447790073
        }, {
          "1": "4三鼎01",
          "2": 685426.7917023668,
          "3": "信用债",
          "4": 0.002769993065229573
        }, {
          "1": "5临开债",
          "2": 676640.4401650192,
          "3": "信用债",
          "4": 0.00273448506769905
        }, {
          "1": "6华讯01",
          "2": 614990.17198298,
          "3": "信用债",
          "4": 0.0024853398381849607
        }, {
          "1": "7花样03",
          "2": 614990.0028613778,
          "3": "信用债",
          "4": 0.0024853391547193142
        }, {
          "1": "8协信01",
          "2": 614987.6443837617,
          "3": "信用债",
          "4": 0.0024853296234802085
        }, {
          "1": "9三盛03",
          "2": 461240.73328782123,
          "3": "信用债",
          "4": 0.0018639972176101563
        }, {
          "1": "10山钢03",
          "2": 384367.27773985104,
          "3": "信用债",
          "4": 0.0015533310146751303
        }, {
          "1": "11甘公01",
          "2": 324002.01240352966,
          "3": "信用债",
          "4": 0.0013093788254893862
        }, {
          "1": "12新湖债",
          "2": 307493.82219188084,
          "3": "信用债",
          "4": 0.0012426648117401043
        }, {
          "1": "13珠管01",
          "2": 303035.16177009855,
          "3": "信用债",
          "4": 0.0012246461719698726
        }, {
          "1": "14重机债",
          "2": 299103.36126325984,
          "3": "信用债",
          "4": 0.0012087567140880767
        }, {
          "1": "15三鼎01",
          "2": 8163.960979194436,
          "3": "信用债",
          "4": 3.2992750751699765E-5
        }, {
          "1": "16重机债",
          "2": 1475.2323613477674,
          "3": "信用债",
          "4": 5.961808700804324E-6
        }, {
          "1": "17甘公01",
          "2": 723.1485963397557,
          "3": "信用债",
          "4": 2.92243697100979E-6
        }, {
          "1": "18新湖债",
          "2": 707.2357910413259,
          "3": "信用债",
          "4": 2.85812906700224E-6
        }, {
          "1": "19珠管01",
          "2": 153.74691109594042,
          "3": "信用债",
          "4": 6.213324058700521E-7
        }],
      currentPage: 1,
      pageSize: 10,
      pageList: []
    };
  },

  mounted() {
    // this.axios.get("/api/xxx").then(response => {
      // this.rawList = response.data;
      // this.currentChangePage(this.rawList, 1);
    // });
    this.currentChangePage(this.rawList, 1);
  },

  methods: {
    handleSizeChange: function(pageSize) {
      this.pageSize = pageSize;
      this.handleCurrentChange(this.currentPage);
    },
    handleCurrentChange: function(currentPage) {
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