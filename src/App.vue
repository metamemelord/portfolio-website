<template>
  <div id="app">
    <app-header/>
    <router-view></router-view>
    <app-footer/>
  </div>
</template>

<script>
import Homepage from "./components/Home.vue";
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";

export default {
  name: "app",
  data() {
    return {};
  },
  components: {
    appHeader: Header,
    homepage: Homepage,
    appFooter: Footer
  },
  created() {
    var dt = true,
      prev = -1;
    setInterval(() => {
      const hours = new Date().getHours();
      dt = hours > 6 && hours < 20;
      if (prev != dt) {
        if (dt) {
          document.getElementsByTagName("html")[0].className = "";
        } else {
          document.getElementsByTagName("html")[0].className = "dark";
        }
        prev = dt;
      }
    }, 1000);
  }
};
</script>

<style>
* {
  box-sizing: border-box;
}

html {
  font-family: "Open Sans", "Calibri", sans-serif;
  color: #333;
  transition: 0.5s background-color;
}

.dark {
  background: #2e3342;
  color: #eee;
}

body {
  margin: 0;
}

.active {
  color: #228080 !important;
}
</style>
