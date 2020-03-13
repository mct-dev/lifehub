import Vue from "vue";
import Vuex from "vuex";
import {
  SET_IS_ACCEPTING_COMMAND,
  SET_DEBUG_INFO,
  SET_USER_COMMAND_KEYS
} from "./mutation-types";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    debug: true,
    debugInfo: "",
    userCommandKeys: [] as string[],
    isAcceptingCommand: true
  },
  mutations: {
    [SET_DEBUG_INFO](state, newValue: string) {
      state.debugInfo = newValue;
    },
    [SET_IS_ACCEPTING_COMMAND](state, isAccepting: boolean) {
      state.isAcceptingCommand = isAccepting
    },
    [SET_USER_COMMAND_KEYS](state, keys: string[]) {
      state.debugInfo = keys.join(" + ");
      console.log(state.debugInfo);

      state.userCommandKeys = keys;
      if (keys.join("") === "Shift ") {
        state.isAcceptingCommand= false;
      }
    }
  },
  actions: {},
  modules: {}
});
