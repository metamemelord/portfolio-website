<template>
  <div id="app">
    <app-header/>
    <keep-alive>
      <router-view/>
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
  created() {
    var prev = -1;
    const day = this.$store.state.day,
      evening = this.$store.state.evening,
      night = this.$store.state.night;
    setInterval(() => {
      const hour = new Date().getHours();
      if (hour != prev) {
        if (hour >= night || hour < day) {
          document.getElementsByTagName("html")[0].className = "night";
        } else if (hour >= evening && hour < night) {
          document.getElementsByTagName("html")[0].className = "dark";
        } else {
          document.getElementsByTagName("html")[0].className = "";
        }
        prev = hour;
      }
    }, 1500);
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

.night {
  background: black;
  color: rgb(173, 173, 173);
  --shadow-color: rgba(255, 255, 255, 0.15);
}

.dark {
  background: #262d41;
  color: #eee;
}

.light-a {
  color: #333;
}

.dark-a {
  color: #eee;
}

body {
  margin: 0;
}

.active {
  color: rgb(255, 161, 38) !important;
}

.float-animation {
  animation: float-slow 10s linear infinite;
}

@keyframes float-slow {
  0% {
    transform: translateY(0);
  }
  10% {
    transform: translateX(-6px);
  }
  20% {
    transform: translateY(-6px);
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
</style>
