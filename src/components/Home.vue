<template>
  <main>
    <intro-component/>
    <linkedin-component/>
    <github-component/>
  </main>
</template>

<script>
import introComponent from "./Home components/Intro.vue";
import linkedinComponent from "./Home components/Linkedin.vue";
import githubComponent from "./Home components/Github.vue";
export default {
  components: {
    introComponent,
    linkedinComponent,
    githubComponent
  }
};
</script>

<style>
main {
  max-width: 70rem;
  min-width: 20rem;
}

@media screen and (min-width: 550px) {
  main {
    margin: 0 auto;
  }
}
</style>

