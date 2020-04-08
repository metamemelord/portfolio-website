<template>
  <main>
    <intro-component/>
    <linkedin-component/>
    <github-component/>
  </main>
</template>

<script>
import introComponent from "./Home components/Intro";
import linkedinComponent from "./Home components/Linkedin";
import githubComponent from "./Home components/Github/index";
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

