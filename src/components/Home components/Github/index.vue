<template>
  <div id="github" class="github">
    <div class="github__label">
      <a href="https://github.gaurav.dev" target="blank">
        <img
          src="../../../assets/Octocat.webp"
          alt="Github octocat"
          v-infocus="'saturate-image float-animation'"
        >
      </a>
    </div>
    <div class="github__cards">
      <template v-if="!(repos.length)">
        <div class="github__cards-loading-wrapper">
          <div class="loader">
            <div class="circle" id="accent"></div>
            <div class="circle" id="light"></div>
            <div class="circle" id="accent"></div>
            <div class="circle" id="light"></div>
            <div class="circle" id="accent"></div>
          </div>
        </div>
      </template>
      <template v-else>
        <h1 id="repos">Repositories</h1>
        <project-card v-for="repo in ownedRepos" :key="repo.name" :cardData="repo"></project-card>
        <h1 id="repos-forked">Forked</h1>
        <project-card v-for="repo in forkedRepos" :key="repo.name" :cardData="repo"></project-card>
      </template>
    </div>
  </div>
</template>

<script>
import projectCard from "./ProjectCard";
export default {
  data() {
    return {
      repos: []
    };
  },
  methods: {
    loadUsers() {
      this.$http
        .get(
          "api/repos"
        )
        .then(function(responseData) {
          if (responseData.status == 200) {
            this.repos = responseData.body;
            this.repos.concat(responseData.body);
          }
        })
        .catch(() => {
          this.repos = [];
        });
    }
  },
  components: {
    projectCard
  },
  computed: {
    ownedRepos() {
      return this.repos.filter(repo => !repo.fork);
    },
    forkedRepos() {
      return this.repos.filter(repo => repo.fork);
    }
  },
  created() {
    setTimeout(() => {
      this.loadUsers();
    }, 500);
  }
};
</script>

<style>
.github {
  display: flex;
  flex-wrap: wrap;
  margin: 1rem;
  border-radius: 0.3rem;
  box-shadow: 0px 0px 4px 4px var(--shadow-color);
  transition: 0.2s all;
}

.github__label {
  width: 100%;
  display: flex;
  justify-content: center;
}

.github__label img {
  margin-top: 1.4rem;
  max-width: 9rem;
  filter: grayscale(100%);
  -webkit-filter: grayscale(100%);
  -moz-filter: grayscale(100%);
  -ms-filter: grayscale(100%);
  -o-filter: grayscale(100%);
  transition: all 5s;
}

.saturate-image {
  filter: grayscale(0) !important;
  -webkit-filter: grayscale(0) !important;
  -moz-filter: grayscale(0) !important;
  -ms-filter: grayscale(0) !important;
  -o-filter: grayscale(0) !important;
}

.github:hover {
  box-shadow: 0px 0px 3px 2px var(--shadow-color);
}

.github__cards {
  width: 100%;
  padding: 1rem;
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  text-align: center;
}

.github__cards h1 {
  display: block;
  width: 100%;
}

.github__cards-loading-wrapper {
  width: 100%;
  text-align: center;
}

.loader {
  display: flex;
  justify-content: center;
  text-align: center;
  position: relative;
  height: 2.5rem;
  margin: 1rem;
}

.loader .circle {
  position: absolute;
  width: 3rem;
  height: 3rem;
  opacity: 0;
  transform: rotate(225deg);
  animation-iteration-count: infinite;
  animation-name: orbit;
  animation-duration: 5.5s;
}

.loader .circle:after {
  content: "";
  position: absolute;
  width: 0.4rem;
  height: 0.4rem;
  border-radius: 0.4rem;
}

.loader #light:after {
  background: #ccc;
}

.loader #accent:after {
  background: var(--accent-color);
}

.loader .circle:nth-child(2) {
  animation-delay: 240ms;
}

.loader .circle:nth-child(3) {
  animation-delay: 480ms;
}

.loader .circle:nth-child(4) {
  animation-delay: 720ms;
}

.loader .circle:nth-child(5) {
  animation-delay: 960ms;
}

@keyframes orbit {
  0% {
    transform: rotate(225deg);
    opacity: 1;
    animation-timing-function: ease-out;
  }

  7% {
    transform: rotate(345deg);
    animation-timing-function: linear;
  }

  30% {
    transform: rotate(455deg);
    animation-timing-function: ease-in-out;
  }

  39% {
    transform: rotate(690deg);
    animation-timing-function: linear;
  }

  70% {
    transform: rotate(815deg);
    opacity: 1;
    animation-timing-function: ease-out;
  }

  75% {
    transform: rotate(945deg);
    animation-timing-function: ease-out;
  }

  76% {
    transform: rotate(945deg);
    opacity: 0;
  }

  100% {
    transform: rotate(945deg);
    opacity: 0;
  }
}

@media screen and (min-width: 550px) {
  .github {
    max-width: 70rem;
    margin: 1rem;
  }
}

@media screen and (min-width: 73rem) {
  .github {
    margin: 1rem auto;
  }
}
</style>
