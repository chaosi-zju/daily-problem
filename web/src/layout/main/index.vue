<template>
    <div class="container">
      <tags />
      <el-scrollbar wrap-class="scrollbar-wrapper">
        <div class="contents">
          <transition name="fade-transform" mode="out-in">
              <router-view></router-view>
          </transition>
        </div>
       </el-scrollbar>
    </div>
</template>

<script>
import Tags from './tags.vue'
export default {
    components:{
      Tags
    }
}
</script>

<style lang="scss" scoped>
.container {
  height:calc(100vh - 50px);
  position: relative;
  overflow: hidden;
  /*background: #f0f0f0;*/
  width: 100%;
  .contents{
    //height: calc(100vh - 78px);
    -webkit-box-sizing: border-box;
    box-sizing: border-box;
    padding: 10px 20px 60px 20px;
  }
}
</style>