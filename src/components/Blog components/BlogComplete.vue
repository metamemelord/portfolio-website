<template>
  <main>
    <div v-if="getPost._id" class="blog__post-complete" :id="getPost._id">
      <div class="blog__post-complete-closer" @click="closePost()"><i class="fa fa-times" aria-hidden="true"></i></div>
      <h1 v-html="getPost.title" v-if="getPost.title"></h1>
      <template v-if="getPost.subtitle">
        <h4 v-if="getPost.title">{{getPostExcerpt}}</h4>
        <h2 v-else v-html="getPost.subtitle"></h2>
      </template>
      <h4>
        {{getPostDate}} by
        <font>
          <a
            :href="getPostAuthorContact"
            style="color:var(--accent-color);text-decoration: none;"
            @click.stop
            target="blank"
          >{{getPostAuthor}}</a>
        </font>
      </h4>
      <p v-if="getPost.title" v-html="getPost.content"></p>

      <!-- Removing as the these are non-clickable as on now. Will make these directly clickable -->
      <!-- <div class="blog__post-tags-complete" v-if="getPost.tags && getPost.tags.length">
        <b>Tags:</b>
        <span v-for="(tag,idx) in getPost.tags" :key="idx">{{tag|capitalize}}</span>
      </div> -->
      
    </div>
  </main>
</template>

<script>
const moment = require("moment");
export default {
  name: "blog-complete",
  data() {
    return {
      post: {}
    }
  },
  computed: {
    getPost() {
      return this.post
    },
    getPostDate() {
      var postDate = new moment(this.getPost.date);
      return postDate.format("MMMM DD, YYYY");
    },
    getPostAuthor() {
      return this.getPost.author ? this.getPost.author : "Gaurav Saini";
    },
    getPostAuthorContact() {
      const emailPattern = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return this.getPost.author_contact
        ? emailPattern.test(this.getPost.author_contact)
          ? `mailto:${this.getPost.author_contact}`
          : this.getPost.author_contact
        : "mailto:hello@gauravsaini.dev";
    },
    getPostExcerpt() {
      return this.getPost.subtitle.substr(3, 20) + "..."
    }
  },
  methods: {
    closePost() {
      this.$router.push("/blogs");
    }
  },
  created() {
    this.$http.get("/api/wordpress/" + this.$route.params.id).then(res => {
        this.post = res.body;
        this.post.content = this.post.content.split("\n").join("<br>")
      }).catch(() => this.$router.push("/lost"));
  }
};
</script>

<style scoped>
main {
  max-width: 72rem;
  margin: auto;
  margin-top: 7.3rem;
  min-height: calc(100vh - 9rem);
}
.blog__post-complete {
  position: relative;
  margin: 1rem;
  padding: 1rem;
  border-radius: 0.3rem;
  box-shadow: 0px 0px 7px 4px var(--shadow-color);
  padding: 1rem;
  text-align: center;
  transition: 0.2s all;
}

.blog__post-complete:hover {
  box-shadow: 0px 0px 3px 2px var(--shadow-color);
}

.blog__post-complete > h1 {
  color: var(--accent-color);
  font-size: 4rem;
  margin: 0.5rem;
}

.blog__post-complete > h3,
.blog__post-complete > h4 {
  margin: 0.2rem 0;
}

.blog__post-complete > p {
  text-align: justify;
}

.blog__post-tags-complete {
  display: flex;
  flex-wrap: wrap;
  align-items: center;
  justify-content: center;
  align-content: center;
  margin-top: 1.5rem;
  margin-bottom: 1rem;
}
.blog__post-tags-complete > span {
  margin: 0.5rem;
  padding: 0.5rem;
  background: var(--accent-color);
  color: #2e3342;
  border-radius: 0.2rem;
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.blog__post-complete-closer {
  position: absolute;
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0px 0px 3px 2px var(--shadow-color);
  border-radius: 50%;
  right: 1rem;
  top: 1rem;
  height: 1.5rem;
  width: 1.5rem;
  cursor: pointer;
  color: rgb(255, 91, 91);
}

@media screen and (min-width: 550px) {
  main {
    min-height: calc(100vh - 7.3rem);
  }

  .blog__post > p {
    width: 90%;
    margin: auto;
  }
  .blog__post-complete {
    padding: 1.5rem 2rem;
  }
}

@media screen and (min-width: 750px) {
  .blog__post-complete-closer {
    right: 1.5rem;
    top: 1.5rem;
    height: 2.5rem;
    width: 2.5rem;
  }
  .blog__post-complete {
    padding: 1.5rem 5rem;    
  }
}
</style>