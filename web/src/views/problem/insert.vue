<template>
  <div class="insert-problem">
    <el-form :model="insertForm" status-icon :rules="insertRules" ref="insertForm" label-width="100px">
      <el-form-item label="题目标题" :required="true" prop="name">
        <el-input v-model="insertForm.name"></el-input>
      </el-form-item>
      <el-form-item label="题目链接">
        <el-input v-model="insertForm.link"></el-input>
      </el-form-item>
      <el-row :gutter="10">
        <el-col :span="12">
          <el-form-item label="题目类别" :required="true" prop="type">
            <el-input v-model="insertForm.type" label="类别"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="子类别">
            <el-input v-model="insertForm.sub_type" label="类别"></el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="题目内容">
        <mavon-editor :boxShadow="false" :autofocus="false" codeStyle="atom-one-light" v-model="insertForm.content"/>
      </el-form-item>
      <el-form-item label="题目答案">
        <mavon-editor :boxShadow="false" :autofocus="false" codeStyle="atom-one-light" v-model="insertForm.result"/>
      </el-form-item>
      <el-form-item label="是否公开" style="text-align: left">
        <el-radio v-model="insertForm.is_public" :label="true" border size="small">公开</el-radio>
        <el-radio v-model="insertForm.is_public" :label="false" border size="small">个人</el-radio>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submit('insertForm')">提交本题</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {insertProblem} from "@api"

export default {
  name: "insert-problem",
  data() {
    let validateFunc = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("不能为空"))
      } else {
        callback()
      }
    };
    return {
      insertForm: {
        name: "",
        link: "",
        content: "",
        result: "",
        type: "",
        sub_type: "",
        is_public: true
      },
      insertRules: {
        name: [{validator: validateFunc, trigger: "blur"}],
        type: [{validator: validateFunc, trigger: "blur"}],
      }
    }
  },
  methods: {
    submit: function (formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          insertProblem(this.insertForm).then(data=>{
            this.$messages('success', 'success')
            let cur = this.$store.state.curProblem
            cur.id = data.ID
            cur.isInPlan = true
            cur.isInToday = false
            this.$store.commit('SET_CUR_PROBLEM', cur)
            this.$router.push({path: '/problemInfo'})
          })
        }
      })
    }
  }
}
</script>

<style scoped>
.insert-problem {
  width: 90%;
  padding: 40px 0;
}

.insert-problem >>> .el-input__inner {
  height: 33px;
  line-height: 33px;
}

</style>