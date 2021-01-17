<template>
  <div>
    <h1>{{$route.params.name}}</h1>
    <img :src="$route.params.img" width="100%">
    <h2>ingredients</h2>
    <div v-for="(section,idx) in $route.params.sections" :key="idx">
      <div v-for="component in section.components" :key="component.id" :style="{color: changeColor($route.params.hasItems[component.ingredient.name])}">
        {{component.raw_text}}
        <el-button @click.native="buyItem(component)" v-if="!($route.params.hasItems[component.ingredient.name])">Buy</el-button>
      </div>
    </div>
    <h2>instructions</h2>
    <div v-for="instruction in $route.params.instructions" :key="instruction.id">
      {{instruction.display_text}}
    </div>
    <router-link :to="{name:'index'}" class="button">back</router-link>
  </div>
</template>

<script>
export default {
  name: 'recipe',
  props: {
    name: {
      type: String,
      default: ''
    }
  },
  methods: {
    changeColor (hasItem) {
      if (!hasItem) {
        return 'red'
      }
      return 'black'
    },
    buyItem (component) {
      let item = component.ingredient.name
      return this.axios.get('http://localhost:8081/api/market?item=' + item).then(res => {
        if (res.status !== 200) {
          throw new Error('failed to send get request')
        }
      }).catch(err => Promise.reject(err))
    }
  }
}
</script>

<style scoped>

</style>
