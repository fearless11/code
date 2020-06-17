import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

// vuex中文文档 https://vuex.vuejs.org/zh/guide/mutations.html
export default new Vuex.Store({
  // 组件共享状态
  state: {
    count: 0,
    number: 10
  },
  getters: {
    result: state => {
      return state.count
    }
  },
  // mutation 同步函数
  mutations: {
    increment(state) {
      state.count++
    },
    reduce(state,payload) {
      state.number -= payload.num
    }
  },
  // action 执行异步操作
  actions: {
    increment (context) {
      // 提交commit 调用 this.$store.dispatch('increment')
      context.commit('increment')
    },
    waitincrement ({commit}) {
      setTimeout(()=>{
        commit('increment')
      },5000)
    }
  },
  modules: {
  }
})
