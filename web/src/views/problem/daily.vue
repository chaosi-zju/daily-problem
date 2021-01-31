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
            <el-button type="info" @click="jumpToLink(scope.row.link)" class="el-icon-paperclip">原题</el-button>
            <el-button type="info" @click="jumpToLocal(scope.row.ID)" class="el-icon-edit-outline">做题</el-button>
            <el-button type="info" @click="jumpToResult(scope.row.ID)" class="el-icon-document">答案</el-button>
            <el-button type="info" @click="finish(scope.$index)" class="el-icon-circle-check">完成</el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页栏 -->
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
      <!-- 再来几道 -->
      <div style="float:left">
        <el-link icon="el-icon-question" type="warning" v-if="rawList.length === 0"
                 @click="showDialog">不够？再来几道
        </el-link>
      </div>

      <el-dialog title="哪种题再来几道?" :visible.sync="more.dialogVisible" width="30%">
        <el-row :gutter="30" type="flex" justify="center">
          <el-col :span="12">
            <el-select v-model="more.ttype" placeholder="请选择题目类型">
              <el-option v-for="item in more.typeOptions" :key="item" :label="item" :value="item"></el-option>
            </el-select>
          </el-col>
          <el-col :span="12">
            <el-input v-model="more.num" placeholder="请输入题目数量"></el-input>
          </el-col>
        </el-row>
        <span slot="footer" class="dialog-footer">
          <el-button class="more-btn" @click="more.dialogVisible = false">取 消</el-button>
          <el-button class="more-btn" type="primary" @click="getMoreProblem">确 定</el-button>
        </span>
      </el-dialog>
    </div>
  </div>
</template>

<script>
import {dailyProblem, finishProblem, getAllTypes, moreProblem} from "@api";
import InsertProblem from "@/views/problem/insert";

export default {
  components: {InsertProblem},
  data() {
    return {
      rawList: [],
      currentPage: 1,
      pageSize: 10,
      pageList: [],
      more: {
        dialogVisible: false,
        typeOptions: [],
        ttype: '',
        num: 1,
      }
    };
  },
  mounted() {
    dailyProblem().then(data => {
      this.rawList = data ? data : []
      this.currentChangePage(this.rawList, 1)
    })
  },

  methods: {
    jumpToLink: function (link) {
      window.open(link)
    },
    jumpToLocal: function (id) {
      this.$store.commit('SET_COLLAPSE', true)
      this.$store.commit('SET_FULLSCREEN', true)
      this.$router.push({path: "/problemDo", query: {problem_id: id}})
    },
    jumpToResult: function (id) {
      let cur = this.$store.state.curProblem
      cur.id = id
      cur.isInPlan = true
      this.$store.commit('SET_CUR_PROBLEM', cur)
      this.$router.push({path: "/problemInfo"})
    },
    finish: function (idx) {
      this.$confirm("完成后今日将不再展示该题，确认完成？", '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        finishProblem({problem_id: this.pageList[idx].ID}).then(() => {
          this.$messages('success', 'success')
          this.rawList.splice((this.currentPage - 1) * this.pageSize + idx, 1)
          this.currentChangePage(this.rawList, this.currentPage)
        })
      });
    },
    showDialog: function () {
      getAllTypes().then(data => {
        this.more.typeOptions = data ? data : []
        this.more.dialogVisible = true
      })
    },
    getMoreProblem: function () {
      if (this.more.ttype === "") {
        this.$messages('error', '请选择题目类型')
        return
      }
      moreProblem({type: this.more.ttype, num: this.more.num}).then(data => {
        this.rawList = data ? data : []
        this.currentChangePage(this.rawList, 1)
        this.more.dialogVisible = false
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
.el-button--info {
  width: auto;
  height: 25px;
  font-size: 10px;
  padding: 3px 6px;
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

.more-btn {
  width: auto;
  height: 27px;
  font-size: 13px;
  padding: 5px 10px;
  margin: 5px 5px;
}
</style>

