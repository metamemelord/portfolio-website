import Vue from "vue";
import Vuex from "vuex";

import blogPosts from "./modules/posts";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    day: 7,
    afternoon: 13,
    evening: 19,
    night: 1
  },
  modules: {
    blogPosts
  }
});
