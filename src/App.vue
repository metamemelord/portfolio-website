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
  components: {
    appHeader: Header,
    homepage: Homepage,
    appFooter: Footer
  }
};
</script>

<style>
* {
  box-sizing: border-box;
}
#app {
  width: 65%;
  height: 100vh;
  margin: 0px auto;
}
</style>
