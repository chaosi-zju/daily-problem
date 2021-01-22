<template>
  <div>hello problem {{ problemId }}</div>
</template>


<script>
import {getProblemByID} from '@api'

export default {
  data() {
    return {
      problemId: 0,
    };
  },

  mounted() {
    this.problemId = this.$route.query.problem_id
    getProblemByID({problem_id: this.problemId}).then(data => {
      console.log(data)
    })
  },

  methods: {}
};
</script>