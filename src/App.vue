<template>
  <div id="app">
    <app-header />
    <transition name="fade" mode="out-in">
      <div v-if="contactMeDialog" class="page-shield"></div>
    </transition>
    <keep-alive exclude="blog-complete">
      <transition name="fade" mode="out-in">
        <router-view />
      </transition>
    </keep-alive>
    <contact-me />
    <app-footer />
  </div>
</template>

<script>
import Header from "./components/Header.vue";
import Footer from "./components/Footer.vue";
import contactMe from "./components/ContactMe.vue";

export default {
  name: "app",
  components: {
    appHeader: Header,
    appFooter: Footer,
    contactMe
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
          document.body.className = "night-font";
        } else if (hour >= evening && hour < night) {
          document.getElementsByTagName("html")[0].className = "dark";
          document.body.className = "night-font";
        } else {
          document.getElementsByTagName("html")[0].className = "";
          document.body.className = "";
        }
        prev = hour;
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
  transition: 0.5s background-color;
  font-display: swap;
}
.night {
  background: black;
  --background-color: black;
  color: #fff;
  --text-color: #eee;
  --shadow-color: rgba(255, 255, 255, 0.15);
}
.night-font {
  color: #eee;
  --text-color: #eee;
}
.dark {
  background: #262d41;
  --background-color: #262d41;
  color: #eee;
  --text-color: #eee;
}
.light-a {
  color: #333;
}
.dark-a {
  color: #eee;
}
body {
  margin: 0;
  transition: 2s background-color, color;
}
.active {
  color: var(--accent-color) !important;
}
.float-animation {
  animation: float-slow 10s linear infinite;
}
.page-shield {
  background: rgba(0, 0, 0, 0.5);
  position: fixed;
  width: 100vw;
  top: 0;
  height: 100vh;
  z-index: 2;
  overflow: hidden;
}

@keyframes float-slow {
  0% {
    transform: translateY(0);
  }
  10% {
    transform: translateX(-6px);
  }
  20% {
    transform: translateY(-4px);
  }
  30% {
    transform: translateX(4px) translateY(2px);
  }
  40% {
    transform: translateX(-2px) translateY(6px);
  }
  50% {
    transform: translateX(2px) translateY(-2px);
  }
  60% {
    transform: translateX(+4px) translateY(-1px);
  }
  70% {
    transform: translateX(-4px);
  }
  80% {
    transform: translateX(+4px) translateY(+1px);
  }
  90% {
    transform: translateX(-4px) translateY(+2px);
  }
  100% {
    transform: translateX(+4px) translateY(-2px);
  }
}
.fade-enter-active {
  animation: fade-in 100ms ease-out forwards;
}
.fade-leave-active {
  animation: fade-out 100ms ease-out forwards;
}
@keyframes fade-in {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
@keyframes fade-out {
  from {
    opacity: 1;
  }
  to {
    opacity: 0;
  }
}
</style>