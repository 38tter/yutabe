<template>
  <div id="app">
    <div id="header"/>
    <div>
      <h1>夕ご飯何食べる？（ゆうたべ）</h1>
      <h3>食材を入力してください（カンマ区切り）</h3>
    </div>
    <el-form>
      <el-form-item>
        <textarea v-model="address" placeholder="Input your address"></textarea>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click.native="fetchRecipes">
          送信
        </el-button>
      </el-form-item>
    </el-form>
    <el-table
    current-row-key
    stripe
    highlight-current-row>
      <el-table-column label="レシピ">
        <template slot-scope="scope">
          <span v-for="recipe in recipes" :key="recipe" class="spec">
            {{recipe}}
          </span>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>

export default {
  name: 'index',
  data: function () {
    return {
      products: [],
      address: '',
      ingredients: '',
      recipes: []
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
    },
    fetchRecipes () {
      const query = this.ingredients
      this.axios.get('https://recipe-puppy.p.rapidapi.com/?' + query.replace(',', '%2C'), {
        headers: {
          'X-RapidAPI-Host': 'recipe-puppy.p.rapidapi.com',
          'X-RapidAPI-Key': process.env.VUE_APP_RECIPE_PUPPY_API_KEY,
          'useQueryString': true
        },
        data: {}
      }).then(
        res => {
          if (res.status !== 200) {
            throw new Error('response error')
          } else {
            console.log(res.data)
            res.data.results.forEach((r) => {
              console.log(r.title)
              this.recipes.push(r.title)
              console.log(this.recipes)
            })
          }
        }
      )
    }
  }
}
</script>

<style scoped>

</style>
