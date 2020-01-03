<template>
  <transition name="fade" mode="out-in">
    <component :is="activeComponent" :postId="postId" @click.native="onlyTitle=!onlyTitle"></component>
  </transition>
</template>

<script>
import blogTitle from "./BlogTitle.vue";
import blogComplete from "./BlogComplete.vue";
export default {
  props: ["blogPost"],
  data() {
    return {
      onlyTitle: true
    };
  },
  computed: {
    postId() {
      return this.blogPost._id;
    },
    activeComponent() {
      return this.onlyTitle ? "blog-title" : "blog-complete";
    }
  },
  components: {
    blogTitle,
    blogComplete
  }
};
</script>

<style>
.blog__post {
  padding: 1rem;
  margin: 1rem;
  border-radius: 0.3rem;
  box-shadow: 0px 0px 7px 4px rgba(0, 0, 0, 0.15);
  transition: 0.2s all;
  cursor: pointer;
}

.blog__post:hover {
  box-shadow: 0px 0px 3px 2px rgba(0, 0, 0, 0.15);
}

.blog__post > h1 {
  color: rgb(255, 161, 38);
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

.fade-enter-active {
  animation: fade-in 0.3s ease-out forwards;
}
.fade-leave-active {
  animation: fade-out 0.3s ease-out forwards;
}

@keyframes fade-in {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
@keyframes fade-out {
  from {
    opacity: 1;
  }
  to {
    opacity: 0;
  }
}
</style>