<template>
  <div class="do-main">
    <iframe :src="problem.link" class="frame"></iframe>
  </div>
</template>

<script>

import {getProblemByID} from "@api";

export default {
  data () {
    return {
      problem: {}
    }
  },
  mounted() {
    let problemId = this.$route.query.problem_id
    getProblemByID({problem_id: problemId}).then(data => {
      this.problem = data
    })
  },
}
</script>

<style scoped>
.do-main{
  width: 100%;
  height: 100%
}
.frame{
  width: 100%;
  height: 2000px;
  border: 0;
}
</style>