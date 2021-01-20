<template>
  <div class="login-container">
    <vue-particles color="#fff" :particlesNumber='60' :moveSpeed='1.5' :lineOpacity='0.5' class="bg"></vue-particles>
    <div class="login-form">
      <el-row :gutter="20">
        <el-col :lg="6" :sm="10" class="aa">
          <h3>{{ $t('login.system') }}</h3>
          <el-form :model="loginForm" status-icon :rules="loginRules" ref="ruleForm" label-width="100px"
                   class="login-ruleForm">
            <el-form-item :label="$t('login.username')" prop="username">
              <el-input v-model="loginForm.username"></el-input>
            </el-form-item>
            <el-form-item :label="$t('login.password')" prop="password">
              <el-input type="password" v-model="loginForm.password" autocomplete="off" show-password></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="submitForm('ruleForm')">提交</el-button>
              <el-button @click="resetForm('ruleForm')">重置</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import {login} from "@api";
import axios from "axios";

export default {
  name: "login",
  data() {
    let validatePass = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入密码"));
      } else {
        callback();
      }
    };
    let validateName = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入用户名"));
      } else {
        callback();
      }
    };
    return {
      loginForm: {
        password: "",
        username: ""
      },
      loginRules: {
        password: [{validator: validatePass, trigger: "blur"}],
        username: [{validator: validateName, trigger: "blur"}]
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
            if (valid) {
              // login(this.loginForm).then(res => {
              //   console.log(res)
              //   //提交数据到vuex
              //   this.$store.commit("COMMIT_TOKEN", res);
              //   this.$message('success', res.message)
              //   this.$router.push({
              //     path: "/home"
              //   });
              // }).catch(err => {
              //   this.$message("error", err.message);
              // });
              axios.post('/api/login', this.loginForm)
                  .then(function (response) {
                    console.log(response);
                  })
                  .catch(function (error) {
                    console.log(error);
                  });
            }
            //这里模拟管理员以及用户两种权限,一般的都是登陆后接口传过来
            // let roles=[]
            // roles.push(this.ruleForm2.username)
            // this.$store.commit("COMMIT_ROLE", roles);
            // this.$router.push({
            //       path: "/home"
            //     });
            //   login(this.ruleForm2)
            //       .then(res => {
            //         console.log(res)
            //         //提交数据到vuex
            //         this.$store.commit("COMMIT_TOKEN", res);
            //         this.$message('success', res.message)
            //         this.$router.push({
            //           path: "/"
            //         });
            //       })
            //       .catch(err => {
            //         this.$message("error", err.message);
            //       });
            // } else {
            //   return false;
            // }
          }
      );
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    }
  }
};
</script>

<style lang="scss" scoped>
.bg {
  position: fixed;
  z-index: -1;
  width: 100%;
  height: 100%;
}

.login-container {
  background: #2d3a4b;
  width: 100%;
  height: 100%;
  position: fixed;
  display: table;

  .login-form {
    display: table-cell;
    vertical-align: middle;
    text-align: center;
    color: white;
    font-size: 18px;

    .aa {
      margin: auto;
      float: none;
    }

    h3 {
      line-height: 60px;
      margin-left: 100px
    }

    .acount {
      text-align: left
    }

  }
}
</style>

