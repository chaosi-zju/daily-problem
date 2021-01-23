<template>
  <div class="login-container">
    <vue-particles color="#fff" :particlesNumber='100' :moveSpeed='1.5' :lineOpacity='0.5' class="bg"></vue-particles>
    <div class="login-form" id="login-form">
      <el-row type="flex" justify="center">
        <el-col :xs="3" :sm="5" :md="6" :lg="7" :xl="8"></el-col>
        <el-col :xs="16" :sm="12" :md="10" :lg="8" :xl="6" class="aa">
          <h3>{{ $t('login.system') }}</h3>
          <el-form :model="loginForm" status-icon :rules="loginRules" ref="ruleForm" label-width="100px">
            <el-form-item :label="$t('login.username')" :required="true" prop="name" class="item">
              <el-input v-model="loginForm.name"></el-input>
            </el-form-item>
            <el-form-item :label="$t('login.password')" :required="true" prop="password" class="item">
              <el-input type="password" v-model="loginForm.password" autocomplete="off" show-password></el-input>
            </el-form-item>
            <el-form-item>
              <el-row :gutter="0" type="flex" justify="center">
                <el-col :span="12">
                  <el-button type="primary" @click="loginFunc('ruleForm')">登录</el-button>
                </el-col>
                <el-col :span="12">
                  <el-button type="info" @click="registerFunc('ruleForm')">注册</el-button>
                </el-col>
              </el-row>
            </el-form-item>
          </el-form>
        </el-col>
        <el-col :xs="5" :sm="7" :md="8" :lg="9" :xl="10"></el-col>
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
            this.$confirm('注册成功，是否立即登录？', '成功', {
              confirmButtonText: '立即登录',
              cancelButtonText: '暂不登录',
              type: 'success'
            }).then(() => {
              this.loginFunc('ruleForm')
            })
          }).catch(() => {
            this.loginForm.password = temp
          })
        }
      });
    }
  }
};
</script>

<style lang="scss">
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

      .el-form-item__label{
        color: wheat;
      }
    }

    h3 {
      line-height: 60px;
      margin-left: 100px
    }
  }

  //.item {
  //
  //}
}
</style>

