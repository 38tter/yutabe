var app = new Vue({
    delimiters: ['[',']'],
    el: '#app',

    data: {
        products: [],
        productName: '',
        productMemo: '',
        message: 'hello, world!'
    },

    // computed: {
    //     computedProducts() {
    //         return this.products
    //     },
    // },
    //
    // created: function() {
    //     this.fetchProducts()
    // },
    //
    // methods: {
    //     fetchProducts() {
    //         axios.get('/products').then(res=>{if(res.status!=200){
    //             throw new Error('response error')
    //         } else {
    //             alert(res)
    //         }
    //         })
    //     }
    // }
})