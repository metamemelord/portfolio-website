import Vue from "vue";
import Vuex from "vuex";

import blogPosts from "./modules/posts";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    day: 7,
    evening: 17,
    night: 20,
    contactMeDialog: false,
  },
  actions: {
    setContactMeDialog: ({ commit }, contactMeDialog) => {
      commit("SET_CONTACT_ME_DIALOG", contactMeDialog);
    },
  },
  mutations: {
    SET_CONTACT_ME_DIALOG(state, contactMeDialog) {
      state.contactMeDialog = contactMeDialog;
    }
  },
  getters: {
    contactMeDialog: state => {
      return state.contactMeDialog;
    }
  },
  modules: {
    blogPosts
  }
});
