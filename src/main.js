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

Vue.directive("infocus", {
  isLiteral: true,
  inserted: (el, binding, vnode) => {
    let f = () => {
      let rect = el.getBoundingClientRect();
      let inView =
        rect.width > 0 &&
        rect.height > 0 &&
        rect.top >= 0 &&
        rect.bottom <=
          (window.innerHeight || document.documentElement.clientHeight);
      if (inView) {
        let classesToAdd = binding.value.split(" ");
        for (let idx = 0; idx < classesToAdd.length; idx++) {
          el.classList.add(classesToAdd[idx]);
        }
        window.removeEventListener("scroll", f);
      }
    };
    window.addEventListener("scroll", f);
    f();
  }
});

Vue.config.productionTip = false;

new Vue({
  store,
  router,
  render: h => h(App)
}).$mount("#app");
