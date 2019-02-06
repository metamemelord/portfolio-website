import Vue from "vue";
import Vuex from "vuex";

import blogPosts from "./modules/posts";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    day: 6,
    night: 20
  },
  modules: {
    blogPosts
  }
});
