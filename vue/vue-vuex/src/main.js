import Vue from 'vue'
import App from './App.vue'
import store from './store'

Vue.config.productionTip = false

new Vue({
  // 把store注入到所有子组件，通过this.$store可以操作
  store,
  render: h => h(App)
}).$mount('#app')
