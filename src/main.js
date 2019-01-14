import Vue from "vue";
import VueRouter from "vue-router";
import Vuex from "vuex";
import VueResource from "vue-resource";
import App from "./App.vue";
import store from "./store/store";
import router from "./routes";

Vue.use(VueResource);
Vue.use(Vuex);
Vue.use(VueRouter);

Vue.config.productionTip = false;

new Vue({
  store,
  router,
  render: h => h(App)
}).$mount("#app");
