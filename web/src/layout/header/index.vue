<template>
  <div class="head-container clearfix">
    <div class="header-left">
      <showAside :toggle-click="toggleClick"/>
      <breadcrumb />
    </div>

    <div class="header-center" v-show="tomato_time_str!==''">
      <span style="font-size: 13px; letter-spacing: 1px; color: #3a8ee6">番茄开启 请专注 ：</span>
      <span>{{ tomato_time_str }}</span>
    </div>

    <div class="header-right">
      <div class="header-user-con">
        <!-- 全屏显示 -->
        <div class="btn-bell" @click="startTomato">
          <el-tooltip effect="dark" content="开启番茄钟 🍅" placement="bottom">
            <i class="el-icon-video-play" v-if="!tomato_run"></i>
            <i class="el-icon-video-pause" v-else></i>
          </el-tooltip>
        </div>
        <!-- 全屏显示 -->
        <div class="btn-fullscreen" @click="handleFullScreen">
          <el-tooltip effect="dark" :content="fullscreen?$t('header.cancelFullScreen'):$t('header.fullScreen')" placement="bottom">
            <i class="el-icon-rank"></i>
          </el-tooltip>
        </div>
        <!-- 多语言 -->
        <select-lang></select-lang>
        <!-- 消息中心 -->
        <div class="btn-bell">
          <el-tooltip effect="dark" :content="$t('header.message')" placement="bottom">
            <router-link to="/home">
             <i class="el-icon-bell"></i>
             </router-link>
          </el-tooltip>
          <span class="btn-bell-badge" v-if="message"></span>
        </div>
        <!-- 用户名下拉菜单 -->
        <el-dropdown class="avatar-container" trigger="click">
          <div class="avatar-wrapper">
            <img
              src="https://ss0.bdstatic.com/70cFvHSh_Q1YnxGkpoWK1HF6hhy/it/u=3266090804,66355162&fm=26&gp=0.jpg"
              class="user-avatar"
            >
            {{username }}<i class="el-icon-caret-bottom"/>
          </div>
          <el-dropdown-menu slot="dropdown" class="user-dropdown">
            <router-link class="inlineBlock" to="/home">
              <el-dropdown-item>{{$t('route.home')}}</el-dropdown-item>
            </router-link>
            <el-dropdown-item>{{$t('header.setting')}}</el-dropdown-item>
            <el-dropdown-item divided>
              <span style="display:block;" @click="logout">{{$t('header.logout')}}</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>

<script>
import showAside from "./showAside";
import selectLang from './selectLang'
import breadcrumb from './breadCrumb'
export default {
  name:'header',
  components: {
    showAside,
    selectLang,
    breadcrumb
  },
  data() {
    return {
      tomato_time_str: "",
      tomato_run: false,
      tomato_time_init: 300, // 60s
      fullscreen: this.$store.state.isFullScreen,
      message: 2,
      username: this.$store.state.user.name
    };
  },
  watch:{
    isFullScreen(val, oldVal){
      this.fullscreen = val
      let element = document.documentElement;
      if (oldVal === false && val === true) {
        if (element.requestFullscreen) {
          element.requestFullscreen();
        } else if (element.webkitRequestFullScreen) {
          element.webkitRequestFullScreen();
        } else if (element.mozRequestFullScreen) {
          element.mozRequestFullScreen();
        } else if (element.msRequestFullscreen) {
          // IE11
          element.msRequestFullscreen();
        }
      } else {
        if (document.exitFullscreen) {
          document.exitFullscreen();
        } else if (document.webkitCancelFullScreen) {
          document.webkitCancelFullScreen();
        } else if (document.mozCancelFullScreen) {
          document.mozCancelFullScreen();
        } else if (document.msExitFullscreen) {
          document.msExitFullscreen();
        }
      }
    }
  },
  computed: {
    isCollapse: {
      get: function() {
        return this.$store.state.isCollapse;
      },
      set: function(newValue) {
        this.$store.commit("SET_COLLAPSE", newValue);
      }
    },
    isFullScreen(){
      return this.$store.state.isFullScreen
    }
  },
  methods: {
    toggleClick() {
      this.isCollapse = !this.isCollapse;
    },
    // 用户名下拉菜单选择事件
    logout(command) {
      this.$store.commit('TAGS_LIST',[])
      this.$store.commit('SET_BREAD',['home'])
      this.$store.commit('COMMIT_TOKEN', '')
      this.$store.commit('COMMIT_USER', '')
      this.$router.push("/login");
    },
    // 全屏事件
    handleFullScreen() {
      let isFullScreen = this.$store.state.isFullScreen
      this.$store.commit('SET_FULLSCREEN', !isFullScreen)
    },
    startTomato() {
      if (this.tomato_time_str === "") {
        let ts = this.tomato_time_init
        this.tomato_run = true
        let timer = setInterval(() => {
          if (this.tomato_run) {
            ts--;
            this.tomato_time_str = ('0' + Math.floor(ts / 60)).slice(-2) + ':' + ('0' + ts % 60).slice(-2)
            if (ts <= 0) {
              this.tomato_time_str = ""
              this.tomato_run = false
              clearInterval(timer);
            }
          }
        }, 1000)
      }else{
        this.tomato_run = !this.tomato_run
      }
    }
  }
};
</script>

<style lang="scss" scoped>
.head-container {
  height: 52px;
  line-height: 50px;
  -webkit-box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.12),
    0 0 3px 0 rgba(0, 0, 0, 0.04);
  box-shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.12), 0 0 3px 0 rgba(0, 0, 0, 0.04);
  border-bottom: 1px solid #f0f0f0;
}
.header-left {
  float: left;
}
.header-center {
  float: left;
  padding-left: 30%;
  position: relative;
  letter-spacing: 5px;
  font-weight: bolder;
  font-size: 21px;
  color: red;
}
.header-right {
  float: right;
  padding-right: 50px;
}
.header-user-con {
  display: flex;
  height: 50px;
  align-items: center;
}
.btn-fullscreen {
  transform: rotate(45deg);
  margin-right: 5px;
  font-size: 24px;
}

.btn-fullscreen {
  position: relative;
  width: 30px;
  height: 30px;
  text-align: center;
  border-radius: 15px;
  cursor: pointer;
  margin-bottom: 10px;
}
.btn-bell{
  position: relative;
  width: 30px;
  height: 30px;
  text-align: center;
  border-radius: 15px;
  cursor: pointer;
  font-size: 24px;
  margin-right: 20px;
  margin-bottom: 15px;
}
.btn-bell-badge {
  position: absolute;
  right: 0;
  top: 8px;
  width: 8px;
  height: 8px;
  border-radius: 4px;
  background: #f56c6c;
}
.btn-bell .el-icon-bell {
  color: #666;
}
.user-name {
  margin-left: 10px;
}
.user-avator {
  margin-left: 20px;
}
.user-avator img {
  display: block;
  width: 40px;
  height: 40px;
  border-radius: 50%;
}
.el-dropdown-link {
  color: #fff;
  cursor: pointer;
}
.el-dropdown-menu__item {
  text-align: center;
}
.avatar-container {
  height: 50px;
  display: inline-block;
  // position: absolute;
  // right: 35px;
  .avatar-wrapper {
    cursor: pointer;
    margin-top: 5px;
    position: relative;
    line-height: initial;
    .user-avatar {
      width: 40px;
      height: 40px;
      border-radius: 10px;
    }
    .el-icon-caret-bottom {
      position: absolute;
      right: -20px;
      top: 25px;
      font-size: 12px;
    }
  }
}
</style>


