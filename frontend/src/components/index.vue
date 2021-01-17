<template>
  <div id="app">
    <div id="header"/>
    <div>
      <h1>夕ご飯何作る？（ゆうつく）</h1>
      <h3>食材を入力してください（カンマ区切り）</h3>
    </div>
    <el-form>
      <el-form-item>
        <textarea type="textarea" :rows="2" v-model="ingredients" placeholder="Input your ingredients" width="100%"></textarea>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click.native="fetchRecipes">
          送信
        </el-button>
      </el-form-item>
    </el-form>
    <el-row>
      <el-col :span="8" v-for="recipe in recipes" :key="recipe.name" :offset="2">
        <el-card :body-style="{ padding: '0px' }">
          <img :src="recipe.thumbnail_url" class="image" width="75%">
          <div style="padding: 14px;">
            <span>{{recipe.name}}</span>
            <span>{{recipe.keywords}}</span>
            <div class="bottom clearfix">
              <el-button type="text" class="button">{{recipe.href}}</el-button>
            </div>
            <div>
              <router-link :to="{name:'recipe',params:{name:'hoge',img:'fuga',description:'heke'}}" class="button">detail</router-link>
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
      }).catch(err => Promise.reject(err))
    },
    fetchRecipes () {
      this.clearRecipes()
      const url = 'https://tasty.p.rapidapi.com/recipes/list?tags=under_30_minutes&q=' + this.ingredients + '&from=0&size=2'
      this.axios.get(url, {
        headers: {
          'X-RapidAPI-Host': 'tasty.p.rapidapi.com',
          'X-RapidAPI-Key': process.env.VUE_APP_TASTY_API_KEY,
          'useQueryString': 'useQueryString'
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
    },
    clearRecipes () {
      this.recipes = []
    }
  }
}
</script>

<style scoped>

</style>
