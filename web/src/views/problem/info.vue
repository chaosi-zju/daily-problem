<template>
    <div>
      <no-ssr>
        <mavon-editor
            :subfield="false"
            :boxShadow="false"
            codeStyle="atom-one-light"
            defaultOpen="preview"
            :toolbarsFlag="false"
            v-model="content"
            class="mavonEditor"
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
      this.content = "## " + data.name + "\n[OJ链接](" + data.link + ")\n\n" + data.content + "### 解答\n" + data.result
    })
  },

  methods: {}
};
</script>

<style scoped>
.mavonEditor {
  margin: 20px 40px;
  padding: 20px 40px;
  border: 1px solid #E1E4E8;
  border-radius: 6px;
}
</style>
