import Vue from 'vue'
import Vuex from 'vuex'
import state from "./state";
import getters from "./getters";
import mutations from "./mutations";
import actions from "./actions";
//引入vuex 数据持久化插件
import createPersistedState from "vuex-persistedstate"
Vue.use(Vuex)

export default new Vuex.Store({
  state,
  getters,
  mutations,
  actions,
  plugins: [createPersistedState({
    reducer(val) {
      return {
        // 只储存state中的token
        token: val.token,
        user: val.user,
        curProblem: val.curProblem,
        curDoProblemId: val.curDoProblemId,
        lang:val.lang,
        tagsList: val.tagsList,
        breadList:val.breadList
      }
    }
  })]
})