<template>
  <div class="tech" id="tech">
    <div class="tech__label">
      <a
        class="light-a"
        v-infocus="'float-animation saturated-tech'"
        style="cursor:pointer;"
      >
        <h1>
          <i class="fa fa-code"></i>
        </h1>
      </a>
    </div>
    <div class="tech__content">
      <div v-for="(tech,idx) in technologies" :key="idx" class="tech__content__image-container">
        <a :href="tech.url" target="blank" style="text-decoration:none;outline: none;">
          <i v-if="tech.media_type=='font'" :class="tech.css_class" style="outline: none;" />
          <img v-else-if="tech.media_type=='img'" :src="tech.src" style="outline: none;" :alt="tech.name" />
        </a>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  data() {
    return {
      technologies: []
    }
  },
  created() {
    var dt = true,
      prev = -1;
    setInterval(() => {
      const hours = new Date().getHours();
      dt = hours > this.$store.state.day && hours < this.$store.state.night;
      if (prev != dt) {
        var label;
        if (dt) {
          label = document
            .getElementById("tech")
            .getElementsByTagName("a")[0];
          label.classList.remove("dark-a");
          label.classList.add("light-a");
        } else {
          label = document
            .getElementById("tech")
            .getElementsByTagName("a")[0];
          label.classList.remove("light-a");
          label.classList.add("dark-a");
        }
        prev = dt;
      }
    }, 1000);
  },
  beforeMount() {
    this.$http.get("/api/technologies").then(res => {
        this.technologies = res.body;
      }).catch(() => void 0);
  }
};
</script>
<style>
.tech {
  display: flex;
  flex-wrap: wrap;
  margin: 1rem;
  border-radius: 0.3rem;
  box-shadow: 0px 0px 4px 4px var(--shadow-color);
  transition: 0.2s all;
  justify-content: center;
}

.tech__label {
  width: 100%;
  display: flex;
  justify-content: center;
}

.tech__label h1 {
  margin: 0;
  font-size: 6rem;
  transition: all 5s;
  color: inherit;
}

.saturated-tech {
  color: var(--accent-color) !important;
}

.tech:hover {
  box-shadow: 0px 0px 3px 2px var(--shadow-color);
}

.tech__content {
  width: 100%;
  padding: 1rem;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
}

.tech__content__image-container {
  flex-basis: 20%;  
  margin: 1rem;
  display: flex;
  justify-content: center;
  align-items: center;
}

.tech__content__image-container i,
.tech__content__image-container img {
  font-size: 4rem;
  max-width: 3.5rem;
  cursor: pointer;
  transition: filter 0.2s linear;
}

.tech__content__image-container span {
  position: absolute;
  top: -2rem;
}

@media screen and (min-width: 550px) {
  .tech {
    max-width: 70rem;
    margin: 1rem;
  }

  .tech__content__image-container i,
  .tech__content__image-container img {
    filter: grayscale(100);
  }

  .tech__content__image-container i:hover,
  .tech__content__image-container img:hover {
    filter: grayscale(0);
  }
}

@media screen and (min-width: 73rem) {
  .tech {
    margin: 1rem auto;
  }

  .tech__content__image-container {
    flex-basis: 10%;
    text-align: center; 
    margin: 1rem;
  }

  .tech__content__image-container i,
  .tech__content__image-container img {
    font-size: 5rem;
    max-width: 4.5rem;
    filter: grayscale(100);
  }

  .tech__content__image-container i:hover,
  .tech__content__image-container img:hover {
    filter: grayscale(0);
  }
}

@media screen and (min-width: 45rem) {
  .tech__content__image-container i,
  .tech__content__image-container img {
    font-size: 4.5rem;
    max-width: 4rem;
  }
}
</style>
