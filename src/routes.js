import Vue from "vue";
import VueRouter from "vue-router";

import store from "./store";

import Home from "./components/Home.vue";
import Blogs from "./components/Blogs.vue";

Vue.use(VueRouter);

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
    component: Home
  }
];

export default new VueRouter({ mode: "history", routes });
