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
    var dt = 0,
      prev = -1;
    const day = this.$store.state.day,
      afternoon = this.$store.state.afternoon,
      evening = this.$store.state.evening,
      night = this.$store.state.night;
    setInterval(() => {
      const hours = new Date().getHours();
      if (dt != prev) {
        if (hours >= night && hours < day) {
          document.getElementsByTagName("html")[0].className = "night";
          dt = 1;
        } else if (hours >= afternoon && hours < evening) {
          document.getElementsByTagName("html")[0].className = "night";
          // document.getElementsByTagName("html")[0].className = "afternoon";
          dt = 2;
        } else if (hours >= evening && hours < night) {
          document.getElementsByTagName("html")[0].className = "dark";
        } else {
          document.getElementsByTagName("html")[0].className = "";
          dt = 0;
        }
        prev = dt;
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
  box-shadow: 0 0 10px white !important;
}

.dark {
  background: #262d41;
  color: #eee;
}

.afternoon {
  background:rgb(255, 224, 183);
  color: #333;
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
