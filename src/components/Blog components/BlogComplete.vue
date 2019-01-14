<template>
  <div class="blog__post" :id="getPost.id">
    <h1>{{getPost.title}}</h1>
    <h3 v-if="getPost.subtitle">{{getPost.subtitle}}</h3>
    <h3>{{getPostDate}}</h3>
    <p></p>
    <div class="blog__post-tags" v-if="getPost.tags">
      <b>Tags:</b>
      <span v-for="(tag,idx) in getPost.tags" :key="idx">{{tag}}</span>
    </div>
  </div>
</template>

<script>
const months = {
  0: "Janurary",
  1: "February",
  2: "March",
  3: "April",
  4: "May",
  5: "June",
  6: "July",
  7: "August",
  8: "September",
  9: "October",
  10: "November",
  11: "December"
};
export default {
  props: ["postId"],
  computed: {
    getPost() {
      return this.$store.getters.blogPostById(this.postId);
    },
    getPostDate() {
      var postDate = new Date(this.getPost.date);
      return `${
        months[postDate.getMonth()]
      } ${postDate.getDay()}, ${postDate.getFullYear()}`;
    }
  },
  created() {
    let vm = this;
    setTimeout(function() {
      var post = document.getElementById(vm.getPost.id);
      var contentNode = post.childNodes[3];
      var content = vm.getPost.content.split("\n").join("<br>");
      contentNode.innerHTML = content;
    }, 5);
  }
};
</script>

<style scoped>
.blog__post {
  cursor: unset;
}

.blog__post > h1 {
  color: rgb(255, 161, 38);
  font-size: 3rem;
  margin: 0.5rem;
  text-align: center;
}

.blog__post > h3 {
  margin: 0.2rem 0;
  text-align: center;
}

.blog__post > p {
  text-align: justify;
}
.blog__post-tags {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: center;
  align-content: center;
  margin-top: 1.5rem;
  margin-bottom: 1rem;
}
.blog__post-tags > span {
  margin: 0.5rem;
  padding: 0.5rem;
  background: rgb(255, 161, 38);
  color: #2e3342;
  border-radius: 0.2rem;
}
</style>

