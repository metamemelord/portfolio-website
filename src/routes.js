import Vue from "vue";
import VueRouter from "vue-router";
import Vuex from "vuex";

import Home from "./components/Home.vue";
import Blogs from "./components/Blog components/Blogs.vue";
import ErrorPage from "./components/404.vue";

Vue.use(VueRouter);
Vue.use(Vuex);

const routes = [
  {
    path: "/",
    component: Home
  },
  {
    path: "/blogs",
    component: Blogs
  },
  {
    path: "*",
    component: ErrorPage
  }
];

export default new VueRouter({ mode: "history", routes });
