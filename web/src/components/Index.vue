<template>
  <div class="hello">

    <input id="points" type="text" v-model="number" placeholder="Enter a number" class="form-control">
    <button @click="test">Submit</button>

    <p><b>number * 2 = </b>: {{ result }}</p>

  </div>
</template>

<script>
export default {
  name: 'Index',
  data: () => ({
    number: 0,
    result: ''
  }),
  methods: {
    async test() {

      const {body: data} = await this.$http.post('/api/test', {number: this.number})
      this.result = data

    }
  }
}
</script>
