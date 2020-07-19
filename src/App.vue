<template>
  <div id="app">
    <app-header/>
    <div v-if="contactMeDialog" class="page-shield"></div>
    <keep-alive exclude="blog-complete">
      <transition name="fade" mode="out-in">
        <router-view/>
      </transition>
    </keep-alive>
    <app-footer/>
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";

export default {
  name: "app",
  components: {
    appHeader: Header,
    appFooter: Footer
  },
  computed: {
    contactMeDialog() {
      return this.$store.state.contactMeDialog;
    }
  },
  created() {
    const day = this.$store.state.day,
      evening = this.$store.state.evening,
      night = this.$store.state.night;
      
    let prev = -1;
    setInterval(() => {
      const hour = new Date().getHours();
      if (hour != prev) {
        if (hour >= night || hour < day) {
          document.getElementsByTagName("html")[0].className = "night";
          document.body.className = "night-font"
        } else if (hour >= evening && hour < night) {
          document.getElementsByTagName("html")[0].className = "dark";
          document.body.className = "night-font"
        } else {
          document.getElementsByTagName("html")[0].className = "";
          document.body.className = ""
        }
        prev = hour;
      }
    }, 1000);
  }
};
</script>
