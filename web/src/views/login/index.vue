<template>
  <div class="login-container">
    <vue-particles color="#fff" :particlesNumber='60' :moveSpeed='1.5' :lineOpacity='0.5' class="bg"></vue-particles>
    <div class="login-form" id="login-form">
      <el-row :gutter="20">
        <el-col :lg="6" :sm="10" class="aa">
          <h3>{{ $t('login.system') }}</h3>
          <el-form :model="loginForm" status-icon :rules="loginRules" ref="ruleForm" label-width="100px">
            <el-form-item :label="$t('login.username')" :required="true" prop="name">
              <el-input v-model="loginForm.name"></el-input>
            </el-form-item>
            <el-form-item :label="$t('login.password')" :required="true" prop="password">
              <el-input type="password" v-model="loginForm.password" autocomplete="off" show-password></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary" @click="loginFunc('ruleForm')" class="btn">登录</el-button>
              <el-button type="info" @click="registerFunc('ruleForm')" class="btn">注册</el-button>
            </el-form-item>
          </el-form>
        </el-col>
      </el-row>
    </div>
  </div>
</template>

<script>
import {login, register} from "@api";
import md5 from 'js-md5';

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
    loginFunc(formName) {
      this.$refs[formName].validate(valid => {
            if (valid) {
              let temp = this.loginForm.password
              this.loginForm.password = md5(temp)
              login(this.loginForm).then(data => {
                this.$store.commit("COMMIT_TOKEN", data.token);
                this.$store.commit("COMMIT_USER", data.user);
                this.$messages('success', '登录成功')
                this.$router.push({path: "/home"});
              }).catch(() => {
                this.loginForm.password = temp
              });
            } else {
              console.log(valid)
            }
          }
      );
    },
    registerFunc(formName) {
      this.$refs[formName].validate(valid => {
        if (valid) {
          let temp = this.loginForm.password
          this.loginForm.password = md5(temp)
          register(this.loginForm).then(() => {
            this.loginForm.password = temp
            this.$messages('success', '注册成功，请点击登录')
          }).catch(() => {
            this.loginForm.password = temp
          })
        }
      });
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
      font-weight: bolder;
    }

    h3 {
      line-height: 60px;
      margin-left: 100px
    }

    .btn {
      margin-left: 35px;
      margin-right: 35px;
    }
  }
}
</style>

