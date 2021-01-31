<template>
  <div class="do-main">
    <iframe :src="problem.link" class="frame" id="frame"></iframe>
  </div>
</template>

<script>

import {getProblemByID} from "@api";

export default {
  data() {
    return {
      problem: {}
    }
  },
  mounted() {
    let problemId = this.$route.query.problem_id
    getProblemByID({problem_id: problemId}).then(data => {
      this.problem = data
      if (data.link.indexOf('leetcode.com') !== -1) {
        this.$message({
          message: '该网站似乎不支持iframe嵌入，请前往原网站做题',
          type: 'warning',
          showClose: true,
          duration: 8000
        })
      }
    })
    // iframe-宽高自适应显示
    const frame = document.getElementById('frame');
    const deviceWidth = document.documentElement.clientWidth;
    const deviceHeight = document.documentElement.clientHeight;
    frame.style.width = (Number(deviceWidth) - 10) + 'px';
    frame.style.height = (Number(deviceHeight) - 5) + 'px';
  },
}
</script>

<style scoped>
.do-main {
  width: 100%;
  height: 100%
}

.frame {
  width: 100%;
  height: 100%;
  border: 0;
}
</style>