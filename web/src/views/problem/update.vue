<template>
  <div class="update-problem">
    <el-form :model="updateForm" status-icon :rules="updateRules" ref="updateForm" label-width="100px">
      <el-form-item label="题目标题" :required="true" prop="name">
        <el-input v-model="updateForm.name"></el-input>
      </el-form-item>
      <el-form-item label="题目链接">
        <el-input v-model="updateForm.link"></el-input>
      </el-form-item>
      <el-row :gutter="10">
        <el-col :span="12">
          <el-form-item label="题目类别" :required="true" prop="type">
            <el-input v-model="updateForm.type" label="类别"></el-input>
          </el-form-item>
        </el-col>
        <el-col :span="12">
          <el-form-item label="子类别">
            <el-input v-model="updateForm.sub_type" label="类别"></el-input>
          </el-form-item>
        </el-col>
      </el-row>
      <el-form-item label="题目内容">
        <mavon-editor :boxShadow="false" codeStyle="atom-one-light" v-model="updateForm.content"/>
      </el-form-item>
      <el-form-item label="题目答案">
        <mavon-editor :boxShadow="false" codeStyle="atom-one-light" v-model="updateForm.result"/>
      </el-form-item>
      <el-form-item label="是否公开" style="text-align: left">
        <el-radio v-model="updateForm.is_public" :label="true" border size="small">公开</el-radio>
        <el-radio v-model="updateForm.is_public" :label="false" border size="small">个人</el-radio>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submit('updateForm')">确定更新</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import {getProblemByID, updateProblem} from "@api"

export default {
  name: "update-problem",
  data() {
    let validateFunc = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("不能为空"))
      } else {
        callback()
      }
    };
    return {
      updateForm: {
        id: 0,
        name: "",
        link: "",
        content: "",
        result: "",
        type: "",
        sub_type: "",
        is_public: false
      },
      updateRules: {
        name: [{validator: validateFunc, trigger: "blur"}],
        type: [{validator: validateFunc, trigger: "blur"}],
      }
    }
  },
  mounted() {
    this.updateForm.id = this.$store.state.curProblem.id
    getProblemByID({problem_id: this.updateForm.id}).then(data => {
      this.updateForm = data
    })
  },
  methods: {
    submit: function (formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          updateProblem(this.updateForm).then(data => {
            this.$messages('success', 'success')
            this.$router.push({path: '/problemInfo'})
          })
        }
      })
    }
  }
}
</script>

<style scoped>
.update-problem {
  width: 90%;
  padding: 40px 0;
}

.update-problem >>> .el-input__inner {
  height: 33px;
  line-height: 33px;
}

</style>