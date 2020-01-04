<template>
  <div class="blog__post">
    <h1 v-html="getPost.title" v-if="getPost.title"></h1>
    <h1 v-html="previewContent" v-else></h1>
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
  text-align: center;
}
</style>