<template>
  <div class="mavonEditor">
    <el-row type="flex" justify="end">
      <el-button icon="el-icon-edit" @click="edit" v-if="isCreator">编辑</el-button>
      <el-button icon="el-icon-circle-check" @click="finish" v-if="isInToday">完成</el-button>
      <el-button icon="el-icon-delete" @click="removePlan" v-if="isInPlan">移出学习计划</el-button>
      <el-button icon="el-icon-circle-plus-outline" @click="addPlan" v-if="!isInPlan">加入学习计划</el-button>
    </el-row>

    <mavon-editor
        :subfield="false"
        :boxShadow="false"
        :autofocus="false"
        previewBackground="#ffffff"
        codeStyle="atom-one-light"
        defaultOpen="preview"
        :toolbarsFlag="false"
        style="border: 0"
        v-model="content"
    />
  </div>
</template>


<script>
import {addToProblemPlan, getProblemByID, removeFromProblemPlan} from '@api'

export default {
  name: "info-problem",
  data() {
    return {
      problemId: 0,
      content: "",
      isCreator: false,
    };
  },
  computed: {
    isInPlan() {
      return this.$store.state.curProblem.isInPlan
    },
    isInToday() {
      return this.$store.state.curProblem.isInToday
    }
  },
  mounted() {
    this.problemId = this.$store.state.curProblem.id
    getProblemByID({problem_id: this.problemId}).then(data => {
      this.content = "## " + data.name + "\n[OJ链接](" + data.link + ")\n\n" + data.content + "\n### 解答\n" + data.result
      this.isCreator = (this.$store.state.user.ID === data.creator_id)
    })
  },

  methods: {
    edit: function () {
      this.$router.push({path: "/problemUpdate"})
    },
    finish: function () {
      this.$confirm("完成后今日将不再展示该题，确认完成？", '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        finishProblem({problem_id: this.problemId}).then(() => {
          this.$messages('success', 'success')
        })
      });
    },
    removePlan: function () {
      this.$confirm('移出即默认完成了此题并不再做此题，确认移出？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        removeFromProblemPlan({problem_id: this.problemId}).then(() => {
          let cur = this.$store.state.curProblem
          cur.isInPlan = false
          this.$store.commit('SET_CUR_PROBLEM', cur)
          this.$messages('success', 'success')
        })
      });
    },
    addPlan: function (){
      addToProblemPlan({problem_id: this.problemId}).then(() => {
        let cur = this.$store.state.curProblem
        cur.isInPlan = true
        this.$store.commit('SET_CUR_PROBLEM', cur)
        this.$messages('success', 'success')
      })
    }
  }
};
</script>

<style scoped>
.mavonEditor {
  margin: 20px 40px;
  padding: 20px 80px;
  border: 1px solid #E1E4E8;
  border-radius: 6px;
}

.mavonEditor >>> .el-button {
  width: auto;
  height: 30px;
  font-size: 12px;
  padding: 6px 10px;
}
</style>
