<template>
  <div class="blog__post">
    <h1 v-html="getPost.title" v-if="getPost.title"></h1>
    <h1 v-html="previewContent" v-else></h1>
    <div class="blog__post-tags" v-if="getPost.tags && getPost.tags.length">
      <span v-for="(tag,idx) in getPost.tags" :key="idx">{{tag}}</span>
    </div>
  </div>
</template>

<script>
export default {
  props: ["postId"],
  computed: {
    getPost() {
      return this.$store.getters.blogPostById(this.postId);
    },
    previewContent() {
      var post = this.getPost;
      var content = post.content;
      var trimmedContent = content.substr(0, 50);
      var length = trimmedContent.length;
      var result = "";
      for (var i = 0; i < length; i++) {
        if (trimmedContent[i] == "<") {
          i++;
          while (i < length && trimmedContent[i] != ">") {
            i++;
          }
          i++;
          if (i == length) break;
        }
        result += trimmedContent[i];
      }
      return result + "...";
    }
  }
};
</script>

<style scoped>
.blog__post {
  display: flex;
  flex-flow: column;
  align-items: center;
}

.blog__post > h1 {
  color: rgb(255, 161, 38);
  font-size: 3rem;
  margin: 0.8rem;
  margin-bottom: 0.3rem;
  text-align: center;
}

.blog__post-tags {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: center;
  align-content: center;
  margin: 0;
}
.blog__post-tags > span {
  margin: 0.5rem;
  padding: 0.5rem;
  background: rgb(255, 161, 38);
  color: #2e3342;
  border-radius: 0.2rem;
}
</style>