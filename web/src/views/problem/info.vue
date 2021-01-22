<template>
  <div class="mavonEditor">
    <no-ssr>
      <mavon-editor
          :subfield="false"
          :boxShadow="false"
          defaultOpen="preview"
          :toolbarsFlag="false"
          v-model="content"
      />
    </no-ssr>
  </div>
</template>


<script>
import {getProblemByID} from '@api'

export default {
  data() {
    return {
      problemId: 0,
      content: ""
    };
  },
  mounted() {
    this.problemId = this.$route.query.problem_id
    getProblemByID({problem_id: this.problemId}).then(data => {
      this.content = "## " + data.name + "\n[OJ链接](" + data.link + ")\n" + data.content + "### 解答\n" + data.result
    })
  },

  methods: {}
};
</script>

<style scoped>
.mavonEditor {
  width: 100%;
  height: 100%;
}
</style>
