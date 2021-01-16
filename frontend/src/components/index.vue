<template>
  <div id="app">
    <div id="header"/>
    <div>
      <h1>夕ご飯何食べる？（ゆうたべ）</h1>
      <h3>食材を入力してください（カンマ区切り）</h3>
    </div>
    <el-form>
      <el-form-item>
        <textarea v-model="ingredients" placeholder="Input your ingredients"></textarea>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click.native="fetchRecipes">
          送信
        </el-button>
      </el-form-item>
    </el-form>
    <el-row>
      <el-col :span="8" v-for="recipe in recipes" :key="recipe.title" :offset="2">
        <el-card :body-style="{ padding: '0px' }">
          <img :src="recipe.thumbnail" class="image">
          <div style="padding: 14px;">
            <span>{{recipe.title}}</span>
            <div class="bottom clearfix">
              <el-button type="text" class="button">{{recipe.href}}</el-button>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
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
      console.log(query)
      const url = 'https://recipe-puppy.p.rapidapi.com/?'
      this.axios.get(url, {
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
            res.data.results.forEach((r) => {
              this.recipes.push(r)
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
