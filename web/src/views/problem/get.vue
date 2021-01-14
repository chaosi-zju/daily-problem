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
          "name": "1协信01",
          "marketValue": 691861.0999317318,
          "type": "信用债",
          "ratio": 0.0027959958264152343
        }, {
          "name": "2朗诗01",
          "marketValue": 690131.4471819025,
          "type": "信用债",
          "ratio": 0.002789005836849196
        }, {
          "name": "3三盛01",
          "marketValue": 688816.9110920322,
          "type": "信用债",
          "ratio": 0.0027836934447790073
        }, {
          "name": "4三鼎01",
          "marketValue": 685426.7917023668,
          "type": "信用债",
          "ratio": 0.002769993065229573
        }, {
          "name": "5临开债",
          "marketValue": 676640.4401650192,
          "type": "信用债",
          "ratio": 0.00273448506769905
        }, {
          "name": "6华讯01",
          "marketValue": 614990.17198298,
          "type": "信用债",
          "ratio": 0.0024853398381849607
        }, {
          "name": "7花样03",
          "marketValue": 614990.0028613778,
          "type": "信用债",
          "ratio": 0.0024853391547193142
        }, {
          "name": "8协信01",
          "marketValue": 614987.6443837617,
          "type": "信用债",
          "ratio": 0.0024853296234802085
        }, {
          "name": "9三盛03",
          "marketValue": 461240.73328782123,
          "type": "信用债",
          "ratio": 0.0018639972176101563
        }, {
          "name": "10山钢03",
          "marketValue": 384367.27773985104,
          "type": "信用债",
          "ratio": 0.0015533310146751303
        }, {
          "name": "11甘公01",
          "marketValue": 324002.01240352966,
          "type": "信用债",
          "ratio": 0.0013093788254893862
        }, {
          "name": "12新湖债",
          "marketValue": 307493.82219188084,
          "type": "信用债",
          "ratio": 0.0012426648117401043
        }, {
          "name": "13珠管01",
          "marketValue": 303035.16177009855,
          "type": "信用债",
          "ratio": 0.0012246461719698726
        }, {
          "name": "14重机债",
          "marketValue": 299103.36126325984,
          "type": "信用债",
          "ratio": 0.0012087567140880767
        }, {
          "name": "15三鼎01",
          "marketValue": 8163.960979194436,
          "type": "信用债",
          "ratio": 3.2992750751699765E-5
        }, {
          "name": "16重机债",
          "marketValue": 1475.2323613477674,
          "type": "信用债",
          "ratio": 5.961808700804324E-6
        }, {
          "name": "17甘公01",
          "marketValue": 723.1485963397557,
          "type": "信用债",
          "ratio": 2.92243697100979E-6
        }, {
          "name": "18新湖债",
          "marketValue": 707.2357910413259,
          "type": "信用债",
          "ratio": 2.85812906700224E-6
        }, {
          "name": "19珠管01",
          "marketValue": 153.74691109594042,
          "type": "信用债",
          "ratio": 6.213324058700521E-7
        }],
      currentPage: 1,
      pageSize: 10,
      pageList: []
    };
  },

  mounted() {
    this.axios.get("/api/xxx").then(response => {
      // this.rawList = response.data;
      this.currentChangePage(this.rawList, 1);
    });
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