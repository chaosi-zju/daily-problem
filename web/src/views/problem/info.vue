<template>
  <div class="mavonEditor">
    <el-row type="flex" justify="end" class="title-btn">
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

    <el-row type="flex" justify="end" class="bottom-btn">
      <el-button icon="el-icon-arrow-left" @click="pre" v-if="JSON.stringify(preProblem)!=='{}'">
        上一题：{{ preProblem.name }}
      </el-button>
      <el-button @click="next" v-if="JSON.stringify(nextProblem)!=='{}'">
        下一题：{{ nextProblem.name }}
        <i class="el-icon-arrow-right"/>
      </el-button>
    </el-row>

  </div>
</template>


<script>
import {addToProblemPlan, finishProblem, getProblemByID, removeFromProblemPlan} from '@api'

export default {
  name: "info-problem",
  data() {
    return {
      problemId: 0,
      isInToday: false,
      content: "",
      isCreator: false,
      preProblem: {},
      nextProblem: {}
    };
  },
  computed: {
    isInPlan() {
      return this.$store.state.curProblem.isInPlan
    }
  },
  mounted() {
    this.problemId = this.$store.state.curProblem.id
    this.isInToday = this.$store.state.curProblem.isInToday
    if (this.isInToday) {
      let idx = this.$store.state.curProblem.idx
      let len = this.$store.state.curProblem.rawList.length
      if (idx - 1 >= 0) {
        this.preProblem = this.$store.state.curProblem.rawList[idx - 1]
      }
      if (idx + 1 < len) {
        this.nextProblem = this.$store.state.curProblem.rawList[idx + 1]
      }
    }
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
          this.isInToday = false
          let cur = this.$store.state.curProblem
          cur.isInToday = false
          this.$store.commit('SET_CUR_PROBLEM', cur)
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
          this.isInToday = false
          let cur = this.$store.state.curProblem
          cur.isInPlan = false
          cur.isInToday = false
          this.$store.commit('SET_CUR_PROBLEM', cur)
          this.$messages('success', 'success')
        })
      });
    },
    addPlan: function () {
      addToProblemPlan({problem_id: this.problemId}).then(() => {
        let cur = this.$store.state.curProblem
        cur.isInPlan = true
        this.$store.commit('SET_CUR_PROBLEM', cur)
        this.$messages('success', 'success')
      })
    },
    pre: function () {
      let cur = this.$store.state.curProblem
      if (!cur.isInToday || !cur.isInPlan) {
        cur.rawList.splice(cur.idx, 1)
      }
      cur.id = this.preProblem.ID
      cur.idx = (cur.idx - 1) % cur.rawList.length
      cur.isInPlan = true
      cur.isInToday = true
      this.$store.commit('SET_CUR_PROBLEM', cur)
      location.reload()
    },
    next: function () {
      let cur = this.$store.state.curProblem
      if (!cur.isInToday || !cur.isInPlan) {
        cur.rawList.splice(cur.idx, 1)
        cur.idx = cur.idx % cur.rawList.length
      } else {
        cur.idx = (cur.idx + 1) % cur.rawList.length
      }
      cur.id = this.nextProblem.ID
      cur.isInPlan = true
      cur.isInToday = true
      this.$store.commit('SET_CUR_PROBLEM', cur)
      location.reload()
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

.title-btn >>> .el-button {
  width: auto;
  height: 30px;
  font-size: 12px;
  padding: 6px 10px;
}

.bottom-btn >>> .el-button {
  width: auto;
  height: 40px;
  font-size: 13px;
  padding: 10px 10px;
  margin-top: 20px;;
  max-width: 500px;
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
  display: block;
}


</style>
