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
        <loading
        :active.sync="isLoading"
        :can-cancel="true"
        :on-cancel="onCancel"
        :is-full-page="fullPage"
        ></loading>
        <el-button type="primary" @click.native="fetchRecipes" @click.prevent="doAjax">
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
              <router-link :to="{
                name: 'recipe',
                params: {
                  name: recipe.name,
                  img: recipe.thumbnail_url,
                  description:recipe.description,
                  sections:recipe.sections,
                  instructions:recipe.instructions,
                  hasItems: hasItems
                }}" class="button">detail</router-link>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>

import Loading from 'vue-loading-overlay'

export default {
  name: 'index',
  data: function () {
    return {
      products: [],
      address: '',
      ingredients: '',
      recipes: [],
      hasItems: {},
      isLoading: false,
      fullPage: true
    }
  },

  components: {
    'loading': Loading
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
    doAjax () {
      this.isLoading = true
      setTimeout(function () {
        this.isLoading = false
        console.log('load off')
      }, 1000)
    },
    onCancel () {
      console.log('User cancelled the loader.')
    },
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
              this.saveRecipes(r)
              this.checkItems(r)
            })
            this.isLoading = false
          }
        }
      )
    },
    clearRecipes () {
      this.recipes = []
    },
    checkItems (recipe) {
      if (recipe.sections) {
        recipe.sections.forEach((s) => {
          if (s.components) {
            s.components.forEach((c) => {
              let item = c['ingredient']['name']
              this.hasItems[item] = this.ingredients.includes(item)
            })
          }
        })
      }
    },
    saveRecipes (recipe) {
      let texts = []
      if (recipe.instructions) {
        recipe.instructions.forEach((instruction) => {
          texts.push(instruction.display_text)
        })
      }
      let ingredients = []
      if (recipe.sections) {
        recipe.sections.forEach((s) => {
          if (s.components) {
            s.components.forEach((c) => {
              ingredients.push({'name': c['ingredient']['name']})
            })
          }
        })
      }
      const message = {
        'name': recipe.name,
        'image_url': recipe.thumbnail_url,
        'description': recipe.description,
        'tasty_id': recipe.id,
        'instruction': {
          'texts': texts
        },
        'ingredients': ingredients
      }
      let messageJson = JSON.stringify(message)
      return this.axios.post('http://localhost:8081/api/recipe', messageJson).then(res => {
        if (res.status !== 201) {
          throw new Error('failed to save recipe')
        }
      }).catch(err => Promise.reject(err))
    }
  }
}
</script>

<style scoped>

</style>
