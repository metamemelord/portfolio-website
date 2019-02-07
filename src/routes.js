import Vue from "vue";
import VueRouter from "vue-router";
import Vuex from "vuex";

Vue.use(VueRouter);
Vue.use(Vuex);

const routes = [
  {
    path: "/",
    component: () => import("./components/Home")
  },
  {
    path: "/blogs",
    component: () => import("./components/Blog components/Blogs")
  },
  {
    path: "*",
    component: () => import("./components/404")
  }
];

export default new VueRouter({ mode: "history", routes });
