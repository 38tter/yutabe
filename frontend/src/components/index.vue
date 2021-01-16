<template>
  <div id="app">
    <div id="header"/>
    <div>
      <h1>夕ご飯何食べる？（ゆうたべ）</h1>
      <h2>
        {{message}}
      </h2>
      <h2>{{products}}</h2>
    </div>
  </div>
</template>

<script>
export default {
  name: 'index',
  data: function () {
    return {
      products: [],
      // productName: '',
      // productMemo: '',
      message: 'hello, world!'
    }
  },

  computed: {
    computedProducts () {
      return this.products
    }
  },

  created: function () {
    this.fetchProducts()
  },

  methods: {
    fetchProducts () {
      this.axios.get('/products').then(
        res => {
          if (res.status !== 200) {
            throw new Error('response error')
          } else {
            var resultProducts = res.data
            this.products = resultProducts
          }
        })
    }
  }
}
</script>

<style scoped>

</style>
