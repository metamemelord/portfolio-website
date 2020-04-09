<template>
  <main>
    <intro-component/>
    <technologies-component/>
    <linkedin-component/>
    <github-component/>
  </main>
</template>

<script>
import introComponent from "./Home components/Intro";
import technologiesComponent from "./Home components/Technologies";
import linkedinComponent from "./Home components/Linkedin";
import githubComponent from "./Home components/Github/index";
export default {
  components: {
    introComponent,
    technologiesComponent,
    linkedinComponent,
    githubComponent
  }
};
</script>

<style>
main {
  margin-top: 5rem;
  max-width: 70rem;
  min-width: 20rem;
}

@media screen and (min-width: 550px) {
  main {
    position: relative;
    margin: 0 auto;
    margin-top: 5rem;
  }
}
</style>

