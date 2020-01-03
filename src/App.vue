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
    this.$http.get("api/blogs").then(res => {
        if (res.status == 200) {
          this.$store.dispatch("setPosts", res.body)
        }
      }).catch(err => {
        this.$store.dispatch("initPosts")
      });
    var dt = true,
      prev = -1;
    setInterval(() => {
      const hours = new Date().getHours();
      dt = hours > this.$store.state.day && hours < this.$store.state.night;
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
