import Vue from 'vue'
import Vuex from 'vuex'
import {
  SET_SHOW_COMMAND_INPUT,
  SET_DEBUG_INFO,
  SET_USER_COMMAND_KEYS
} from './mutation-types'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    debug: true,
    debugInfo: '',
    userCommandKeys: [] as string[],
    showCommandInput: false
  },
  mutations: {
    [SET_DEBUG_INFO] (state, newValue: string) {
      state.debugInfo = newValue
    },
    [SET_SHOW_COMMAND_INPUT] (state, show: boolean) {
      state.showCommandInput = show
    },
    [SET_USER_COMMAND_KEYS] (state, keys: string[]) {
      state.debugInfo = keys.join(' + ')
      state.userCommandKeys = keys
      if (keys.join('') === 'Shift ') {
        state.showCommandInput = true
      }
    }
  },
  actions: {},
  modules: {}
})
