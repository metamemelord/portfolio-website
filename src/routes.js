import Vue from "vue";
import VueRouter from "vue-router";
import Vuex from "vuex";
import store from "./store/store"

Vue.use(VueRouter);
Vue.use(Vuex);

const routes = [
  {
    path: "/",
    component: () => import("./components/Home")
  },
  {
    path: "/blog",
    redirect: "/blogs"
  },
  {
    path: "/blog/:id",
    component: () => import("./components/Blog components/BlogComplete")
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

const router = new VueRouter({ mode: "history", routes });
router.beforeEach((_, __, next) => {
  store.state.contactMeDialog = false;
  next();
});

export default router;
