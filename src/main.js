import Vue from "vue";
import VueRouter from "vue-router";
import Vuex from "vuex";
import VueResource from "vue-resource";
import VueTypedJs from 'vue-typed-js'
import App from "./App.vue";
import store from "./store/store";
import router from "./routes";
import "normalize.css"

Vue.use(VueResource);
Vue.use(Vuex);
Vue.use(VueRouter);
Vue.use(VueTypedJs);

Vue.directive("infocus", {
  isLiteral: true,
  inserted: (el, binding) => {
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

Vue.filter('capitalize', function (value) {
  if (!value) return ''
  value = value.toString()
  return value.charAt(0).toUpperCase() + value.slice(1)
})

Vue.config.productionTip = false;

new Vue({
  store,
  router,
  render: h => h(App)
}).$mount("#app");
