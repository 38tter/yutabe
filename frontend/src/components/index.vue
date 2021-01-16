<template>
  <div id="app">
    <div id="header"/>
    <div>
      <h1>夕ご飯何食べる？（ゆうたべ）</h1>
      <h3>住所を入力してください</h3>
    </div>
    <el-form>
      <el-form-item>
        <textarea v-model="address" placeholder="Input your address"></textarea>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click.native="postAddress">
          送信
        </el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'index',
  data: function () {
    return {
      products: [],
      address: ''
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
      this.axios.get('http://localhost:8081/api/products').then(
        res => {
          if (res.status !== 200) {
            throw new Error('response error')
          } else {
            var resultProducts = res.data
            this.products = resultProducts
          }
        })
    },
    postAddress () {
      const message = {'address': this.address}
      let messageJson = JSON.stringify(message)
      return this.axios.post('http://localhost:8081/api/address', messageJson).then(res => {
        console.log(res.status)
      }).catch(err => Promise.reject(err))
    }
  }
}
</script>

<style scoped>

</style>
