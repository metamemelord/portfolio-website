<template>
  <main>
    <blog-post v-for="(blogPost, idx) in blogPosts" :key="idx" :blogPost="blogPost"></blog-post>
  </main>
</template>

<script>
import blogPost from "./Blog.vue";

export default {
  components: {
    blogPost
  },
  computed: {
    blogPosts() {
      return this.$store.getters.blogPosts;
    }
  }
};
</script>
<style>
main {
  max-width: 70rem;
}
</style>