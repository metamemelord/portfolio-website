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
  },
  beforeMount() {
  	this.$http.get("api/wordpress").then(res => {
		if (res.status == 200) this.$store.dispatch("setPosts", res.body)
		else throw "Wordpress call failed"
	}).catch(err => this.$store.dispatch("initPosts"));
  }
};
</script>
<style>
main {
  max-width: 70rem;
  margin: auto;
  margin-top: 5rem;
}
</style>
