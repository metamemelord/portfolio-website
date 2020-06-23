<template>
  <blog-title :postId="postId" @click.native="loadPost"></blog-title>
</template>

<script>
import blogTitle from "./BlogTitle.vue";

export default {
  props: ["blogPost"],
  computed: {
    postId() {
      return this.blogPost._id;
    }
  },
  methods: {
    loadPost() {
      this.$router.push("/blog/" + this.blogPost._id);
    }
  },
  components: {
    blogTitle
  }
};
</script>

<style>
.blog__post {
  padding: 1rem;
  margin: 1rem;
  border-radius: 0.3rem;
  box-shadow: 0px 0px 7px 4px var(--shadow-color);
  transition: 0.2s all;
  cursor: pointer;
}

.blog__post:hover {
  box-shadow: 0px 0px 3px 2px var(--shadow-color);
}

.blog__post > h1 {
  color: var(--accent-color);
  font-size: 3rem;
  margin: 0.8rem;
  text-align: center;
}

.blog__post > h3 {
  margin: 0.2rem 0;
}

.blog__post > p {
  text-align: center;
}
</style>