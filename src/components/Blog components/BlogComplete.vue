<template>
  <div class="blog__post" :id="getPost._id">
    <h1>{{getPost.title}}</h1>
    <h3 v-if="getPost.subtitle">{{getPost.subtitle}}</h3>
    <h4>
      {{getPostDate}} by
      <font>
        <a
          :href="getPostAuthorContact"
          style="color:rgb(255, 161, 38);text-decoration: none;"
          @click.stop
          target="blank"
        >{{getPostAuthor}}</a>
      </font>
    </h4>
    <p></p>
    <div class="blog__post-tags" v-if="getPost.tags">
      <b>Tags:</b>
      <span v-for="(tag,idx) in getPost.tags" :key="idx">{{tag}}</span>
    </div>
  </div>
</template>

<script>
const moment = require("moment");
export default {
  props: ["postId"],
  computed: {
    getPost() {
      return this.$store.getters.blogPostById(this.postId);
    },
    getPostDate() {
      var postDate = new moment(this.getPost.date);
      return postDate.format("MMMM DD, YYYY");
    },
    getPostAuthor() {
      return this.getPost.author ? this.getPost.author : "metamemelord";
    },
    getPostAuthorContact() {
      const emailPattern = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return this.getPost.author_contact
        ? emailPattern.test(this.getPost.author_contact)
          ? `mailto:${this.getPost.author_contact}`
          : this.getPost.author_contact
        : "mailto:me@metamemelord.com";
    }
  },
  created() {
    let vm = this;
    setTimeout(function() {
      var post = document.getElementById(vm.getPost._id);
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

.blog__post > h3,
.blog__post > h4 {
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

