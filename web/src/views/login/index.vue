<template>
  <div class="login-container">
    <vue-particles color="#fff" :particlesNumber='60' :moveSpeed='1.5' :lineOpacity='0.5' class="bg"></vue-particles>
    <div class="login-form">
      <el-row :gutter="20">
        <el-col :lg="6" :sm="10" class="aa">
          <h3>{{ $t('login.system') }}</h3>
          <el-form :model="loginForm" status-icon :rules="loginRules" ref="ruleForm" label-width="100px"
                   class="login-ruleForm">
            <el-form-item :label="$t('login.username')" prop="name">
              <el-input v-model="loginForm.name"></el-input>
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

export default {
  name: "login",
  data() {
    let validateName = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入用户名"));
      } else {
        callback();
      }
    };
    let validatePass = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请输入密码"));
      } else {
        callback();
      }
    };
    return {
      loginForm: {
        name: "",
        password: "",
      },
      loginRules: {
        name: [{validator: validateName, trigger: "blur"}],
        password: [{validator: validatePass, trigger: "blur"}],
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate(valid => {
            if (valid) {
              login(this.loginForm).then(data => {
                this.$store.commit("COMMIT_TOKEN", data.token);
                this.$store.commit("COMMIT_USER", data.user);
                this.$messages('success', 'success')
                this.$router.push({path: "/home"});
              }).catch(err => {
                this.$messages("error", err.message);
              });
            } else {
              console.log("err")
              return false;
            }
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

